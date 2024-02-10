package models

type InvoiceStatus int

const (
	InvoiceStatusDraft  = 0
	InvoiceStatusClosed = 1
)

// InvoicesInsertRequest represents the request structure for inserting a new invoice.
type InvoicesInsertRequest struct {
	CompanyID      int              `json:"company_id"`
	Date           Time             `json:"date"`
	ExpirationDate Time             `json:"expiration_date"`
	DocumentSetID  int              `json:"document_set_id"`
	CustomerID     int              `json:"customer_id"`
	Products       []InvoiceProduct `json:"products"`
	Status         *InvoiceStatus   `json:"status,omitempty"`
	Notes          *string          `json:"notes,omitempty"`
}

// InvoiceProduct represents the product information of an invoice.
type InvoiceProduct struct {
	ProductID       int                 `json:"product_id"`
	Name            string              `json:"name"`
	Summary         *string             `json:"summary,omitempty"`
	Quantity        float64             `json:"qty"`
	Price           float64             `json:"price"`
	ExemptionReason *string             `json:"exemption_reason,omitempty"`
	Taxes           []InvoiceProductTax `json:"taxes,omitempty"`
}

// InvoiceProductTax represents the product tax information of an invoice.
type InvoiceProductTax struct {
	TaxID      int      `json:"tax_id"`
	Value      *float64 `json:"value,omitempty"`
	Order      *int     `json:"order,omitempty"`
	Cumulative *int     `json:"cumulative,omitempty"`
}

// InvoicesInsertResponse represents the response structure for the inserting a new invoice.
type InvoicesInsertResponse struct {
	Valid      int `json:"valid"`       // 1 for valid, 0 for not valid
	DocumentID int `json:"document_id"` // The ID of the invoice
}

// InvoicesGetAllRequest represents the request structure for getting all invoices.
type InvoicesGetAllRequest struct {
	CompanyID  int  `json:"company_id"`
	CustomerID *int `json:"customer_id,omitempty"`
	Qty        *int `json:"qty,omitempty"`    // Defaults to 50
	Offset     *int `json:"offset,omitempty"` // Defaults to 0
}

// InvoiceEntry represents a single invoice entry in the InvoicesGetAllResponse.
type InvoiceEntry struct {
	DocumentID     int                 `json:"document_id"`
	DocumentTypeID int                 `json:"document_type_id"`
	DocumentSetID  int                 `json:"document_set_id"`
	Number         int                 `json:"number"`
	Date           Time                `json:"date"`
	ExpirationDate Time                `json:"expiration_date"`
	EntityNumber   string              `json:"entity_number"`
	EntityName     string              `json:"entity_name"`
	EntityVAT      string              `json:"entity_vat"`
	EntityAddress  string              `json:"entity_address"`
	EntityCity     string              `json:"entity_city"`
	EntityZipCode  string              `json:"entity_zip_code"`
	GrossValue     float64             `json:"gross_value"`
	TaxesValue     float64             `json:"taxes_value"`
	NetValue       float64             `json:"net_value"`
	Status         InvoiceStatus       `json:"status"`
	DocumentType   InvoiceDocumentType `json:"document_type"`
	DocumentSet    InvoiceDocumentSet  `json:"document_set"`
}

type InvoiceDocumentType struct {
	DocumentTypeID int    `json:"document_type_id"`
	SaftType       string `json:"saft_type"`
}

type InvoiceDocumentSet struct {
	DocumentSetID int    `json:"document_set_id"`
	Name          string `json:"name"`
}

// InvoicesGetAllResponse represents the response for getting all invoices.
type InvoicesGetAllResponse []InvoiceEntry

// InvoicesUpdateRequest represents the request structure for updating a invoice.
type InvoicesUpdateRequest struct {
	CompanyID int `json:"company_id"`
	InvoiceID int `json:"invoice_id"`
}

// InvoicesUpdateResponse represents the response structure for the updating a invoice.
type InvoicesUpdateResponse struct {
	Valid     int `json:"valid"`      // 1 for valid, 0 for not valid
	InvoiceID int `json:"invoice_id"` // The ID of the invoice
}

// InvoicesDeleteRequest represents the request structure for deleting a invoice.
type InvoicesDeleteRequest struct {
	CompanyID  int `json:"company_id"`
	DocumentID int `json:"document_id"`
}

// InvoicesDeleteResponse represents the response structure for the deleting a invoice.
type InvoicesDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
