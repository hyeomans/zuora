package zuora

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

type ActionsService struct {
	config       *Config
	tokenService *TokenService
	errorHandler errors.RequestHandler
}

func newActionsService(config *Config, tokenService *TokenService, errorHandler errors.RequestHandler) *ActionsService {
	return &ActionsService{
		config:       config,
		tokenService: tokenService,
		errorHandler: errorHandler,
	}
}

func (s *ActionsService) Query(ctx context.Context, querier Querier) ([]byte, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprint(s.config.BaseURL, "/v1/action/query")
	query := querier.Build()

	if err != nil {
		return nil, s.errorHandler.BadRequest(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(query))

	if err != nil {
		return nil, s.errorHandler.BadRequest(err)
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
		return nil, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	return body, err
}
