package so24

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valuetechdev/24sevenoffice-go/so24/company24"
)

func TestCompany(t *testing.T) {
	require := require.New(t)

	c, err := NewAuthenticatedClient(&Credentials{
		ApplicationId: os.Getenv("TWENTYFOURSEVEN_API_APPLICATIONID"),
		Username:      os.Getenv("TWENTYFOURSEVEN_API_USERNAME"),
		Password:      os.Getenv("TWENTYFOURSEVEN_API_PASSWORD"),
		PayrollAPI:    os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"),
	})
	require.NoError(err)
	require.NotEmpty(c.SessionId)

	now := time.Now() // 13:00
	_, err = c.Company.SaveCompanies(&company24.SaveCompanies{
		Companies: &company24.ArrayOfCompany{Company: []*company24.Company{{Id: 22, Name: "CW Kunde fra test lol"}}},
	}) // 13:01
	require.NoError(err)

	changedAfter := EnsureValidTimeForSOAP(now)
	require.Equal(now.Unix(), changedAfter.ToGoTime().Unix())

	cres, err := c.Company.GetCompanies(&company24.GetCompanies{SearchParams: &company24.CompanySearchParameters{ChangedAfter: changedAfter}})
	require.NoError(err)
	require.NotNil(cres.GetCompaniesResult)
	require.NotNil(cres.GetCompaniesResult.Company)
	require.NotEmpty(cres.GetCompaniesResult.Company)

	res, err := c.Company.GetCompanySyncList(&company24.GetCompanySyncList{
		SyncSearchParameters: &company24.SyncSearchParameters{ChangedAfter: changedAfter, Page: 1}},
	)
	require.NoError(err)
	require.NotNil(res.GetCompanySyncListResult)
	require.NotNil(res.GetCompanySyncListResult.Items)
	require.NotNil(res.GetCompanySyncListResult.Items.SyncCompany)
	require.NotEmpty(res.GetCompanySyncListResult.Items.SyncCompany)
}
