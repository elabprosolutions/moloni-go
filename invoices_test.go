package moloni_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type InvoicesTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *InvoicesTestSuite) SetupSuite() {
	client, err := moloni.NewClient(
		moloni.LoadCredentialsFromEnv(),
		moloni.DisplayHumanErrors(),
	)
	s.Require().NoError(err)

	s.client = client
}

func (s *InvoicesTestSuite) TearDownSuite() {
	CleanupInvoices(s.T(), s.client)
	CleanupCustomers(s.T(), s.client)
	CleanupDocumentSets(s.T(), s.client)
	CleanupProducts(s.T(), s.client)
	CleanupTaxes(s.T(), s.client)
}

func TestInvoicesTestSuite(t *testing.T) {
	suite.Run(t, new(InvoicesTestSuite))
}

func (s *InvoicesTestSuite) TestInsertInvoice() {
	product, productID := s.insertProduct()
	// _, documentSetID := s.insertDocumentSet()
	_, customerID := s.insertCustomer()

	resp, err := s.client.Invoices.Insert(models.InvoicesInsertRequest{
		CompanyID: CompanyID,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.DocumentID)
}

func (s *InvoicesTestSuite) TestGetAllInvoices() {
	product, productID := s.insertProduct()
	// _, documentSetID := s.insertDocumentSet()
	customer, customerID := s.insertCustomer()

	insertReq := models.InvoicesInsertRequest{
		CompanyID: CompanyID,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	}
	insertResp, err := s.client.Invoices.Insert(insertReq)
	s.NoError(err)

	resp, err := s.client.Invoices.GetAll(models.InvoicesGetAllRequest{
		CompanyID: CompanyID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertInvoicesGetAllResponseContainsInvoiceWithID(resp, insertResp.DocumentID, insertReq, customer)
}

func (s *InvoicesTestSuite) TestGetAllInvoicesFilterByCustomerID() {
	product, productID := s.insertProduct()
	// _, documentSetID := s.insertDocumentSet()
	customer, customerID := s.insertCustomer()

	insertReq := models.InvoicesInsertRequest{
		CompanyID: CompanyID,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	}
	insertResp, err := s.client.Invoices.Insert(insertReq)
	s.NoError(err)

	resp, err := s.client.Invoices.GetAll(models.InvoicesGetAllRequest{
		CompanyID:  CompanyID,
		CustomerID: moloni.Int(customerID),
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Len(*resp, 1)
	s.assertInvoicesGetAllResponseContainsInvoiceWithID(resp, insertResp.DocumentID, insertReq, customer)
}

func (s *InvoicesTestSuite) TestUpdateInvoice() {
	product, productID := s.insertProduct()
	// _, documentSetID := s.insertDocumentSet()
	_, customerID := s.insertCustomer()

	insertReq := models.InvoicesInsertRequest{
		CompanyID: CompanyID,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	}
	insertResp, err := s.client.Invoices.Insert(insertReq)
	s.NoError(err)

	req := models.InvoicesUpdateRequest{
		CompanyID:  CompanyID,
		DocumentID: insertResp.DocumentID,
		Date: models.Time{
			Time: time.Now().AddDate(1, 1, 1),
		},
		ExpirationDate: models.Time{
			Time: time.Now().AddDate(2, 2, 2),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	}
	resp, err := s.client.Invoices.Update(req)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.Equal(insertResp.DocumentID, resp.DocumentID)

	invoice, err := s.client.Invoices.GetOne(models.InvoicesGetOneRequest{
		CompanyID:  CompanyID,
		DocumentID: moloni.Int(insertResp.DocumentID),
	})
	s.NoError(err)

	s.EqualDateWithoutTime(req.Date, invoice.Date)
	s.EqualDateWithoutTime(req.ExpirationDate, invoice.ExpirationDate)
}

func (s *InvoicesTestSuite) TestDeleteInvoice() {
	product, productID := s.insertProduct()
	// _, documentSetID := s.insertDocumentSet()
	_, customerID := s.insertCustomer()

	insertReq := models.InvoicesInsertRequest{
		CompanyID: CompanyID,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: 711748,
		CustomerID:    customerID,
		Products: []models.InvoiceProduct{
			{
				ProductID: productID,
				Name:      product.Name,
				Summary:   product.Summary,
				Quantity:  10,
				Price:     product.Price,
				Taxes: []models.InvoiceProductTax{
					{
						TaxID: product.Taxes[0].TaxID,
					},
				},
			},
		},
	}
	insertResp, err := s.client.Invoices.Insert(insertReq)
	s.NoError(err)

	resp, err := s.client.Invoices.Delete(models.InvoicesDeleteRequest{
		CompanyID:  CompanyID,
		DocumentID: insertResp.DocumentID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	invoice, err := s.client.Invoices.GetOne(models.InvoicesGetOneRequest{
		CompanyID:  CompanyID,
		DocumentID: moloni.Int(insertResp.DocumentID),
	})
	s.NoError(err)
	s.Nil(invoice)
}

func (s *InvoicesTestSuite) assertInvoicesGetAllResponseContainsInvoiceWithID(resp *models.InvoicesGetAllResponse, documentID int, expected models.InvoicesInsertRequest, expectedCustomer models.CustomersInsertRequest) {
	s.NotNil(resp, "InvoicesGetAllResponse should not be nil")

	var found *models.InvoiceEntry
	for _, invoice := range *resp {
		if invoice.DocumentID == documentID {
			found = &invoice
			break
		}
	}

	s.NotNil(found, "Invoice should be present in the InvoicesGetAllResponse")

	s.EqualDateWithoutTime(expected.Date, found.Date)
	s.EqualDateWithoutTime(expected.ExpirationDate, found.ExpirationDate)
	s.Equal(-1, found.Number)
	s.Equal(expectedCustomer.Name, found.EntityName)
	s.Equal(expectedCustomer.Number, found.EntityNumber)
	s.Equal(expectedCustomer.VAT, found.EntityVAT)
	s.Equal(expectedCustomer.Address, found.EntityAddress)
	s.Equal(expectedCustomer.City, found.EntityCity)
	s.Equal("0000-000", found.EntityZipCode)
	s.NotZero(found.GrossValue)
	s.NotZero(found.TaxesValue)
	s.NotZero(found.NetValue)
}

func (s *InvoicesTestSuite) EqualDateWithoutTime(expected models.Time, actual models.Time) {
	s.Equal(expected.Year(), actual.Year())
	s.Equal(expected.Month(), actual.Month())
	s.Equal(expected.Day(), actual.Day())
}

func (s *InvoicesTestSuite) insertTax() (models.TaxesInsertRequest, int) {
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

	return insertTaxReq, insertTax.TaxID
}

// func (s *InvoicesTestSuite) insertDocumentSet() (models.DocumentSetsInsertRequest, int) {
// 	req := models.DocumentSetsInsertRequest{
// 		CompanyID: CompanyID,
// 		Name:      IntegrationTestRandomName(),
// 	}
// 	resp, err := s.client.DocumentSets.Insert(req)
// 	s.Require().NoError(err)

// 	documentSetID, err := strconv.Atoi(resp.DocumentSetID)
// 	s.Require().NoError(err)

// 	return req, documentSetID
// }

func (s *InvoicesTestSuite) insertCustomer() (models.CustomersInsertRequest, int) {
	req := models.CustomersInsertRequest{
		CompanyID:        CompanyID,
		VAT:              GenerateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             IntegrationTestRandomName(),
		LanguageID:       1,
		Address:          "Test",
		City:             "Test",
		CountryID:        1,
		SalesmanID:       moloni.Int(0),
		MaturityDateID:   0,
		PaymentDay:       moloni.Int(0),
		Discount:         moloni.Float64(0),
		CreditLimit:      moloni.Float64(0),
		PaymentMethodID:  0,
		DeliveryMethodID: 0,
	}
	resp, err := s.client.Customers.Insert(req)
	s.Require().NoError(err)
	s.Require().NoError(err)

	return req, resp.CustomerID
}

func (s *InvoicesTestSuite) insertProduct() (models.ProductsInsertRequest, int) {
	insertTaxReq, taxID := s.insertTax()

	insertProductReq := models.ProductsInsertRequest{
		CompanyID:  CompanyID,
		Name:       IntegrationTestRandomName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     2720315,
		HasStock:   0,
		Taxes: []models.ProductTax{
			{
				TaxID:      taxID,
				Value:      float64(insertTaxReq.Value),
				Order:      1,
				Cumulative: 0,
			},
		},
	}
	insertProduct, err := s.client.Products.Insert(insertProductReq)
	s.Require().NoError(err)

	return insertProductReq, insertProduct.ProductID
}
