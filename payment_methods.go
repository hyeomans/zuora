package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type paymentMethods struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

func newPaymentMethods(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *paymentMethods {
	return &paymentMethods{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
	}
}

// GetPaymentMethodSnapshot A Payment Method Snapshot is a copy of the particular Payment Method used in a
// transaction. If the Payment Method is deleted, the Payment Method Snapshot continues to retain the
// data used in each of the past transactions.
// More info at: https://www.zuora.com/developer/api-reference/#operation/Object_GETPaymentMethodSnapshot
func (t *paymentMethods) GetPaymentMethodSnapshot(ctx context.Context, snapshotID string) ([]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/object/payment-method-snapshot/%v", t.baseURL, snapshotID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
	}

	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")

	if ctx.Value(ContextKeyZuoraEntityIds) != nil {
		req.Header.Add("Zuora-Entity-Ids", ctx.Value(ContextKeyZuoraEntityIds).(string))
	}

	if ctx.Value(ContextKeyZuoraTrackID) != nil {
		req.Header.Add("Zuora-Track-Id", ctx.Value(ContextKeyZuoraTrackID).(string))
	}

	res, err := t.http.Do(req.WithContext(ctx))

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to read body response into memory: %v", err)}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error on HTTP request. Response code: %v - Error message %v", res.StatusCode, string(body))}
	}

	jsonResponse := Response{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	if !jsonResponse.Success {
		errorResponse := errorResponse{}

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json error response. Error: %v. Raw JSON: %v", err, string(body))}
		}

		return nil, errorResponse
	}

	return body, nil
}
