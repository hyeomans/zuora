package zuora

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DescribeService struct {
	config       *Config
	tokenService *TokenService
}

func newDescribeService(config *Config, tokenService *TokenService) *DescribeService {
	return &DescribeService{
		config:       config,
		tokenService: tokenService,
	}
}

func (s *DescribeService) Model(ctx context.Context, object string) (string, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return "", err
	}

	documentsURL := fmt.Sprint(s.config.BaseURL, "/v1/describe/", object)
	req, err := http.NewRequest(http.MethodGet, documentsURL, nil)

	if err != nil {
		return "", err
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
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
