package zuora

// AdjustmentSourceType supported by Zuora
type AdjustmentSourceType string

//AdjustmentTypeInvoiceDetail --
var AdjustmentTypeInvoiceDetail AdjustmentSourceType = "InvoiceDetail"

//AdjustmentTypeTax --
var AdjustmentTypeTax AdjustmentSourceType = "Tax"

// AdjustmentType Indicates whether the adjustment credits or debits the invoice item amount.
type AdjustmentType string

// AdjustmentTypeCredit --
var AdjustmentTypeCredit AdjustmentType = "Credit"

// AdjustmentTypeDebit --
var AdjustmentTypeDebit AdjustmentType = "Debit"

//InvoiceItemAdjustmentCreatePayload use it when calling Actions.Create endpoint.
type InvoiceItemAdjustmentCreatePayload struct {
	// The accounting code for the invoice item. Accounting codes group transactions that contain similar accounting attributes.
	AccountingCode string `json:"AccountingCode,omitempty"`

	// Required. The date when the invoice item adjustment is applied. This date must be the same as the invoice's date or later.
	AdjustmentDate string `json:"AdjustmentDate"`

	// A unique string to identify an individual invoice item adjustment.
	AdjustmentNumber string `json:"AdjustmentNumber,omitempty"`

	// Required. The amount of the invoice item adjustment. The value of Amount must be positive.
	// Use the required parameter Type to either credit or charge (debit) this amount on the invoice.
	Amount float64 `json:"Amount"`

	// Use this field to record comments about the invoice item adjustment.
	Comment string `json:"Comment,omitempty"`

	// Records the deferred accounting code in the finance system.
	DeferredRevenueAccount string `json:"DeferredRevenueAccount,omitempty"`

	// Required. The ID of the invoice associated with the adjustment. The adjustment invoice item is in this invoice. This field is optional if you specify a value for the InvoiceNumber field.
	InvoiceID string `json:"InvoiceId"`

	// Required. The unique identification number for the invoice that contains the invoice item. This field is optional if you specify a value for the InvoiceId field.
	InvoiceNumber string `json:"InvoiceNumber"`

	// A code identifying the reason for the transaction. Must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code.
	ReasonCode string `json:"ReasonCode,omitempty"`

	// Records the recognized accounting code in the finance system.
	RecognizedRevenueAccount string `json:"RecognizedRevenueAccount,omitempty"`

	// A code to reference an object external to Zuora. For example, you can use this field to reference a case number in an external system.
	ReferenceID string `json:"ReferenceId,omitempty"`

	// Required. The ID of the item specified in the SourceType field.
	SourceID string `json:"SourceId"`

	// Required. The type of adjustment.
	SourceType AdjustmentSourceType `json:"SourceType"`

	// Required. Indicates whether the adjustment credits or debits the invoice item amount.
	Type AdjustmentType `json:"Type"`
}

// InvoiceItemAdjustment use this for GET requests
type InvoiceItemAdjustment struct {
	InvoiceItemAdjustmentCreatePayload
	AccountID               string `json:"AccountId,omitempty"`
	CancelledByID           string `json:"CancelledById,omitempty"`
	CancelledDate           string `json:"CancelledDate,omitempty"`
	CreatedByID             string `json:"CreatedById,omitempty"`
	CreatedDate             string `json:"CreatedDate,omitempty"`
	CustomerName            string `json:"CustomerName"`
	CustomerNumber          string `json:"CustomerNumber"`
	ID                      string `json:"Id,omitempty"`
	InvoiceItemName         string `json:"InvoiceItemName"`
	ServiceEndDate          string `json:"ServiceEndDate,omitempty"`
	ServiceStartDate        string `json:"ServiceStartDate,omitempty"`
	Status                  string `json:"Status,omitempty"`
	TransferredToAccounting string `json:"TransferredToAccounting,omitempty"`
	UpdatedByID             string `json:"UpdatedById,omitempty"`
	UpdatedDate             string `json:"UpdatedDate,omitempty"`
}
