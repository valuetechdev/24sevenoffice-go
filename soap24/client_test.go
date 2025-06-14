package soap24

import (
	"os"
	"testing"
	"time"

	"github.com/hooklift/gowsdl/soap"
	"github.com/stretchr/testify/require"
	"github.com/valuetechdev/24sevenoffice-go/soap24/account24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/auth24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/client24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/company24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/invoice24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/person24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/product24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/project24"
)

var applicationId = auth24.Guid(os.Getenv("TFSO_SOAP_APPLICATIONID"))
var credentials = auth24.Credential{
	ApplicationId: &applicationId,
	Username:      os.Getenv("TFSO_SOAP_USERNAME"),
	Password:      os.Getenv("TFSO_SOAP_PASSWORD"),
}

func TestClientInitialization(t *testing.T) {
	require := require.New(t)

	c := New(credentials)
	require.False(c.IsSessionIdValid(), "client should not have valid sessionId after init")
	require.NoError(c.CheckAuth(), "client should be able to check auth after init")
	require.True(c.IsSessionIdValid(), "client should have a valid sessionId checkAuth")
}

func TestServices(t *testing.T) {
	require := require.New(t)

	c := New(credentials)
	require.NoError(c.CheckAuth(), "client should be able to check auth after init")

	changedAfter := soap.XSDDateTime(soap.CreateXsdDateTime(time.Now(), true))

	_, err := c.Account.GetAccountList(&account24.GetAccountList{})
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
}
