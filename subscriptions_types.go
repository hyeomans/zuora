package zuora

//SubscriptionChargeUpdateDetails is part of SubscriptionUpdate payload inside SubscriptionUpdate
type SubscriptionChargeUpdateDetails []struct {
	Quantity         int    `json:"quantity"`
	RatePlanChargeID string `json:"ratePlanChargeId"`
}

//SubscriptionUpdate is part of SubscriptionUpdate payload
type SubscriptionUpdate struct {
	ChargeUpdateDetails   []SubscriptionChargeUpdateDetails `json:"chargeUpdateDetails"`
	ContractEffectiveDate string                            `json:"contractEffectiveDate"`
	RatePlanID            string                            `json:"ratePlanId"`
}

//SubscriptionUpdateMinimalPayload represents the minimal payload for updating a subscription
type SubscriptionUpdateMinimalPayload struct {
	Updates []SubscriptionUpdate `json:"update"`
}

//SubscriptionUpdateFullPayload represents the full payload declared in Zuora documentation
type SubscriptionUpdateFullPayload struct {
	AutoRenew             bool                 `json:"autoRenew"`
	Collect               bool                 `json:"collect"`
	CurrentTerm           string               `json:"currentTerm"`
	CurrentTermPeriodType string               `json:"currentTermPeriodType"`
	Notes                 string               `json:"notes"`
	RenewalSetting        string               `json:"renewalSetting"`
	RenewalTerm           string               `json:"renewalTerm"`
	RenewalTermPeriodType string               `json:"renewalTermPeriodType"`
	RunBilling            bool                 `json:"runBilling"`
	TermType              string               `json:"termType"`
	Updates               []SubscriptionUpdate `json:"update"`
}

//SubscriptionUpdateResponse once a subscription is updated, it contains useful information
type SubscriptionUpdateResponse struct {
	Success        bool    `json:"success"`
	SubscriptionID string  `json:"subscriptionId"`
	TotalDeltaMrr  float64 `json:"totalDeltaMrr"`
	TotalDeltaTcv  float64 `json:"totalDeltaTcv"`
}

//SubscriptionCancellationPayload payload to cancel a subscription
type SubscriptionCancellationPayload struct {
	CancellationPolicy CancellationPolicy `json:"cancellationPolicy"`
}

//SubscriptionCancellationResponse response when cancelling a subscription
type SubscriptionCancellationResponse struct {
	Success        bool    `json:"success"`
	SubscriptionID string  `json:"subscriptionId"`
	CancelledDate  string  `json:"cancelledDate"`
	TotalDeltaMrr  float64 `json:"totalDeltaMrr"`
	TotalDeltaTcv  float64 `json:"totalDeltaTcv"`
	InvoiceID      string  `json:"invoiceId"`
}

//CancellationPolicy Cancellation method
type CancellationPolicy string

//EndOfCurrentTerm is a valid cancellation policy in payload
const EndOfCurrentTerm CancellationPolicy = "EndOfCurrentTerm"

//EndOfLastInvoicePeriod is a valid cancellation policy in payload
const EndOfLastInvoicePeriod CancellationPolicy = "EndOfLastInvoicePeriod"

//SpecificDate is a valid cancellation policy in payload
const SpecificDate CancellationPolicy = "SpecificDate"
