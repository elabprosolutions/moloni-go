package moloni

import "github.com/elabprosolutions/moloni-go/models"

type Products struct {
	client *Client
}

func (products *Products) Insert(req models.ProductsInsertRequest) (*models.ProductsInsertResponse, error) {
	var resp models.ProductsInsertResponse
	err := products.client.Call("/v1/products/insert/", req, &resp)
	return &resp, err
}

func (products *Products) GetAll(req models.ProductsGetAllRequest) (*models.ProductsGetAllResponse, error) {
	var resp models.ProductsGetAllResponse
	err := products.client.Call("/v1/products/getAll/", req, &resp)
	return &resp, err
}

func (products *Products) Update(req models.ProductsUpdateRequest) (*models.ProductsUpdateResponse, error) {
	var resp models.ProductsUpdateResponse
	err := products.client.Call("/v1/products/update/", req, &resp)
	return &resp, err
}

func (products *Products) Delete(req models.ProductsDeleteRequest) (*models.ProductsDeleteResponse, error) {
	var resp models.ProductsDeleteResponse
	err := products.client.Call("/v1/products/delete/", req, &resp)
	return &resp, err
}
