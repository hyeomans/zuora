package zuora

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
	// Required.
	Amount float64 `json:"Amount"`

	// Required. Specifies if the refund is electronic or external.
	Type RefundType `json:"Type"`

	AccountID string `json:"AccountId,omitempty"`

	Comment string `json:"Comment,omitempty"`
	// Indicates how an external refund was issued to a customer. This field is only required if the Type field is set to External.
	MethodType string `json:"MethodType,omitempty"`

	// The unique ID of the payment method that the customer used to make the payment.
	// This field is only required if you create a non-referenced refund.
	PaymentID string `json:"PaymentId,omitempty"`

	// The unique ID of the payment method that the customer used to make the payment.
	// Specify a value for this field only if you're creating an electronic non-referenced refund.
	PaymentMethodID string `json:"PaymentMethodID,omitempty"`

	// ReasonCode A code identifying the reason for the transaction. Must be an existing reason code or empty.
	// If you do not specify a value, Zuora uses the default reason code.
	ReasonCode RefundReasonCode `json:"ReasonCode,omitempty"`

	// The date of the refund, in yyyy-mm-dd format.
	// The date of the refund cannot be before the payment date.
	// This field is only required if the Type field is set to External.
	// Zuora automatically generates this field for electronic refunds.
	RefundDate string `json:"RefundDate,omitempty"`

	// Use this field to pass gateway options.
	GatewayOptionData *GatewayOptionData `json:"GatewayOptionData,omitempty"`

	// RefundInvoicePaymentData Container for the refund invoice payment data.
	// This field is only required if you apply a full or partical refund against a payment attached to muliple invoices.
	RefundInvoicePaymentData *RefundInvoicePaymentData `json:"RefundInvoicePaymentData,omitempty"`

	// A payment gateway-specific field that maps Zuora to other gateways .
	SoftDescriptor string `json:"SoftDescriptor,omitempty"`

	// A payment gateway-specific field that maps Zuora to other gateways .
	SoftDescriptorPhone string `json:"SoftDescriptorPhone,omitempty"`

	// Specifies whether the refund is a refund payment or a credit balance.
	// This field is only required if you create a non-referenced refund.
	// If you creating an non-referenced refund, then set this value to CreditBalance
	SourceType string `json:"SourceType,omitempty"`
}

// RefundInvoicePaymentData --
type RefundInvoicePaymentData struct {
	RefundInvoicePayment []RefundInvoicePayment `json:"RefundInvoicePayment,omitempty"`
}

// GatewayOptionData --
type GatewayOptionData struct {
	GatewayOption []GatewayOption `json:"GatewayOption,omitempty"`
}

//GatewayOption --
type GatewayOption struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//RefundInvoicePayment --
type RefundInvoicePayment struct {
	CreatedByID      string  `json:"CreatedById,omitempty"`
	CreatedDate      string  `json:"CreatedDate,omitempty"`
	ID               string  `json:"Id,omitempty"`
	InvoiceID        string  `json:"InvoiceId,omitempty"`
	InvoicePaymentID string  `json:"InvoicePaymentId"`
	RefundAmount     float64 `json:"RefundAmount"`
	RefundID         string  `json:"RefundId"`
	UpdatedByID      string  `json:"UpdatedById,omitempty"`
	UpdatedDate      string  `json:"UpdatedDate,omitempty"`
}

// Refund A refund returns money to a customer - as opposed to a credit, which creates a customer credit balance
// that may be applied to reduce the amount owed to you. For instance,
// refunds are used when a customer cancels service and is no longer your customer.
// Refunds can also represent processed payments that are reversed, such as a chargeback or a direct debit payment reversal.
type Refund struct {
	AccountID               string  `json:"AccountId,omitempty"`
	AccountingCode          string  `json:"AccountingCode,omitempty"`
	Amount                  float64 `json:"Amount,omitempty"`
	CancelledOn             string  `json:"CancelledOn,omitempty"`
	Comment                 string  `json:"Comment,omitempty"`
	CreatedByID             string  `json:"CreatedById,omitempty"`
	CreatedDate             string  `json:"CreatedDate,omitempty"`
	Gateway                 string  `json:"Gateway,omitempty"`
	GatewayResponse         string  `json:"GatewayResponse,omitempty"`
	GatewayResponseCode     string  `json:"GatewayResponseCode,omitempty"`
	GatewayState            string  `json:"GatewayState"`
	ID                      string  `json:"Id,omitempty"`
	MarkedForSubmissionOn   string  `json:"MarkedForSubmissionOn,omitempty"`
	MethodType              string  `json:"MethodType,omitempty"`
	PaymentID               string  `json:"PaymentId,omitempty"`
	PaymentMethodID         string  `json:"PaymentMethodId,omitempty"`
	PaymentMethodSnapshotID string  `json:"PaymentMethodSnapshotId,omitempty"`
	ReasonCode              string  `json:"ReasonCode,omitempty"`
	ReferenceID             string  `json:"ReferenceID,omitempty"`
	RefundDate              string  `json:"RefundDate,omitempty"`
	RefundNumber            string  `json:"RefundNumber,omitempty"`
	RefundTransactionTime   string  `json:"RefundTransactionTime,omitempty"`
	SecondRefundReferenceID string  `json:"SecondRefundReferenceId,omitempty"`
	SettledOn               string  `json:"SettledOn,omitempty"`
	SoftDescriptor          string  `json:"SoftDescriptor,omitempty"`
	SoftDescriptorPhone     string  `json:"SoftDescriptorPhone,omitempty"`
	SourceType              string  `json:"SourceType,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	SubmittedOn             string  `json:"SubmittedOn,omitempty"`
	TransferredToAccounting string  `json:"TransferredToAccounting,omitempty"`
	Type                    string  `json:"Type"`
	UpdatedByID             string  `json:"UpdatedById,omitempty"`
	UpdatedDate             string  `json:"UpdatedDate,omitempty"`
}

//RefundCreateResonse --
type RefundCreateResonse struct {
	Response
	ID string `json:"Id"`
}

// InvoicePayment intermediary table between Refunds & Refunds.
// You can use PaymentID from Refund to search all Invoices using a ZOQL query.
type InvoicePayment struct {
	Amount       float64 `json:"Amount"`
	CreatedByID  string  `json:"CreatedById,omitempty"`
	CreatedDate  string  `json:"CreatedDate,omitempty"`
	ID           string  `json:"Id,omitempty"`
	InvoiceID    string  `json:"InvoiceId"`
	PaymentID    string  `json:"PaymentId,omitempty"`
	RefundAmount float64 `json:"RefundAmount,omitempty"`
	UpdatedByID  string  `json:"UpdatedById,omitempty"`
	UpdatedDate  string  `json:"UpdatedDate,omitempty"`
}
