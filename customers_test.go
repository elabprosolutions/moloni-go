package moloni_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/suite"
)

type CustomersTestSuite struct {
	suite.Suite
	client *moloni.Client
}

func (s *CustomersTestSuite) SetupSuite() {
	client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
	s.Require().NoError(err)

	s.client = client
}

func (s *CustomersTestSuite) TearDownSuite() {
	CleanupCustomers(s.T(), s.client)
}

func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}

func (s *CustomersTestSuite) TestInsertCustomer() {
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
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.CustomerID)
}

func (s *CustomersTestSuite) TestGetAllCustomers() {
	insertReq := models.CustomersInsertRequest{
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.GetAll(models.CustomersGetAllRequest{
		CompanyID: CompanyID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertCustomersGetAllResponseContainsCustomerWithID(resp, insertResp.CustomerID)
}

func (s *CustomersTestSuite) TestUpdateCustomer() {
	insertReq := models.CustomersInsertRequest{
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.Update(models.CustomersUpdateRequest{
		CompanyID:        CompanyID,
		CustomerID:       insertResp.CustomerID,
		VAT:              GenerateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             "IntegrationTest Customer Updated",
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
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.Equal(insertResp.CustomerID, resp.CustomerID)

	customer, err := s.findCustomerWithID(insertResp.CustomerID)
	s.NoError(err)

	s.Equal("IntegrationTest Customer Updated", customer.Name)
}

func (s *CustomersTestSuite) TestDeleteCustomer() {
	insertReq := models.CustomersInsertRequest{
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.Delete(models.CustomersDeleteRequest{
		CompanyID:  CompanyID,
		CustomerID: insertResp.CustomerID,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)

	customer, err := s.findCustomerWithID(insertResp.CustomerID)
	s.NoError(err)
	s.Nil(customer)
}

func (s *CustomersTestSuite) findCustomerWithID(customerID int) (*models.CustomerEntry, error) {
	customeres, err := s.client.Customers.GetAll(models.CustomersGetAllRequest{
		CompanyID: CompanyID,
	})
	if err != nil {
		return nil, err
	}

	for _, customer := range *customeres {
		if customer.CustomerID == customerID {
			return &customer, nil
		}
	}

	return nil, nil
}

func (s *CustomersTestSuite) assertCustomersGetAllResponseContainsCustomerWithID(resp *models.CustomersGetAllResponse, customerID int) {
	s.NotNil(resp, "CustomersGetAllResponse should not be nil")

	found := false
	for _, customer := range *resp {
		if customer.CustomerID == customerID {
			found = true
			break
		}
	}

	s.True(found, "Customer should be present in the CustomersGetAllResponse")
}
