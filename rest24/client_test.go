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

var cachedClient *Rest24Client

func getClient(t *testing.T) *Rest24Client {
	if cachedClient != nil {
		return cachedClient
	}
	cachedClient = New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})
	require.NoError(t, cachedClient.Authenticate(), "client should authenticate")
	return cachedClient
}

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

	c := getClient(t)

	cPostRequest := CustomerPostRequest{
		Email: &EmailsDto{
			Billing: u.R("billing@example.org"),
			Contact: u.R("contact@example.org"),
		},
		Phone: u.R("+47-12345678"),

		Address: &AddressesDto{
			Visit: &VisitAddress{
				Street:             u.R("Torgallmenningen 1"),
				PostalArea:         u.R("Bergen"),
				PostalCode:         u.R("5009"),
				CountryCode:        u.R("NO"),
				CountrySubdivision: u.R("Vestland"),
			},
		},
		IsSupplier: u.R(false),
	}
	err := cPostRequest.FromCustomerPostRequest1(CustomerPostRequest1{
		IsCompany: CustomerPostRequest1IsCompany(false),
		Person: FirstnameLastnameDto{
			FirstName: u.R("John"),
			LastName:  u.R("Doe"),
		},
	})
	require.NoError(err)

	res, err := c.CreateCustomerWithResponse(t.Context(), cPostRequest)
	require.NoError(err)
	require.NotNil(res.JSON200, "no private customer was created")
}

func TestCreateCompanyCustomer(t *testing.T) {
	require := require.New(t)

	c := getClient(t)

	cPostRequest := CustomerPostRequest{
		Email: &EmailsDto{
			Billing: u.R("billing@example.org"),
			Contact: u.R("contact@example.org"),
		},
		Phone: u.R("+47-12345678"),
		Address: &AddressesDto{
			Visit: &VisitAddress{
				Street:             u.R("Torgallmenningen 1"),
				PostalArea:         u.R("Bergen"),
				PostalCode:         u.R("5009"),
				CountryCode:        u.R("NO"),
				CountrySubdivision: u.R("Vestland"),
			},
		},
		IsSupplier:         u.R(false),
		OrganizationNumber: u.R("123456789"),
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

func TestRetrieveProductUnits(t *testing.T) {
	require := require.New(t)
	c := getClient(t)
	res, err := c.GetUnitsWithResponse(t.Context())
	require.NoError(err)
	require.NotNil(res.JSON200, "did not fetch product units")
}

func TestCreateProduct(t *testing.T) {
	require := require.New(t)

	c := getClient(t)

	cPostRequest := CustomerPostRequest{
		Email: &EmailsDto{
			Billing: u.R("billing@example.org"),
			Contact: u.R("contact@example.org"),
		},
		Phone: u.R("+47-12345678"),
		Address: &AddressesDto{
			Visit: &VisitAddress{
				Street:             u.R("Torgallmenningen 1"),
				PostalArea:         u.R("Bergen"),
				PostalCode:         u.R("5009"),
				CountryCode:        u.R("NO"),
				CountrySubdivision: u.R("Vestland"),
			},
		},
		IsSupplier:         u.R(true),
		OrganizationNumber: u.R("123456789"),
	}
	err := cPostRequest.FromCustomerPostRequest0(CustomerPostRequest0{
		IsCompany: CustomerPostRequest0IsCompany(true),
		Name:      "ACME INC.",
	})
	require.NoError(err)
	resCustomer, err := c.CreateCustomerWithResponse(t.Context(), cPostRequest)
	require.NoError(err)
	require.NotNil(resCustomer.JSON200, "no customer was created")

	pPostRequest := ProductRequestPost{
		Name:         u.R("Badeball"),
		Number:       u.R(u.RandSeq(10)),
		Type:         u.R(ProductTypeEnum(Default)),
		Status:       u.R(ProductStatusEnumActive),
		Description:  u.R("Rund og flytende"),
		CostPrice:    u.R(float32(30)),
		IndirectCost: u.R(float32(20)),
		SalesPrice:   u.R(float32(100)),
		Stock: &StockDto{
			IsManaged: u.R(true),
			Quantity:  u.R(float32(129)),
			Location:  u.R("B-301"),
		},
		Category: &CategoryRequest{Id: u.R(-1)},
		// Optional fields
		// Ean: u.R("0123456789000"),
		// EanAlternative: u.R("0123456789022345678901234"),
		// SupplierProduct: &SupplierProductDto{
		// 	ItemCode: u.R("T1234"),
		// 	Number:   u.R("1234"),
		// 	Name:     u.R("Cheap beach ball"),
		// 	Price:    u.R(float32(5)),
		// },
		// Units:    &UnitsRequest{Id: u.R(float32(4))},
		// Supplier: &SupplierRequest{Id: u.R(float32(u.SafeDeref(resCustomer.JSON200.Id)))},
	}
	resProduct, err := c.CreateProductWithResponse(t.Context(), pPostRequest)
	require.NoError(err)
	require.NotNil(resProduct.JSON201, "no product was created")
}
