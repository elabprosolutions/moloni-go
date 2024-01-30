package models

// ProductType represents the type of product.
type ProductType int

// Enum values for ProductType
const (
	ProductTypeProduct ProductType = 1
	ProductTypeService ProductType = 2
	ProductTypeOthers  ProductType = 3
)

// ATProductCategory represents the category of a product when it has stock.
type ATProductCategory string

// Enum values for ATProductCategory
const (
	ATProductCategoryMerchandise           ATProductCategory = "M"
	ATProductCategoryRawMaterials          ATProductCategory = "P"
	ATProductCategoryFinishedGoods         ATProductCategory = "A"
	ATProductCategorySubProductsWaste      ATProductCategory = "S"
	ATProductCategoryProductsInDevelopment ATProductCategory = "T"
)

// ProductsInsertRequest represents the request structure for inserting a new product.
type ProductsInsertRequest struct {
	CompanyID         int                `json:"company_id"`
	CategoryID        int                `json:"category_id"`
	Type              ProductType        `json:"type"`
	Name              string             `json:"name"`
	Summary           *string            `json:"summary,omitempty"`
	Reference         string             `json:"reference"`
	EAN               *string            `json:"ean,omitempty"`
	Price             float64            `json:"price"`
	UnitID            int                `json:"unit_id"`
	HasStock          int                `json:"has_stock"`
	Stock             *float64           `json:"stock,omitempty"`
	MinimumStock      *float64           `json:"minimum_stock,omitempty"`
	PosFavorite       *int               `json:"pos_favorite,omitempty"`
	ATProductCategory *ATProductCategory `json:"at_product_category,omitempty"`
	ExemptionReason   *string            `json:"exemption_reason,omitempty"`
	Taxes             []ProductTax       `json:"taxes,omitempty"`
	Suppliers         []ProductSupplier  `json:"suppliers,omitempty"`
	Properties        []ProductProperty  `json:"properties,omitempty"`
	Warehouses        []ProductWarehouse `json:"warehouses,omitempty"`
}

// ProductTax represents a tax item associated with a product.
type ProductTax struct {
	TaxID      int     `json:"tax_id"`
	Value      float64 `json:"value"`
	Order      int     `json:"order"`
	Cumulative int     `json:"cumulative"`
}

// ProductSupplier represents a supplier item for a product.
type ProductSupplier struct {
	SupplierID   int     `json:"supplier_id"`
	CostPrice    float64 `json:"cost_price"`
	Referenceint *string `json:"referenceint,omitempty"`
}

// ProductProperty represents a property of a product.
type ProductProperty struct {
	PropertyID int    `json:"property_id"`
	Value      string `json:"value"`
}

// ProductWarehouse represents warehouse information for a product.
type ProductWarehouse struct {
	WarehouseID int     `json:"warehouse_id"`
	Stock       float64 `json:"stock"`
}

// ProductsInsertResponse represents the response structure for the inserting a new product.
type ProductsInsertResponse struct {
	Valid     int `json:"valid"`      // 1 for valid, 0 for not valid
	ProductID int `json:"product_id"` // The ID of the product
}

// ProductsGetAllRequest represents the request structure for getting all products.
type ProductsGetAllRequest struct {
	CompanyID  int  `json:"company_id"`
	CategoryID *int `json:"category_id,omitempty"`
	Qty        *int `json:"qty,omitempty"`    // Defaults to 50
	Offset     *int `json:"offset,omitempty"` // Defaults to 0
}

// ProductEntry represents a single product entry in the ProductsGetAllResponse.
type ProductEntry struct {
	ProductID         int                      `json:"product_id"`
	Type              ProductType              `json:"type"`
	Name              string                   `json:"name"`
	Reference         string                   `json:"reference"`
	EAN               *string                  `json:"ean,omitempty"`
	Price             float64                  `json:"price"`
	Stock             *float64                 `json:"stock,omitempty"`
	PosFavorite       *int                     `json:"pos_favorite,omitempty"`
	ATProductCategory *ATProductCategory       `json:"at_product_category,omitempty"`
	Image             *string                  `json:"image,omitempty"`
	UnitID            int                      `json:"unit_id"`
	MeasurementUnit   MeasurementUnit          `json:"measurement_unit"`
	Taxes             []ProductEntryTax        `json:"taxes,omitempty"`
	Properties        []ProductEntryProperty   `json:"properties,omitempty"`
	PriceClasses      []ProductEntryPriceClass `json:"price_classes,omitempty"`
	Warehouses        []ProductEntryWarehouse  `json:"warehouses,omitempty"`
}

// MeasurementUnit defines the unit of measurement for a product.
type MeasurementUnit struct {
	UnitID    int     `json:"unit_id"`
	Name      *string `json:"name,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
}

// ProductEntryTax represents the tax information associated with a product.
type ProductEntryTax struct {
	ProductID  int       `json:"product_id"`
	TaxID      int       `json:"tax_id"`
	Value      float64   `json:"value"`
	Order      int       `json:"order"`
	Cumulative int       `json:"cumulative"`
	Tax        *TaxEntry `json:"tax,omitempty"`
}

// ProductEntryProperty represents a property of a product.
type ProductEntryProperty struct {
	PropertyID int           `json:"property_id"`
	Value      string        `json:"value"`
	Property   PropertyEntry `json:"property"`
}

// PropertyEntry provides details about a specific property.
type PropertyEntry struct {
	PropertyID int    `json:"property_id"`
	Title      string `json:"title"`
}

// ProductEntryPriceClass represents the price class information of a product.
type ProductEntryPriceClass struct {
	ProductPriceClassID int             `json:"product_price_class_id"`
	PriceClassID        int             `json:"price_class_id"`
	Value               float64         `json:"value"`
	PriceClass          PriceClassEntry `json:"price_class"`
}

// PriceClassEntry provides details about a specific price class.
type PriceClassEntry struct {
	PriceClassID   int    `json:"price_class_id"`
	Title          string `json:"title"`
	LastModifiedBy int    `json:"lastmodifiedby"`
	LastModified   int    `json:"lastmodified"`
}

// ProductEntryWarehouse represents the warehouse information for a product.
type ProductEntryWarehouse struct {
	ProductID   int     `json:"product_id"`
	WarehouseID int     `json:"warehouse_id"`
	Stock       float64 `json:"stock"`
}

// ProductsGetAllResponse represents the response for getting all products.
type ProductsGetAllResponse []ProductEntry

// ProductsUpdateRequest represents the request structure for updating a product.
type ProductsUpdateRequest struct {
	CompanyID         int                `json:"company_id"`
	ProductID         int                `json:"product_id"`
	CategoryID        int                `json:"category_id"`
	Type              ProductType        `json:"type"`
	Name              string             `json:"name"`
	Summary           *string            `json:"summary,omitempty"`
	Reference         string             `json:"reference"`
	EAN               *string            `json:"ean,omitempty"`
	Price             float64            `json:"price"`
	UnitID            int                `json:"unit_id"`
	HasStock          int                `json:"has_stock"`
	Stock             *float64           `json:"stock,omitempty"`
	MinimumStock      *float64           `json:"minimum_stock,omitempty"`
	PosFavorite       *int               `json:"pos_favorite,omitempty"`
	ATProductCategory *ATProductCategory `json:"at_product_category,omitempty"`
	ExemptionReason   *string            `json:"exemption_reason,omitempty"`
	Taxes             []ProductTax       `json:"taxes,omitempty"`
	Suppliers         []ProductSupplier  `json:"suppliers,omitempty"`
	Properties        []ProductProperty  `json:"properties,omitempty"`
	Warehouses        []ProductWarehouse `json:"warehouses,omitempty"`
}

// ProductsUpdateResponse represents the response structure for the updating a product.
type ProductsUpdateResponse struct {
	Valid     int `json:"valid"`      // 1 for valid, 0 for not valid
	ProductID int `json:"product_id"` // The ID of the product
}

// ProductsDeleteRequest represents the request structure for deleting a product.
type ProductsDeleteRequest struct {
	CompanyID int `json:"company_id"`
	ProductID int `json:"product_id"`
}

// ProductsDeleteResponse represents the response structure for the deleting a product.
type ProductsDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
