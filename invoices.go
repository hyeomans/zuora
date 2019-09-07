package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type invoices struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
	isPce              bool
}

func newInvoices(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string, isPce bool) *invoices {
	return &invoices{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
		isPce:              isPce,
	}
}

//InvoiceFile --
type InvoiceFile struct {
	ID            string `json:"id"`
	VersionNumber int64  `json:"versionNumber"`
	PdfFileURL    string `json:"pdfFileUrl"`
}

//InvoiceFilesResponse --
type InvoiceFilesResponse struct {
	InvoiceFiles []InvoiceFile `json:"invoiceFiles"`
	Response
}

// GetInvoice More info at: https://www.zuora.com/developer/API-Reference/#operation/Object_GETInvoice
func (t *invoices) GetInvoice(ctx context.Context, invoiceID string) (Invoice, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return Invoice{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/object/invoice/%v", t.baseURL, invoiceID)
	} else {
		url = fmt.Sprintf("%v/v1/object/invoice/%v", t.baseURL, invoiceID)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return Invoice{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return Invoice{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
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
			return Invoice{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return Invoice{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := Invoice{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return Invoice{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	return jsonResponse, nil
}

// GetInvoiceFiles Retrieves the information about all PDF files of a specified invoice.
// Invoice PDF files are returned in reverse chronological order by the value of the versionNumber field.
// More info: https://www.zuora.com/developer/API-Reference/#operation/GET_InvoiceFiles
func (t *invoices) GetInvoiceFiles(ctx context.Context, invoiceID string, pageSize int) (InvoiceFilesResponse, error) {
	if pageSize == 0 {
		pageSize = 20 //Default value accepted by Zuora
	}

	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return InvoiceFilesResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	var url string

	if t.isPce {
		url = fmt.Sprintf("%v:19016/v1/invoices/%v/files?pageSize=%v", t.baseURL, invoiceID, pageSize)
	} else {
		url = fmt.Sprintf("%v/v1/invoices/%v/files?pageSize=%v", t.baseURL, invoiceID, pageSize)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return InvoiceFilesResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
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
		return InvoiceFilesResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
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
			return InvoiceFilesResponse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return InvoiceFilesResponse{}, responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	jsonResponse := InvoiceFilesResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return InvoiceFilesResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	if !jsonResponse.Success {
		errorResponse := errorResponse{}

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return InvoiceFilesResponse{}, responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json error response. Error: %v. Raw JSON: %v", err, string(body))}
		}

		return InvoiceFilesResponse{}, errorResponse
	}

	return jsonResponse, nil
}
