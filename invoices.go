package moloni

import "github.com/elabprosolutions/moloni-go/models"

type Invoices struct {
	client *Client
}

func (invoices *Invoices) Insert(req models.InvoicesInsertRequest) (*models.InvoicesInsertResponse, error) {
	var resp models.InvoicesInsertResponse
	err := invoices.client.Call("/v1/invoices/insert/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (invoices *Invoices) GetAll(req models.InvoicesGetAllRequest) (*models.InvoicesGetAllResponse, error) {
	var resp models.InvoicesGetAllResponse
	err := invoices.client.Call("/v1/invoices/getAll/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (invoices *Invoices) GetOne(req models.InvoicesGetOneRequest) (*models.InvoicesGetOneResponse, error) {
	var resp models.InvoicesGetOneResponse
	err := invoices.client.Call("/v1/invoices/getOne/", req, &resp)
	if err != nil {
		return nil, err
	}
	if resp == (models.InvoicesGetOneResponse{}) {
		return nil, nil
	}

	return &resp, nil
}

func (invoices *Invoices) Update(req models.InvoicesUpdateRequest) (*models.InvoicesUpdateResponse, error) {
	var resp models.InvoicesUpdateResponse
	err := invoices.client.Call("/v1/invoices/update/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (invoices *Invoices) Delete(req models.InvoicesDeleteRequest) (*models.InvoicesDeleteResponse, error) {
	var resp models.InvoicesDeleteResponse
	err := invoices.client.Call("/v1/invoices/delete/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
