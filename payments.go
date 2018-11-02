package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

//PaymentsService Access all methods related to Payments
type PaymentsService struct {
	config         *Config
	tokenService   *TokenService
	actionsService *ActionsService
	errorHandler   errors.RequestHandler
}

func newPaymentsService(config *Config, tokenService *TokenService, actionsService *ActionsService, errorHandler errors.RequestHandler) *PaymentsService {
	return &PaymentsService{
		config:         config,
		tokenService:   tokenService,
		errorHandler:   errorHandler,
		actionsService: actionsService,
	}
}

func (s *PaymentsService) ByIdThroughObject(ctx context.Context, id string) (Payment, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return Payment{}, err
	}

	url := fmt.Sprintf("%v/v1/object/payment/%v", s.config.BaseURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return Payment{}, s.errorHandler.BadRequest(err)
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
		return Payment{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return Payment{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := Payment{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return Payment{}, s.errorHandler.BadRequest(err)
	}

	return jsonResponse, err
}
