package moloni_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	s.Cleanup()
}

func TestCustomersTestSuite(t *testing.T) {
	suite.Run(t, new(CustomersTestSuite))
}

func (s *CustomersTestSuite) Cleanup() {
	resp, err := s.client.Customers.GetAll(models.CustomersGetAllRequest{
		CompanyID: 5,
	})
	s.Require().NoError(err)

	for _, customer := range *resp {
		if strings.Contains(customer.Name, "IntegrationTest") {
			_, err = s.client.Customers.Delete(models.CustomersDeleteRequest{
				CompanyID:  5,
				CustomerID: customer.CustomerID,
			})
			s.Require().NoError(err)
		}
	}
}

func (s *CustomersTestSuite) TestInsertCustomer() {
	zeroInt := 0
	zeroFloat := float64(0)
	req := models.CustomersInsertRequest{
		CompanyID:        5,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             s.integrationTestCustomerName(),
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
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(1, resp.Valid)
	s.NotZero(resp.CustomerID)
}

func (s *CustomersTestSuite) TestGetAllCustomers() {
	zeroInt := 0
	zeroFloat := float64(0)
	insertReq := models.CustomersInsertRequest{
		CompanyID:        5,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             s.integrationTestCustomerName(),
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.GetAll(models.CustomersGetAllRequest{
		CompanyID: 5,
	})
	s.NoError(err)
	s.NotNil(resp)
	s.assertCustomersGetAllResponseContainsCustomerWithID(resp, insertResp.CustomerID)
}

func (s *CustomersTestSuite) TestUpdateCustomer() {
	zeroInt := 0
	zeroFloat := float64(0)
	insertReq := models.CustomersInsertRequest{
		CompanyID:        5,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             s.integrationTestCustomerName(),
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.Update(models.CustomersUpdateRequest{
		CompanyID:        5,
		CustomerID:       insertResp.CustomerID,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             "IntegrationTest Customer Updated",
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
	zeroInt := 0
	zeroFloat := float64(0)
	insertReq := models.CustomersInsertRequest{
		CompanyID:        5,
		VAT:              generateRandomNIF(),
		Number:           strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:             s.integrationTestCustomerName(),
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
	insertResp, err := s.client.Customers.Insert(insertReq)
	s.Require().NoError(err)
	s.Require().NotNil(insertResp)

	resp, err := s.client.Customers.Delete(models.CustomersDeleteRequest{
		CompanyID:  5,
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
		CompanyID: 5,
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

func (s *CustomersTestSuite) integrationTestCustomerName() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("IntegrationTest%d", timestamp)
}

func generateRandomNIF() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	nif := "3" + fmt.Sprintf("%07d", rng.Intn(10000000))

	controlDigit := calculateControlDigit(nif)

	return nif + fmt.Sprintf("%d", controlDigit)
}

func calculateControlDigit(nif string) int {
	sum := 0
	for i := 0; i < 8; i++ {
		digit := int(nif[i] - '0')
		sum += digit * (9 - i)
	}

	remainder := sum % 11
	controlDigit := 11 - remainder

	if controlDigit >= 10 {
		return 0
	}
	return controlDigit
}
