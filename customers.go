package moloni

import "github.com/elabprosolutions/moloni-go/models"

type Customers struct {
	client *Client
}

func (taxes *Customers) Insert(req models.CustomersInsertRequest) (*models.CustomersInsertResponse, error) {
	var resp models.CustomersInsertResponse
	err := taxes.client.Call("/v1/customers/insert/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (customers *Customers) GetAll(req models.CustomersGetAllRequest) (*models.CustomersGetAllResponse, error) {
	var resp models.CustomersGetAllResponse
	err := customers.client.Call("/v1/customers/getAll/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (customers *Customers) Update(req models.CustomersUpdateRequest) (*models.CustomersUpdateResponse, error) {
	var resp models.CustomersUpdateResponse
	err := customers.client.Call("/v1/customers/update/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (customers *Customers) Delete(req models.CustomersDeleteRequest) (*models.CustomersDeleteResponse, error) {
	var resp models.CustomersDeleteResponse
	err := customers.client.Call("/v1/customers/delete/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
