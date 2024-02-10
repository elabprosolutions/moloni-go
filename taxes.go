package moloni

import "github.com/elabprosolutions/moloni-go/models"

type Taxes struct {
	client *Client
}

func (taxes *Taxes) Insert(req models.TaxesInsertRequest) (*models.TaxesInsertResponse, error) {
	var resp models.TaxesInsertResponse
	err := taxes.client.Call("/v1/taxes/insert/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (taxes *Taxes) GetAll(req models.TaxesGetAllRequest) (*models.TaxesGetAllResponse, error) {
	var resp models.TaxesGetAllResponse
	err := taxes.client.Call("/v1/taxes/getAll/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (taxes *Taxes) Update(req models.TaxesUpdateRequest) (*models.TaxesUpdateResponse, error) {
	var resp models.TaxesUpdateResponse
	err := taxes.client.Call("/v1/taxes/update/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (taxes *Taxes) Delete(req models.TaxesDeleteRequest) (*models.TaxesDeleteResponse, error) {
	var resp models.TaxesDeleteResponse
	err := taxes.client.Call("/v1/taxes/delete/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
