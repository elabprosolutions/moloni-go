package moloni

import "github.com/elabprosolutions/moloni-go/models"

type Invoices struct {
	client *Client
}

func (invoices *Invoices) Insert(req models.InvoicesInsertRequest) (*models.InvoicesInsertResponse, error) {
	var resp models.InvoicesInsertResponse
	err := invoices.client.Call("/v1/invoices/insert/", req, &resp)
	return &resp, err
}

func (invoices *Invoices) GetAll(req models.InvoicesGetAllRequest) (*models.InvoicesGetAllResponse, error) {
	var resp models.InvoicesGetAllResponse
	err := invoices.client.Call("/v1/invoices/getAll/", req, &resp)
	return &resp, err
}

func (invoices *Invoices) Update(req models.InvoicesUpdateRequest) (*models.InvoicesUpdateResponse, error) {
	var resp models.InvoicesUpdateResponse
	err := invoices.client.Call("/v1/invoices/update/", req, &resp)
	return &resp, err
}

func (invoices *Invoices) Delete(req models.InvoicesDeleteRequest) (*models.InvoicesDeleteResponse, error) {
	var resp models.InvoicesDeleteResponse
	err := invoices.client.Call("/v1/invoices/delete/", req, &resp)
	return &resp, err
}
