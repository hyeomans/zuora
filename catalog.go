package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type catalogService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

func newCatalogService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *catalogService {
	return &catalogService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
	}
}

func (t *catalogService) GetProduct(ctx context.Context, pageSize int) (*ProductsResponse, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	if pageSize == 0 {
		pageSize = 10 //Default value
	}

	url := fmt.Sprintf("%v/v1/catalog/products?pageSize=%v", t.baseURL, pageSize)

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

	if ctx.Value(ContextKeyZuoraVersion) != nil {
		req.Header.Add("zuora-version", ctx.Value(ContextKeyZuoraVersion).(string))
	}

	res, err := t.http.Do(req.WithContext(ctx))
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

		return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory: %v", err)}
	}

	jsonResponse := ProductsResponse{}

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

	return &jsonResponse, nil
}

func (t *catalogService) GetProductNextPage(ctx context.Context, nextPageURI string) (*ProductsResponse, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v%v", t.baseURL, nextPageURI)

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

	if ctx.Value(ContextKeyZuoraVersion) != nil {
		req.Header.Add("zuora-version", ctx.Value(ContextKeyZuoraVersion).(string))
	}

	res, err := t.http.Do(req.WithContext(ctx))
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

		return nil, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory: %v", err)}
	}

	jsonResponse := ProductsResponse{}

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

	return &jsonResponse, nil
}
