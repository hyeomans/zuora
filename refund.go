package zuora

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type refundService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
	isPce              bool
}

func newRefundService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string, isPce bool) *refundService {
	return &refundService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
		isPce:              isPce,
	}
}

func (t *refundService) Create(ctx context.Context, refundCreatePayload interface{}) (RefundCreateResonse, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/object/refund", t.baseURL)
	} else {
		url = fmt.Sprintf("%v/v1/object/refund", t.baseURL)
	}

	j, err := json.Marshal(refundCreatePayload)

	if err != nil {
		return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to convert refundCreatePayload: %v", err)}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))

	if err != nil {
		return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
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
			return RefundCreateResonse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return RefundCreateResonse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := RefundCreateResonse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	if !jsonResponse.Success {
		errorResponse := errorResponse{}

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return RefundCreateResonse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json error response. Error: %v. Raw JSON: %v", err, string(body))}
		}

		return RefundCreateResonse{}, errorResponse
	}

	return jsonResponse, nil
}
