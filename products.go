package zuora

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

type ProductsService struct {
	config         *Config
	tokenService   *TokenService
	actionsService *ActionsService
	errorHandler   errors.RequestHandler
}

func newProductsService(config *Config, tokenService *TokenService, actionsService *ActionsService, errorHandler errors.RequestHandler) *ProductsService {
	return &ProductsService{
		config:         config,
		tokenService:   tokenService,
		errorHandler:   errorHandler,
		actionsService: actionsService,
	}
}

func (s *ProductsService) Get(ctx context.Context, objectID string) (Product, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return Product{}, err
	}

	documentsURL := fmt.Sprint(s.config.BaseURL, "/v1/object/product/", objectID)
	req, err := http.NewRequest(http.MethodGet, documentsURL, nil)

	if err != nil {
		return Product{}, s.errorHandler.BadRequest(err)
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
		return Product{}, s.errorHandler.BadRequest(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return Product{}, s.errorHandler.InvalidResponse(body, res.StatusCode)
	}

	jsonResponse := Product{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return Product{}, s.errorHandler.BadRequest(err)
	}

	return jsonResponse, err
}
