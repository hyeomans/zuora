package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

type BillingDocumentsService struct {
	config       *Config
	tokenService *TokenService
	errorHandler errors.RequestHandler
}

type jsonResponse struct {
	Documents []BillingDocuments `json:"documents"`
	Success   bool               `json:"success"`
}

type BillingDocuments struct {
	ID             string `json:"id"`
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	DocumentDate   string `json:"documentDate"`
	Amount         int    `json:"amount"`
	Balance        int    `json:"balance"`
	AccountID      string `json:"accountId"`
	Status         string `json:"status"`
}

func newBillingDocumentsService(config *Config, tokenService *TokenService, errorHandler errors.RequestHandler) *BillingDocumentsService {
	return &BillingDocumentsService{
		config:       config,
		tokenService: tokenService,
		errorHandler: errorHandler,
	}
}

func (s *BillingDocumentsService) Get(ctx context.Context, accountID string, queryParams ...string) ([]BillingDocuments, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return []BillingDocuments{}, err
	}

	documentsURL := fmt.Sprint(s.config.BaseURL, "/v1/billing-documents")
	req, err := http.NewRequest(http.MethodGet, documentsURL, nil)

	if err != nil {
		return []BillingDocuments{}, s.errorHandler.BadRequest(err)
	}

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token.AccessToken))

	res, err := s.config.HTTPClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return []BillingDocuments{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := jsonResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return []BillingDocuments{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return []BillingDocuments{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse.Documents, err
}
