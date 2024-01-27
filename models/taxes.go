package models

// TaxType represents the type of tax.
type TaxType int

const (
	TaxTypePercentage TaxType = 1
	TaxTypeFixedValue TaxType = 2
	TaxTypeVariable   TaxType = 3
)

// SaftType represents the SAFT type of tax.
type SaftType int

const (
	SaftTypeValueAdded     SaftType = 1
	SaftTypeDirectTax      SaftType = 2
	SaftTypeNoneOfTheAbove SaftType = 3
)

// VATType represents the VAT type.
type VATType string

const (
	VATTypeReduced      VATType = "RED"
	VATTypeIntermediate VATType = "INT"
	VATTypeNormal       VATType = "NOR"
	VATTypeExempt       VATType = "ISE"
	VATTypeOther        VATType = "OUT"
)

// TaxesInsertRequest represents the request structure for inserting a new tax.
type TaxesInsertRequest struct {
	CompanyID       int      `json:"company_id"`                 // Mandatory
	Name            string   `json:"name"`                       // Mandatory
	Value           int      `json:"value"`                      // Mandatory
	Type            TaxType  `json:"type"`                       // Mandatory
	SaftType        SaftType `json:"saft_type"`                  // Mandatory
	VATType         *VATType `json:"vat_type,omitempty"`         // Mandatory if SaftType is SaftTypeValueAdded
	StampTax        *string  `json:"stamp_tax,omitempty"`        // Mandatory if SaftType is SaftTypeDirectTax
	ExemptionReason *string  `json:"exemption_reason,omitempty"` // Mandatory if tax value is 0
	FiscalZone      string   `json:"fiscal_zone"`                // Mandatory
	ActiveByDefault int      `json:"active_by_default"`          // Mandatory
}

// TaxesInsertResponse represents the response structure for the inserting a new tax.
type TaxesInsertResponse struct {
	Valid int `json:"valid"`  // 1 for valid, 0 for not valid
	TaxID int `json:"tax_id"` // The ID of the tax
}

// TaxesGetAllRequest represents the request structure for getting all taxes.
type TaxesGetAllRequest struct {
	CompanyID       int      `json:"company_id"`                  // Mandatory
	CountryID       *int     `json:"country_id,omitempty"`        // Optional
	FiscalZone      *string  `json:"fiscal_zone,omitempty"`       // Optional
	Value           *float64 `json:"value,omitempty"`             // Optional
	Type            *int     `json:"type,omitempty"`              // Optional
	ActiveByDefault *int     `json:"active_by_default,omitempty"` // Optional
	WithInvisible   *int     `json:"with_invisible,omitempty"`    // Optional
}

// TaxEntry represents a single tax entry in the TaxesGetAllResponse.
type TaxEntry struct {
	TaxID           int      `json:"tax_id"`
	Name            string   `json:"name"`
	Value           float64  `json:"value"`
	Type            TaxType  `json:"type"`
	SaftType        SaftType `json:"saft_type"`
	VATType         VATType  `json:"vat_type"`
	StampTax        string   `json:"stamp_tax"`
	ExemptionReason string   `json:"exemption_reason"`
	FiscalZone      string   `json:"fiscal_zone"`
	ActiveByDefault int      `json:"active_by_default"`
}

// TaxesGetAllResponse represents the response for getting all taxes.
type TaxesGetAllResponse []TaxEntry

// TaxesUpdateRequest represents the request structure for updating a tax.
type TaxesUpdateRequest struct {
	CompanyID       int      `json:"company_id"`        // Mandatory
	TaxID           int      `json:"tax_id"`            // Mandatory
	Value           int      `json:"value"`             // Mandatory
	Type            TaxType  `json:"type"`              // Mandatory
	SaftType        SaftType `json:"saft_type"`         // Mandatory
	VATType         VATType  `json:"vat_type"`          // Mandatory
	StampTax        string   `json:"stamp_tax"`         // Mandatory
	ExemptionReason string   `json:"exemption_reason"`  // Mandatory
	FiscalZone      string   `json:"fiscal_zone"`       // Mandatory
	ActiveByDefault int      `json:"active_by_default"` // Mandatory
}

// TaxesUpdateResponse represents the response structure for the updating a tax.
type TaxesUpdateResponse struct {
	Valid int `json:"valid"`  // 1 for valid, 0 for not valid
	TaxID int `json:"tax_id"` // The ID of the tax
}

// TaxesDeleteRequest represents the request structure for deleting a tax.
type TaxesDeleteRequest struct {
	CompanyID int `json:"company_id"` // Mandatory
	TaxID     int `json:"tax_id"`     // Mandatory
}

// TaxesDeleteResponse represents the response structure for the deleting a tax.
type TaxesDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
