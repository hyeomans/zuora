package zuora

import (
	"net/http"

	"github.com/hyeomans/zuora/errors"
)

//API is a container struct with access to all underlying services
type API struct {
	TokenService            *TokenService //TODO: Should this be public?
	BillingDocumentsService *BillingDocumentsService
	ProductsService         *ProductsService
	ActionsService          *ActionsService
	DescribeService         *DescribeService
	ProductRatePlansService *ProductRatePlansService
	AccountsService         *AccountsService
	SubscriptionsService    *SubscriptionsService
	PaymentsService         *PaymentsService
	ObjectModel             ObjectModel
}

//NewAPI helper function to create all required services to interact with Zuora
func NewAPI(httpClient *http.Client, baseURL, clientID, clientSecret string, options ...ConfigOption) *API {
	tokenStore := &memoryTokenStore{}
	zuoraConfig := newConfig(httpClient, baseURL, clientID, clientSecret, tokenStore, options)
	errorRequestHandler := errors.RequestHandler{}
	tokenService := newTokenService(zuoraConfig, errorRequestHandler)
	actionsService := newActionsService(zuoraConfig, tokenService, errorRequestHandler)

	billingDocumentsService := newBillingDocumentsService(zuoraConfig, tokenService, errorRequestHandler)
	productsService := newProductsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	accountsService := newAccountsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	subscriptionsService := newSubscriptionsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	paymentsService := newPaymentsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	productRatePlansService := newProductRatePlansService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	objectModel := newObjectModel()
	describeService := newDescribeService(zuoraConfig, tokenService)

	return &API{
		TokenService:            tokenService,
		BillingDocumentsService: billingDocumentsService,
		ProductsService:         productsService,
		ActionsService:          actionsService,
		DescribeService:         describeService,
		ProductRatePlansService: productRatePlansService,
		AccountsService:         accountsService,
		SubscriptionsService:    subscriptionsService,
		PaymentsService:         paymentsService,
		ObjectModel:             objectModel,
	}
}
