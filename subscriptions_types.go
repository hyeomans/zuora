package zuora

// Subscription has all the possible properties given by Zuora. This comes from
// the Describe endpoint.
type Subscription struct {
	AccountID              string  `json:"accountId"`
	AncestorAccountID      *string `json:"ancestorAccountId,omitempty"`
	AutoRenew              *bool   `json:"autoRenew,omitempty"`
	CancelledDate          *string `json:"cancelledDate,omitempty"`
	ContractAcceptanceDate *string `json:"contractAcceptanceDate,omitempty"`
	ContractEffectiveDate  *string `json:"contractEffectiveDate,omitempty"`
	CpqBundleJSONIDQT      *string `json:"cpqBundleJsonId__QT,omitempty"`
	CreatedByID            *string `json:"createdById,omitempty"`
	CreatedDate            *string `json:"createdDate,omitempty"`
	CreatorAccountID       *string `json:"creatorAccountId,omitempty"`
	CreatorInvoiceOwnerID  *string `json:"creatorInvoiceOwnerId,omitempty"`
	CurrentTerm            *int    `json:"currentTerm,omitempty"`
	CurrentTermPeriodType  *string `json:"currentTermPeriodType,omitempty"`
	ID                     *string `json:"id,omitempty"`
	InitialTerm            *int    `json:"initialTerm,omitempty"`
	InitialTermPeriodType  *string `json:"initialTermPeriodType,omitempty"`
	InvoiceOwnerID         string  `json:"invoiceOwnerId"`
	IsInvoiceSeparate      *bool   `json:"isInvoiceSeparate,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Notes                  *string `json:"notes,omitempty"`
	OpportunityCloseDateQT *string `json:"opportunityCloseDate__QT,omitempty"`
	OpportunityNameQT      *string `json:"opportunityName__QT,omitempty"`
	OriginalCreatedDate    *string `json:"originalCreatedDate,omitempty"`
	OriginalID             *string `json:"originalId,omitempty"`
	PreviousSubscriptionID *string `json:"previousSubscriptionId,omitempty"`
	QuoteBusinessTypeQT    *string `json:"quoteBusinessType__QT,omitempty"`
	QuoteNumberQT          *string `json:"quoteNumber__QT,omitempty"`
	QuoteTypeQT            *string `json:"quoteType__QT,omitempty"`
	RenewalSetting         string  `json:"renewalSetting"`
	RenewalTerm            *int    `json:"renewalTerm,omitempty"`
	RenewalTermPeriodType  *string `json:"renewalTermPeriodType,omitempty"`
	ServiceActivationDate  *string `json:"serviceActivationDate,omitempty"`
	Status                 *string `json:"status,omitempty"`
	SubscriptionEndDate    *string `json:"subscriptionEndDate,omitempty"`
	SubscriptionStartDate  *string `json:"subscriptionStartDate,omitempty"`
	TermEndDate            *string `json:"termEndDate,omitempty"`
	TermStartDate          *string `json:"termStartDate,omitempty"`
	TermType               *string `json:"termType,omitempty"`
	UpdatedByID            *string `json:"updatedById,omitempty"`
	UpdatedDate            *string `json:"updatedDate,omitempty"`
	Version                *int    `json:"version,omitempty"`
}

// RatePlan has all the possible properties given by Zuora.
// Don't confuse RatePlan with ProductRatePlan
type RatePlan struct {
	ID                              *string `json:"id,omitempty"`
	AmendmentID                     *string `json:"amendmentId,omitempty"`
	AmendmentSubscriptionRatePlanID *string `json:"amendmentSubscriptionRatePlanId,omitempty"`
	AmendmentType                   *string `json:"amendmentType,omitempty"`
	CreatedByID                     *string `json:"createdById,omitempty"`
	CreatedDate                     *string `json:"createdDate,omitempty"`
	Name                            string  `json:"name"`
	ProductRatePlanID               *string `json:"productRatePlanId,omitempty"`
	SubscriptionID                  *string `json:"subscriptionId,omitempty"`
	UpdatedByID                     *string `json:"updatedById,omitempty"`
	UpdatedDate                     *string `json:"updatedDate,omitempty"`
}

// RatePlanCharge has all the possible properties given by Zuora.
type RatePlanCharge struct {
	AccountingCode                    *string  `json:"accountingCode,omitempty"`
	ApplyDiscountTo                   *string  `json:"applyDiscountTo,omitempty"`
	BillCycleDay                      *int     `json:"BillCycleDay,omitempty"`
	BillCycleType                     *string  `json:"BillCycleType,omitempty"`
	BillingPeriod                     *string  `json:"BillingPeriod,omitempty"`
	BillingPeriodAlignment            *string  `json:"BillingPeriodAlignment,omitempty"`
	BillingTiming                     *string  `json:"BillingTiming,omitempty"`
	ChargedThroughDate                *string  `json:"chargedThroughDate,omitempty"`
	ChargeModel                       *string  `json:"chargeModel,omitempty"`
	ChargeNumber                      string   `json:"chargeNumber"`
	ChargeType                        string   `json:"chargeType"`
	CreatedByID                       *string  `json:"createdById,omitempty"`
	CreatedDate                       *string  `json:"createdDate,omitempty"`
	Description                       *string  `json:"description,omitempty"`
	DiscountAmount                    *float64 `json:"discountAmount,omitempty"`
	DiscountClass                     *string  `json:"discountClass,omitempty"`
	DiscountLevel                     *string  `json:"discountLevel,omitempty"`
	DiscountPercentage                *float64 `json:"discountPercentage,omitempty"`
	DMRC                              *float64 `json:"dmrc,omitempty"`
	DTCV                              *float64 `json:"dtcv,omitempty"`
	EffectiveEndDate                  *string  `json:"EffectiveEndDate,omitempty"`
	EffectiveStartDate                *string  `json:"EffectiveStartDate,omitempty"`
	EndDateCondition                  *string  `json:"EndDateCondition,omitempty"`
	ID                                *string  `json:"id,omitempty"`
	IncludedUnits                     *float64 `json:"includedUnits,omitempty"`
	IsLastSegment                     *bool    `json:"isLastSegment,omitempty"`
	ListPriceBase                     *string  `json:"ListPriceBase,omitempty"`
	MRR                               *float64 `json:"mrr,omitempty"`
	Name                              string   `json:"name"`
	NumberOfPeriods                   *int     `json:"numberOfPeriods,omitempty"`
	OriginalID                        *string  `json:"originalId,omitempty"`
	OverageCalculationOption          *string  `json:"overageCalculationOption,omitempty"`
	OveragePrice                      *float64 `json:"overagePrice,omitempty"`
	OverageUnusedUnitsCreditOption    *string  `json:"overageUnusedUnitsCreditOption,omitempty"`
	Price                             *float64 `json:"price,omitempty"`
	PriceChangeOption                 *string  `json:"priceChangeOption,omitempty"`
	PriceIncreasePercentage           *float64 `json:"priceIncreasePercentage,omitempty"`
	ProcessedThroughDate              *string  `json:"processedThroughDate,omitempty"`
	ProductRatePlanChargeID           string   `json:"productRatePlanChargeId"`
	Quantity                          *float64 `json:"quantity,omitempty"`
	RatePlanID                        *string  `json:"ratePlanId,omitempty"`
	RatingGroup                       *string  `json:"ratingGroup,omitempty"`
	RevenueRecognitionRuleName        *string  `json:"revenueRecognitionRuleName,omitempty"`
	RevRecCode                        *string  `json:"revRecCode,omitempty"`
	RevRecTriggerCondition            *string  `json:"revRecTriggerCondition,omitempty"`
	RolloverBalance                   *float64 `json:"rolloverBalance,omitempty"`
	Segment                           int      `json:"segment"`
	SpecificBillingPeriod             *int     `json:"specificBillingPeriod,omitempty"`
	SpecificEndDate                   *string  `json:"specificEndDate,omitempty"`
	TCV                               *float64 `json:"tcv,omitempty"`
	TriggerDate                       *string  `json:"triggerDate,omitempty"`
	TriggerEvent                      string   `json:"triggerEvent"`
	UnusedUnitsCreditRates            *float64 `json:"unusedUnitsCreditRates,omitempty"`
	UOM                               string   `json:"uom"`
	UpdatedByID                       *string  `json:"updatedById,omitempty"`
	UpdatedDate                       *string  `json:"updatedDate,omitempty"`
	UpToPeriods                       *int     `json:"upToPeriods,omitempty"`
	UpToPeriodsType                   *string  `json:"upToPeriodsType,omitempty"`
	UsageRecordRatingOption           *string  `json:"usageRecordRatingOption,omitempty"`
	UseDiscountSpecificAccountingCode *bool    `json:"useDiscountSpecificAccountingCode,omitempty"`
	Version                           *int     `json:"version,omitempty"`
	WeeklyBillCycleDay                *string  `json:"weeklyBillCycleDay,omitempty"`
}

// SubscriptionUpdate Update subscription
//
// ** CHECK HERE FOR MORE INFORMATION **
//
// https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription
//
// You will have to defined Add, Remove, Update in a custom struct because they accept custom fields.
type SubscriptionUpdate struct {
	// Add                              []SubscriptionAddRatePlan    `json:"add,omitempty"`
	// Remove                           []SubscriptionRemoveRatePlan `json:"remove,omitempty"`
	// Update                           []SubscriptionUpdateRatePlan `json:"update,omitempty"`
	ApplyCreditBalance               *bool   `json:"applyCreditBalance,omitempty"`
	AutoRenew                        *bool   `json:"autoRenew,omitempty"`
	Collect                          *bool   `json:"collect,omitempty"`
	CurrentTerm                      *int    `json:"currentTerm,omitempty"`
	CurrentTermPeriodType            *string `json:"currentTermPeriodType,omitempty"`
	DocumentDate                     *string `json:"documentDate,omitempty"`
	IncludeExistingDraftDocItems     *bool   `json:"includeExistingDraftDocItems,omitempty"`
	IncludeExistingDraftInvoiceItems *bool   `json:"includeExistingDraftInvoiceItems,omitempty"`
	Invoice                          *bool   `json:"invoice,omitempty"`
	InvoiceCollect                   *bool   `json:"invoiceCollect,omitempty"`
	InvoiceSeparately                *bool   `json:"invoiceSeparately,omitempty"`
	InvoiceTargetDate                *string `json:"invoiceTargetDate,omitempty"`
	Notes                            *string `json:"notes,omitempty"`
	Preview                          *bool   `json:"preview,omitempty"`
	PreviewType                      *string `json:"previewType,omitempty"`
	RenewalSetting                   *string `json:"renewalSetting,omitempty"`
	RenewalTerm                      *int    `json:"renewalTerm,omitempty"`
	RenewalTermPeriodType            *string `json:"renewalTermPeriodType,omitempty"`
	RunBilling                       *bool   `json:"runBilling,omitempty"`
	TargetDate                       *string `json:"targetDate,omitempty"`
	TermStartDate                    *string `json:"termStartDate,omitempty"`
	TermType                         *string `json:"termType,omitempty"`
}

// SubscriptionAddRatePlan Custom struct to add a rate plan whem modifying an existing subscription
type SubscriptionAddRatePlan struct {
	// You will have to define your custom ChargeOverrides given that SubscriptionAddChargeOverride
	// accepts custom properties.
	// ChargeOverrides        []SubscriptionAddChargeOverride `json:"chargeOverrides,omitempty"`
	ContractEffectiveDate  string  `json:"contractEffectiveDate,omitempty"`
	CustomerAcceptanceDate *string `json:"customerAcceptanceDate,omitempty"`
	ProductRatePlanID      string  `json:"productRatePlanId,omitempty"`
	ServiceActivationDate  *string `json:"serviceActivationDate,omitempty"`
}

// SubscriptionAddChargeOverride Custom struct to modify Rate Plan Charges on an existing subscription.
type SubscriptionAddChargeOverride struct {
	ApplyDiscountTo                *string  `json:"applyDiscountTo,omitempty"`
	BillCycleDay                   *string  `json:"billCycleDay,omitempty"`
	BillCycleType                  *string  `json:"billCycleType,omitempty"`
	BillingPeriod                  *string  `json:"billingPeriod,omitempty"`
	BillingPeriodAlignment         *string  `json:"billingPeriodAlignment,omitempty"`
	BillingTiming                  *string  `json:"billingTiming,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	DiscountAmount                 *string  `json:"discountAmount,omitempty"`
	DiscountLevel                  *string  `json:"discountLevel,omitempty"`
	DiscountPercentage             *float64 `json:"discountPercentage,omitempty"`
	EndDateCondition               *string  `json:"endDateCondition,omitempty"`
	IncludedUnits                  *float64 `json:"includedUnits,omitempty"`
	ListPriceBase                  *string  `json:"listPriceBase,omitempty"`
	NumberOfPeriod                 *int     `json:"numberOfPeriod,omitempty"`
	OveragePrice                   *float64 `json:"overagePrice,omitempty"`
	OverageUnusedUnitsCreditOption *string  `json:"overageUnusedUnitsCreditOption,omitempty"`
	Price                          *float64 `json:"price,omitempty"`
	PriceChangeOption              *string  `json:"priceChangeOption,omitempty"`
	PriceIncreasePercentage        *float64 `json:"priceIncreasePercentage,omitempty"`
	ProductRatePlanChargeID        string   `json:"productRatePlanChargeId,omitempty"`
	Quantity                       *float64 `json:"quantity,omitempty"`
	RatingGroup                    *string  `json:"ratingGroup,omitempty"`
	SpecificBillingPeriod          *int     `json:"specificBillingPeriod,omitempty"`
	SpecificEndDate                *string  `json:"specificEndDate,omitempty"`
	Tiers                          []Tier   `json:"tiers,omitempty"`
	TriggerEvent                   *string  `json:"triggerEvent,omitempty"`
	TriggerDate                    *string  `json:"triggerDate,omitempty"`
	UnusedUnitCreditRates          *float64 `json:"unusedUnitCreditRates,omitempty"`
	UpToPeriods                    *int     `json:"upToPeriods,omitempty"`
	UpToPeriodsType                *string  `json:"upToPeriodsType,omitempty"`
}

// SubscriptionRemoveRatePlan Custom struct to remove a rate plan whem modifying an existing subscription
type SubscriptionRemoveRatePlan struct {
	ContractEffectiveDate  string  `json:"contractEffectiveDate,omitempty"`
	CustomerAcceptanceDate *string `json:"customerAcceptanceDate,omitempty"`
	RatePlanID             string  `json:"ratePlanId,omitempty"`
	ServiceActivationDate  *string `json:"serviceActivationDate,omitempty"`
}

// SubscriptionUpdateRatePlan Custom struct to update a rate plan whem modifying an existing subscription
type SubscriptionUpdateRatePlan struct {
	// You will have to define your custom ChargeUpdateDetail given that ChargeUpdateDetails
	// accepts custom properties.
	// ChargeUpdateDetails    []ChargeUpdateDetail `json:"chargeUpdateDetails,omitempty"`
	ContractEffectiveDate  string  `json:"contractEffectiveDate,omitempty"`
	CustomerAcceptanceDate *string `json:"customerAcceptanceDate,omitempty"`
	RatePlanID             string  `json:"ratePlanId,omitempty"`
	ServiceActivationDate  *string `json:"serviceActivationDate,omitempty"`
	SpecificUpdateDate     *string `json:"specificUpdateDate,omitempty"`
}

// ChargeUpdateDetail Custom struct to modify Product Rate Plan charges in an existing subscription.
type ChargeUpdateDetail struct {
	BillingPeriodAlignment  *string  `json:"billingPeriodAlignment,omitempty"`
	Description             *string  `json:"description,omitempty"`
	IncludedUnits           *float64 `json:"includedUnits,omitempty"`
	OveragePrice            *float64 `json:"overagePrice,omitempty"`
	Price                   *float64 `json:"price,omitempty"`
	PriceChangeOption       *string  `json:"priceChangeOption,omitempty"`
	PriceIncreasePercentage *float64 `json:"priceIncreasePercentage,omitempty"`
	Quantity                *float64 `json:"quantity,omitempty"`
	RatePlanChargeID        string   `json:"ratePlanChargeId,omitempty"`
	Tiers                   []Tier   `json:"tiers,omitempty"`
	TriggerDate             *string  `json:"triggerDate,omitempty"`
	TriggerEvent            *string  `json:"triggerEvent,omitempty"`
}

// SubscriptionCancellation is the request body schema to cancell a subscription
// More info at:
// https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription
type SubscriptionCancellation struct {
	ApplyCreditBalance        *bool   `json:"applyCreditBalance,omitempty"`
	CancellationEffectiveDate *string `json:"cancellationEffectiveDate,omitempty"`
	CancellationPolicy        string  `json:"cancellationPolicy"`
	Collect                   *bool   `json:"collect,omitempty"`
	DocumentDate              *string `json:"documentDate,omitempty"`
	Invoice                   *bool   `json:"invoice,omitempty"`
	InvoiceCollect            bool    `json:"invoiceCollect"`
	InvoiceTargetDate         *string `json:"invoiceTargetDate,omitempty"`
	RunBilling                *bool   `json:"runBilling,omitempty"`
	TargetDate                *string `json:"targetDate,omitempty"`
}

// SubscriptionCancellationResponse response when cancelling an Subscription
type SubscriptionCancellationResponse struct {
	CancelledDate  string   `json:"cancelledDate"`
	CreditMemoID   string   `json:"creditMemoId"`
	InvoiceID      string   `json:"invoiceId"`
	PaidAmount     string   `json:"paidAmount"`
	PaymentID      string   `json:"paymentId"`
	SubscriptionID string   `json:"subscriptionId"`
	Success        bool     `json:"success"`
	TotalDeltaMrr  *float64 `json:"totalDeltaMrr"`
	TotalDeltaTcv  *float64 `json:"totalDeltaTcv"`
}
