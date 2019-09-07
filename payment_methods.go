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
	isPce              bool
}

func newPaymentMethods(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string, isPce bool) *paymentMethods {
	return &paymentMethods{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
		isPce:              isPce,
	}
}

//GetPaymentMethod Retrieves a specific Payment Method by ObjectID
// More info at: https://www.zuora.com/developer/api-reference/#operation/Object_GETPaymentMethod
func (t *paymentMethods) GetPaymentMethod(ctx context.Context, objectID string) (PaymentMethod, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/object/payment-method/%v", t.baseURL, objectID)
	} else {
		url = fmt.Sprintf("%v/v1/object/payment-method/%v", t.baseURL, objectID)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		if err != nil {
			return PaymentMethod{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return PaymentMethod{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := PaymentMethod{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	return jsonResponse, nil
}

// GetPaymentMethodSnapshot A Payment Method Snapshot is a copy of the particular Payment Method used in a
// transaction. If the Payment Method is deleted, the Payment Method Snapshot continues to retain the
// data used in each of the past transactions.
// More info at: https://www.zuora.com/developer/api-reference/#operation/Object_GETPaymentMethodSnapshot
func (t *paymentMethods) GetPaymentMethodSnapshot(ctx context.Context, snapshotID string) (PaymentMethod, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/object/payment-method-snapshot/%v", t.baseURL, snapshotID)
	} else {
		url = fmt.Sprintf("%v/v1/object/payment-method-snapshot/%v", t.baseURL, snapshotID)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		if err != nil {
			return PaymentMethod{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return PaymentMethod{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := PaymentMethod{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return PaymentMethod{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	return jsonResponse, nil
}
