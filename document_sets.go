package moloni

import "github.com/elabprosolutions/moloni-go/models"

type DocumentSets struct {
	client *Client
}

func (documentSets *DocumentSets) Insert(req models.DocumentSetsInsertRequest) (*models.DocumentSetsInsertResponse, error) {
	var resp models.DocumentSetsInsertResponse
	err := documentSets.client.Call("/v1/documentSets/insert/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (documentSets *DocumentSets) GetAll(req models.DocumentSetsGetAllRequest) (*models.DocumentSetsGetAllResponse, error) {
	var resp models.DocumentSetsGetAllResponse
	err := documentSets.client.Call("/v1/documentSets/getAll/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (documentSets *DocumentSets) Update(req models.DocumentSetsUpdateRequest) (*models.DocumentSetsUpdateResponse, error) {
	var resp models.DocumentSetsUpdateResponse
	err := documentSets.client.Call("/v1/documentSets/update/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (documentSets *DocumentSets) Delete(req models.DocumentSetsDeleteRequest) (*models.DocumentSetsDeleteResponse, error) {
	var resp models.DocumentSetsDeleteResponse
	err := documentSets.client.Call("/v1/documentSets/delete/", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
