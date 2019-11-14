package zuora

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type actionsService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
	isPce              bool
}

func newActionsService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string, isPce bool) *actionsService {
	return &actionsService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
		isPce:              isPce,
	}
}

// Query The query call sends a query expression by specifying the object to query,
// the fields to retrieve from that object, and any filters to determine whether a
// given object should be queried.
// You can use Zuora Object Query Language (ZOQL) to construct those queries,
// passing them through the queryString.
// https://knowledgecenter.zuora.com/DC_Developers/K_Zuora_Object_Query_Language
// Once the call is made, the API executes the query against the specified object and
// returns a query response object to your application. Your application can then iterate
// through rows in the query response to retrieve information.
//
// Limitations
// This call has the following limitations:
//
// * All ZOQL keywords must be in lower case.
//
// * The number of records returned is limited to 2000 records
//
// * The Invoice Settlement feature is not supported. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. The Orders feature is also not supported.
//
// *The default WSDL version for Actions is 79.
func (t *actionsService) Query(ctx context.Context, zoqlQuery string) ([]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/action/query", t.baseURL)
	} else {
		url = fmt.Sprintf("%v/v1/action/query", t.baseURL)
	}

	var buffer bytes.Buffer
	buffer.WriteString(`{ "queryString" : "`)
	buffer.WriteString(strings.TrimSpace(zoqlQuery))
	buffer.WriteString(`"}`)
	fmt.Println(buffer.String())
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(buffer.String()))

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

	return body, nil
}

func (t *actionsService) Create(ctx context.Context, actionPayload interface{}) ([]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/action/create", t.baseURL)

	j, err := json.Marshal(actionPayload)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to convert empty interface: %v", err)}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))

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

	return body, nil
}
