package zuora

// Product A product is an item or service that your company sells.
// In the subscription economy, a product is generally a service
// that your customers subscribe to rather than a physical item that
// they purchase one time. For example, customers subscribe to a
// service that provides them a car when and where they need a
// car rather than buying a car outright.
//
// A Product object contains all of the items in a specific
// product, especially product rate plans and product rate plan
// charges. Each Product object can have multiple rate plans,
// which in turn can each have multiple rate plan charges.
// Each rate plan charge can have multiple tiers.
// Taken together, all of these items create a single product.
//
// Customers subscribe to products that are in your product
// catalog. Product objects collectively compose your product
// catalog. You create your product catalog by creating products.
// More info at:
// https://knowledgecenter.zuora.com/DC_Developers/G_SOAP_API/E1_SOAP_API_Object_Reference/Product
type Product struct {
	AllowFeatureChanges *bool   `json:"allowFeatureChanges,omitempty"`
	Category            *string `json:"category,omitempty"`
	CreatedByID         *string `json:"createdById,omitempty"`
	CreatedDate         *string `json:"createdDate,omitempty"`
	Description         *string `json:"description,omitempty"`
	EffectiveEndDate    string  `json:"effectiveEndDate"`
	EffectiveStartDate  string  `json:"effectiveStartDate"`
	ID                  *string `json:"id,omitempty"`
	Name                string  `json:"name"`
	SKU                 *string `json:"sku,omitempty"`
	UpdatedByID         *string `json:"updatedById,omitempty"`
	UpdatedDate         *string `json:"updatedDate,omitempty"`
}

// ProductRatePlan A product rate plan is the part of a product that
// your customers subscribe to. Each product can have multiple product
// rate plans, and each product rate plan can have multiple product
// rate plan charges, which are fees for products and their product rate plans.
// Don't confuse a product rate plan and a rate plan. A product rate
// plan is a rate plan that's part of a product in your product
// catalog. A rate plan is the specific rate plan in a subscription.
// More info:
// https://knowledgecenter.zuora.com/DC_Developers/G_SOAP_API/E1_SOAP_API_Object_Reference/ProductRatePlan
type ProductRatePlan struct {
	ActiveCurrencies   *string `json:"activeCurrencies,omitempty"`
	CreatedByID        *string `json:"createdById,omitempty"`
	CreatedDate        *string `json:"createdDate,omitempty"`
	Description        *string `json:"description,omitempty"`
	EffectiveEndDate   *string `json:"effectiveEndDate,omitempty"`
	EffectiveStartDate *string `json:"effectiveStartDate,omitempty"`
	ID                 *string `json:"id,omitempty"`
	Name               string  `json:"name"`
	ProductID          *string `json:"productId,omitempty"`
	UpdatedByID        *string `json:"updatedById,omitempty"`
	UpdatedDate        *string `json:"updatedDate,omitempty"`
}

// ProductRatePlanCharge A product rate plan charge represents a charge model or
// a set of fees associated with a product rate plan.
// A product rate plan charge represents a charge model or a set of fees associated
// with a product rate plan, which is the part of a product that your customers
// subscribe to. Each product rate plan can have multiple product rate plan charges.
// More info at:
// https://knowledgecenter.zuora.com/DC_Developers/G_SOAP_API/E1_SOAP_API_Object_Reference/ProductRatePlanCharge
type ProductRatePlanCharge struct {
	AccountingCode                    *string  `json:"accountingCode,omitempty"`
	ApplyDiscountTo                   *string  `json:"applyDiscountTo,omitempty"`
	BillCycleDay                      *int     `json:"billCycleDay,omitempty"`
	BillCycleType                     *string  `json:"billCycleType,omitempty"`
	BillingPeriod                     *string  `json:"billingPeriod,omitempty"`
	BillingPeriodAlignment            *string  `json:"billingPeriodAlignment,omitempty"`
	BillingTiming                     *string  `json:"billingTiming,omitempty"`
	ChargeModel                       *string  `json:"chargeModel,omitempty"`
	ChargeType                        *string  `json:"chargeType,omitempty"`
	CreatedByID                       *string  `json:"createdById,omitempty"`
	CreatedDate                       *string  `json:"createdDate,omitempty"`
	DefaultQuantity                   *float64 `json:"defaultQuantity,omitempty"`
	DeferredRevenueAccount            *string  `json:"deferredRevenueAccount,omitempty"`
	Description                       *string  `json:"description,omitempty"`
	DiscountClass                     *string  `json:"discountClass,omitempty"`
	DiscountLevel                     *string  `json:"discountLevel,omitempty"`
	EndDateCondition                  *string  `json:"endDateCondition,omitempty"`
	ID                                *string  `json:"id,omitempty"`
	IncludedUnits                     *float64 `json:"includedUnits,omitempty"`
	LegacyRevenueReporting            *bool    `json:"legacyRevenueReporting,omitempty"`
	ListPriceBase                     *string  `json:"listPriceBase,omitempty"`
	MaxQuantity                       *float64 `json:"maxQuantity,omitempty"`
	MinQuantity                       *float64 `json:"minQuantity,omitempty"`
	Name                              string   `json:"name"`
	NumberOfPeriod                    *int     `json:"numberOfPeriod,omitempty"`
	OverageCalculationOption          *string  `json:"overageCalculationOption,omitempty"`
	OverageUnusedUnitsCreditOption    *string  `json:"overageUnusedUnitsCreditOption,omitempty"`
	PriceChangeOption                 *string  `json:"priceChangeOption,omitempty"`
	PriceIncreasePercentage           *float64 `json:"priceIncreasePercentage,omitempty"`
	ProductRatePlanID                 *string  `json:"productRatePlanId,omitempty"`
	RatingGroup                       *string  `json:"ratingGroup,omitempty"`
	RecognizedRevenueAccount          *string  `json:"recognizedRevenueAccount,omitempty"`
	RevenueRecognitionRuleName        *string  `json:"revenueRecognitionRuleName,omitempty"`
	RevRecCode                        *string  `json:"revRecCode,omitempty"`
	RevRecTriggerCondition            *string  `json:"revRecTriggerCondition,omitempty"`
	SmoothingModel                    *string  `json:"smoothingModel,omitempty"`
	SpecificBillingPeriod             *int     `json:"specificBillingPeriod,omitempty"`
	Taxable                           *bool    `json:"taxable,omitempty"`
	TaxCode                           *string  `json:"taxCode,omitempty"`
	TaxMode                           *string  `json:"taxMode,omitempty"`
	TriggerEvent                      *string  `json:"triggerEvent,omitempty"`
	UOM                               *string  `json:"uom,omitempty"`
	UpdatedByID                       *string  `json:"updatedById,omitempty"`
	UpdatedDate                       *string  `json:"updatedDate,omitempty"`
	UpToPeriods                       *int     `json:"upToPeriods,omitempty"`
	UpToPeriodsType                   *string  `json:"upToPeriodsType,omitempty"`
	UsageRecordRatingOption           *string  `json:"usageRecordRatingOption,omitempty"`
	UseDiscountSpecificAccountingCode *bool    `json:"useDiscountSpecificAccountingCode,omitempty"`
	UseTenantDefaultForPriceChange    *bool    `json:"useTenantDefaultForPriceChange,omitempty"`
	WeeklyBillCycleDay                *string  `json:"weeklyBillCycleDay,omitempty"`
}

// ProductRatePlanChargeTier A product rate plan charge tier holds the prices for a product rate plan charge.
// Each product rate plan charge has at least one tier associated with it.
// Use the ProductRatePlanChargeTier object to represent a tier of charges in a ProductRatePlanCharge object.
// Don't confuse a product rate plan charge tier with tiered pricing.
// Tiered pricing is a charge model in which you create different charges or tiers for different quantites of an item.
// The tiers of usage fees in the Product: Family Plan example diagram represented tiered pricing.
type ProductRatePlanChargeTier struct {
	Tier
	Active                  *bool    `json:"active,omitempty"`
	CreatedByID             *string  `json:"createdById,omitempty"`
	CreatedDate             *string  `json:"createdDate,omitempty"`
	Currency                *string  `json:"currency,omitempty"`
	DiscountAmount          *float64 `json:"discountAmount,omitempty"`
	DiscountPercentage      *float64 `json:"discountPercentage,omitempty"`
	ID                      *string  `json:"id,omitempty"`
	IncludedUnits           *float64 `json:"includedUnits,omitempty"`
	IsOveragePrice          bool     `json:"isOveragePrice"`
	OveragePrice            *float64 `json:"overagePrice,omitempty"`
	ProductRatePlanChargeID string   `json:"productRatePlanChargeId"`
	UpdatedByID             *string  `json:"updatedById,omitempty"`
	UpdatedDate             *string  `json:"updatedDate,omitempty"`
}

// Tier Container for Volume, Tiered or Tiered with Overage charge models. Supports the following charge types:
//
// - One-time
//
// - Recurring
//
// - Usage-based
type Tier struct {
	EndingUnit   *float64 `json:"endingUnit,omitempty"`
	Price        float64  `json:"price,omitempty"`
	PriceFormat  *string  `json:"priceFormat,omitempty"`
	StartingUnit *float64 `json:"startingUnit,omitempty"`
	Tier         int      `json:"tier,omitempty"`
}
