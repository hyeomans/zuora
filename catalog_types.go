package zuora

type ProductsResponse struct {
	NextPage string    `json:"nextPage"`
	Products []Product `json:"products"`
	Success  bool      `json:"success"`
}

type Product struct {
	AppVersionC        string            `json:"AppVersion__c"`
	BusinessLineC      string            `json:"BusinessLine__c"`
	DisplayNameC       interface{}       `json:"DisplayName__c"`
	IsPackageC         string            `json:"IsPackage__c"`
	OneTimeDiscountC   interface{}       `json:"OneTimeDiscount__c"`
	PlanMaxQuantityC   interface{}       `json:"PlanMaxQuantity__c"`
	TrialAvailableC    interface{}       `json:"TrialAvailable__c"`
	Category           string            `json:"category"`
	Description        string            `json:"description"`
	EffectiveEndDate   string            `json:"effectiveEndDate"`
	EffectiveStartDate string            `json:"effectiveStartDate"`
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	ProductRatePlans   []ProductRatePlan `json:"productRatePlans"`
	Sku                string            `json:"sku"`
}

type ProductRatePlanCharge struct {
	CompanySizeC                      interface{}        `json:"CompanySize__c"`
	CustomerSuccessManagerAccessC     interface{}        `json:"CustomerSuccessManagerAccess__c"`
	DesignElevationWorkshopsC         interface{}        `json:"DesignElevationWorkshops__c"`
	IntegrationSupportC               interface{}        `json:"IntegrationSupport__c"`
	NumberofSeatsC                    string             `json:"NumberofSeats__c"`
	PrivatecloudC                     interface{}        `json:"Privatecloud__c"`
	SeatsC                            interface{}        `json:"Seats__c"`
	TechnicalSupportC                 interface{}        `json:"TechnicalSupport__c"`
	ApplyDiscountTo                   interface{}        `json:"applyDiscountTo"`
	BillingDay                        string             `json:"billingDay"`
	BillingPeriod                     string             `json:"billingPeriod"`
	BillingPeriodAlignment            string             `json:"billingPeriodAlignment"`
	BillingTiming                     string             `json:"billingTiming"`
	DefaultQuantity                   int64              `json:"defaultQuantity"`
	Description                       string             `json:"description"`
	DiscountClass                     interface{}        `json:"discountClass"`
	DiscountLevel                     interface{}        `json:"discountLevel"`
	EndDateCondition                  string             `json:"endDateCondition"`
	FinanceInformation                FinanceInformation `json:"financeInformation"`
	ID                                string             `json:"id"`
	ListPriceBase                     string             `json:"listPriceBase"`
	Model                             string             `json:"model"`
	Name                              string             `json:"name"`
	NumberOfPeriods                   interface{}        `json:"numberOfPeriods"`
	OverageCalculationOption          interface{}        `json:"overageCalculationOption"`
	OverageUnusedUnitsCreditOption    interface{}        `json:"overageUnusedUnitsCreditOption"`
	PriceChangeOption                 interface{}        `json:"priceChangeOption"`
	PriceIncreasePercentage           interface{}        `json:"priceIncreasePercentage"`
	Pricing                           []Pricing          `json:"pricing"`
	PricingSummary                    []string           `json:"pricingSummary"`
	ProductDiscountApplyDetails       []interface{}      `json:"productDiscountApplyDetails"`
	RevRecCode                        interface{}        `json:"revRecCode"`
	RevRecTriggerCondition            interface{}        `json:"revRecTriggerCondition"`
	RevenueRecognitionRuleName        string             `json:"revenueRecognitionRuleName"`
	SmoothingModel                    interface{}        `json:"smoothingModel"`
	SpecificBillingPeriod             int64              `json:"specificBillingPeriod"`
	TaxCode                           string             `json:"taxCode"`
	TaxMode                           string             `json:"taxMode"`
	Taxable                           bool               `json:"taxable"`
	TriggerEvent                      string             `json:"triggerEvent"`
	Type                              string             `json:"type"`
	UnusedIncludedUnitPrice           interface{}        `json:"unusedIncludedUnitPrice"`
	Uom                               string             `json:"uom"`
	UpToPeriods                       interface{}        `json:"upToPeriods"`
	UpToPeriodsType                   interface{}        `json:"upToPeriodsType"`
	UsageRecordRatingOption           interface{}        `json:"usageRecordRatingOption"`
	UseDiscountSpecificAccountingCode interface{}        `json:"useDiscountSpecificAccountingCode"`
	UseTenantDefaultForPriceChange    bool               `json:"useTenantDefaultForPriceChange"`
}

type Pricing struct {
	Currency           string      `json:"currency"`
	DiscountAmount     interface{} `json:"discountAmount"`
	DiscountPercentage interface{} `json:"discountPercentage"`
	IncludedUnits      int64       `json:"includedUnits"`
	OveragePrice       interface{} `json:"overagePrice"`
	Price              int64       `json:"price"`
	Tiers              interface{} `json:"tiers"`
}

type FinanceInformation struct {
	DeferredRevenueAccountingCode       string `json:"deferredRevenueAccountingCode"`
	DeferredRevenueAccountingCodeType   string `json:"deferredRevenueAccountingCodeType"`
	RecognizedRevenueAccountingCode     string `json:"recognizedRevenueAccountingCode"`
	RecognizedRevenueAccountingCodeType string `json:"recognizedRevenueAccountingCodeType"`
}

type ProductRatePlan struct {
	FrequencyTypeC         string                  `json:"FrequencyType__c"`
	IsPackageC             string                  `json:"IsPackage__c"`
	NumberofSeatsC         string                  `json:"NumberofSeats__c"`
	SFDCGuidedSellingC     string                  `json:"SFDCGuidedSelling__c"`
	Description            string                  `json:"description"`
	EffectiveEndDate       string                  `json:"effectiveEndDate"`
	EffectiveStartDate     string                  `json:"effectiveStartDate"`
	ID                     string                  `json:"id"`
	Name                   string                  `json:"name"`
	ProductRatePlanCharges []ProductRatePlanCharge `json:"productRatePlanCharges"`
	Status                 string                  `json:"status"`
}
