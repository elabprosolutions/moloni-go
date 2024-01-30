package moloni_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type DocumentSetsTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *DocumentSetsTestSuite) SetupSuite() {
	client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
	s.Require().NoError(err)

	s.client = client
}

func (s *DocumentSetsTestSuite) TearDownSuite() {
	s.Cleanup()
}

func TestDocumentSetsTestSuite(t *testing.T) {
	suite.Run(t, new(DocumentSetsTestSuite))
}

func (s *DocumentSetsTestSuite) Cleanup() {
	resp, err := s.client.DocumentSets.GetAll(models.DocumentSetsGetAllRequest{
		CompanyID: 5,
	})
	s.Require().NoError(err)

	for _, ds := range *resp {
		if strings.Contains(ds.Name, "IntegrationTest") {
			_, err = s.client.DocumentSets.Delete(models.DocumentSetsDeleteRequest{
				CompanyID:     5,
				DocumentSetID: ds.DocumentSetID,
			})
			s.Require().NoError(err)
		}
	}
}

func (s *DocumentSetsTestSuite) TestInsertDocumentSet() {
	req := models.DocumentSetsInsertRequest{
		CompanyID: 5,
		Name:      s.integrationTestDocumentSetName(),
	}
	resp, err := s.client.DocumentSets.Insert(req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.DocumentSetID)
}

func (s *DocumentSetsTestSuite) TestGetAllDocumentSets() {
	req := models.DocumentSetsInsertRequest{
		CompanyID: 5,
		Name:      s.integrationTestDocumentSetName(),
	}
	insertResp, err := s.client.DocumentSets.Insert(req)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.DocumentSets.GetAll(models.DocumentSetsGetAllRequest{
		CompanyID: 5,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertDocumentSetsGetAllResponseContainsDocumentSetWithID(resp, insertResp.DocumentSetID)
}

func (s *DocumentSetsTestSuite) TestUpdateDocumentSet() {
	insertReq := models.DocumentSetsInsertRequest{
		CompanyID: 5,
		Name:      s.integrationTestDocumentSetName(),
	}
	insertResp, err := s.client.DocumentSets.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	documentSetID, err := strconv.Atoi(insertResp.DocumentSetID)
	s.Require().NoError(err)

	req := models.DocumentSetsUpdateRequest{
		CompanyID:     5,
		DocumentSetID: documentSetID,
		Name:          s.integrationTestDocumentSetName(),
	}
	resp, err := s.client.DocumentSets.Update(req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.Equal(documentSetID, resp.DocumentSetID)

	ds, err := s.findDocumentSetWithID(resp.DocumentSetID)
	s.NoError(err)

	s.Equal(req.Name, ds.Name)
}

func (s *DocumentSetsTestSuite) TestDeleteDocumentSet() {
	insertReq := models.DocumentSetsInsertRequest{
		CompanyID: 5,
		Name:      s.integrationTestDocumentSetName(),
	}
	insertResp, err := s.client.DocumentSets.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	documentSetID, err := strconv.Atoi(insertResp.DocumentSetID)
	s.Require().NoError(err)

	resp, err := s.client.DocumentSets.Delete(models.DocumentSetsDeleteRequest{
		CompanyID:     5,
		DocumentSetID: documentSetID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	ds, err := s.findDocumentSetWithID(documentSetID)
	s.NoError(err)
	s.Nil(ds)
}

func (s *DocumentSetsTestSuite) findDocumentSetWithID(documentSetID int) (*models.DocumentSetEntry, error) {
	documentSets, err := s.client.DocumentSets.GetAll(models.DocumentSetsGetAllRequest{
		CompanyID: 5,
	})
	if err != nil {
		return nil, err
	}

	for _, ds := range *documentSets {
		if ds.DocumentSetID == documentSetID {
			return &ds, nil
		}
	}

	return nil, nil
}

func (s *DocumentSetsTestSuite) assertDocumentSetsGetAllResponseContainsDocumentSetWithID(resp *models.DocumentSetsGetAllResponse, documentSetID string) {
	s.NotNil(resp, "DocumentSetsGetAllResponse should not be nil")

	found := false
	for _, ds := range *resp {
		if strconv.Itoa(ds.DocumentSetID) == documentSetID {
			found = true
			break
		}
	}

	s.True(found, "Document Set should be present in the DocumentSetsGetAllResponse")
}

func (s *DocumentSetsTestSuite) integrationTestDocumentSetName() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("IntegrationTest%d", timestamp)
}
