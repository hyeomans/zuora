package zuora

import (
	"net/http"
)

//Token represents the OAuth token returned by Zuora.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

//ConfigOption helper function to modify current Config struct
type ConfigOption func(*Config)

//Config is the base configuration to return ZuoraApi
type Config struct {
	HTTPClient   *http.Client
	BaseURL      string
	ClientID     string
	ClientSecret string
	tokenStore   TokenStorer
}

//TokenStorer handles token renewal with two simple methods.
//Token() returns a boolean to indicate a token is valid and if valid, it will return the active token.
//Update() causes a side-effect to update a token in whichever backing store you choose.
type TokenStorer interface {
	Token() (bool, *Token)
	Update(*Token)
}

//Querier One who, or that which, queries actions
type Querier interface {
	Build() string
}

//Product with default fields from Zuora. It does not include custom fields.
type Product struct {
	ID                 string `json:"Id"`
	Category           string `json:"Category,omitempty"`
	CreatedDate        string `json:"CreatedDate"`
	Name               string `json:"Name"`
	SKU                string `json:"SKU"`
	UpdatedDate        string `json:"UpdatedDate"`
	EffectiveStartDate string `json:"EffectiveStartDate"`
	UpdatedByID        string `json:"UpdatedById"`
	EffectiveEndDate   string `json:"EffectiveEndDate"`
	CreatedByID        string `json:"CreatedById"`
	Description        string `json:"Description,omitempty"`
}

//ProductRatePlan with default fields from Zuora. It does not include custom fields.
type ProductRatePlan struct {
	ActiveCurrencies   string `json:"ActiveCurrencies,omitempty"`
	ProductID          string `json:"ProductId"`
	ID                 string `json:"Id"`
	CreatedByID        string `json:"CreatedById"`
	CreatedDate        string `json:"CreatedDate"`
	Name               string `json:"Name"`
	UpdatedDate        string `json:"UpdatedDate"`
	EffectiveStartDate string `json:"EffectiveStartDate"`
	UpdatedByID        string `json:"UpdatedById"`
	EffectiveEndDate   string `json:"EffectiveEndDate"`
	Description        string `json:"Description,omitempty"`
}

type Payment struct {
	PaymentNumber              string  `json:"PaymentNumber"`
	GatewayResponse            string  `json:"GatewayResponse"`
	ID                         string  `json:"Id"`
	UpdatedDate                string  `json:"UpdatedDate"`
	GatewayState               string  `json:"GatewayState"`
	Source                     string  `json:"Source"`
	AccountID                  string  `json:"AccountId"`
	BankIdentificationNumber   string  `json:"BankIdentificationNumber"`
	ReferenceID                string  `json:"ReferenceId"`
	PaymentMethodSnapshotID    string  `json:"PaymentMethodSnapshotId"`
	UpdatedByID                string  `json:"UpdatedById"`
	SubmittedOn                string  `json:"SubmittedOn"`
	Type                       string  `json:"Type"`
	CreatedDate                string  `json:"CreatedDate"`
	RefundAmount               int     `json:"RefundAmount"`
	SourceName                 string  `json:"SourceName"`
	Amount                     float64 `json:"Amount"`
	PaymentMethodID            string  `json:"PaymentMethodId"`
	CreatedByID                string  `json:"CreatedById"`
	Status                     string  `json:"Status"`
	AppliedCreditBalanceAmount int     `json:"AppliedCreditBalanceAmount"`
	Gateway                    string  `json:"Gateway"`
	GatewayResponseCode        string  `json:"GatewayResponseCode"`
	EffectiveDate              string  `json:"EffectiveDate"`
}

type AccountSummary struct {
	Payments []struct {
		PaidInvoices []struct {
			InvoiceNumber        string  `json:"invoiceNumber"`
			AppliedPaymentAmount float64 `json:"appliedPaymentAmount"`
			InvoiceID            string  `json:"invoiceId"`
		} `json:"paidInvoices"`
		PaymentNumber string `json:"paymentNumber"`
		Status        string `json:"status"`
		EffectiveDate string `json:"effectiveDate"`
		ID            string `json:"id"`
		PaymentType   string `json:"paymentType"`
	} `json:"payments"`
	Invoices []struct {
		Amount        float64 `json:"amount"`
		Status        string  `json:"status"`
		InvoiceNumber string  `json:"invoiceNumber"`
		InvoiceDate   string  `json:"invoiceDate"`
		Balance       float64 `json:"balance"`
		ID            string  `json:"id"`
		DueDate       string  `json:"dueDate"`
	} `json:"invoices"`
	Usage []struct {
		UnitOfMeasure string `json:"unitOfMeasure"`
		Quantity      int    `json:"quantity"`
		StartDate     string `json:"startDate"`
	} `json:"usage"`
	BasicInfo struct {
		DefaultPaymentMethod struct {
			CreditCardNumber          string `json:"creditCardNumber"`
			PaymentMethodType         string `json:"paymentMethodType"`
			CreditCardExpirationMonth int    `json:"creditCardExpirationMonth"`
			CreditCardExpirationYear  int    `json:"creditCardExpirationYear"`
			CreditCardType            string `json:"creditCardType"`
			ID                        string `json:"id"`
		} `json:"defaultPaymentMethod"`
		Status                    string   `json:"status"`
		LastInvoiceDate           string   `json:"lastInvoiceDate"`
		LastPaymentAmount         float64  `json:"lastPaymentAmount"`
		BillCycleDay              int      `json:"billCycleDay"`
		InvoiceDeliveryPrefsPrint bool     `json:"invoiceDeliveryPrefsPrint"`
		InvoiceDeliveryPrefsEmail bool     `json:"invoiceDeliveryPrefsEmail"`
		AdditionalEmailAddresses  []string `json:"additionalEmailAddresses"`
		Name                      string   `json:"name"`
		Balance                   float64  `json:"balance"`
		AccountNumber             string   `json:"accountNumber"`
		ID                        string   `json:"id"`
		Currency                  string   `json:"currency"`
		LastPaymentDate           string   `json:"lastPaymentDate"`
	} `json:"basicInfo"`
	SoldToContact struct {
		Fax       string `json:"fax"`
		TaxRegion string `json:"taxRegion"`
		Country   string `json:"country"`
		ZipCode   string `json:"zipCode"`
		County    string `json:"county"`
		LastName  string `json:"lastName"`
		WorkEmail string `json:"workEmail"`
		State     string `json:"state"`
		Address2  string `json:"address2"`
		Address1  string `json:"address1"`
		FirstName string `json:"firstName"`
		ID        string `json:"id"`
		WorkPhone string `json:"workPhone"`
		City      string `json:"city"`
	} `json:"soldToContact"`
	Success       bool `json:"success"`
	Subscriptions []struct {
		TermEndDate           string `json:"termEndDate"`
		TermStartDate         string `json:"termStartDate"`
		Status                string `json:"status"`
		InitialTerm           int    `json:"initialTerm"`
		AutoRenew             bool   `json:"autoRenew"`
		SubscriptionNumber    string `json:"subscriptionNumber"`
		SubscriptionStartDate string `json:"subscriptionStartDate"`
		ID                    string `json:"id"`
		RatePlans             []struct {
			ProductName  string `json:"productName"`
			RatePlanName string `json:"ratePlanName"`
		} `json:"ratePlans"`
		TermType    string `json:"termType"`
		RenewalTerm int    `json:"renewalTerm"`
	} `json:"subscriptions"`
	BillToContact struct {
		Fax       string `json:"fax"`
		TaxRegion string `json:"taxRegion"`
		Country   string `json:"country"`
		ZipCode   string `json:"zipCode"`
		County    string `json:"county"`
		LastName  string `json:"lastName"`
		WorkEmail string `json:"workEmail"`
		State     string `json:"state"`
		Address2  string `json:"address2"`
		Address1  string `json:"address1"`
		FirstName string `json:"firstName"`
		ID        string `json:"id"`
		WorkPhone string `json:"workPhone"`
		City      string `json:"city"`
	} `json:"billToContact"`
}

type SubscriptionResponse struct {
	Subscriptions []struct {
		ID                        string  `json:"id"`
		AccountID                 string  `json:"accountId"`
		AccountNumber             string  `json:"accountNumber"`
		AccountName               string  `json:"accountName"`
		InvoiceOwnerAccountID     string  `json:"invoiceOwnerAccountId"`
		InvoiceOwnerAccountNumber string  `json:"invoiceOwnerAccountNumber"`
		InvoiceOwnerAccountName   string  `json:"invoiceOwnerAccountName"`
		SubscriptionNumber        string  `json:"subscriptionNumber"`
		TermType                  string  `json:"termType"`
		InvoiceSeparately         bool    `json:"invoiceSeparately"`
		ContractEffectiveDate     string  `json:"contractEffectiveDate"`
		ServiceActivationDate     string  `json:"serviceActivationDate"`
		CustomerAcceptanceDate    string  `json:"customerAcceptanceDate"`
		SubscriptionStartDate     string  `json:"subscriptionStartDate"`
		TermStartDate             string  `json:"termStartDate"`
		TermEndDate               string  `json:"termEndDate"`
		InitialTerm               int     `json:"initialTerm"`
		InitialTermPeriodType     string  `json:"initialTermPeriodType"`
		CurrentTerm               int     `json:"currentTerm"`
		CurrentTermPeriodType     string  `json:"currentTermPeriodType"`
		AutoRenew                 bool    `json:"autoRenew"`
		RenewalSetting            string  `json:"renewalSetting"`
		RenewalTerm               int     `json:"renewalTerm"`
		RenewalTermPeriodType     string  `json:"renewalTermPeriodType"`
		ContractedMrr             float64 `json:"contractedMrr"`
		TotalContractedValue      int     `json:"totalContractedValue"`
		Notes                     string  `json:"notes"`
		Status                    string  `json:"status"`
		RatePlans                 []struct {
			ID                string `json:"id"`
			ProductID         string `json:"productId"`
			ProductName       string `json:"productName"`
			ProductSku        string `json:"productSku"`
			ProductRatePlanID string `json:"productRatePlanId"`
			RatePlanName      string `json:"ratePlanName"`
			RatePlanCharges   []struct {
				ID                             string        `json:"id"`
				OriginalChargeID               string        `json:"originalChargeId"`
				ProductRatePlanChargeID        string        `json:"productRatePlanChargeId"`
				Number                         string        `json:"number"`
				Name                           string        `json:"name"`
				Type                           string        `json:"type"`
				Model                          string        `json:"model"`
				Uom                            interface{}   `json:"uom"`
				Version                        int           `json:"version"`
				PricingSummary                 string        `json:"pricingSummary"`
				PriceChangeOption              string        `json:"priceChangeOption"`
				PriceIncreasePercentage        interface{}   `json:"priceIncreasePercentage"`
				Currency                       string        `json:"currency"`
				Price                          int           `json:"price"`
				Tiers                          interface{}   `json:"tiers"`
				IncludedUnits                  interface{}   `json:"includedUnits"`
				OveragePrice                   interface{}   `json:"overagePrice"`
				DiscountPercentage             interface{}   `json:"discountPercentage"`
				DiscountAmount                 interface{}   `json:"discountAmount"`
				ApplyDiscountTo                interface{}   `json:"applyDiscountTo"`
				DiscountLevel                  interface{}   `json:"discountLevel"`
				DiscountClass                  interface{}   `json:"discountClass"`
				DiscountApplyDetails           []interface{} `json:"discountApplyDetails"`
				BillingDay                     string        `json:"billingDay"`
				ListPriceBase                  string        `json:"listPriceBase"`
				BillingPeriod                  string        `json:"billingPeriod"`
				SpecificBillingPeriod          interface{}   `json:"specificBillingPeriod"`
				BillingTiming                  string        `json:"billingTiming"`
				BillingPeriodAlignment         string        `json:"billingPeriodAlignment"`
				Quantity                       int           `json:"quantity"`
				SmoothingModel                 interface{}   `json:"smoothingModel"`
				NumberOfPeriods                interface{}   `json:"numberOfPeriods"`
				OverageCalculationOption       interface{}   `json:"overageCalculationOption"`
				OverageUnusedUnitsCreditOption interface{}   `json:"overageUnusedUnitsCreditOption"`
				UnusedUnitsCreditRates         interface{}   `json:"unusedUnitsCreditRates"`
				UsageRecordRatingOption        interface{}   `json:"usageRecordRatingOption"`
				Segment                        int           `json:"segment"`
				EffectiveStartDate             string        `json:"effectiveStartDate"`
				EffectiveEndDate               string        `json:"effectiveEndDate"`
				ProcessedThroughDate           interface{}   `json:"processedThroughDate"`
				ChargedThroughDate             interface{}   `json:"chargedThroughDate"`
				Done                           bool          `json:"done"`
				TriggerDate                    interface{}   `json:"triggerDate"`
				TriggerEvent                   string        `json:"triggerEvent"`
				EndDateCondition               string        `json:"endDateCondition"`
				UpToPeriodsType                interface{}   `json:"upToPeriodsType"`
				UpToPeriods                    interface{}   `json:"upToPeriods"`
				SpecificEndDate                interface{}   `json:"specificEndDate"`
				Mrr                            float64       `json:"mrr"`
				Dmrc                           float64       `json:"dmrc"`
				Tcv                            int           `json:"tcv"`
				Dtcv                           int           `json:"dtcv"`
				Description                    string        `json:"description"`
			} `json:"ratePlanCharges"`
		} `json:"ratePlans"`
	} `json:"subscriptions"`
	Success bool `json:"success"`
}
