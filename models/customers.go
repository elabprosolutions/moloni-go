package models

// CustomersInsertRequest represents the request structure for inserting a new customer.
type CustomersInsertRequest struct {
	CompanyID        int            `json:"company_id"`
	VAT              string         `json:"vat"`
	Number           string         `json:"number"`
	Name             string         `json:"name"`
	LanguageID       int            `json:"language_id"`
	Address          string         `json:"address"`
	ZipCode          *string        `json:"zip_code,omitempty"`
	City             string         `json:"city"`
	CountryID        int            `json:"country_id"`
	Email            *string        `json:"email,omitempty"`
	Website          *string        `json:"website,omitempty"`
	Phone            *string        `json:"phone,omitempty"`
	Fax              *string        `json:"fax,omitempty"`
	ContactName      *string        `json:"contact_name,omitempty"`
	ContactEmail     *string        `json:"contact_email,omitempty"`
	ContactPhone     *string        `json:"contact_phone,omitempty"`
	Notes            *string        `json:"notes,omitempty"`
	SalesmanID       *int           `json:"salesman_id,omitempty"`
	PriceClassID     *int           `json:"price_class_id,omitempty"`
	MaturityDateID   int            `json:"maturity_date_id"`
	PaymentDay       *int           `json:"payment_day,omitempty"`
	Discount         *float64       `json:"discount,omitempty"`
	CreditLimit      *float64       `json:"credit_limit,omitempty"`
	Copies           []CustomerCopy `json:"copies,omitempty"`
	PaymentMethodID  int            `json:"payment_method_id"`
	DeliveryMethodID int            `json:"delivery_method_id"`
	FieldNotes       *string        `json:"field_notes,omitempty"`
}

type CustomerCopy struct {
	DocumentTypeID int `json:"document_type_id"`
	Copies         int `json:"copies"`
}

// CustomersInsertResponse represents the response structure for the inserting a new customer.
type CustomersInsertResponse struct {
	Valid      int `json:"valid"`       // 1 for valid, 0 for not valid
	CustomerID int `json:"customer_id"` // The ID of the customer
}

// CustomersGetAllRequest represents the request structure for getting all customers.
type CustomersGetAllRequest struct {
	CompanyID int `json:"company_id"`
}

// CustomerEntry represents a single customer entry in the CustomersGetAllResponse.
type CustomerEntry struct {
	CustomerID         int                        `json:"customer_id"`
	Number             string                     `json:"number"`
	Name               string                     `json:"name"`
	VAT                string                     `json:"vat"`
	Address            string                     `json:"address"`
	City               string                     `json:"city"`
	ZipCode            *string                    `json:"zip_code,omitempty"`
	CountryID          int                        `json:"country_id"`
	Email              *string                    `json:"email,omitempty"`
	Website            *string                    `json:"website,omitempty"`
	Phone              *string                    `json:"phone,omitempty"`
	Fax                *string                    `json:"fax,omitempty"`
	ContactName        *string                    `json:"contact_name,omitempty"`
	ContactEmail       *string                    `json:"contact_email,omitempty"`
	ContactPhone       *string                    `json:"contact_phone,omitempty"`
	Notes              *string                    `json:"notes,omitempty"`
	SalesmanID         *int                       `json:"salesman_id,omitempty"`
	Discount           *float64                   `json:"discount,omitempty"`
	CreditLimit        *float64                   `json:"credit_limit,omitempty"`
	MaturityDateID     int                        `json:"maturity_date_id"`
	PaymentDay         *int                       `json:"payment_day,omitempty"`
	FieldNotes         *string                    `json:"field_notes,omitempty"`
	LanguageID         int                        `json:"language_id"`
	PaymentMethodID    int                        `json:"payment_method_id"`
	DeliveryMethodID   int                        `json:"delivery_method_id"`
	Country            CustomerCountry            `json:"country"`
	Language           CustomerLanguage           `json:"language"`
	MaturityDate       CustomerMaturityDate       `json:"maturity_date"`
	PaymentMethod      CustomerPaymentMethod      `json:"payment_method"`
	DeliveryMethod     CustomerDeliveryMethod     `json:"delivery_method"`
	Salesman           CustomerSalesman           `json:"salesman"`
	AlternateAddresses []CustomerAlternateAddress `json:"alternate_addresses"`
	Copies             []CustomerCopy             `json:"copies"`
	AssociatedTaxes    []CustomerTax              `json:"associated_taxes"`
	PriceClass         CustomerPriceClass         `json:"price_class"`
}

// CustomerCountry represents the country-specific information for a customer.
type CustomerCountry struct {
	CountryID int    `json:"country_id"`
	Country   string `json:"country"`
	ISO31661  string `json:"iso_3166_1"`
}

// CustomerLanguage represents the language preferences of a customer.
type CustomerLanguage struct {
	LanguageID int    `json:"language_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

// CustomerMaturityDate holds information related to payment maturity dates for a customer.
type CustomerMaturityDate struct {
	MaturityDateID     int     `json:"maturity_date_id"`
	Name               string  `json:"name"`
	Days               int     `json:"days"`
	AssociatedDiscount float64 `json:"associated_discount"`
}

// CustomerPaymentMethod details the payment methods used by a customer.
type CustomerPaymentMethod struct {
	PaymentMethodID int    `json:"payment_method_id"`
	Name            string `json:"name"`
}

// CustomerDeliveryMethod provides information about the delivery methods preferred by a customer.
type CustomerDeliveryMethod struct {
	DeliveryMethodID int    `json:"delivery_method_id"`
	Name             string `json:"name"`
}

// CustomerSalesman contains details about the salesman associated with a customer.
type CustomerSalesman struct {
	SalesmanID     int     `json:"salesman_id"`
	Number         string  `json:"number"`
	Name           string  `json:"name"`
	BaseCommission float64 `json:"base_commission"`
}

// CustomerAlternateAddress represents alternate addresses associated with a customer.
type CustomerAlternateAddress struct {
	AddressID   int    `json:"address_id"`
	Designation string `json:"designation"`
	Code        string `json:"code"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	CountryID   int    `json:"country_id"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	ContactName string `json:"contact_name"`
}

// CustomerTax details the tax information associated with a customer.
type CustomerTax struct {
	TaxID int `json:"tax_id"`
}

// CustomerPriceClass represents the pricing classification for a customer.
type CustomerPriceClass struct {
	PriceClassID int    `json:"price_class_id"`
	Title        string `json:"title"`
}

// CustomersGetAllResponse represents the response for getting all customers.
type CustomersGetAllResponse []CustomerEntry

// CustomersUpdateRequest represents the request structure for updating a customer.
type CustomersUpdateRequest struct {
	CompanyID        int            `json:"company_id"`
	CustomerID       int            `json:"customer_id"`
	VAT              string         `json:"vat"`
	Number           string         `json:"number"`
	Name             string         `json:"name"`
	LanguageID       int            `json:"language_id"`
	Address          string         `json:"address"`
	ZipCode          *string        `json:"zip_code,omitempty"`
	City             string         `json:"city"`
	CountryID        int            `json:"country_id"`
	Email            *string        `json:"email,omitempty"`
	Website          *string        `json:"website,omitempty"`
	Phone            *string        `json:"phone,omitempty"`
	Fax              *string        `json:"fax,omitempty"`
	ContactName      *string        `json:"contact_name,omitempty"`
	ContactEmail     *string        `json:"contact_email,omitempty"`
	ContactPhone     *string        `json:"contact_phone,omitempty"`
	Notes            *string        `json:"notes,omitempty"`
	SalesmanID       *int           `json:"salesman_id,omitempty"`
	PriceClassID     *int           `json:"price_class_id,omitempty"`
	MaturityDateID   int            `json:"maturity_date_id"`
	PaymentDay       *int           `json:"payment_day,omitempty"`
	Discount         *float64       `json:"discount,omitempty"`
	CreditLimit      *float64       `json:"credit_limit,omitempty"`
	Copies           []CustomerCopy `json:"copies,omitempty"`
	PaymentMethodID  int            `json:"payment_method_id"`
	DeliveryMethodID int            `json:"delivery_method_id"`
	FieldNotes       *string        `json:"field_notes,omitempty"`
}

// CustomersUpdateResponse represents the response structure for the updating a customer.
type CustomersUpdateResponse struct {
	Valid      int `json:"valid"`       // 1 for valid, 0 for not valid
	CustomerID int `json:"customer_id"` // The ID of the customer
}

// CustomersDeleteRequest represents the request structure for deleting a customer.
type CustomersDeleteRequest struct {
	CompanyID  int `json:"company_id"`
	CustomerID int `json:"customer_id"`
}

// CustomersDeleteResponse represents the response structure for the deleting a customer.
type CustomersDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
