package zuora

//GatewayOption --
type GatewayOption struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//RefundInvoicePayment --
type RefundInvoicePayment struct {
	InvoiceID    string `json:"InvoiceId"`
	RefundAmount string `json:"RefundAmount"`
}

//RefundType --
type RefundType string

//ElectronigRefundType --
var ElectronigRefundType RefundType = "Electronic"

//ExternalRefundType --
var ExternalRefundType RefundType = "External"

//RefundReasonCode --
type RefundReasonCode string

//RefundReasonCodeCustomerSatisfaction --
var RefundReasonCodeCustomerSatisfaction RefundReasonCode = "Customer Satisfaction"

//RefundReasonCodeChargeback --
var RefundReasonCodeChargeback RefundReasonCode = "Chargeback"

//RefundReasonCodePaymentReversal --
var RefundReasonCodePaymentReversal RefundReasonCode = "Payment Reversal"

//RefundCreatePayload --
type RefundCreatePayload struct {
	// AccountID The ID of the account associated with this refund.
	// This field is only required if you create a non-referenced refund.
	// Don't specify a value for any other type of refund; Zuora associates the refund automatically with the account from the associated payment. Character limit: 32 Values: a valid account ID
	AccountID string `json:"AccountId,omitempty"`
	// Amount - REQUIRED
	// The amount of the refund.
	// The amount can't exceed the amount of the associated payment.
	// If the original payment was applied to a single invoice, then you can create a partial refund.
	// However, if the payment was applies to multiple invoices, then you can only make a partial
	// refund through the web-based UI, not through the API.
	// Character limit: 16 Values: a valid currency amount
	Amount  float64 `json:"Amount"`
	Comment string  `json:"Comment,omitempty"`
	// GatewayOptionData A field used to pass gateway options.
	// Zuora allows you to pass in special gateway-specific parameters for payments that go through the
	// Adyen, Autorize.et, CyberSource, Merchant eSolutions, Orbital (Chase Paymentech), QValent, Vantiv, and Verifi gateways.
	// For each of these special parameters, you supply the name-value pair and Zuora passes it to the gateway.
	// This allows you to add functionality that's supported by a specific gateway but currently not supported by Zuora.
	GatewayOptionData *struct {
		GatewayOption []GatewayOption `json:"GatewayOption,omitempty"`
	} `json:"GatewayOptionData,omitempty"`
	// MethodType Indicates how an external refund was issued to a customer. This field is only required if the Type field is set to External.
	MethodType string `json:"MethodType,omitempty"`
	// PaymentID The unique ID of the payment method that the customer used to make the payment.
	// This field is only required if you create a non-referenced refund. Character limit: 32 V**alues**: a valid payment method ID
	PaymentID string `json:"PaymentId"`
	// ReasonCode A code identifying the reason for the transaction. Must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code. Character limit: 32 V**alues**: a valid reason code
	ReasonCode RefundReasonCode `json:"ReasonCode,omitempty"`
	// RefundDate The date of the refund, in yyyy-mm-dd format.
	// The date of the refund cannot be before the payment date.
	// This field is only required if the Type field is set to External.
	// Zuora automatically generates this field for electronic refunds. Character limit: 29
	RefundDate string `json:"RefundDate,omitempty"`

	// RefundInvoicePaymentData Container for the refund invoice payment data. This field is only required if you apply a full or partical refund against a payment attached to muliple invoices.
	RefundInvoicePaymentData *struct {
		RefundInvoicePayment []RefundInvoicePayment `json:"RefundInvoicePayment,omitempty"`
	} `json:"RefundInvoicePaymentData,omitempty"`

	// SoftDescriptor A payment gateway-specific field that maps Zuora to other gateways .
	SoftDescriptor string `json:"SoftDescriptor,omitempty"`
	// SoftDescriptorPhone A payment gateway-specific field that maps Zuora to other gateways .
	SoftDescriptorPhone string `json:"SoftDescriptorPhone,omitempty"`
	// SourceType Specifies whether the refund is a refund payment or a credit balance.
	// This field is only required if you create a non-referenced refund.
	// If you creating an non-referenced refund, then set this value to CreditBalance
	SourceType string `json:"SourceType,omitempty"`
	// Type Specifies if the refund is electronic or external.
	Type RefundType `json:"Type"`
}

//RefundCreateResonse --
type RefundCreateResonse struct {
	Response
	ID string `json:"Id"`
}
