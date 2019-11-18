package zuora

//InvoiceFile --
type InvoiceFile struct {
	ID            string `json:"id"`
	VersionNumber int64  `json:"versionNumber"`
	PdfFileURL    string `json:"pdfFileUrl"`
}

//InvoiceFilesResponse --
type InvoiceFilesResponse struct {
	InvoiceFiles []InvoiceFile `json:"invoiceFiles"`
	Response
}

// TaxationItem --
type TaxationItem struct {
	Data []struct {
		Balance            float64     `json:"balance"`
		PaymentAmount      float64     `json:"paymentAmount"`
		CreditAmount       float64     `json:"creditAmount"`
		ID                 string      `json:"id"`
		TaxAmount          float64     `json:"taxAmount"`
		Name               string      `json:"name"`
		ExemptAmount       float64     `json:"exemptAmount"`
		Jurisdiction       string      `json:"jurisdiction"`
		LocationCode       string      `json:"locationCode"`
		TaxCode            string      `json:"taxCode"`
		TaxCodeDescription interface{} `json:"taxCodeDescription"`
		TaxDate            string      `json:"taxDate"`
		TaxRate            float64     `json:"taxRate"`
		TaxRateDescription string      `json:"taxRateDescription"`
		TaxRateType        string      `json:"taxRateType"`
	} `json:"data"`
}

// InvoiceItem --
type InvoiceItem struct {
	ID                string       `json:"id"`
	SubscriptionName  string       `json:"subscriptionName"`
	SubscriptionID    string       `json:"subscriptionId"`
	ServiceStartDate  string       `json:"serviceStartDate"`
	ServiceEndDate    string       `json:"serviceEndDate"`
	Balance           float64      `json:"balance,omitempty"`
	ChargeAmount      float64      `json:"chargeAmount"`
	ChargeDescription string       `json:"chargeDescription"`
	ChargeName        string       `json:"chargeName"`
	ChargeID          string       `json:"chargeId"`
	ProductName       string       `json:"productName"`
	Quantity          float64      `json:"quantity"`
	TaxAmount         float64      `json:"taxAmount"`
	UnitOfMeasure     string       `json:"unitOfMeasure"`
	AppliedToItemID   interface{}  `json:"appliedToItemId"`
	TaxationItems     TaxationItem `json:"taxationItems"`
}

//InvoiceItemZoql this includes all the fields that can be retrieved when
// calling InvoiceItem through ZOQL query.
type InvoiceItemZoql struct {
	AccountingCode         string  `json:"AccountingCode,omitempty"`
	AppliedToChargeNumber  string  `json:"AppliedToChargeNumber,omitempty"`
	AppliedToInvoiceItemID string  `json:"AppliedToInvoiceItemId,omitempty"`
	Balance                float64 `json:"Balance,omitempty"`
	ChargeAmount           float64 `json:"ChargeAmount,omitempty"`
	ChargeDate             string  `json:"ChargeDate,omitempty"`
	ChargeDescription      string  `json:"ChargeDescription,omitempty"`
	ChargeID               string  `json:"ChargeId,omitempty"`
	ChargeName             string  `json:"ChargeName,omitempty"`
	ChargeNumber           string  `json:"ChargeNumber,omitempty"`
	ChargeType             string  `json:"ChargeType,omitempty"`
	CreatedByID            string  `json:"CreatedById,omitempty"`
	CreatedDate            string  `json:"CreatedDate,omitempty"`
	ID                     string  `json:"Id,omitempty"`
	InvoiceID              string  `json:"InvoiceId,omitempty"`
	ProcessingType         string  `json:"ProcessingType,omitempty"`
	ProductDescription     string  `json:"ProductDescription,omitempty"`
	ProductID              string  `json:"ProductId,omitempty"`
	ProductName            string  `json:"ProductName,omitempty"`
	Quantity               float64 `json:"Quantity,omitempty"`
	RatePlanChargeID       string  `json:"RatePlanChargeId,omitempty"`
	RevRecCode             string  `json:"RevRecCode,omitempty"`
	RevRecStartDate        string  `json:"RevRecStartDate,omitempty"`
	RevRecTriggerCondition string  `json:"RevRecTriggerCondition,omitempty"`
	ServiceEndDate         string  `json:"ServiceEndDate,omitempty"`
	ServiceStartDate       string  `json:"ServiceStartDate,omitempty"`
	SKU                    string  `json:"SKU,omitempty"`
	SubscriptionID         string  `json:"SubscriptionId,omitempty"`
	SubscriptionNumber     string  `json:"SubscriptionNumber,omitempty"`
	TaxAmount              float64 `json:"TaxAmount,omitempty"`
	TaxCode                string  `json:"TaxCode,omitempty"`
	TaxExemptAmount        float64 `json:"TaxExemptAmount,omitempty"`
	TaxMode                string  `json:"TaxMode,omitempty"`
	UnitPrice              float64 `json:"UnitPrice,omitempty"`
	UOM                    string  `json:"UOM,omitempty"`
	UpdatedByID            string  `json:"UpdatedById,omitempty"`
	UpdatedDate            string  `json:"UpdatedDate,omitempty"`
}

// InvoiceItemsResponse --
type InvoiceItemsResponse struct {
	InvoiceItems []InvoiceItem `json:"invoiceItems"`
	Response
}
