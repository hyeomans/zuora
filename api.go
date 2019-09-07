package zuora

// V1 All the available REST endpoints in Zuora
type V1 struct {
	ActionsService       *actionsService
	AccountsService      *accountsService
	CatalogService       *catalogService
	SubscriptionsService *subscriptionsService
	DescribeService      *describeService
	PaymentMethods       *paymentMethods
	Invoices             *invoices
}

//API is a container struct with access to all underlying services
type API struct {
	V1          V1
	ObjectModel ObjectModel
}

//NewAPI helper function to create all required services to interact with Zuora
func NewAPI(httpClient Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *API {
	return &API{
		V1: V1{
			AccountsService:      newAccountsService(httpClient, authHeaderProvider, baseURL),
			CatalogService:       newCatalogService(httpClient, authHeaderProvider, baseURL),
			SubscriptionsService: newSubscriptionsService(httpClient, authHeaderProvider, baseURL),
			DescribeService:      newDescribeService(httpClient, authHeaderProvider, baseURL),
			ActionsService:       newActionsService(httpClient, authHeaderProvider, baseURL, false),
			PaymentMethods:       newPaymentMethods(httpClient, authHeaderProvider, baseURL, false),
			Invoices:             newInvoices(httpClient, authHeaderProvider, baseURL, false),
		},
		ObjectModel: newObjectModel(),
	}
}

//NewPCEAPI helper function to create all required services to interact with Zuora Production Copy Environment (PCE)
func NewPCEAPI(httpClient Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *API {
	return &API{
		V1: V1{
			AccountsService:      newAccountsService(httpClient, authHeaderProvider, baseURL),
			CatalogService:       newCatalogService(httpClient, authHeaderProvider, baseURL),
			SubscriptionsService: newSubscriptionsService(httpClient, authHeaderProvider, baseURL),
			DescribeService:      newDescribeService(httpClient, authHeaderProvider, baseURL),
			ActionsService:       newActionsService(httpClient, authHeaderProvider, baseURL, true),
			PaymentMethods:       newPaymentMethods(httpClient, authHeaderProvider, baseURL, true),
			Invoices:             newInvoices(httpClient, authHeaderProvider, baseURL, true),
		},
		ObjectModel: newObjectModel(),
	}
}
