package moloni_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type ProductsTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *ProductsTestSuite) SetupSuite() {
	client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
	s.Require().NoError(err)

	s.client = client
}

func (s *ProductsTestSuite) TearDownSuite() {
	CleanupProducts(s.T(), s.client)
}

func TestProductsTestSuite(t *testing.T) {
	suite.Run(t, new(ProductsTestSuite))
}

func (s *ProductsTestSuite) TestInsertProduct() {
	vatType := models.VATTypeNormal
	insertTaxReq := models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		FiscalZone:      "PT",
		VATType:         &vatType,
		ActiveByDefault: 0,
	}
	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
	s.Require().NoError(err)

	resp, err := s.client.Products.Insert(models.ProductsInsertRequest{
		CompanyID:  CompanyID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      insertTax.TaxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.ProductID)
}

func (s *ProductsTestSuite) TestGetAllProducts() {
	vatType := models.VATTypeNormal
	insertTaxReq := models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		FiscalZone:      "PT",
		VATType:         &vatType,
		ActiveByDefault: 0,
	}
	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
	s.Require().NoError(err)

	insertResp, err := s.client.Products.Insert(models.ProductsInsertRequest{
		CompanyID:  CompanyID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      insertTax.TaxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Products.GetAll(models.ProductsGetAllRequest{
		CompanyID:  CompanyID,
		CategoryID: moloni.Int(CategoryID),
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertProductsGetAllResponseContainsProductWithID(resp, insertResp.ProductID)
}

func (s *ProductsTestSuite) TestUpdateProduct() {
	vatType := models.VATTypeNormal
	insertTaxReq := models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		FiscalZone:      "PT",
		VATType:         &vatType,
		ActiveByDefault: 0,
	}
	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
	s.Require().NoError(err)

	insertResp, err := s.client.Products.Insert(models.ProductsInsertRequest{
		CompanyID:  CompanyID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      insertTax.TaxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	req := models.ProductsUpdateRequest{
		CompanyID:  CompanyID,
		ProductID:  insertResp.ProductID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      insertTax.TaxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	}
	resp, err := s.client.Products.Update(req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.Equal(insertResp.ProductID, resp.ProductID)

	product, err := s.findProductWithID(insertResp.ProductID)
	s.NoError(err)

	s.Equal(req.Name, product.Name)
}

func (s *ProductsTestSuite) TestDeleteProduct() {
	vatType := models.VATTypeNormal
	insertTaxReq := models.TaxesInsertRequest{
		CompanyID:       CompanyID,
		Name:            IntegrationTestRandomName(),
		Value:           5,
		Type:            models.TaxTypePercentage,
		SaftType:        models.SaftTypeValueAdded,
		FiscalZone:      "PT",
		VATType:         &vatType,
		ActiveByDefault: 0,
	}
	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
	s.Require().NoError(err)

	insertResp, err := s.client.Products.Insert(models.ProductsInsertRequest{
		CompanyID:  CompanyID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      insertTax.TaxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Products.Delete(models.ProductsDeleteRequest{
		CompanyID: CompanyID,
		ProductID: insertResp.ProductID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	product, err := s.findProductWithID(insertResp.ProductID)
	s.NoError(err)
	s.Nil(product)
}

func (s *ProductsTestSuite) findProductWithID(productID int) (*models.ProductEntry, error) {
	productes, err := s.client.Products.GetAll(models.ProductsGetAllRequest{
		CompanyID:  CompanyID,
		CategoryID: moloni.Int(CategoryID),
	})
	if err != nil {
		return nil, err
	}

	for _, product := range *productes {
		if product.ProductID == productID {
			return &product, nil
		}
	}

	return nil, nil
}

func (s *ProductsTestSuite) assertProductsGetAllResponseContainsProductWithID(resp *models.ProductsGetAllResponse, productID int) {
	s.NotNil(resp, "ProductsGetAllResponse should not be nil")

	found := false
	for _, product := range *resp {
		if product.ProductID == productID {
			found = true
			break
		}
	}

	s.True(found, "Product should be present in the ProductsGetAllResponse")
}
