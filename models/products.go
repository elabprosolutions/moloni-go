package models

// ProductsInsertRequest represents the request structure for inserting a new product.
type ProductsInsertRequest struct {
}

// ProductsInsertResponse represents the response structure for the inserting a new product.
type ProductsInsertResponse struct {
	Valid     int `json:"valid"`      // 1 for valid, 0 for not valid
	ProductID int `json:"product_id"` // The ID of the product
}

// ProductsGetAllRequest represents the request structure for getting all products.
type ProductsGetAllRequest struct {
}

// ProductEntry represents a single product entry in the ProductsGetAllResponse.
type ProductEntry struct {
}

// ProductsGetAllResponse represents the response for getting all products.
type ProductsGetAllResponse []ProductEntry

// ProductsUpdateRequest represents the request structure for updating a product.
type ProductsUpdateRequest struct {
}

// ProductsUpdateResponse represents the response structure for the updating a product.
type ProductsUpdateResponse struct {
	Valid     int `json:"valid"`      // 1 for valid, 0 for not valid
	ProductID int `json:"product_id"` // The ID of the product
}

// ProductsDeleteRequest represents the request structure for deleting a product.
type ProductsDeleteRequest struct {
	CompanyID int `json:"company_id"` // Mandatory
	ProductID int `json:"product_id"` // Mandatory
}

// ProductsDeleteResponse represents the response structure for the deleting a product.
type ProductsDeleteResponse struct {
	Valid int `json:"valid"` // 1 for valid, 0 for not valid
}
