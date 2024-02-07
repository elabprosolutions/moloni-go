package moloni_test

import (
	"fmt"
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
	s.Cleanup()
}

func TestInvoicesTestSuite(t *testing.T) {
	suite.Run(t, new(InvoicesTestSuite))
}

func (s *InvoicesTestSuite) Cleanup() {
	// resp, err := s.client.Invoices.GetAll(models.InvoicesGetAllRequest{
	// 	CompanyID: 5,
	// })
	// s.Require().NoError(err)

	// for _, invoice := range *resp {
	// 	if strings.Contains(invoice.Name, "IntegrationTest") {
	// 		_, err = s.client.Invoices.Delete(models.InvoicesDeleteRequest{
	// 			CompanyID: 5,
	// 			InvoiceID: invoice.InvoiceID,
	// 		})
	// 		s.Require().NoError(err)
	// 	}
	// }
}

func (s *InvoicesTestSuite) TestInsertInvoice() {
	product, productID := s.insertProduct()
	_, documentSetID := s.insertDocumentSet()
	_, customerID := s.insertCustomer()

	resp, err := s.client.Invoices.Insert(models.InvoicesInsertRequest{
		CompanyID: 5,
		Date: models.Time{
			Time: time.Now(),
		},
		ExpirationDate: models.Time{
			Time: time.Now(),
		},
		DocumentSetID: documentSetID,
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

func (s *InvoicesTestSuite) insertTax() (models.TaxesInsertRequest, int) {
	vatType := models.VATTypeNormal
	insertTaxReq := models.TaxesInsertRequest{
		CompanyID:       5,
		Name:            s.integrationTestInvoiceName(),
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

func (s *InvoicesTestSuite) insertDocumentSet() (models.DocumentSetsInsertRequest, int) {
	req := models.DocumentSetsInsertRequest{
		CompanyID: 5,
		Name:      s.integrationTestInvoiceName(),
	}
	resp, err := s.client.DocumentSets.Insert(req)
	s.Require().NoError(err)

	documentSetID, err := strconv.Atoi(resp.DocumentSetID)
	s.Require().NoError(err)

	return req, documentSetID
}

func (s *InvoicesTestSuite) insertCustomer() (models.CustomersInsertRequest, int) {
	zeroInt := 0
	zeroFloat := float64(0)
	req := models.CustomersInsertRequest{
		CompanyID:        5,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             s.integrationTestInvoiceName(),
		LanguageID:       1,
		Address:          "Test",
		City:             "Test",
		CountryID:        1,
		SalesmanID:       &zeroInt,
		MaturityDateID:   zeroInt,
		PaymentDay:       &zeroInt,
		Discount:         &zeroFloat,
		CreditLimit:      &zeroFloat,
		PaymentMethodID:  zeroInt,
		DeliveryMethodID: zeroInt,
	}
	resp, err := s.client.Customers.Insert(req)
	s.Require().NoError(err)
	s.Require().NoError(err)

	return req, resp.CustomerID
}

func (s *InvoicesTestSuite) insertProduct() (models.ProductsInsertRequest, int) {
	insertTaxReq, taxID := s.insertTax()

	insertProductReq := models.ProductsInsertRequest{
		CompanyID:  5,
		Name:       s.integrationTestInvoiceName(),
		CategoryID: CategoryID,
		Type:       models.ProductTypeProduct,
		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Price:      10,
		UnitID:     1,
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

// func (s *InvoicesTestSuite) TestGetAllInvoices() {
// 	vatType := models.VATTypeNormal
// 	insertTaxReq := models.TaxesInsertRequest{
// 		CompanyID:       5,
// 		Name:            s.integrationTestInvoiceName(),
// 		Value:           5,
// 		Type:            models.TaxTypePercentage,
// 		SaftType:        models.SaftTypeValueAdded,
// 		FiscalZone:      "PT",
// 		VATType:         &vatType,
// 		ActiveByDefault: 0,
// 	}
// 	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
// 	s.Require().NoError(err)

// 	insertResp, err := s.client.Invoices.Insert(models.InvoicesInsertRequest{
// 		CompanyID:  5,
// 		Name:       s.integrationTestInvoiceName(),
// 		CategoryID: CategoryID,
// 		Type:       models.InvoiceTypeInvoice,
// 		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
// 		Price:      10,
// 		UnitID:     1,
// 		HasStock:   0,
// 		Taxes: []models.InvoiceTax{
// 			{
// 				TaxID:      insertTax.TaxID,
// 				Value:      float64(insertTaxReq.Value),
// 				Order:      1,
// 				Cumulative: 0,
// 			},
// 		},
// 	})
// 	s.Require().NoError(err)
// 	s.Require().NotNil(insertResp)

// 	resp, err := s.client.Invoices.GetAll(models.InvoicesGetAllRequest{
// 		CompanyID:  5,
// 		CategoryID: &CategoryID,
// 	})
// 	s.NoError(err)
// 	s.NotNil(resp)
// 	s.assertInvoicesGetAllResponseContainsInvoiceWithID(resp, insertResp.InvoiceID)
// }

// func (s *InvoicesTestSuite) TestUpdateInvoice() {
// 	vatType := models.VATTypeNormal
// 	insertTaxReq := models.TaxesInsertRequest{
// 		CompanyID:       5,
// 		Name:            s.integrationTestInvoiceName(),
// 		Value:           5,
// 		Type:            models.TaxTypePercentage,
// 		SaftType:        models.SaftTypeValueAdded,
// 		FiscalZone:      "PT",
// 		VATType:         &vatType,
// 		ActiveByDefault: 0,
// 	}
// 	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
// 	s.Require().NoError(err)

// 	insertResp, err := s.client.Invoices.Insert(models.InvoicesInsertRequest{
// 		CompanyID:  5,
// 		Name:       s.integrationTestInvoiceName(),
// 		CategoryID: CategoryID,
// 		Type:       models.InvoiceTypeInvoice,
// 		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
// 		Price:      10,
// 		UnitID:     1,
// 		HasStock:   0,
// 		Taxes: []models.InvoiceTax{
// 			{
// 				TaxID:      insertTax.TaxID,
// 				Value:      float64(insertTaxReq.Value),
// 				Order:      1,
// 				Cumulative: 0,
// 			},
// 		},
// 	})
// 	s.Require().NoError(err)
// 	s.Require().NotNil(insertResp)

// 	req := models.InvoicesUpdateRequest{
// 		CompanyID:  5,
// 		InvoiceID:  insertResp.InvoiceID,
// 		Name:       s.integrationTestInvoiceName(),
// 		CategoryID: CategoryID,
// 		Type:       models.InvoiceTypeInvoice,
// 		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
// 		Price:      10,
// 		UnitID:     1,
// 		HasStock:   0,
// 		Taxes: []models.InvoiceTax{
// 			{
// 				TaxID:      insertTax.TaxID,
// 				Value:      float64(insertTaxReq.Value),
// 				Order:      1,
// 				Cumulative: 0,
// 			},
// 		},
// 	}
// 	resp, err := s.client.Invoices.Update(req)
// 	s.NoError(err)
// 	s.NotNil(resp)
// 	s.Equal(1, resp.Valid)
// 	s.Equal(insertResp.InvoiceID, resp.InvoiceID)

// 	invoice, err := s.findInvoiceWithID(insertResp.InvoiceID)
// 	s.NoError(err)

// 	s.Equal(req.Name, invoice.Name)
// }

// func (s *InvoicesTestSuite) TestDeleteInvoice() {
// 	vatType := models.VATTypeNormal
// 	insertTaxReq := models.TaxesInsertRequest{
// 		CompanyID:       5,
// 		Name:            s.integrationTestInvoiceName(),
// 		Value:           5,
// 		Type:            models.TaxTypePercentage,
// 		SaftType:        models.SaftTypeValueAdded,
// 		FiscalZone:      "PT",
// 		VATType:         &vatType,
// 		ActiveByDefault: 0,
// 	}
// 	insertTax, err := s.client.Taxes.Insert(insertTaxReq)
// 	s.Require().NoError(err)

// 	insertResp, err := s.client.Invoices.Insert(models.InvoicesInsertRequest{
// 		CompanyID:  5,
// 		Name:       s.integrationTestInvoiceName(),
// 		CategoryID: CategoryID,
// 		Type:       models.InvoiceTypeInvoice,
// 		Reference:  strconv.FormatInt(time.Now().UnixNano(), 10),
// 		Price:      10,
// 		UnitID:     1,
// 		HasStock:   0,
// 		Taxes: []models.InvoiceTax{
// 			{
// 				TaxID:      insertTax.TaxID,
// 				Value:      float64(insertTaxReq.Value),
// 				Order:      1,
// 				Cumulative: 0,
// 			},
// 		},
// 	})
// 	s.Require().NoError(err)
// 	s.Require().NotNil(insertResp)

// 	resp, err := s.client.Invoices.Delete(models.InvoicesDeleteRequest{
// 		CompanyID: 5,
// 		InvoiceID: insertResp.InvoiceID,
// 	})
// 	s.NoError(err)
// 	s.NotNil(resp)
// 	s.Equal(1, resp.Valid)

// 	invoice, err := s.findInvoiceWithID(insertResp.InvoiceID)
// 	s.NoError(err)
// 	s.Nil(invoice)
// }

// func (s *InvoicesTestSuite) findInvoiceWithID(invoiceID int) (*models.InvoiceEntry, error) {
// 	invoicees, err := s.client.Invoices.GetAll(models.InvoicesGetAllRequest{
// 		CompanyID:  5,
// 		CategoryID: &CategoryID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, invoice := range *invoicees {
// 		if invoice.InvoiceID == invoiceID {
// 			return &invoice, nil
// 		}
// 	}

// 	return nil, nil
// }

func (s *InvoicesTestSuite) assertInvoicesGetAllResponseContainsInvoiceWithID(resp *models.InvoicesGetAllResponse, invoiceID int) {
	s.NotNil(resp, "InvoicesGetAllResponse should not be nil")

	found := false
	for _, invoice := range *resp {
		if invoice.InvoiceID == invoiceID {
			found = true
			break
		}
	}

	s.True(found, "Invoice should be present in the InvoicesGetAllResponse")
}

func (s *InvoicesTestSuite) integrationTestInvoiceName() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("IntegrationTest%d", timestamp)
}
