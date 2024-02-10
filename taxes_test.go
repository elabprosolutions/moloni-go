package moloni_test

import (
	"testing"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type TaxesTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *TaxesTestSuite) SetupSuite() {
	client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
	s.Require().NoError(err)

	s.client = client
}

func (s *TaxesTestSuite) TearDownSuite() {
	CleanupTaxes(s.T(), s.client)
}

func TestTaxesTestSuite(t *testing.T) {
	suite.Run(t, new(TaxesTestSuite))
}

func (s *TaxesTestSuite) TestInsertTax() {
	vatType := models.VATTypeNormal

	resp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		VATType:         &vatType,
		FiscalZone:      "PT",
		ActiveByDefault: 0,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.TaxID)
}

func (s *TaxesTestSuite) TestGetAllTaxes() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		VATType:         &vatType,
		FiscalZone:      "PT",
		ActiveByDefault: 0,
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Taxes.GetAll(models.TaxesGetAllRequest{
		CompanyID: CompanyID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertTaxesGetAllResponseContainsTaxWithID(resp, insertResp.TaxID)
}

func (s *TaxesTestSuite) TestUpdateTax() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		VATType:         &vatType,
		FiscalZone:      "PT",
		ActiveByDefault: 0,
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Taxes.Update(models.TaxesUpdateRequest{
		CompanyID:       CompanyID,
		TaxID:           insertResp.TaxID,
		Name:            "IntegrationTest Tax Updated",
		Value:           6,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		VATType:         &vatType,
		FiscalZone:      "PT",
		ActiveByDefault: 0,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.Equal(insertResp.TaxID, resp.TaxID)

	tax, err := s.findTaxWithID(insertResp.TaxID)
	s.NoError(err)

	s.Equal("IntegrationTest Tax Updated", tax.Name)
	s.Equal(float64(6), tax.Value)
}

func (s *TaxesTestSuite) TestDeleteTax() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		VATType:         &vatType,
		FiscalZone:      "PT",
		ActiveByDefault: 0,
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Taxes.Delete(models.TaxesDeleteRequest{
		CompanyID: CompanyID,
		TaxID:     insertResp.TaxID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	tax, err := s.findTaxWithID(insertResp.TaxID)
	s.NoError(err)
	s.Nil(tax)
}

func (s *TaxesTestSuite) findTaxWithID(taxID int) (*models.TaxEntry, error) {
	taxes, err := s.client.Taxes.GetAll(models.TaxesGetAllRequest{
		CompanyID: CompanyID,
	})
	if err != nil {
		return nil, err
	}

	for _, tax := range *taxes {
		if tax.TaxID == taxID {
			return &tax, nil
		}
	}

	return nil, nil
}

func (s *TaxesTestSuite) assertTaxesGetAllResponseContainsTaxWithID(resp *models.TaxesGetAllResponse, taxID int) {
	s.NotNil(resp, "TaxesGetAllResponse should not be nil")

	found := false
	for _, tax := range *resp {
		if tax.TaxID == taxID {
			found = true
			break
		}
	}

	s.True(found, "Tax should be present in the TaxesGetAllResponse")
}
