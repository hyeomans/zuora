package zuora

import "github.com/hyeomans/zuora/errors"

type ProductRatePlansService struct {
	config         *Config
	tokenService   *TokenService
	actionsService *ActionsService
	errorHandler   errors.RequestHandler
}

func newProductRatePlansService(config *Config, tokenService *TokenService, actionsService *ActionsService, errorHandler errors.RequestHandler) *ProductRatePlansService {
	return &ProductRatePlansService{
		config:         config,
		tokenService:   tokenService,
		errorHandler:   errorHandler,
		actionsService: actionsService,
	}
}
