//go:generate go run generate24.go

package so24

import (
	"fmt"
	"os"

	"github.com/hooklift/gowsdl/soap"
	"github.com/valuetechdev/api-client-24so/pkg/client/account24"
	"github.com/valuetechdev/api-client-24so/pkg/client/auth24"
	"github.com/valuetechdev/api-client-24so/pkg/client/company24"
	"github.com/valuetechdev/api-client-24so/pkg/client/invoice24"
	"github.com/valuetechdev/api-client-24so/pkg/client/product24"
)

const (
	AUTH_URL    = "https://api.24sevenoffice.com/authenticate/v001/authenticate.asmx"
	INVOICE_URL = "https://api.24sevenoffice.com/Economy/InvoiceOrder/V001/InvoiceService.asmx"
	ACCOUNT_URL = "https://api.24sevenoffice.com/Economy/Account/V004/Accountservice.asmx"
	PRODUCT_URL = "https://api.24sevenoffice.com/Logistics/Product/V001/ProductService.asmx"
	COMPANY_URL = "https://api.24sevenoffice.com/CRM/Company/V001/CompanyService.asmx"
)

// holds the different soap clients for the different services
type Client struct {
	SessionId string
	Auth      auth24.AuthenticateSoap
	Invoice   invoice24.InvoiceServiceSoap
	Product   product24.ProductServiceSoap
	Account   account24.AccountServiceSoap
	Company   company24.CompanyServiceSoap
}

// panic if missing credentials
func NewAuthenticatedClient() (*Client, error) {

	// map for http headers. all clients share the headers map
	// for simplicity sake
	headers := make(map[string]string)

	// add headers to auth client for
	// later we add session id to make all calls authenticated
	withHeaders := soap.WithHTTPHeaders(headers)

	authService := auth24.NewAuthenticateSoap(
		soap.NewClient(AUTH_URL, withHeaders),
	)

	credentials, err := getCredentials()
	if err != nil {
		return nil, err
	}

	res, err := authService.Login(&auth24.Login{
		Credential: credentials,
	})
	if err != nil {
		return nil, err
	}

	sessionId := res.LoginResult
	// add sessionId to http headers to allow for authenticated requests
	headers["cookie"] = fmt.Sprintf("ASP.NET_SessionId=%s", sessionId)

	// create the required services
	invoiceService := invoice24.NewInvoiceServiceSoap(
		soap.NewClient(INVOICE_URL, withHeaders),
	)
	productService := product24.NewProductServiceSoap(
		soap.NewClient(PRODUCT_URL, withHeaders),
	)
	accountService := account24.NewAccountServiceSoap(
		soap.NewClient(ACCOUNT_URL, withHeaders),
	)
	companyService := company24.NewCompanyServiceSoap(
		soap.NewClient(COMPANY_URL, withHeaders),
	)

	client := Client{
		SessionId: sessionId,
		Auth:      authService,
		Invoice:   invoiceService,
		Product:   productService,
		Account:   accountService,
		Company:   companyService,
	}

	return &client, nil
}

func getCredentials() (*auth24.Credential, error) {
	so24username, ok := os.LookupEnv("TWENTYFOURSEVEN_API_USERNAME")
	if !ok {
		return nil, fmt.Errorf("missing TWENTYFOURSEVEN_API_USERNAME")
	}
	so24password, ok := os.LookupEnv("TWENTYFOURSEVEN_API_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing TWENTYFOURSEVEN_API_PASSWORD")
	}
	so24appId, ok := os.LookupEnv("TWENTYFOURSEVEN_API_APPLICATIONID")
	if !ok {
		return nil, fmt.Errorf("missing TWENTYFOURSEVEN_API_APPLICATIONID")
	}

	so24appGuid := auth24.Guid(so24appId)

	return &auth24.Credential{
		ApplicationId: &so24appGuid,
		Password:      so24password,
		Username:      so24username,
	}, nil
}
