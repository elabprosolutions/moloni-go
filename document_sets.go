package moloni

import "github.com/elabprosolutions/moloni-go/models"

type DocumentSets struct {
	backend Backend
}

func (documentSets *DocumentSets) Insert(req models.DocumentSetsInsertRequest) (*models.DocumentSetsInsertResponse, error) {
	var resp models.DocumentSetsInsertResponse
	err := documentSets.backend.Call("/v1/documentSets/insert", req, &resp)
	return &resp, err
}

func (documentSets *DocumentSets) GetAll(req models.DocumentSetsGetAllRequest) (*models.DocumentSetsGetAllResponse, error) {
	var resp models.DocumentSetsGetAllResponse
	err := documentSets.backend.Call("/v1/documentSets/getAll", req, &resp)
	return &resp, err
}

func (documentSets *DocumentSets) Update(req models.DocumentSetsUpdateRequest) (*models.DocumentSetsUpdateResponse, error) {
	var resp models.DocumentSetsUpdateResponse
	err := documentSets.backend.Call("/v1/documentSets/update", req, &resp)
	return &resp, err
}

func (documentSets *DocumentSets) Delete(req models.DocumentSetsDeleteRequest) (*models.DocumentSetsDeleteResponse, error) {
	var resp models.DocumentSetsDeleteResponse
	err := documentSets.backend.Call("/v1/documentSets/delete", req, &resp)
	return &resp, err
}
