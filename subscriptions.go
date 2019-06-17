package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type subscriptionsService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

func newSubscriptionsService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *subscriptionsService {
	return &subscriptionsService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
	}
}

//ByKey This REST API reference describes how to retrieve detailed information about a specified subscription in the latest version.
// https://www.zuora.com/developer/api-reference/#operation/GET_SubscriptionsByKey
// Possible values for subscriptionKey are: subscription number or subscription ID
// An array of bytes is returned allowing the consumer to extend the defined type with custom fields.
func (t *subscriptionsService) ByKey(ctx context.Context, subscriptionKey string) ([]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", t.baseURL, subscriptionKey)

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

	return body, nil
}
