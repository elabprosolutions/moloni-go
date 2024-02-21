package moloni_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/elabprosolutions/moloni-go"
	"github.com/elabprosolutions/moloni-go/models"
	"github.com/stretchr/testify/require"
)

const (
	CompanyID     = 5
	CategoryID    = 7818474
	DocumentSetID = 714385
	UnitID        = 2730166
	CleanupPrefix = "IntegrationTest"
)

func CleanupTaxes(t *testing.T, client *moloni.Client) {
	resp, err := client.Taxes.GetAll(models.TaxesGetAllRequest{
		CompanyID: CompanyID,
	})
	require.NoError(t, err)

	for _, tax := range *resp {
		if strings.Contains(tax.Name, CleanupPrefix) {
			_, err = client.Taxes.Delete(models.TaxesDeleteRequest{
				CompanyID: CompanyID,
				TaxID:     tax.TaxID,
			})
			require.NoError(t, err)
		}
	}
}

func CleanupDocumentSets(t *testing.T, client *moloni.Client) {
	resp, err := client.DocumentSets.GetAll(models.DocumentSetsGetAllRequest{
		CompanyID: CompanyID,
	})
	require.NoError(t, err)

	for _, ds := range *resp {
		if strings.Contains(ds.Name, CleanupPrefix) {
			_, err = client.DocumentSets.Delete(models.DocumentSetsDeleteRequest{
				CompanyID:     CompanyID,
				DocumentSetID: ds.DocumentSetID,
			})
			require.NoError(t, err)
		}
	}
}

func CleanupCustomers(t *testing.T, client *moloni.Client) {
	resp, err := client.Customers.GetAll(models.CustomersGetAllRequest{
		CompanyID: CompanyID,
	})
	require.NoError(t, err)

	for _, customer := range *resp {
		if strings.Contains(customer.Name, CleanupPrefix) {
			_, err = client.Customers.Delete(models.CustomersDeleteRequest{
				CompanyID:  CompanyID,
				CustomerID: customer.CustomerID,
			})
			require.NoError(t, err)
		}
	}
}

func CleanupProducts(t *testing.T, client *moloni.Client) {
	resp, err := client.Products.GetAll(models.ProductsGetAllRequest{
		CompanyID:  CompanyID,
		CategoryID: moloni.Int(CategoryID),
	})
	require.NoError(t, err)

	for _, product := range *resp {
		if strings.Contains(product.Name, CleanupPrefix) {
			_, err = client.Products.Delete(models.ProductsDeleteRequest{
				CompanyID: CompanyID,
				ProductID: product.ProductID,
			})
			require.NoError(t, err)
		}
	}
}

func CleanupInvoices(t *testing.T, client *moloni.Client) {
	resp, err := client.Invoices.GetAll(models.InvoicesGetAllRequest{
		CompanyID: CompanyID,
	})
	require.NoError(t, err)

	for _, invoice := range *resp {
		if strings.Contains(invoice.EntityName, CleanupPrefix) {
			_, err = client.Invoices.Delete(models.InvoicesDeleteRequest{
				CompanyID:  CompanyID,
				DocumentID: invoice.DocumentID,
			})
			require.NoError(t, err)
		}
	}
}

func GenerateRandomNIF() string {
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

func IntegrationTestRandomName() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s%d", CleanupPrefix, timestamp)
}
