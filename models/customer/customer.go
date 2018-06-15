package customer

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	customerEnum "github.com/chargebee/chargebee-go/models/customer/enum"
)

type Customer struct {
	Id                     string                        `json:"id"`
	FirstName              string                        `json:"first_name"`
	LastName               string                        `json:"last_name"`
	Email                  string                        `json:"email"`
	Phone                  string                        `json:"phone"`
	Company                string                        `json:"company"`
	VatNumber              string                        `json:"vat_number"`
	AutoCollection         enum.AutoCollection           `json:"auto_collection"`
	NetTermDays            int32                         `json:"net_term_days"`
	AllowDirectDebit       bool                          `json:"allow_direct_debit"`
	CreatedAt              int64                         `json:"created_at"`
	CreatedFromIp          string                        `json:"created_from_ip"`
	Taxability             enum.Taxability               `json:"taxability"`
	EntityCode             enum.EntityCode               `json:"entity_code"`
	ExemptNumber           string                        `json:"exempt_number"`
	ResourceVersion        int64                         `json:"resource_version"`
	UpdatedAt              int64                         `json:"updated_at"`
	Locale                 string                        `json:"locale"`
	ConsolidatedInvoicing  bool                          `json:"consolidated_invoicing"`
	BillingDate            int32                         `json:"billing_date"`
	BillingDateMode        enum.BillingDateMode          `json:"billing_date_mode"`
	BillingDayOfWeek       customerEnum.BillingDayOfWeek `json:"billing_day_of_week"`
	BillingDayOfWeekMode   enum.BillingDayOfWeekMode     `json:"billing_day_of_week_mode"`
	CardStatus             customerEnum.CardStatus       `json:"card_status"`
	FraudFlag              customerEnum.FraudFlag        `json:"fraud_flag"`
	PrimaryPaymentSourceId string                        `json:"primary_payment_source_id"`
	BackupPaymentSourceId  string                        `json:"backup_payment_source_id"`
	BillingAddress         *BillingAddress               `json:"billing_address"`
	ReferralUrls           []*ReferralUrl                `json:"referral_urls"`
	Contacts               []*Contact                    `json:"contacts"`
	PaymentMethod          *PaymentMethod                `json:"payment_method"`
	InvoiceNotes           string                        `json:"invoice_notes"`
	PreferredCurrencyCode  string                        `json:"preferred_currency_code"`
	PromotionalCredits     int32                         `json:"promotional_credits"`
	UnbilledCharges        int32                         `json:"unbilled_charges"`
	RefundableCredits      int32                         `json:"refundable_credits"`
	ExcessPayments         int32                         `json:"excess_payments"`
	Balances               []*Balance                    `json:"balances"`
	MetaData               json.RawMessage               `json:"meta_data"`
	Deleted                bool                          `json:"deleted"`
	RegisteredForGst       bool                          `json:"registered_for_gst"`
	CustomField            map[string]interface{}        `json:"custom_field"`
	Object                 string                        `json:"object"`
}
type BillingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type ReferralUrl struct {
	ExternalCustomerId         string              `json:"external_customer_id"`
	ReferralSharingUrl         string              `json:"referral_sharing_url"`
	CreatedAt                  int64               `json:"created_at"`
	UpdatedAt                  int64               `json:"updated_at"`
	ReferralCampaignId         string              `json:"referral_campaign_id"`
	ReferralAccountId          string              `json:"referral_account_id"`
	ReferralExternalCampaignId string              `json:"referral_external_campaign_id"`
	ReferralSystem             enum.ReferralSystem `json:"referral_system"`
	Object                     string              `json:"object"`
}
type Contact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Label            string `json:"label"`
	Enabled          bool   `json:"enabled"`
	SendAccountEmail bool   `json:"send_account_email"`
	SendBillingEmail bool   `json:"send_billing_email"`
	Object           string `json:"object"`
}
type PaymentMethod struct {
	Type             enum.Type                        `json:"type"`
	Gateway          enum.Gateway                     `json:"gateway"`
	GatewayAccountId string                           `json:"gateway_account_id"`
	Status           customerEnum.PaymentMethodStatus `json:"status"`
	ReferenceId      string                           `json:"reference_id"`
	Object           string                           `json:"object"`
}
type Balance struct {
	PromotionalCredits  int32  `json:"promotional_credits"`
	ExcessPayments      int32  `json:"excess_payments"`
	RefundableCredits   int32  `json:"refundable_credits"`
	UnbilledCharges     int32  `json:"unbilled_charges"`
	CurrencyCode        string `json:"currency_code"`
	BalanceCurrencyCode string `json:"balance_currency_code"`
	Object              string `json:"object"`
}
type CreateRequestParams struct {
	Id                    string                      `json:"id,omitempty"`
	FirstName             string                      `json:"first_name,omitempty"`
	LastName              string                      `json:"last_name,omitempty"`
	Email                 string                      `json:"email,omitempty"`
	PreferredCurrencyCode string                      `json:"preferred_currency_code,omitempty"`
	Phone                 string                      `json:"phone,omitempty"`
	Company               string                      `json:"company,omitempty"`
	AutoCollection        enum.AutoCollection         `json:"auto_collection,omitempty"`
	NetTermDays           *int32                      `json:"net_term_days,omitempty"`
	AllowDirectDebit      *bool                       `json:"allow_direct_debit,omitempty"`
	VatNumber             string                      `json:"vat_number,omitempty"`
	RegisteredForGst      *bool                       `json:"registered_for_gst,omitempty"`
	Taxability            enum.Taxability             `json:"taxability,omitempty"`
	Locale                string                      `json:"locale,omitempty"`
	EntityCode            enum.EntityCode             `json:"entity_code,omitempty"`
	ExemptNumber          string                      `json:"exempt_number,omitempty"`
	MetaData              map[string]interface{}      `json:"meta_data,omitempty"`
	ConsolidatedInvoicing *bool                       `json:"consolidated_invoicing,omitempty"`
	Card                  *CreateCardParams           `json:"card,omitempty"`
	PaymentMethod         *CreatePaymentMethodParams  `json:"payment_method,omitempty"`
	BillingAddress        *CreateBillingAddressParams `json:"billing_address,omitempty"`
	CreatedFromIp         string                      `json:"created_from_ip,omitempty"`
	InvoiceNotes          string                      `json:"invoice_notes,omitempty"`
}
type CreateCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	FirstName        string       `json:"first_name,omitempty"`
	LastName         string       `json:"last_name,omitempty"`
	Number           string       `json:"number,omitempty"`
	ExpiryMonth      *int32       `json:"expiry_month,omitempty"`
	ExpiryYear       *int32       `json:"expiry_year,omitempty"`
	Cvv              string       `json:"cvv,omitempty"`
	BillingAddr1     string       `json:"billing_addr1,omitempty"`
	BillingAddr2     string       `json:"billing_addr2,omitempty"`
	BillingCity      string       `json:"billing_city,omitempty"`
	BillingStateCode string       `json:"billing_state_code,omitempty"`
	BillingState     string       `json:"billing_state,omitempty"`
	BillingZip       string       `json:"billing_zip,omitempty"`
	BillingCountry   string       `json:"billing_country,omitempty"`
	IpAddress        string       `json:"ip_address,omitempty"`
}
type CreatePaymentMethodParams struct {
	Type             enum.Type    `json:"type,omitempty"`
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	ReferenceId      string       `json:"reference_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	IssuingCountry   string       `json:"issuing_country,omitempty"`
}
type CreateBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}

type UpdateRequestParams struct {
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Email                 string                 `json:"email,omitempty"`
	PreferredCurrencyCode string                 `json:"preferred_currency_code,omitempty"`
	Phone                 string                 `json:"phone,omitempty"`
	Company               string                 `json:"company,omitempty"`
	AutoCollection        enum.AutoCollection    `json:"auto_collection,omitempty"`
	AllowDirectDebit      *bool                  `json:"allow_direct_debit,omitempty"`
	NetTermDays           *int32                 `json:"net_term_days,omitempty"`
	Taxability            enum.Taxability        `json:"taxability,omitempty"`
	Locale                string                 `json:"locale,omitempty"`
	EntityCode            enum.EntityCode        `json:"entity_code,omitempty"`
	ExemptNumber          string                 `json:"exempt_number,omitempty"`
	InvoiceNotes          string                 `json:"invoice_notes,omitempty"`
	MetaData              map[string]interface{} `json:"meta_data,omitempty"`
	FraudFlag             customerEnum.FraudFlag `json:"fraud_flag,omitempty"`
	ConsolidatedInvoicing *bool                  `json:"consolidated_invoicing,omitempty"`
}

type UpdatePaymentMethodRequestParams struct {
	PaymentMethod *UpdatePaymentMethodPaymentMethodParams `json:"payment_method,omitempty"`
}
type UpdatePaymentMethodPaymentMethodParams struct {
	Type             enum.Type    `json:"type"`
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	ReferenceId      string       `json:"reference_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	IssuingCountry   string       `json:"issuing_country,omitempty"`
}

type UpdateBillingInfoRequestParams struct {
	BillingAddress   *UpdateBillingInfoBillingAddressParams `json:"billing_address,omitempty"`
	VatNumber        string                                 `json:"vat_number,omitempty"`
	RegisteredForGst *bool                                  `json:"registered_for_gst,omitempty"`
}
type UpdateBillingInfoBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type ContactsForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type AssignPaymentRoleRequestParams struct {
	PaymentSourceId string    `json:"payment_source_id"`
	Role            enum.Role `json:"role"`
}

type AddContactRequestParams struct {
	Contact *AddContactContactParams `json:"contact,omitempty"`
}
type AddContactContactParams struct {
	Id               string `json:"id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}

type UpdateContactRequestParams struct {
	Contact *UpdateContactContactParams `json:"contact,omitempty"`
}
type UpdateContactContactParams struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}

type DeleteContactRequestParams struct {
	Contact *DeleteContactContactParams `json:"contact,omitempty"`
}
type DeleteContactContactParams struct {
	Id string `json:"id"`
}

type AddPromotionalCreditsRequestParams struct {
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}

type DeductPromotionalCreditsRequestParams struct {
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}

type SetPromotionalCreditsRequestParams struct {
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}

type RecordExcessPaymentRequestParams struct {
	Transaction *RecordExcessPaymentTransactionParams `json:"transaction,omitempty"`
	Comment     string                                `json:"comment,omitempty"`
}
type RecordExcessPaymentTransactionParams struct {
	Amount          *int32             `json:"amount"`
	CurrencyCode    string             `json:"currency_code,omitempty"`
	Date            *int64             `json:"date"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
}

type CollectPaymentRequestParams struct {
	Amount                      *int32                                   `json:"amount,omitempty"`
	InvoiceAllocations          []*CollectPaymentInvoiceAllocationParams `json:"invoice_allocations,omitempty"`
	PaymentSourceId             string                                   `json:"payment_source_id,omitempty"`
	PaymentMethod               *CollectPaymentPaymentMethodParams       `json:"payment_method,omitempty"`
	Card                        *CollectPaymentCardParams                `json:"card,omitempty"`
	ReplacePrimaryPaymentSource *bool                                    `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                                    `json:"retain_payment_source,omitempty"`
}
type CollectPaymentInvoiceAllocationParams struct {
	InvoiceId        string `json:"invoice_id,omitempty"`
	AllocationAmount *int32 `json:"allocation_amount,omitempty"`
}
type CollectPaymentPaymentMethodParams struct {
	Type             enum.Type `json:"type,omitempty"`
	GatewayAccountId string    `json:"gateway_account_id,omitempty"`
	ReferenceId      string    `json:"reference_id,omitempty"`
	TmpToken         string    `json:"tmp_token,omitempty"`
}
type CollectPaymentCardParams struct {
	GatewayAccountId string `json:"gateway_account_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Number           string `json:"number,omitempty"`
	ExpiryMonth      *int32 `json:"expiry_month,omitempty"`
	ExpiryYear       *int32 `json:"expiry_year,omitempty"`
	Cvv              string `json:"cvv,omitempty"`
	BillingAddr1     string `json:"billing_addr1,omitempty"`
	BillingAddr2     string `json:"billing_addr2,omitempty"`
	BillingCity      string `json:"billing_city,omitempty"`
	BillingStateCode string `json:"billing_state_code,omitempty"`
	BillingState     string `json:"billing_state,omitempty"`
	BillingZip       string `json:"billing_zip,omitempty"`
	BillingCountry   string `json:"billing_country,omitempty"`
}

type DeleteRequestParams struct {
	DeletePaymentMethod *bool `json:"delete_payment_method,omitempty"`
}

type MoveRequestParams struct {
	IdAtFromSite string `json:"id_at_from_site"`
	FromSite     string `json:"from_site"`
}

type ChangeBillingDateRequestParams struct {
	BillingDate          *int32                        `json:"billing_date,omitempty"`
	BillingDateMode      enum.BillingDateMode          `json:"billing_date_mode,omitempty"`
	BillingDayOfWeek     customerEnum.BillingDayOfWeek `json:"billing_day_of_week,omitempty"`
	BillingDayOfWeekMode enum.BillingDayOfWeekMode     `json:"billing_day_of_week_mode,omitempty"`
}