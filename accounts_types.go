package zuora

// Account has all the possible properties given by Zuora. This comes from
// the Describe endpoint.
type Account struct {
	ID                           *string  `json:"Id,omitempty"`
	AccountNumber                *string  `json:"AccountNumber,omitempty"`
	AdditionalEmailAddresses     *string  `json:"AdditionalEmailAddresses,omitempty"`
	AllowInvoiceEdit             *bool    `json:"AllowInvoiceEdit,omitempty"`
	AutoPay                      *bool    `json:"AutoPay,omitempty"`
	Balance                      *float64 `json:"Balance,omitempty"`
	Batch                        *string  `json:"Batch,omitempty"`
	BcdSettingOption             string   `json:"BcdSettingOption"`
	BillCycleDay                 int      `json:"BillCycleDay"`
	BillToID                     *string  `json:"BillToId,omitempty"`
	CommunicationProfileID       *string  `json:"CommunicationProfileId,omitempty"`
	CreatedByID                  *string  `json:"CreatedById,omitempty"`
	CreatedDate                  *string  `json:"CreatedDate,omitempty"`
	CreditBalance                *float64 `json:"CreditBalance,omitempty"`
	CrmID                        *string  `json:"CrmId,omitempty"`
	Currency                     string   `json:"Currency"`
	CustomerServiceRepName       *string  `json:"CustomerServiceRepName,omitempty"`
	DefaultPaymentMethodID       *string  `json:"DefaultPaymentMethodId,omitempty"`
	InvoiceDeliveryPrefsEmail    *bool    `json:"InvoiceDeliveryPrefsEmail,omitempty"`
	InvoiceDeliveryPrefsPrint    *bool    `json:"InvoiceDeliveryPrefsPrint,omitempty"`
	InvoiceTemplateID            *string  `json:"InvoiceTemplateId,omitempty"`
	LastInvoiceDate              *string  `json:"LastInvoiceDate,omitempty"`
	Mrr                          *float64 `json:"Mrr,omitempty"`
	Name                         string   `json:"Name"`
	Notes                        *string  `json:"Notes,omitempty"`
	ParentID                     *string  `json:"ParentId,omitempty"`
	PaymentGateway               *string  `json:"PaymentGateway,omitempty"`
	PaymentTerm                  *string  `json:"PaymentTerm,omitempty"`
	PurchaseOrderNumber          *string  `json:"PurchaseOrderNumber,omitempty"`
	SalesRepName                 *string  `json:"SalesRepName,omitempty"`
	SequenceSetID                *string  `json:"SequenceSetId,omitempty"`
	SoldToID                     *string  `json:"SoldToId,omitempty"`
	Status                       string   `json:"Status"`
	TaxCompanyCode               *string  `json:"TaxCompanyCode,omitempty"`
	TaxExemptCertificateID       *string  `json:"TaxExemptCertificateID,omitempty"`
	TaxExemptCertificateType     *string  `json:"TaxExemptCertificateType,omitempty"`
	TaxExemptDescription         *string  `json:"TaxExemptDescription,omitempty"`
	TaxExemptEffectiveDate       *string  `json:"TaxExemptEffectiveDate,omitempty"`
	TaxExemptEntityUseCode       *string  `json:"TaxExemptEntityUseCode,omitempty"`
	TaxExemptExpirationDate      *string  `json:"TaxExemptExpirationDate,omitempty"`
	TaxExemptIssuingJurisdiction *string  `json:"TaxExemptIssuingJurisdiction,omitempty"`
	TaxExemptStatus              *string  `json:"TaxExemptStatus,omitempty"`
	TotalDebitMemoBalance        *float64 `json:"TotalDebitMemoBalance,omitempty"`
	TotalInvoiceBalance          *float64 `json:"TotalInvoiceBalance,omitempty"`
	UnappliedBalance             float64  `json:"UnappliedBalance"`
	UnappliedCreditMemoAmount    *float64 `json:"UnappliedCreditMemoAmount,omitempty"`
	UpdatedByID                  *string  `json:"UpdatedById,omitempty"`
	UpdatedDate                  *string  `json:"UpdatedDate,omitempty"`
	VATId                        *string  `json:"VATId,omitempty"`
}

// AccountUpdateResponse CRUD responses have this response.
type AccountUpdateResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"Id"`
}
