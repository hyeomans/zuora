package zuora

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

//SubscriptionsService Access all methods related to Subscriptions
type SubscriptionsService struct {
	config         *Config
	tokenService   *TokenService
	actionsService *ActionsService
	errorHandler   errors.RequestHandler
}

func newSubscriptionsService(config *Config, tokenService *TokenService, actionsService *ActionsService, errorHandler errors.RequestHandler) *SubscriptionsService {
	return &SubscriptionsService{
		config:         config,
		tokenService:   tokenService,
		errorHandler:   errorHandler,
		actionsService: actionsService,
	}
}

//ByKey This REST API reference describes how to retrieve detailed information about a
//specified subscription in the latest version.
func (s *SubscriptionsService) ByKey(ctx context.Context, accountKey string) (SubscriptionResponse, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return SubscriptionResponse{}, err
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", s.config.BaseURL, accountKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
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
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return SubscriptionResponse{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := SubscriptionResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return SubscriptionResponse{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}

//ByAccount Retrieves all subscriptions associated with the specified account. Zuora only
//returns the latest version of the subscriptions.
//Subscription data is returned in reverse chronological order based on updatedDate.
//accountKey possible values are: accountNumber or accountID
func (s *SubscriptionsService) ByAccount(ctx context.Context, accountKey string) (SubscriptionResponse, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return SubscriptionResponse{}, err
	}

	url := fmt.Sprintf("%v/v1/subscriptions/accounts/%v", s.config.BaseURL, accountKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
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
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return SubscriptionResponse{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := SubscriptionResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionResponse{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return SubscriptionResponse{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}

//Update contains the minimal update payload to apply to an accountKey
func (s *SubscriptionsService) Update(ctx context.Context, subscriptionKey string, subscriptionUpdatePayload SubscriptionUpdateMinimalPayload) (SubscriptionUpdateResponse, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return SubscriptionUpdateResponse{}, err
	}
	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(subscriptionUpdatePayload)

	if err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", s.config.BaseURL, subscriptionKey)
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
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
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return SubscriptionUpdateResponse{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := SubscriptionUpdateResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return SubscriptionUpdateResponse{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}

//UpdateFull Use this call to make the following kinds of changes to a subscription:
// Add a note
// Change the renewal term or auto-renewal flag
// Change the term length or change between evergreen and termed
// Add a new product rate plan
// Remove an existing subscription rate plan
// Change the quantity or price of an existing subscription rate plan
//Notes:
// This feature is unavailable if you have the Orders feature enabled. See Orders Migration Guidance for more information.
// The Update Subscription call creates a new subscription, which has the old subscription number but a new subscription ID. The old subscription is canceled but remains in the system.
// In one request, this call can make:
// Up to 9 combined add, update, and remove changes
// No more than 1 change to terms & conditions
// Updates are performed in the following sequence:
// First change the notes on the existing subscription, if requested.
// Then change the terms and conditions, if requested.
// Then perform the remaining amendments based upon the effective dates specified. If multiple amendments have the same contract-effective dates, then execute adds before updates, and updates before removes.
// The update operation is atomic. If any of the updates fails, the entire operation is rolled back.
// The response of the Update Subscription call is based on the REST API minor version you set in the request header. The response structure might be different if you use different minor version numbers.
// If you have the Invoice Settlement feature enabled, we recommend that you set the zuora-version parameter to 207.0 or later. Otherwise, an error is returned.
func (s *SubscriptionsService) UpdateFull(ctx context.Context, subscriptionKey string, subscriptionUpdatePayload SubscriptionUpdateFullPayload) (SubscriptionUpdateResponse, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return SubscriptionUpdateResponse{}, err
	}
	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(subscriptionUpdatePayload)

	if err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", s.config.BaseURL, subscriptionKey)
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
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
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return SubscriptionUpdateResponse{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := SubscriptionUpdateResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionUpdateResponse{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return SubscriptionUpdateResponse{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}

//Cancel use this to cancel an active subscription.
func (s *SubscriptionsService) Cancel(ctx context.Context, subscriptionKey string, subscriptionCancellationPayload SubscriptionCancellationPayload) (SubscriptionCancellationResponse, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return SubscriptionCancellationResponse{}, err
	}

	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(subscriptionCancellationPayload)

	if err != nil {
		return SubscriptionCancellationResponse{}, s.errorHandler.BadRequest(err)
	}

	url := fmt.Sprintf("%v/v1/subscriptions/%v", s.config.BaseURL, subscriptionKey)
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return SubscriptionCancellationResponse{}, s.errorHandler.BadRequest(err)
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
		return SubscriptionCancellationResponse{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return SubscriptionCancellationResponse{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := SubscriptionCancellationResponse{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return SubscriptionCancellationResponse{}, s.errorHandler.BadRequest(err)
	}

	if !jsonResponse.Success {
		return SubscriptionCancellationResponse{}, s.errorHandler.ValidRequestError(body, res.StatusCode)
	}

	return jsonResponse, err
}
