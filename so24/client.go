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
	"github.com/valuetechdev/api-client-24so/so24/payroll24"
	"github.com/valuetechdev/api-client-24so/so24/person24"
	"github.com/valuetechdev/api-client-24so/so24/product24"
	"github.com/valuetechdev/api-client-24so/so24/project24"
	"github.com/valuetechdev/api-client-24so/so24/transaction24"
)

const (
	account_url     = "https://api.24sevenoffice.com/Economy/Account/V004/Accountservice.asmx"
	auth_url        = "https://api.24sevenoffice.com/authenticate/v001/authenticate.asmx"
	client_url      = "https://api.24sevenoffice.com/Client/V001/ClientService.asmx"
	company_url     = "https://api.24sevenoffice.com/CRM/Company/V001/CompanyService.asmx"
	invoice_url     = "https://api.24sevenoffice.com/Economy/InvoiceOrder/V001/InvoiceService.asmx"
	person_url      = "https://webservices.24sevenoffice.com/CRM/Contact/PersonService.asmx"
	product_url     = "https://api.24sevenoffice.com/Logistics/Product/V001/ProductService.asmx"
	project_url     = "https://webservices.24sevenoffice.com/Project/V001/ProjectService.asmx"
	transaction_url = "https://api.24sevenoffice.com/Economy/Accounting/V001/TransactionService.asmx"
)

// holds the different soap clients for the different services
type Client struct {
	SessionId   string
	Account     account24.AccountServiceSoap
	Auth        auth24.AuthenticateSoap
	Client      client24.ClientServiceSoap
	Company     company24.CompanyServiceSoap
	Invoice     invoice24.InvoiceServiceSoap
	Person      person24.PersonServiceSoap
	Product     product24.ProductServiceSoap
	Project     project24.ProjectServiceSoap
	Payroll     payroll24.PayrollService
	Transaction transaction24.TransactionServiceSoap
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
	clientService := client24.NewClientServiceSoap(
		soap.NewClient(client_url, withHeaders),
	)
	personService := person24.NewPersonServiceSoap(
		soap.NewClient(person_url, withHeaders),
	)
	projectService := project24.NewProjectServiceSoap(
		soap.NewClient(project_url, withHeaders),
	)
	transactionService := transaction24.NewTransactionServiceSoap(
		soap.NewClient(transaction_url, withHeaders),
	)
	payrollService, err := payroll24.New()
	if err != nil {
		return nil, err
	}

	client := Client{
		SessionId:   sessionId,
		Account:     accountService,
		Auth:        authService,
		Client:      clientService,
		Company:     companyService,
		Invoice:     invoiceService,
		Payroll:     *payrollService,
		Person:      personService,
		Product:     productService,
		Project:     projectService,
		Transaction: transactionService,
	}

	return &client, nil
}

func getCredentials() (*auth24.Credential, error) {
	so24username := requireEnv("TWENTYFOURSEVEN_API_USERNAME")
	so24password := requireEnv("TWENTYFOURSEVEN_API_PASSWORD")
	so24appId := requireEnv("TWENTYFOURSEVEN_API_APPLICATIONID")

	so24appGuid := auth24.Guid(so24appId)

	return &auth24.Credential{
		ApplicationId: &so24appGuid,
		Password:      so24password,
		Username:      so24username,
	}, nil
}

func requireEnv(env string) string {
	v, ok := os.LookupEnv(env)
	if !ok {
		panic(fmt.Sprintf("%s is not set", env))
	}

	return v
}
