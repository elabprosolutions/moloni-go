package models

// DocumentSetsTemplate represents the detailed information of a template associated with a document set.
type DocumentSetsTemplate struct {
	TemplateID         int    `json:"template_id"`
	Name               string `json:"name"`
	BusinessName       string `json:"business_name"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	City               string `json:"city"`
	ZipCode            string `json:"zip_code"`
	CountryID          int    `json:"country_id"`
	Phone              string `json:"phone"`
	Fax                string `json:"fax"`
	Website            string `json:"website"`
	Notes              string `json:"notes"`
	DocumentsFootnote  string `json:"documents_footnote"`
	EmailSenderName    string `json:"email_sender_name"`
	EmailSenderAddress string `json:"email_sender_address"`
	Image              string `json:"image"`
}

// DocumentSetsInsertRequest represents the request structure for inserting a new document set.
type DocumentSetsInsertRequest struct {
	CompanyID              int    `json:"company_id"`                          // Mandatory
	Name                   string `json:"name"`                                // Mandatory
	CashVATSchemeIndicator *int   `json:"cash_vat_scheme_indicator,omitempty"` // Optional
	ActiveByDefault        *int   `json:"active_by_default,omitempty"`         // Optional
	TemplateID             *int   `json:"template_id,omitempty"`               // Optional
}

// DocumentSetsInsertResponse represents the response structure for the inserting a new document set.
type DocumentSetsInsertResponse struct {
	Valid         int    `json:"valid"`           // 1 for valid, 0 for not valid
	DocumentSetID string `json:"document_set_id"` // The ID of the document set
}

// DocumentSetsGetAllRequest represents the request structure for getting all document sets.
type DocumentSetsGetAllRequest struct {
	CompanyID int `json:"company_id"` // Mandatory
}

// DocumentSetEntry represents a single document set entry in the DocumentSetGetAllResponse.
type DocumentSetEntry struct {
	DocumentSetID          int                   `json:"document_set_id"`                     // Mandatory
	Name                   string                `json:"name"`                                // Mandatory
	CashVATSchemeIndicator *int                  `json:"cash_vat_scheme_indicator,omitempty"` // Optional
	ActiveByDefault        *int                  `json:"active_by_default,omitempty"`         // Optional
	TemplateID             *int                  `json:"template_id,omitempty"`               // Optional
	ImgGr1                 *string               `json:"img_gr_1,omitempty"`                  // Optional
	Template               *DocumentSetsTemplate `json:"template,omitempty"`                  // Optional
}

// DocumentSetsGetAllResponse represents the response for getting all document sets.
type DocumentSetsGetAllResponse []DocumentSetEntry

// DocumentSetsUpdateRequest represents the request structure for updating a document set.
type DocumentSetsUpdateRequest struct {
	CompanyID              int    `json:"company_id"`                // Mandatory
	DocumentSetID          int    `json:"document_set_id"`           // Mandatory
	Name                   string `json:"name"`                      // Mandatory
	CashVATSchemeIndicator *int   `json:"cash_vat_scheme_indicator"` // Optional
	ActiveByDefault        *int   `json:"active_by_default"`         // Optional
	TemplateID             *int   `json:"template_id"`               // Optional
}

// DocumentSetsUpdateResponse represents the response structure for the updating a document set.
type DocumentSetsUpdateResponse struct {
	Valid         int `json:"valid"`           // 1 for valid, 0 for not valid
	DocumentSetID int `json:"document_set_id"` // The ID of the document set
}

// DocumentSetsDeleteRequest represents the request structure for deleting a document set.
type DocumentSetsDeleteRequest struct {
	CompanyID     int `json:"company_id"`      // Mandatory
	DocumentSetID int `json:"document_set_id"` // Mandatory
}

// DocumentSetsDeleteResponse represents the response structure for the deleting a document set.
type DocumentSetsDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
