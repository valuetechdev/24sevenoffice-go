package so24

import (
	"context"
	"testing"
	"time"

	"github.com/hooklift/gowsdl/soap"
	"github.com/stretchr/testify/assert"
	"github.com/valuetechdev/api-client-24so/so24/account24"
	"github.com/valuetechdev/api-client-24so/so24/auth24"
	"github.com/valuetechdev/api-client-24so/so24/client24"
	"github.com/valuetechdev/api-client-24so/so24/company24"
	"github.com/valuetechdev/api-client-24so/so24/invoice24"
	"github.com/valuetechdev/api-client-24so/so24/person24"
	"github.com/valuetechdev/api-client-24so/so24/product24"
	"github.com/valuetechdev/api-client-24so/so24/project24"
)

func TestClientInitialization(t *testing.T) {
	assert := assert.New(t)

	c, err := NewAuthenticatedClient()
	assert.NoError(err)

	assert.NotEmpty(c.SessionId)
}

func TestServices(t *testing.T) {
	assert := assert.New(t)

	c, err := NewAuthenticatedClient()
	assert.NoError(err)

	changedAfter := soap.XSDDateTime(soap.CreateXsdDateTime(time.Now(), true))

	_, err = c.Account.GetAccountList(&account24.GetAccountList{})
	assert.NoError(err, "GetAccountList")
	_, err = c.Auth.GetIdentities(&auth24.GetIdentities{})
	assert.NoError(err, "GetIdentities")
	_, err = c.Client.GetClientInformation(&client24.GetClientInformation{})
	assert.NoError(err, "GetClientInformation")
	_, err = c.Company.GetCompanies(&company24.GetCompanies{SearchParams: &company24.CompanySearchParameters{
		ChangedAfter: changedAfter,
	}})
	assert.NoError(err, "GetCompanies")
	_, err = c.Invoice.GetInvoices(&invoice24.GetInvoices{SearchParams: &invoice24.InvoiceSearchParameters{
		ChangedAfter: changedAfter,
	}})
	assert.NoError(err, "GetInvoices")
	_, err = c.Person.GetPersons(&person24.GetPersons{
		PersonSearch: &person24.PersonSearchParameters{
			ChangedAfter: changedAfter,
		},
	})
	assert.NoError(err, "GetPersons")
	_, err = c.Person.GetPersons(&person24.GetPersons{
		PersonSearch: &person24.PersonSearchParameters{
			ChangedAfter: changedAfter,
		},
	})
	assert.NoError(err, "GetPersons")
	_, err = c.Product.GetProducts(&product24.GetProducts{
		SearchParams: &product24.ProductSearchParameters{DateChanged: changedAfter},
	})
	assert.NoError(err, "GetProducts")
	_, err = c.Project.GetProjectList(&project24.GetProjectList{Ps: &project24.ProjectSearch{ChangedAfter: changedAfter}})
	assert.NoError(err, "GetProjectList")

	a, err := c.Payroll.Client.GetAbsenceV2WithResponse(context.TODO())
	assert.NoError(err, "GetAbsenceV2EmpIdWithResponse")
	assert.Empty(a.JSON200)
}
