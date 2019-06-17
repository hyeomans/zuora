package zuora

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type actionsService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

func newActionsService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *actionsService {
	return &actionsService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
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
func (t *actionsService) Query(ctx context.Context, querier Querier) (*[]byte, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return nil, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/action/query", t.baseURL)
	query := querier.Build()

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(query))

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

	return &body, nil
}
