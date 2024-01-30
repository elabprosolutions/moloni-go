package moloni_test

import (
	"strings"
	"testing"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type TaxTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *TaxTestSuite) SetupSuite() {
	client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
	s.Require().NoError(err)

	s.client = client
}

func (s *TaxTestSuite) TearDownSuite() {
	s.Cleanup()
}

func TestTaxTestSuite(t *testing.T) {
	suite.Run(t, new(TaxTestSuite))
}

func (s *TaxTestSuite) Cleanup() {
	resp, err := s.client.Taxes.GetAll(models.TaxesGetAllRequest{
		CompanyID: 5,
	})
	s.Require().NoError(err)

	for _, tax := range *resp {
		if strings.Contains(tax.Name, "Integration Tests") {
			_, err = s.client.Taxes.Delete(models.TaxesDeleteRequest{
				CompanyID: 5,
				TaxID:     tax.TaxID,
			})
			s.Require().NoError(err)
		}
	}
}

func (s *TaxTestSuite) TestInsertTax() {
	vatType := models.VATTypeNormal

	resp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       5,
		Name:            "Integration Tests Tax",
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

func (s *TaxTestSuite) TestGetAllTaxes() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       5,
		Name:            "Integration Tests Tax",
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
		CompanyID: 5,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertTaxesGetAllResponseContainsTaxWithID(resp, insertResp.TaxID)
}

func (s *TaxTestSuite) TestUpdateTax() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       5,
		Name:            "Integration Tests Tax",
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
		CompanyID:       5,
		TaxID:           insertResp.TaxID,
		Name:            "Integration Tests Tax Updated",
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

	s.Equal("Integration Tests Tax Updated", tax.Name)
	s.Equal(float64(6), tax.Value)
}

func (s *TaxTestSuite) TestDeleteTax() {
	vatType := models.VATTypeNormal

	insertResp, err := s.client.Taxes.Insert(models.TaxesInsertRequest{
		CompanyID:       5,
		Name:            "Integration Tests Tax",
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
		CompanyID: 5,
		TaxID:     insertResp.TaxID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	tax, err := s.findTaxWithID(insertResp.TaxID)
	s.NoError(err)
	s.Nil(tax)
}

func (s *TaxTestSuite) findTaxWithID(taxID int) (*models.TaxEntry, error) {
	taxes, err := s.client.Taxes.GetAll(models.TaxesGetAllRequest{
		CompanyID: 5,
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

func (s *TaxTestSuite) assertTaxesGetAllResponseContainsTaxWithID(resp *models.TaxesGetAllResponse, taxID int) {
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
