package rest24

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/require"
	u "github.com/valuetechdev/24sevenoffice-go/rest24/internal"
	"golang.org/x/oauth2"
)

// SO24 project: "WooCommerce-integrasjon (DEMO)"
const orgId = "543819716587312"

func TestClientInitialization(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})

	require.False(c.IsTokenValid(), "token should be invalid")
	require.NoError(c.Authenticate(), "client should authenticate")

	// Verify token properties
	token := c.GetToken()
	require.NotNil(token, "token should not be nil")
	require.NotEmpty(token.AccessToken, "access token should not be empty")
	require.Equal("Bearer", token.TokenType, "token type should be Bearer")

	t.Logf("Authentication successful. Token expires: %v", token.Expiry)

	res, err := c.GetAccountsWithResponse(context.Background(), &GetAccountsParams{})
	require.NoError(err, "GetAccounts should not error")
	require.NotNil(res, "GetAccounts should not be nil")
	require.Equal(http.StatusOK, res.StatusCode(), "GetAccounts status should be OK")
	require.NotNil(res.JSON200, "GetAccounts JSON200 should not be nil")
}

func TestClientTokenManagement(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       "test-client",
		ClientSecret:   "test-secret",
		OrganizationId: "test-org",
	})

	testToken := &oauth2.Token{
		AccessToken: "test-access-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Hour),
	}
	c.SetToken(testToken)
	require.Equal(testToken.AccessToken, c.GetToken().AccessToken, "access token should match")
}

func TestTransactionLines(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})
	require.NoError(c.Authenticate(), "client should authenticate")

	params := &GetTransactionlinesParams{
		DateFrom: types.Date{Time: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)},
		DateTo:   types.Date{Time: time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC)},
	}
	res, err := c.GetTransactionlinesWithResponse(context.Background(), params)
	require.NoError(err, "GetTransactionlines should not error")
	require.NotNil(res, "GetTransactionlines should not be nil")
	require.Equal(http.StatusOK, res.StatusCode(), "GetTransactionlines status should be OK")
	require.NotNil(res.JSON200, "GetTransactionlines JSON200 should not be nil")
	require.NotEmpty(*res.JSON200, "GetTransactionlines JSON200 should not be empty")
}

func TestCreatePrivateCustomer(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})
	require.NoError(c.Authenticate(), "client should authenticate")

	cPostRequest := CustomerPostRequest{
		Email: &EmailsDto{
			Billing: u.Ref("billing@example.org"),
			Contact: u.Ref("contact@example.org"),
		},
		Phone: u.Ref("+47-12345678"),

		Address: &AddressesDto{
			Visit: &VisitAddress{
				Street:             u.Ref("Torgallmenningen 1"),
				PostalArea:         u.Ref("Bergen"),
				PostalCode:         u.Ref("5009"),
				CountryCode:        u.Ref("NO"),
				CountrySubdivision: u.Ref("Vestland"),
			},
		},
		IsSupplier: u.Ref(false),
	}
	err := cPostRequest.FromCustomerPostRequest1(CustomerPostRequest1{
		IsCompany: CustomerPostRequest1IsCompany(false),
		Person: FirstnameLastnameDto{
			FirstName: u.Ref("John"),
			LastName:  u.Ref("Doe"),
		},
	})
	require.NoError(err)

	res, err := c.CreateCustomerWithResponse(t.Context(), cPostRequest)
	require.NoError(err)
	require.NotNil(res.JSON200, "no private customer was created")
}

func TestCreateCompanyCustomer(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})
	require.NoError(c.Authenticate(), "client should authenticate")

	cPostRequest := CustomerPostRequest{
		Email: &EmailsDto{
			Billing: u.Ref("billing@example.org"),
			Contact: u.Ref("contact@example.org"),
		},
		Phone: u.Ref("+47-12345678"),
		Address: &AddressesDto{
			Visit: &VisitAddress{
				Street:             u.Ref("Torgallmenningen 1"),
				PostalArea:         u.Ref("Bergen"),
				PostalCode:         u.Ref("5009"),
				CountryCode:        u.Ref("NO"),
				CountrySubdivision: u.Ref("Vestland"),
			},
		},
		IsSupplier:         u.Ref(false),
		OrganizationNumber: u.Ref("123456789"),
	}
	err := cPostRequest.FromCustomerPostRequest0(CustomerPostRequest0{
		IsCompany: CustomerPostRequest0IsCompany(true),
		Name:      "ACME INC.",
	})
	require.NoError(err)

	res, err := c.CreateCustomerWithResponse(t.Context(), cPostRequest)
	require.NoError(err)
	require.NotNil(res.JSON200, "no company customer was created")
}
