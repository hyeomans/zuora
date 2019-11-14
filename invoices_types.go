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

// InvoiceItemsResponse --
type InvoiceItemsResponse struct {
	InvoiceItems []InvoiceItem `json:"invoiceItems"`
	Response
}
