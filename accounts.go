package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type accountsService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

func newAccountsService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *accountsService {
	return &accountsService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
	}
}

//Summary Retrieves detailed information about the specified customer account.
// The response includes the account information and a summary of the account’s subscriptions, invoices, payments,
// and usages for the last six recently updated subscriptions.
// Returns only the six most recent subscriptions based on the subscription updatedDate. Within those
//subscriptions, there may be many rate plans and many rate plan charges. These items are subject to the maximum
//limit on the array size.
//NOTE: An array of bytes is returned allowing the consumer to extend the defined type with custom fields.
func (t *accountsService) Summary(ctx context.Context, objectID string) (*[]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/accounts/%v/summary", t.baseURL, objectID)

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
	defer res.Body.Close()

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

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

		return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory: %v", err)}
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

	return &body, nil
}
