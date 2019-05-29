package zuora

type SummaryResponse struct {
	BasicInfo     SummaryBasicInfo      `json:"basicInfo"`
	BillToContact SummaryContact        `json:"billToContact"`
	Invoices      []SummaryInvoice      `json:"invoices"`
	Payments      []SummaryPayment      `json:"payments"`
	SoldToContact SummaryContact        `json:"soldToContact"`
	Subscriptions []SummarySubscription `json:"subscriptions"`
	Success       bool                  `json:"success"`
	TaxInfo       SummaryTaxInfo        `json:"taxInfo"`
	Usage         []SummaryUsage        `json:"usage"`
}

type SummaryContact struct {
	Address1  string      `json:"address1"`
	Address2  interface{} `json:"address2"`
	City      string      `json:"city"`
	Country   string      `json:"country"`
	County    interface{} `json:"county"`
	Fax       interface{} `json:"fax"`
	FirstName string      `json:"firstName"`
	ID        string      `json:"id"`
	LastName  string      `json:"lastName"`
	State     string      `json:"state"`
	TaxRegion interface{} `json:"taxRegion"`
	WorkEmail string      `json:"workEmail"`
	WorkPhone interface{} `json:"workPhone"`
	ZipCode   string      `json:"zipCode"`
}

type SummaryInvoice struct {
	Amount        float64 `json:"amount"`
	Balance       int64   `json:"balance"`
	DueDate       string  `json:"dueDate"`
	ID            string  `json:"id"`
	InvoiceDate   string  `json:"invoiceDate"`
	InvoiceNumber string  `json:"invoiceNumber"`
	Status        string  `json:"status"`
}

type SummarySubscription struct {
	AppVersionC                 string            `json:"AppVersion__c"`
	BusinessLineC               string            `json:"BusinessLine__c"`
	CpqBundleJSONIDQT           interface{}       `json:"CpqBundleJsonId__QT"`
	EmployeesFromAccountObjectC interface{}       `json:"Employees_From_Account_Object__c"`
	OpportunityCloseDateQT      interface{}       `json:"OpportunityCloseDate__QT"`
	OpportunityNameQT           interface{}       `json:"OpportunityName__QT"`
	PONumberC                   interface{}       `json:"PO_Number__c"`
	QuoteBusinessTypeQT         interface{}       `json:"QuoteBusinessType__QT"`
	QuoteNumberQT               interface{}       `json:"QuoteNumber__QT"`
	QuoteTypeQT                 interface{}       `json:"QuoteType__QT"`
	QuoteOpportunityIDC         interface{}       `json:"Quote_Opportunity_ID__c"`
	ResellerCompanyNameC        interface{}       `json:"ResellerCompanyName__c"`
	SegmentC                    interface{}       `json:"Segment__c"`
	AutoRenew                   bool              `json:"autoRenew"`
	ID                          string            `json:"id"`
	InitialTerm                 interface{}       `json:"initialTerm"`
	RatePlans                   []SummaryRatePlan `json:"ratePlans"`
	RenewalTerm                 interface{}       `json:"renewalTerm"`
	Status                      string            `json:"status"`
	SubscriptionNumber          string            `json:"subscriptionNumber"`
	SubscriptionStartDate       string            `json:"subscriptionStartDate"`
	TermEndDate                 interface{}       `json:"termEndDate"`
	TermStartDate               string            `json:"termStartDate"`
	TermType                    string            `json:"termType"`
}

type SummaryPaidInvoice struct {
	AppliedPaymentAmount float64 `json:"appliedPaymentAmount"`
	InvoiceID            string  `json:"invoiceId"`
	InvoiceNumber        string  `json:"invoiceNumber"`
}

type SummaryDefaultPaymentMethod struct {
	CreditCardExpirationMonth int64  `json:"creditCardExpirationMonth"`
	CreditCardExpirationYear  int64  `json:"creditCardExpirationYear"`
	CreditCardNumber          string `json:"creditCardNumber"`
	CreditCardType            string `json:"creditCardType"`
	ID                        string `json:"id"`
	PaymentMethodType         string `json:"paymentMethodType"`
}

type SummaryBasicInfo struct {
	DomainC                             string                      `json:"Domain__c"`
	InvisionTeamIDC                     string                      `json:"InvisionTeamId__c"`
	InvisionCompanyIDC                  interface{}                 `json:"Invision_Company_ID__c"`
	LastQuoteResellerCompanyNameC       interface{}                 `json:"Last_Quote_Reseller_Company_Name__c"`
	SubdomainC                          interface{}                 `json:"Subdomain__c"`
	TECHolderC                          interface{}                 `json:"TECHolder__c"`
	TaxExemptCertificateExpirationDateC interface{}                 `json:"TaxExemptCertificateExpirationDate__c"`
	TaxExemptCertificateStateC          interface{}                 `json:"TaxExemptCertificateState__c"`
	AccountNumber                       string                      `json:"accountNumber"`
	AdditionalEmailAddresses            []interface{}               `json:"additionalEmailAddresses"`
	Balance                             int64                       `json:"balance"`
	Batch                               string                      `json:"batch"`
	BillCycleDay                        int64                       `json:"billCycleDay"`
	Currency                            string                      `json:"currency"`
	DefaultPaymentMethod                SummaryDefaultPaymentMethod `json:"defaultPaymentMethod"`
	ID                                  string                      `json:"id"`
	InvoiceDeliveryPrefsEmail           bool                        `json:"invoiceDeliveryPrefsEmail"`
	InvoiceDeliveryPrefsPrint           bool                        `json:"invoiceDeliveryPrefsPrint"`
	LastInvoiceDate                     string                      `json:"lastInvoiceDate"`
	LastPaymentAmount                   float64                     `json:"lastPaymentAmount"`
	LastPaymentDate                     string                      `json:"lastPaymentDate"`
	Name                                string                      `json:"name"`
	Status                              string                      `json:"status"`
}

type SummaryPayment struct {
	EffectiveDate string               `json:"effectiveDate"`
	ID            string               `json:"id"`
	PaidInvoices  []SummaryPaidInvoice `json:"paidInvoices"`
	PaymentNumber string               `json:"paymentNumber"`
	PaymentType   string               `json:"paymentType"`
	Status        string               `json:"status"`
}

type SummaryRatePlan struct {
	ProductID         string `json:"productId"`
	ProductName       string `json:"productName"`
	ProductRatePlanID string `json:"productRatePlanId"`
	ProductSku        string `json:"productSku"`
	RatePlanName      string `json:"ratePlanName"`
}

type SummaryUsage struct {
	Quantity      string `json:"quantity"`
	StartDate     string `json:"startDate"`
	UnitOfMeasure string `json:"unitOfMeasure"`
}

type SummaryTaxInfo struct {
	VATId                     string      `json:"VATId"`
	CompanyCode               interface{} `json:"companyCode"`
	ExemptCertificateID       interface{} `json:"exemptCertificateId"`
	ExemptCertificateType     interface{} `json:"exemptCertificateType"`
	ExemptDescription         interface{} `json:"exemptDescription"`
	ExemptEffectiveDate       interface{} `json:"exemptEffectiveDate"`
	ExemptEntityUseCode       interface{} `json:"exemptEntityUseCode"`
	ExemptExpirationDate      interface{} `json:"exemptExpirationDate"`
	ExemptIssuingJurisdiction interface{} `json:"exemptIssuingJurisdiction"`
	ExemptStatus              interface{} `json:"exemptStatus"`
}
