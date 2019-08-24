package zuora

// Account has all the possible properties given by Zuora. This comes from
// the Describe endpoint.
type Account struct {
	ID            *string `json:"id,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	// AdditionalEmailAddresses show as part of GetById response as a string but
	// as an array in Get response but inside `billingAndPayment`.
	// TODO: Check with ZOQL
	// AdditionalEmailAddresses     *string  `json:"additionalEmailAddresses,omitempty"`
	AllowInvoiceEdit             *bool    `json:"allowInvoiceEdit,omitempty"`
	AutoPay                      *bool    `json:"autoPay,omitempty"`
	Balance                      *float64 `json:"balance,omitempty"`
	Batch                        *string  `json:"batch,omitempty"`
	BcdSettingOption             string   `json:"bcdSettingOption"`
	BillCycleDay                 int      `json:"billCycleDay"`
	BillToID                     *string  `json:"billToId,omitempty"`
	CommunicationProfileID       *string  `json:"communicationProfileId,omitempty"`
	CreatedByID                  *string  `json:"createdById,omitempty"`
	CreatedDate                  *string  `json:"createdDate,omitempty"`
	CreditBalance                *float64 `json:"creditBalance,omitempty"`
	CrmID                        *string  `json:"crmId,omitempty"`
	Currency                     string   `json:"currency"`
	CustomerServiceRepName       *string  `json:"customerServiceRepName,omitempty"`
	DefaultPaymentMethodID       *string  `json:"defaultPaymentMethodId,omitempty"`
	InvoiceDeliveryPrefsEmail    *bool    `json:"invoiceDeliveryPrefsEmail,omitempty"`
	InvoiceDeliveryPrefsPrint    *bool    `json:"invoiceDeliveryPrefsPrint,omitempty"`
	InvoiceTemplateID            *string  `json:"invoiceTemplateId,omitempty"`
	LastInvoiceDate              *string  `json:"lastInvoiceDate,omitempty"`
	Mrr                          *float64 `json:"mrr,omitempty"`
	Name                         string   `json:"name"`
	Notes                        *string  `json:"notes,omitempty"`
	ParentID                     *string  `json:"parentId,omitempty"`
	PaymentGateway               *string  `json:"paymentGateway,omitempty"`
	PaymentTerm                  *string  `json:"paymentTerm,omitempty"`
	PurchaseOrderNumber          *string  `json:"purchaseOrderNumber,omitempty"`
	SalesRepName                 *string  `json:"salesRepName,omitempty"`
	SequenceSetID                *string  `json:"sequenceSetId,omitempty"`
	SoldToID                     *string  `json:"soldToId,omitempty"`
	Status                       string   `json:"status"`
	TaxCompanyCode               *string  `json:"taxCompanyCode,omitempty"`
	TaxExemptCertificateID       *string  `json:"taxExemptCertificateID,omitempty"`
	TaxExemptCertificateType     *string  `json:"taxExemptCertificateType,omitempty"`
	TaxExemptDescription         *string  `json:"taxExemptDescription,omitempty"`
	TaxExemptEffectiveDate       *string  `json:"taxExemptEffectiveDate,omitempty"`
	TaxExemptEntityUseCode       *string  `json:"taxExemptEntityUseCode,omitempty"`
	TaxExemptExpirationDate      *string  `json:"taxExemptExpirationDate,omitempty"`
	TaxExemptIssuingJurisdiction *string  `json:"taxExemptIssuingJurisdiction,omitempty"`
	TaxExemptStatus              *string  `json:"taxExemptStatus,omitempty"`
	TotalDebitMemoBalance        *float64 `json:"totalDebitMemoBalance,omitempty"`
	TotalInvoiceBalance          *float64 `json:"totalInvoiceBalance,omitempty"`
	UnappliedBalance             float64  `json:"unappliedBalance"`
	UnappliedCreditMemoAmount    *float64 `json:"unappliedCreditMemoAmount,omitempty"`
	UpdatedByID                  *string  `json:"updatedById,omitempty"`
	UpdatedDate                  *string  `json:"updatedDate,omitempty"`
	VATId                        *string  `json:"vatId,omitempty"`
}

// AccountUpdate has all the possible properties given by Zuora. This comes from
// the Describe endpoint.
type AccountUpdate struct {
	AdditionalEmailAddresses  []string              `json:"additionalEmailAddresses,omitempty"`
	AutoPay                   *bool                 `json:"autoPay,omitempty"`
	Batch                     *string               `json:"batch,omitempty"`
	BillToContact             *Contact              `json:"billToContact,omitempty"`
	CommunicationProfileID    *string               `json:"communicationProfileId,omitempty"`
	CreditMemoTemplateID      *string               `json:"creditMemoTemplateId,omitempty"`
	CrmID                     *string               `json:"crmId,omitempty"`
	DebitMemoTemplateID       *string               `json:"debitMemoTemplateId,omitempty"`
	InvoiceDeliveryPrefsEmail *bool                 `json:"invoiceDeliveryPrefsEmail,omitempty"`
	InvoiceDeliveryPrefsPrint *bool                 `json:"invoiceDeliveryPrefsPrint,omitempty"`
	InvoiceTemplateID         *string               `json:"invoiceTemplateId,omitempty"`
	Name                      string                `json:"name,omitempty"`
	Notes                     *string               `json:"notes,omitempty"`
	ParentID                  *string               `json:"parentId,omitempty"`
	PaymentGateway            *string               `json:"paymentGateway,omitempty"`
	SalesRepName              *string               `json:"salesRepName,omitempty"`
	SequenceSetID             *string               `json:"sequenceSetId,omitempty"`
	SoldToContact             *Contact              `json:"soldToContact,omitempty"`
	Tagging                   *string               `json:"tagging,omitempty"`
	TaxInfo                   *AccountUpdateTaxInfo `json:"taxInfo,omitempty"`
}

// AccountUpdateTaxInfo Container for tax exempt information, used to establish the tax exempt status of a customer account.
type AccountUpdateTaxInfo struct {
	VATId                     *string `json:"vatId,omitempty"`
	CompanyCode               *string `json:"companyCode,omitempty"`
	ExemptCertificateID       *string `json:"exemptCertificateID,omitempty"`
	ExemptCertificateType     *string `json:"exemptCertificateType,omitempty"`
	ExemptDescription         *string `json:"exemptDescription,omitempty"`
	ExemptEffectiveDate       *string `json:"exemptEffectiveDate,omitempty"`
	ExemptEntityUseCode       *string `json:"exemptEntityUseCode,omitempty"`
	ExemptExpirationDate      *string `json:"exemptExpirationDate,omitempty"`
	ExemptIssuingJurisdiction *string `json:"exemptIssuingJurisdiction,omitempty"`
	ExemptStatus              *string `json:"exemptStatus,omitempty"`
}

// AccountUpdateResponse CRUD responses have this response.
type AccountUpdateResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}
