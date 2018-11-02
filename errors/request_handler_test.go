package errors_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hyeomans/zuora/errors"
)

func TestInvalidResponseHandler(t *testing.T) {
	tests := []struct {
		name             string
		inputStatusCode  int
		expectedResponse error
		input            []byte
	}{
		{"PlainError", http.StatusBadRequest, fmt.Errorf("zuora returned error - statusCode: 400 message: plain error"), []byte(`plain error`)},
		{"AuthenticationError", http.StatusUnauthorized, fmt.Errorf("zuora returned error - statusCode: 401 message: Authentication error"), []byte(`{"message":"Authentication error"}`)},
		{"BadRequest", http.StatusBadRequest, fmt.Errorf(`zuora returned error - statusCode: 400 message: {"message":"Error - invalid action: queryy"}`), []byte(`{"message":"Error - invalid action: queryy"}`)},
		{"NotFound", http.StatusNotFound, fmt.Errorf(`zuora returned error - statusCode: 404 message: object not found`), []byte(`{"records":[],"size":0,"done":true}`)},
		{"InternalServerError", http.StatusInternalServerError, fmt.Errorf(`zuora returned error - statusCode: 500 message: {"faultcode":"fns:INVALID_TYPE","faultstring":"invalid type specified: productt","detail":{"MalformedQueryFault":{"FaultCode":"INVALID_TYPE","FaultMessage":"invalid type specified: productt"}}}`), []byte(`{"faultcode":"fns:INVALID_TYPE","faultstring":"invalid type specified: productt","detail":{"MalformedQueryFault":{"FaultCode":"INVALID_TYPE","FaultMessage":"invalid type specified: productt"}}}`)},
	}

	type temporary interface {
		Temporary() bool
	}

	for idx, tc := range tests {
		tt := tc
		i := idx
		tf := func(t *testing.T) {
			t.Parallel()
			t.Logf("\tTest: %d\t Name: %v", i, tt.name)
			requestHandler := errors.RequestHandler{}
			result := requestHandler.InvalidResponse(tt.input, tt.inputStatusCode)
			if result.Error() != tt.expectedResponse.Error() {
				t.Fatalf("Expected: %v but got: %v", tt.expectedResponse, result)
			}

			temporary, ok := result.(temporary)
			if !ok {
				t.Fatalf("It should support the temporary interface: %v", ok)
			} else {
				t.Logf("Implements temporary interface: %v %v", temporary.Temporary(), ok)
			}
		}

		t.Run(tt.name, tf)
	}
}

func TestValidRequestError(t *testing.T) {
	tests := []struct {
		name             string
		inputStatusCode  int
		isTemporary      bool
		expectedResponse error
		input            []byte
	}{
		{"PlainError", http.StatusBadRequest, false, fmt.Errorf("zuora returned an error (status: 400): invalid character 'p' looking for beginning of value"), []byte(`plain error`)},
		{"BadRequest", http.StatusBadRequest, false, fmt.Errorf(`zuora returned an error (status: 400): code: 53100020 message:  {com.zuora.constraints.either_or_both} || code: 53100320 message: 'termtype' value should be one of: termed, evergreen`), []byte(`{"success":false,"processId":"3F7EA3FD706C7E7C","reasons":[{"code":53100020,"message":" {com.zuora.constraints.either_or_both}"},{"code":53100320,"message":"'termType' value should be one of: TERMED, EVERGREEN"}]}`)},
		{"429 takes priority", http.StatusBadRequest, true, fmt.Errorf(`zuora returned an error (status: 429): code: 53100020 message:  {com.zuora.constraints.either_or_both} || code: 53100370 message: 'termtype' value should be one of: termed, evergreen`), []byte(`{"success":false,"processId":"3F7EA3FD706C7E7C","reasons":[{"code":53100020,"message":" {com.zuora.constraints.either_or_both}"},{"code":53100370,"message":"'termType' value should be one of: TERMED, EVERGREEN"}]}`)},
	}

	type temporary interface {
		Temporary() bool
	}

	for idx, tc := range tests {
		tt := tc
		i := idx
		tf := func(t *testing.T) {
			t.Parallel()
			t.Logf("\tTest: %d\t Name: %v", i, tt.name)
			requestHandler := errors.RequestHandler{}
			result := requestHandler.ValidRequestError(tt.input, tt.inputStatusCode)

			if result.Error() != tt.expectedResponse.Error() {
				t.Fatalf("Expected: %v but got: %v", tt.expectedResponse, result)
			}

			temporary, ok := result.(temporary)

			if !ok {
				t.Fatalf("It should support the temporary interface: %v", ok)
			} else {
				t.Logf("Implements temporary interface")
			}
			if temporary.Temporary() != tt.isTemporary {
				t.Fatalf("Expected error to be Temporary: %v", tt.input)
			}
		}

		t.Run(tt.name, tf)
	}
}
