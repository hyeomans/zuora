package zuora

import (
	"bytes"
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

// ByKey This REST API reference describes how to retrieve detailed information about a specified subscription in the latest version.
// https://www.zuora.com/developer/api-reference/#operation/GET_SubscriptionsByKey
// Possible values for subscriptionKey are: subscription number or subscription ID
// NOTE: Why return a raw array of bytes? You can take advantage of binding to your custom struct with custom properties.
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		if err != nil {
			return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
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

// Update Use this call to make the following kinds of changes to a subscription:
//     - Add a note
//
//     - Change the renewal term or auto-renewal flag
//
//     - Change the term length or change between evergreen and termed
//
//     - Add a new product rate plan
//
//     - Remove an existing subscription rate plan
//
//     - Change the quantity or price of an existing subscription rate plan
func (t *subscriptionsService) Update(ctx context.Context, subscriptionKey string, subscriptionUpdate interface{}) (Response, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", t.baseURL, subscriptionKey)

	j, err := json.Marshal(subscriptionUpdate)

	if err != nil {
		return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to convert empty interface: %v", err)}
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(j))

	if err != nil {
		return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
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
			return Response{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return Response{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := Response{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	if !jsonResponse.Success {
		errorResponse := errorResponse{}

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return Response{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json error response. Error: %v. Raw JSON: %v", err, string(body))}
		}

		return Response{}, errorResponse
	}

	return jsonResponse, nil
}

// Cancel This REST API reference describes how to cancel an active subscription.
// Note: This feature is unavailable if you have the Orders feature enabled. See Orders Migration Guidance for more information.
func (t *subscriptionsService) Cancel(ctx context.Context, subscriptionKey string, subscriptionCancellation SubscriptionCancellation) (SubscriptionCancellationResponse, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v/cancel", t.baseURL, subscriptionKey)

	j, err := json.Marshal(subscriptionCancellation)

	if err != nil {
		return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to convert empty interface: %v", err)}
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(j))

	if err != nil {
		return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
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
			return SubscriptionCancellationResponse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return SubscriptionCancellationResponse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := SubscriptionCancellationResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	if !jsonResponse.Success {
		errorResponse := errorResponse{}

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return SubscriptionCancellationResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json error response. Error: %v. Raw JSON: %v", err, string(body))}
		}

		return SubscriptionCancellationResponse{}, errorResponse
	}

	return jsonResponse, nil
}
