package zuora

import (
	"net/http"
)

// Doer is a common interface for Http clients
// https://golang.org/pkg/net/http/#Client.Do
type Doer interface {
	Do(request *http.Request) (*http.Response, error)
}

//ConfigOption helper function to modify current Config struct
type ConfigOption func(*Config)

//Config is the base configuration to return ZuoraApi
type Config struct {
	HTTPClient   *http.Client
	BaseURL      string
	ClientID     string
	ClientSecret string
	tokenStore   TokenStorer
}

//Querier One who, or that which, queries actions
type Querier interface {
	Build() string
}

//ContextKey is a helper to define Zuora context values
type ContextKey string

func (c ContextKey) String() string {
	return "zuora context key " + string(c)
}

//ContextKeyZuoraEntityIds will be added as a header on requests
const ContextKeyZuoraEntityIds = ContextKey("Zuora-Entity-Ids")

//ContextKeyZuoraTrackID will be addeed as a header on requests
const ContextKeyZuoraTrackID = ContextKey("Zuora-Track-Id")

//ContextKeyZuoraVersion will be added as header when an endpoint supports it.
const ContextKeyZuoraVersion = ContextKey("zuora-version")

//Response is a generic catch all struct to check if a response was successfull
type Response struct {
	Success bool `json:"success"`
}

// PaymentMethod Payment methods are the ways in which customers pay for
// their subscriptions. Your customers can choose a payment method
// from your company's list of preferred payment methods.
// The PaymentMethod object represents payment method
// details associated with a customer account.
// Electronic payment methods include credit cards, debit cards,
// bank transfers, and third-party processors, such as PayPal.
// Non-electronic methods, which must be made outside of Z-Payments,
// are called external methods, and include checks, cash, and wire transfers.
// You define payment methods in the web-based UI. The methods that you
// define are available for you to use for individual customer accounts in the API.
// More info at:
// https://knowledgecenter.zuora.com/DC_Developers/G_SOAP_API/E1_SOAP_API_Object_Reference/PaymentMethod
type PaymentMethod struct {
	AccountID                      *string `json:"accountId,omitempty"`
	AchAbaCode                     *string `json:"achAbaCode,omitempty"`
	AchAccountName                 *string `json:"achAccountName,omitempty"`
	AchAccountNumberMask           *string `json:"achAccountNumberMask,omitempty"`
	AchAccountType                 *string `json:"achAccountType,omitempty"`
	AchAddress1                    *string `json:"achAddress1,omitempty"`
	AchAddress2                    *string `json:"achAddress2,omitempty"`
	AchBankName                    *string `json:"achBankName,omitempty"`
	AchCity                        *string `json:"achCity,omitempty"`
	AchCountry                     *string `json:"achCountry,omitempty"`
	AchPostalCode                  *string `json:"achPostalCode,omitempty"`
	AchState                       *string `json:"achState,omitempty"`
	Active                         *bool   `json:"active,omitempty"`
	BankBranchCode                 *string `json:"bankBranchCode,omitempty"`
	BankCheckDigit                 *string `json:"bankCheckDigit,omitempty"`
	BankCity                       *string `json:"bankCity,omitempty"`
	BankCode                       *string `json:"bankCode,omitempty"`
	BankIdentificationNumber       *string `json:"bankIdentificationNumber,omitempty"`
	BankName                       *string `json:"bankName,omitempty"`
	BankPostalCode                 *string `json:"bankPostalCode,omitempty"`
	BankStreetName                 *string `json:"bankStreetName,omitempty"`
	BankStreetNumber               *string `json:"bankStreetNumber,omitempty"`
	BankTransferAccountName        *string `json:"bankTransferAccountName,omitempty"`
	BankTransferAccountNumberMask  *string `json:"bankTransferAccountNumberMask,omitempty"`
	BankTransferAccountType        *string `json:"bankTransferAccountType,omitempty"`
	BankTransferType               *string `json:"bankTransferType,omitempty"`
	BusinessIdentificationCode     *string `json:"businessIdentificationCode,omitempty"`
	City                           *string `json:"city,omitempty"`
	CompanyName                    *string `json:"companyName,omitempty"`
	Country                        *string `json:"country,omitempty"`
	CreatedByID                    *string `json:"createdById,omitempty"`
	CreatedDate                    *string `json:"createdDate,omitempty"`
	CreditCardAddress1             *string `json:"creditCardAddress1,omitempty"`
	CreditCardAddress2             *string `json:"creditCardAddress2,omitempty"`
	CreditCardCity                 *string `json:"creditCardCity,omitempty"`
	CreditCardCountry              *string `json:"creditCardCountry,omitempty"`
	CreditCardExpirationMonth      *int    `json:"creditCardExpirationMonth,omitempty"`
	CreditCardExpirationYear       *int    `json:"creditCardExpirationYear,omitempty"`
	CreditCardHolderName           *string `json:"creditCardHolderName,omitempty"`
	CreditCardMaskNumber           *string `json:"creditCardMaskNumber,omitempty"`
	CreditCardNumber               *string `json:"creditCardNumber,omitempty"`
	CreditCardPostalCode           *string `json:"creditCardPostalCode,omitempty"`
	CreditCardSecurityCode         *string `json:"creditCardSecurityCode,omitempty"`
	CreditCardState                *string `json:"creditCardState,omitempty"`
	CreditCardType                 *string `json:"creditCardType,omitempty"`
	DeviceSessionID                *string `json:"deviceSessionId,omitempty"`
	Email                          *string `json:"email,omitempty"`
	ExistingMandate                *string `json:"existingMandate,omitempty"`
	FirstName                      *string `json:"firstName,omitempty"`
	IBAN                           *string `json:"iban,omitempty"`
	ID                             *string `json:"id,omitempty"`
	IdentityNumber                 *string `json:"identityNumber,omitempty"`
	IPAddress                      *string `json:"iPAddress,omitempty"`
	IsCompany                      bool    `json:"isCompany"`
	LastFailedSaleTransactionDate  *string `json:"lastFailedSaleTransactionDate,omitempty"`
	LastName                       *string `json:"lastName,omitempty"`
	LastTransactionDateTime        *string `json:"lastTransactionDateTime,omitempty"`
	LastTransactionStatus          *string `json:"lastTransactionStatus,omitempty"`
	MandateCreationDate            *string `json:"mandateCreationDate,omitempty"`
	MandateID                      *string `json:"mandateID,omitempty"`
	MandateReceived                *string `json:"mandateReceived,omitempty"`
	MandateUpdateDate              *string `json:"mandateUpdateDate,omitempty"`
	MaxConsecutivePaymentFailures  *int    `json:"maxConsecutivePaymentFailures,omitempty"`
	MitConsentAgreementRef         *string `json:"mitConsentAgreementRef,omitempty"`
	MitConsentAgreementSrc         *string `json:"mitConsentAgreementSrc,omitempty"`
	MitNetworkTransactionID        *string `json:"mitNetworkTransactionId,omitempty"`
	MitProfileAction               *string `json:"mitProfileAction,omitempty"`
	MitProfileAgreedOn             *string `json:"mitProfileAgreedOn,omitempty"`
	MitProfileType                 *string `json:"mitProfileType,omitempty"`
	Name                           *string `json:"name,omitempty"`
	NumConsecutiveFailures         *int    `json:"numConsecutiveFailures,omitempty"`
	PaymentMethodStatus            *string `json:"paymentMethodStatus,omitempty"`
	PaymentRetryWindow             *int    `json:"paymentRetryWindow,omitempty"`
	PaypalBaid                     *string `json:"paypalBaid,omitempty"`
	PaypalEmail                    *string `json:"paypalEmail,omitempty"`
	PaypalPreapprovalKey           *string `json:"paypalPreapprovalKey,omitempty"`
	PaypalType                     *string `json:"paypalType,omitempty"`
	Phone                          *string `json:"phone,omitempty"`
	PostalCode                     *string `json:"postalCode,omitempty"`
	SecondTokenID                  *string `json:"secondTokenId,omitempty"`
	SkipValidation                 *bool   `json:"skipValidation,omitempty"`
	State                          *string `json:"state,omitempty"`
	StreetName                     *string `json:"streetName,omitempty"`
	StreetNumber                   *string `json:"streetNumber,omitempty"`
	TokenID                        *string `json:"tokenId,omitempty"`
	TotalNumberOfErrorPayments     int     `json:"totalNumberOfErrorPayments"`
	TotalNumberOfProcessedPayments int     `json:"totalNumberOfProcessedPayments"`
	Type                           string  `json:"type"`
	UpdatedByID                    *string `json:"updatedById,omitempty"`
	UpdatedDate                    *string `json:"updatedDate,omitempty"`
	UseDefaultRetryRule            bool    `json:"useDefaultRetryRule"`
}

// Contact The Contact object defines the customer who holds an account
// or who is otherwise a person to contact about an account. An Account
// object requires a contact for the BillToId and SoldToId fields before
// the account can be active. The Contact object provides the attributes
// that these Account object fields need.
type Contact struct {
	AccountID      string  `json:"accountId"`
	Address1       *string `json:"address1,omitempty"`
	Address2       *string `json:"address2,omitempty"`
	City           *string `json:"city,omitempty"`
	Country        *string `json:"country,omitempty"`
	County         *string `json:"county,omitempty"`
	CreatedByID    *string `json:"createdById,omitempty"`
	CreatedDate    *string `json:"createdDate,omitempty"`
	Description    *string `json:"description,omitempty"`
	Fax            *string `json:"fax,omitempty"`
	FirstName      string  `json:"firstName"`
	HomePhone      *string `json:"HomePhone,omitempty"`
	ID             *string `json:"id,omitempty"`
	LastName       string  `json:"lastName"`
	MobilePhone    *string `json:"mobilePhone,omitempty"`
	NickName       *string `json:"nickName,omitempty"`
	OtherPhone     *string `json:"otherPhone,omitempty"`
	OtherPhoneType *string `json:"otherPhoneType,omitempty"`
	PersonalEmail  *string `json:"personalEmail,omitempty"`
	PostalCode     *string `json:"postalCode,omitempty"`
	State          *string `json:"state,omitempty"`
	TaxRegion      *string `json:"taxRegion,omitempty"`
	UpdatedByID    *string `json:"updatedById,omitempty"`
	UpdatedDate    *string `json:"updatedDate,omitempty"`
	WorkEmail      *string `json:"workEmail,omitempty"`
	WorkPhone      *string `json:"workPhone,omitempty"`
}

// Invoice An invoice  represents a bill to a customer.
// You generate invoices from a bill run, then email
// them as PDFs to your customers in batches or individually,
// or print them and send them to customers through postal
// mail. You can also generate a preview of an invoice
// before you activate a new subscription or amendment.
// The Invoice object provides information about customers'
// accounts for invoices, including dates, the status, and
// amounts. It is created at the account level, and can include
// all of the charges for multiple subscriptions for an account.
type Invoice struct {
	AccountID                     string   `json:"accountId"`
	AdjustmentAmount              float64  `json:"adjustmentAmount"`
	Amount                        *float64 `json:"amount,omitempty"`
	AmountWithoutTax              *float64 `json:"amountWithoutTax,omitempty"`
	AutoPay                       *bool    `json:"autoPay,omitempty"`
	Balance                       *float64 `json:"balance,omitempty"`
	BillRunID                     *string  `json:"billRunId,omitempty"`
	BillToContactSnapshotID       *string  `json:"billToContactSnapshotId,omitempty"`
	Body                          string   `json:"body"`
	Comments                      *string  `json:"comments,omitempty"`
	CreatedByID                   *string  `json:"createdById,omitempty"`
	CreatedDate                   *string  `json:"createdDate,omitempty"`
	CreditBalanceAdjustmentAmount float64  `json:"creditBalanceAdjustmentAmount"`
	DueDate                       *string  `json:"dueDate,omitempty"`
	ID                            *string  `json:"id,omitempty"`
	IncludesOneTime               *bool    `json:"includesOneTime,omitempty"`
	IncludesRecurring             *bool    `json:"includesRecurring,omitempty"`
	IncludesUsage                 *bool    `json:"includesUsage,omitempty"`
	InvoiceDate                   *string  `json:"invoiceDate,omitempty"`
	InvoiceNumber                 *string  `json:"invoiceNumber,omitempty"`
	LastEmailSentDate             *string  `json:"lastEmailSentDate,omitempty"`
	PaymentAmount                 float64  `json:"paymentAmount"`
	PostedBy                      *string  `json:"postedBy,omitempty"`
	PostedDate                    *string  `json:"postedDate,omitempty"`
	RefundAmount                  float64  `json:"refundAmount"`
	RegenerateInvoicePDF          *bool    `json:"regenerateInvoicePDF,omitempty"`
	Reversed                      *bool    `json:"reversed,omitempty"`
	SoldToContactSnapshotID       *string  `json:"soldToContactSnapshotId,omitempty"`
	Source                        *string  `json:"source,omitempty"`
	SourceID                      *string  `json:"sourceId,omitempty"`
	Status                        *string  `json:"status,omitempty"`
	TargetDate                    *string  `json:"targetDate,omitempty"`
	TaxAmount                     float64  `json:"taxAmount"`
	TaxExemptAmount               float64  `json:"taxExemptAmount"`
	TaxMessage                    *string  `json:"taxMessage,omitempty"`
	TaxStatus                     *string  `json:"taxStatus,omitempty"`
	TransferredToAccounting       *string  `json:"transferredToAccounting,omitempty"`
	UpdatedByID                   *string  `json:"updatedById,omitempty"`
	UpdatedDate                   *string  `json:"updatedDate,omitempty"`
}

// Usage Usage is the amount of resources a customer uses.
// You track the usage of metered resources, then charge based
// on the amount that your customers consume.
// Usage is always billed in arrears. For example, you bill
// in November for usage consumed in October. You can bill usage
// on a recurring, monthly, quarterly, semi-annual, and annual basis.
// Use the Usage object to import the quantity of units that
// customers use of a product, such as the number of page loads on a wiki.
type Usage struct {
	AccountID          *string  `json:"accountId,omitempty"`
	AccountNumber      *string  `json:"accountNumber,omitempty"`
	AncestorAccountID  *string  `json:"ancestorAccountId,omitempty"`
	ChargeID           *string  `json:"chargeId,omitempty"`
	ChargeNumber       *string  `json:"chargeNumber,omitempty"`
	CreatedByID        *string  `json:"createdById,omitempty"`
	CreatedDate        *string  `json:"createdDate,omitempty"`
	Description        *string  `json:"description,omitempty"`
	EndDateTime        *string  `json:"endDateTime,omitempty"`
	ID                 *string  `json:"id,omitempty"`
	ImportID           *string  `json:"importId,omitempty"`
	InvoiceID          *string  `json:"invoiceId,omitempty"`
	InvoiceNumber      *string  `json:"invoiceNumber,omitempty"`
	Quantity           *float64 `json:"quantity,omitempty"`
	RbeStatus          *string  `json:"rbeStatus,omitempty"`
	SourceName         *string  `json:"sourceName,omitempty"`
	SourceType         *string  `json:"sourceType,omitempty"`
	StartDateTime      *string  `json:"startDateTime,omitempty"`
	SubmissionDateTime *string  `json:"submissionDateTime,omitempty"`
	SubscriptionID     *string  `json:"subscriptionId,omitempty"`
	SubscriptionNumber *string  `json:"subscriptionNumber,omitempty"`
	UOM                *string  `json:"uom,omitempty"`
	UpdatedByID        *string  `json:"updatedById,omitempty"`
	UpdatedDate        *string  `json:"updatedDate,omitempty"`
}

// Payment A payment is the money that customers send to pay for invoices related to their subscriptions.
// The Payment object holds all of the information about an individual payment,
// including the payment amount and to which invoices the payment was applied to.
// More info at:
// https://knowledgecenter.zuora.com/DC_Developers/G_SOAP_API/E1_SOAP_API_Object_Reference/Payment
type Payment struct {
	// Something something
	AccountID                  *string  `json:"accountId,omitempty"`
	AccountingCode             *string  `json:"accountingCode,omitempty"`
	Amount                     float64  `json:"amount"`
	AppliedAmount              float64  `json:"appliedAmount"`
	AppliedCreditBalanceAmount float64  `json:"appliedCreditBalanceAmount"`
	AppliedInvoiceAmount       *float64 `json:"appliedInvoiceAmount,omitempty"`
	AuthTransactionID          *string  `json:"authTransactionId,omitempty"`
	BankIdentificationNumber   *string  `json:"bankIdentificationNumber,omitempty"`
	CancelledOn                *string  `json:"cancelledOn,omitempty"`
	Comment                    *string  `json:"comment,omitempty"`
	CreatedByID                *string  `json:"createdById,omitempty"`
	CreatedDate                *string  `json:"createdDate,omitempty"`
	Currency                   *string  `json:"currency,omitempty"`
	EffectiveDate              string   `json:"effectiveDate"`
	Gateway                    *string  `json:"gateway,omitempty"`
	// GatewayOrderID A merchant-specified natural key value that can be passed
	// to the payment gateway when an electronic payment is created.
	// A gateway is an online service provider that connects
	// an online shopping cart to a payment processor. Gateways
	// check duplicates on the gateway order ID to ensure that
	// the merchant didn't accidentally enter the same transaction
	// twice. This ID can also be used to do reconciliation and tie
	// the payment to a natural key in external systems. For example,
	// a shopping cart order ID.
	// The source of this ID varies by merchant. Some merchants
	// use their shopping cart order IDs, and others use something different.
	// Merchants use this ID to track transactions in their eCommerce systems.
	// Gateways usually do a uniqueness check on this value to prevent
	// multiple submissions of the same transaction, such as when a customer
	// clicks the Pay button twice during a single checkout process, inadvertently
	// sending two identical orders to Zuora and the gateway.  A uniqueness check prevents DuplicateOrderID exceptions.
	// If not provided, Zuora will default this value to the PaymentNumber.
	GatewayOrderID           *string  `json:"gatewayOrderId,omitempty"`
	GatewayResponse          string   `json:"gatewayResponse"`
	GatewayResponseCode      string   `json:"gatewayResponseCode"`
	GatewayState             string   `json:"gatewayState"`
	ID                       *string  `json:"id,omitempty"`
	InvoiceID                *string  `json:"invoiceId,omitempty"`
	InvoiceNumber            *string  `json:"invoiceNumber,omitempty"`
	MarkedForSubmissionOn    *string  `json:"markedForSubmissionOn,omitempty"`
	PaymentMethodID          *string  `json:"paymentMethodId,omitempty"`
	PaymentMethodSnapshotID  *string  `json:"paymentMethodSnapshotId,omitempty"`
	PaymentNumber            string   `json:"paymentNumber"`
	ReferencedPaymentID      *string  `json:"referencedPaymentID,omitempty"`
	ReferenceID              *string  `json:"referenceId,omitempty"`
	RefundAmount             *float64 `json:"refundAmount,omitempty"`
	SecondPaymentReferenceID *string  `json:"secondPaymentReferenceId,omitempty"`
	SettledOn                *string  `json:"settledOn,omitempty"`
	SoftDescriptor           *string  `json:"softDescriptor,omitempty"`
	SoftDescriptorPhone      *string  `json:"softDescriptorPhone,omitempty"`
	Source                   *string  `json:"source,omitempty"`
	SourceName               *string  `json:"sourceName,omitempty"`
	Status                   string   `json:"status"`
	SubmittedOn              *string  `json:"submittedOn,omitempty"`
	TransferredToAccounting  *string  `json:"transferredToAccounting,omitempty"`
	Type                     string   `json:"type"`
	UnappliedAmount          float64  `json:"unappliedAmount"`
	UpdatedByID              *string  `json:"updatedById,omitempty"`
	UpdatedDate              *string  `json:"updatedDate,omitempty"`
}
