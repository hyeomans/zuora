package zuora

type Subscription struct {
	AccountID              string  `json:"AccountId"`
	AncestorAccountID      *string `json:"AncestorAccountId,omitempty"`
	AutoRenew              *bool   `json:"AutoRenew,omitempty"`
	CancelledDate          *string `json:"CancelledDate,omitempty"`
	ContractAcceptanceDate *string `json:"ContractAcceptanceDate,omitempty"`
	ContractEffectiveDate  *string `json:"ContractEffectiveDate,omitempty"`

	CreatedByID           *string  `json:"CreatedById,omitempty"`
	CreatedDate           *string  `json:"CreatedDate,omitempty"`
	CreatorAccountID      *string  `json:"CreatorAccountId,omitempty"`
	CreatorInvoiceOwnerID *string  `json:"CreatorInvoiceOwnerId,omitempty"`
	ContractedMrr         *float64 `json:"contractedMrr"`
	CurrentTerm           *string  `json:"CurrentTerm,omitempty"`
	CurrentTermPeriodType *string  `json:"CurrentTermPeriodType,omitempty"`

	ID                    *string `json:"Id,omitempty"`
	InitialTerm           *string `json:"InitialTerm,omitempty"`
	InitialTermPeriodType *string `json:"InitialTermPeriodType,omitempty"`
	InvoiceOwnerID        string  `json:"InvoiceOwnerId"`
	IsInvoiceSeparate     *bool   `json:"IsInvoiceSeparate,omitempty"`
	Name                  *string `json:"Name,omitempty"`
	Notes                 *string `json:"Notes,omitempty"`

	OriginalCreatedDate *string `json:"OriginalCreatedDate,omitempty"`
	OriginalID          *string `json:"OriginalId,omitempty"`

	PreviousSubscriptionID *string `json:"PreviousSubscriptionId,omitempty"`

	RenewalSetting        string  `json:"RenewalSetting"`
	RenewalTerm           *string `json:"RenewalTerm,omitempty"`
	RenewalTermPeriodType *string `json:"RenewalTermPeriodType,omitempty"`

	ServiceActivationDate *string `json:"ServiceActivationDate,omitempty"`
	Status                *string `json:"Status,omitempty"`
	SubscriptionEndDate   *string `json:"SubscriptionEndDate,omitempty"`
	SubscriptionStartDate *string `json:"SubscriptionStartDate,omitempty"`
	TermEndDate           *string `json:"TermEndDate,omitempty"`
	TermStartDate         *string `json:"TermStartDate,omitempty"`
	TermType              *string `json:"TermType,omitempty"`
	UpdatedByID           *string `json:"UpdatedById,omitempty"`
	UpdatedDate           *string `json:"UpdatedDate,omitempty"`
	Version               *string `json:"Version,omitempty"`

	CpqBundleJSONID      *string `json:"CpqBundleJsonId__QT,omitempty"`
	OpportunityCloseDate *string `json:"OpportunityCloseDate__QT,omitempty"`
	OpportunityName      *string `json:"OpportunityName__QT,omitempty"`
	QuoteBusinessType    *string `json:"QuoteBusinessType__QT,omitempty"`
	QuoteNumber          *string `json:"QuoteNumber__QT,omitempty"`
	QuoteType            *string `json:"QuoteType__QT,omitempty"`
}

type SubscriptionRatePlanCharge struct {
	ApplyDiscountTo                interface{}   `json:"applyDiscountTo"`
	BillingDay                     string        `json:"billingDay"`
	BillingPeriod                  string        `json:"billingPeriod"`
	BillingPeriodAlignment         string        `json:"billingPeriodAlignment"`
	BillingTiming                  string        `json:"billingTiming"`
	ChargedThroughDate             string        `json:"chargedThroughDate"`
	Currency                       string        `json:"currency"`
	Description                    string        `json:"description"`
	DiscountAmount                 interface{}   `json:"discountAmount"`
	DiscountApplyDetails           []interface{} `json:"discountApplyDetails"`
	DiscountClass                  interface{}   `json:"discountClass"`
	DiscountLevel                  interface{}   `json:"discountLevel"`
	DiscountPercentage             interface{}   `json:"discountPercentage"`
	Dmrc                           float64       `json:"dmrc"`
	Done                           bool          `json:"done"`
	Dtcv                           interface{}   `json:"dtcv"`
	EffectiveEndDate               interface{}   `json:"effectiveEndDate"`
	EffectiveStartDate             string        `json:"effectiveStartDate"`
	EndDateCondition               string        `json:"endDateCondition"`
	ID                             string        `json:"id"`
	IncludedUnits                  interface{}   `json:"includedUnits"`
	ListPriceBase                  string        `json:"listPriceBase"`
	Model                          string        `json:"model"`
	Mrr                            float64       `json:"mrr"`
	Name                           string        `json:"name"`
	Number                         string        `json:"number"`
	NumberOfPeriods                interface{}   `json:"numberOfPeriods"`
	OriginalChargeID               string        `json:"originalChargeId"`
	OverageCalculationOption       interface{}   `json:"overageCalculationOption"`
	OveragePrice                   interface{}   `json:"overagePrice"`
	OverageUnusedUnitsCreditOption interface{}   `json:"overageUnusedUnitsCreditOption"`
	Price                          float64       `json:"price"`
	PriceChangeOption              string        `json:"priceChangeOption"`
	PriceIncreasePercentage        interface{}   `json:"priceIncreasePercentage"`
	PricingSummary                 string        `json:"pricingSummary"`
	ProcessedThroughDate           string        `json:"processedThroughDate"`
	ProductRatePlanChargeID        string        `json:"productRatePlanChargeId"`
	Quantity                       float64       `json:"quantity"`
	Segment                        int64         `json:"segment"`
	SmoothingModel                 interface{}   `json:"smoothingModel"`
	SpecificBillingPeriod          interface{}   `json:"specificBillingPeriod"`
	SpecificEndDate                interface{}   `json:"specificEndDate"`
	Tcv                            interface{}   `json:"tcv"`
	Tiers                          interface{}   `json:"tiers"`
	TriggerDate                    interface{}   `json:"triggerDate"`
	TriggerEvent                   string        `json:"triggerEvent"`
	Type                           string        `json:"type"`
	UnusedUnitsCreditRates         interface{}   `json:"unusedUnitsCreditRates"`
	Uom                            string        `json:"uom"`
	UpToPeriods                    interface{}   `json:"upToPeriods"`
	UpToPeriodsType                interface{}   `json:"upToPeriodsType"`
	UsageRecordRatingOption        interface{}   `json:"usageRecordRatingOption"`
	Version                        int64         `json:"version"`
}

type RatePlanCharge struct {
	AccountingCode                    *string  `json:"AccountingCode,omitempty"`
	ApplyDiscountTo                   *string  `json:"ApplyDiscountTo,omitempty"`
	BillCycleDay                      *int     `json:"BillCycleDay,omitempty"`
	BillCycleType                     *string  `json:"BillCycleType,omitempty"`
	BillingPeriod                     *string  `json:"BillingPeriod,omitempty"`
	BillingPeriodAlignment            *string  `json:"BillingPeriodAlignment,omitempty"`
	BillingTiming                     *string  `json:"BillingTiming,omitempty"`
	ChargedThroughDate                *string  `json:"ChargedThroughDate,omitempty"`
	ChargeModel                       *string  `json:"ChargeModel,omitempty"`
	ChargeNumber                      string   `json:"ChargeNumber"`
	ChargeType                        string   `json:"ChargeType"`
	CreatedByID                       *string  `json:"CreatedById,omitempty"`
	CreatedDate                       *string  `json:"CreatedDate,omitempty"`
	Description                       *string  `json:"Description,omitempty"`
	DiscountAmount                    *float64 `json:"DiscountAmount,omitempty"`
	DiscountClass                     *string  `json:"DiscountClass,omitempty"`
	DiscountLevel                     *string  `json:"DiscountLevel,omitempty"`
	DiscountPercentage                *float64 `json:"DiscountPercentage,omitempty"`
	DMRC                              *float64 `json:"DMRC,omitempty"`
	DTCV                              *float64 `json:"DTCV,omitempty"`
	EffectiveEndDate                  *string  `json:"EffectiveEndDate,omitempty"`
	EffectiveStartDate                *string  `json:"EffectiveStartDate,omitempty"`
	EndDateCondition                  *string  `json:"EndDateCondition,omitempty"`
	ID                                *string  `json:"Id,omitempty"`
	IncludedUnits                     *float64 `json:"IncludedUnits,omitempty"`
	IsLastSegment                     *bool    `json:"IsLastSegment,omitempty"`
	ListPriceBase                     *string  `json:"ListPriceBase,omitempty"`
	MRR                               *float64 `json:"MRR,omitempty"`
	Name                              string   `json:"Name"`
	NumberOfPeriods                   *int     `json:"NumberOfPeriods,omitempty"`
	OriginalID                        *string  `json:"OriginalId,omitempty"`
	OverageCalculationOption          *string  `json:"OverageCalculationOption,omitempty"`
	OveragePrice                      *float64 `json:"OveragePrice,omitempty"`
	OverageUnusedUnitsCreditOption    *string  `json:"OverageUnusedUnitsCreditOption,omitempty"`
	Price                             *float64 `json:"Price,omitempty"`
	PriceChangeOption                 *string  `json:"PriceChangeOption,omitempty"`
	PriceIncreasePercentage           *float64 `json:"PriceIncreasePercentage,omitempty"`
	ProcessedThroughDate              *string  `json:"ProcessedThroughDate,omitempty"`
	ProductRatePlanChargeID           string   `json:"ProductRatePlanChargeId"`
	Quantity                          *float64 `json:"Quantity,omitempty"`
	RatePlanID                        *string  `json:"RatePlanId,omitempty"`
	RatingGroup                       *string  `json:"RatingGroup,omitempty"`
	RevenueRecognitionRuleName        *string  `json:"RevenueRecognitionRuleName,omitempty"`
	RevRecCode                        *string  `json:"RevRecCode,omitempty"`
	RevRecTriggerCondition            *string  `json:"RevRecTriggerCondition,omitempty"`
	RolloverBalance                   *float64 `json:"RolloverBalance,omitempty"`
	Segment                           int      `json:"Segment"`
	SpecificBillingPeriod             *int     `json:"SpecificBillingPeriod,omitempty"`
	SpecificEndDate                   *string  `json:"SpecificEndDate,omitempty"`
	TCV                               *float64 `json:"TCV,omitempty"`
	TriggerDate                       *string  `json:"TriggerDate,omitempty"`
	TriggerEvent                      string   `json:"TriggerEvent"`
	UnusedUnitsCreditRates            *float64 `json:"UnusedUnitsCreditRates,omitempty"`
	UOM                               string   `json:"UOM"`
	UpdatedByID                       *string  `json:"UpdatedById,omitempty"`
	UpdatedDate                       *string  `json:"UpdatedDate,omitempty"`
	UpToPeriods                       *int     `json:"UpToPeriods,omitempty"`
	UpToPeriodsType                   *string  `json:"UpToPeriodsType,omitempty"`
	UsageRecordRatingOption           *string  `json:"UsageRecordRatingOption,omitempty"`
	UseDiscountSpecificAccountingCode *bool    `json:"UseDiscountSpecificAccountingCode,omitempty"`
	Version                           *int     `json:"Version,omitempty"`
	WeeklyBillCycleDay                *string  `json:"WeeklyBillCycleDay,omitempty"`

	// Number_of_Seats__c                *string  `json:"Number_of_Seats__c,omitempty"`
}

type SubscriptionRatePlan struct {
	ID                string                       `json:"id"`
	ProductID         string                       `json:"productId"`
	ProductName       string                       `json:"productName"`
	ProductRatePlanID string                       `json:"productRatePlanId"`
	ProductSku        string                       `json:"productSku"`
	RatePlanCharges   []SubscriptionRatePlanCharge `json:"ratePlanCharges"`
	RatePlanName      string                       `json:"ratePlanName"`
}
