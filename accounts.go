package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

type AccountsService struct {
	config         *Config
	tokenService   *TokenService
	actionsService *ActionsService
	errorHandler   errors.RequestHandler
}

func newAccountsService(config *Config, tokenService *TokenService, actionsService *ActionsService, errorHandler errors.RequestHandler) *AccountsService {
	return &AccountsService{
		config:         config,
		tokenService:   tokenService,
		errorHandler:   errorHandler,
		actionsService: actionsService,
	}
}

//Summary Retrieves detailed information about the specified customer account.
//The response includes the account information and a summary of the account’s subscriptions, invoices, payments,
//and usages for the last six recently updated subscriptions.
//NOTE: Returns only the six most recent subscriptions based on the subscription updatedDate. Within those
//subscriptions, there may be many rate plans and many rate plan charges. These items are subject to the maximum
//limit on the array size.
func (s *AccountsService) Summary(ctx context.Context, objectID string) (AccountSummary, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return AccountSummary{}, err
	}

	url := fmt.Sprintf("%v/v1/accounts/%v/summary", s.config.BaseURL, objectID)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return AccountSummary{}, s.errorHandler.BadRequest(err)
	}

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	if ctx.Value("Zuora-Entity-Ids") != nil {
		req.Header.Add("Zuora-Entity-Ids", ctx.Value("Zuora-Entity-Ids").(string))
	}

	if ctx.Value("Zuora-Track-Id") != nil {
		req.Header.Add("Zuora-Track-Id", ctx.Value("Zuora-Track-Id").(string))
	}

	res, err := s.config.HTTPClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return AccountSummary{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return AccountSummary{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := AccountSummary{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return AccountSummary{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return AccountSummary{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}
