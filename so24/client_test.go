package so24

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hooklift/gowsdl/soap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valuetechdev/24sevenoffice-go/so24/account24"
	"github.com/valuetechdev/24sevenoffice-go/so24/auth24"
	"github.com/valuetechdev/24sevenoffice-go/so24/client24"
	"github.com/valuetechdev/24sevenoffice-go/so24/company24"
	"github.com/valuetechdev/24sevenoffice-go/so24/invoice24"
	"github.com/valuetechdev/24sevenoffice-go/so24/person24"
	"github.com/valuetechdev/24sevenoffice-go/so24/product24"
	"github.com/valuetechdev/24sevenoffice-go/so24/project24"
)

func TestClientInitialization(t *testing.T) {
	assert := assert.New(t)

	c, err := NewAuthenticatedClient(&Credentials{
		ApplicationId: os.Getenv("TWENTYFOURSEVEN_API_APPLICATIONID"),
		Username:      os.Getenv("TWENTYFOURSEVEN_API_USERNAME"),
		Password:      os.Getenv("TWENTYFOURSEVEN_API_PASSWORD"),
		PayrollAPI:    os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"),
	})
	assert.NoError(err)

	assert.NotEmpty(c.SessionId)
}

func TestClientInitializationWithoutPayroll(t *testing.T) {
	assert := assert.New(t)

	c, err := NewAuthenticatedClient(&Credentials{
		ApplicationId: os.Getenv("TWENTYFOURSEVEN_API_APPLICATIONID"),
		Username:      os.Getenv("TWENTYFOURSEVEN_API_USERNAME"),
		Password:      os.Getenv("TWENTYFOURSEVEN_API_PASSWORD"),
		PayrollAPI:    os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"),
	})
	assert.NoError(err)

	assert.NotEmpty(c.SessionId)
}

func TestServices(t *testing.T) {
	require := require.New(t)

	c, err := NewAuthenticatedClient(&Credentials{
		ApplicationId: os.Getenv("TWENTYFOURSEVEN_API_APPLICATIONID"),
		Username:      os.Getenv("TWENTYFOURSEVEN_API_USERNAME"),
		Password:      os.Getenv("TWENTYFOURSEVEN_API_PASSWORD"),
		PayrollAPI:    os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"),
	})
	require.NoError(err)

	changedAfter := soap.XSDDateTime(soap.CreateXsdDateTime(time.Now(), true))

	_, err = c.Account.GetAccountList(&account24.GetAccountList{})
	require.NoError(err, "GetAccountList")
	_, err = c.Auth.GetIdentities(&auth24.GetIdentities{})
	require.NoError(err, "GetIdentities")
	_, err = c.Client.GetClientInformation(&client24.GetClientInformation{})
	require.NoError(err, "GetClientInformation")
	_, err = c.Company.GetCompanies(&company24.GetCompanies{SearchParams: &company24.CompanySearchParameters{
		ChangedAfter: changedAfter,
	}})
	require.NoError(err, "GetCompanies")
	_, err = c.Invoice.GetInvoices(&invoice24.GetInvoices{SearchParams: &invoice24.InvoiceSearchParameters{
		ChangedAfter: changedAfter,
	}})
	require.NoError(err, "GetInvoices")
	_, err = c.Person.GetPersons(&person24.GetPersons{
		PersonSearch: &person24.PersonSearchParameters{
			ChangedAfter: changedAfter,
		},
	})
	require.NoError(err, "GetPersons")

	var tmp person24.TriState = person24.TriStateNone
	_, err = c.Person.GetPersons(&person24.GetPersons{
		PersonSearch: &person24.PersonSearchParameters{
			ChangedAfter: changedAfter,
			Email:        "lol",
			IsEmployee:   &tmp,
		},
	})
	require.NoError(err, "GetPersons")
	_, err = c.Product.GetProducts(&product24.GetProducts{
		SearchParams: &product24.ProductSearchParameters{DateChanged: changedAfter},
	})
	require.NoError(err, "GetProducts")
	_, err = c.Project.GetProjectList(&project24.GetProjectList{Ps: &project24.ProjectSearch{ChangedAfter: changedAfter}})
	require.NoError(err, "GetProjectList")

	a, err := c.Payroll.Client.GetAbsenceV2WithResponse(context.TODO())
	require.NoError(err, "GetAbsenceV2EmpIdWithResponse")
	require.Empty(a.JSON200)
}
