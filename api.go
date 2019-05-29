package zuora

type V1 struct {
	ActionsService  *actionsService
	AccountsService *accountsService
	// BillingDocumentsService *BillingDocumentsService
	// ProductsService         *ProductsService
	// ActionsService          *ActionsService
	// DescribeService         *DescribeService
	// ProductRatePlansService *ProductRatePlansService
	// AccountsService         *AccountsService
	// SubscriptionsService    *SubscriptionsService
	// PaymentsService         *PaymentsService
	// ObjectModel             ObjectModel
}

//API is a container struct with access to all underlying services
type API struct {
	V1 V1
}

//NewAPI helper function to create all required services to interact with Zuora
func NewAPI(httpClient Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *API {
	// billingDocumentsService := newBillingDocumentsService(zuoraConfig, tokenService, errorRequestHandler)
	// productsService := newProductsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	// accountsService := newAccountsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	// subscriptionsService := newSubscriptionsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	// paymentsService := newPaymentsService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	// productRatePlansService := newProductRatePlansService(zuoraConfig, tokenService, actionsService, errorRequestHandler)
	// objectModel := newObjectModel()
	// describeService := newDescribeService(zuoraConfig, tokenService)

	return &API{
		V1: V1{
			ActionsService:  newActionsService(httpClient, authHeaderProvider, baseURL),
			AccountsService: newAccountsService(httpClient, authHeaderProvider, baseURL),
		},
		// BillingDocumentsService: billingDocumentsService,
		// ProductsService:         productsService,
		// ActionsService:          actionsService,
		// DescribeService:         describeService,
		// ProductRatePlansService: productRatePlansService,
		// AccountsService:         accountsService,
		// SubscriptionsService:    subscriptionsService,
		// PaymentsService:         paymentsService,
		// ObjectModel:             objectModel,
	}
}
