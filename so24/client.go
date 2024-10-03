//go:generate go run generate24.go

package so24

import (
	"fmt"
	"os"

	"github.com/hooklift/gowsdl/soap"
	"github.com/valuetechdev/api-client-24so/so24/account24"
	"github.com/valuetechdev/api-client-24so/so24/auth24"
	"github.com/valuetechdev/api-client-24so/so24/client24"
	"github.com/valuetechdev/api-client-24so/so24/company24"
	"github.com/valuetechdev/api-client-24so/so24/invoice24"
	"github.com/valuetechdev/api-client-24so/so24/product24"
)

const (
	auth_url    = "https://api.24sevenoffice.com/authenticate/v001/authenticate.asmx"
	invoice_url = "https://api.24sevenoffice.com/Economy/InvoiceOrder/V001/InvoiceService.asmx"
	account_url = "https://api.24sevenoffice.com/Economy/Account/V004/Accountservice.asmx"
	product_url = "https://api.24sevenoffice.com/Logistics/Product/V001/ProductService.asmx"
	company_url = "https://api.24sevenoffice.com/CRM/Company/V001/CompanyService.asmx"
)

// holds the different soap clients for the different services
type Client struct {
	SessionId string
	Account   account24.AccountServiceSoap
	Auth      auth24.AuthenticateSoap
	Client    client24.ClientServiceSoap
	Company   company24.CompanyServiceSoap
	Invoice   invoice24.InvoiceServiceSoap
	Product   product24.ProductServiceSoap
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
		soap.NewClient(auth_url, withHeaders),
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
		soap.NewClient(invoice_url, withHeaders),
	)
	productService := product24.NewProductServiceSoap(
		soap.NewClient(product_url, withHeaders),
	)
	accountService := account24.NewAccountServiceSoap(
		soap.NewClient(account_url, withHeaders),
	)
	companyService := company24.NewCompanyServiceSoap(
		soap.NewClient(company_url, withHeaders),
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
