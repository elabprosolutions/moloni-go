package moloni

import "github.com/0gener/moloni-go/models"

type Taxes struct {
	backend Backend
}

func (taxes *Taxes) Insert(req models.TaxesInsertRequest) (*models.TaxesInsertResponse, error) {
	var resp models.TaxesInsertResponse
	err := taxes.backend.Call("/v1/taxes/insert", req, &resp)
	return &resp, err
}

func (taxes *Taxes) GetAll(req models.TaxesGetAllRequest) (*models.TaxesGetAllResponse, error) {
	var resp models.TaxesGetAllResponse
	err := taxes.backend.Call("/v1/taxes/getAll", req, &resp)
	return &resp, err
}

func (taxes *Taxes) Update(req models.TaxesUpdateRequest) (*models.TaxesUpdateResponse, error) {
	var resp models.TaxesUpdateResponse
	err := taxes.backend.Call("/v1/taxes/update", req, &resp)
	return &resp, err
}

func (taxes *Taxes) Delete(req models.TaxesDeleteRequest) (*models.TaxesDeleteResponse, error) {
	var resp models.TaxesDeleteResponse
	err := taxes.backend.Call("/v1/taxes/delete", req, &resp)
	return &resp, err
}
