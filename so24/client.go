//go:generate go run generate24.go

package so24

import (
	"fmt"
	"os"
	"time"

	"github.com/hooklift/gowsdl/soap"
	"github.com/valuetechdev/api-client-24so/so24/account24"
	"github.com/valuetechdev/api-client-24so/so24/attachment24"
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
	attachment_url  = "https://webservices.24sevenoffice.com/Economy/Accounting/Accounting_V001/AttachmentService.asmx"
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
	Attachment  attachment24.AttachmentServiceSoap
	Auth        auth24.AuthenticateSoap
	Client      client24.ClientServiceSoap
	Company     company24.CompanyServiceSoap
	Invoice     invoice24.InvoiceServiceSoap
	Person      person24.PersonServiceSoap
	Product     product24.ProductServiceSoap
	Project     project24.ProjectServiceSoap
	Payroll     payroll24.PayrollService
	Transaction transaction24.TransactionServiceSoap
	headers     map[string]string
}

// panic if missing credentials
func NewAuthenticatedClient(payrollAPI string) (*Client, error) {
	c, err := NewClient(payrollAPI)
	if err != nil {
		return nil, err
	}

	err = c.CheckAuth(nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewClient(payrollAPI string) (*Client, error) {
	// map for http headers. all clients share the headers map
	// for simplicity sake
	headers := make(map[string]string)

	authService := auth24.NewAuthenticateSoap(
		newSoapClient(auth_url, headers),
	)
	invoiceService := invoice24.NewInvoiceServiceSoap(
		newSoapClient(invoice_url, headers),
	)
	productService := product24.NewProductServiceSoap(
		newSoapClient(product_url, headers),
	)
	accountService := account24.NewAccountServiceSoap(
		newSoapClient(account_url, headers),
	)
	companyService := company24.NewCompanyServiceSoap(
		newSoapClient(company_url, headers),
	)
	clientService := client24.NewClientServiceSoap(
		newSoapClient(client_url, headers),
	)
	personService := person24.NewPersonServiceSoap(
		newSoapClient(person_url, headers),
	)
	projectService := project24.NewProjectServiceSoap(
		newSoapClient(project_url, headers),
	)
	transactionService := transaction24.NewTransactionServiceSoap(
		newSoapClient(transaction_url, headers),
	)
	attachmentService := attachment24.NewAttachmentServiceSoap(
		newSoapClient(attachment_url, headers),
	)
	payrollService, err := payroll24.New(payrollAPI)
	if err != nil {
		return nil, err
	}

	client := Client{
		Account:     accountService,
		Attachment:  attachmentService,
		Auth:        authService,
		Client:      clientService,
		Company:     companyService,
		Invoice:     invoiceService,
		Person:      personService,
		Product:     productService,
		Project:     projectService,
		Transaction: transactionService,
		headers:     headers,
	}

	if payrollService != nil {
		client.Payroll = *payrollService
	}

	return &client, nil
}

func newSoapClient(url string, headers map[string]string) *soap.Client {
	// add headers to auth client for
	// later we add session id to make all calls authenticated
	withHeaders := soap.WithHTTPHeaders(headers)

	// handle timeouts
	withTLSHandshakeTimeout := soap.WithTLSHandshakeTimeout(10 * time.Second)
	withRequestTimeout := soap.WithRequestTimeout(60 * time.Second)

	return soap.NewClient(url, withHeaders, withTLSHandshakeTimeout, withRequestTimeout)
}

func (c *Client) CheckAuth(callback func(sessionId string) error) error {
	r, err := c.Auth.HasSession(&auth24.HasSession{})
	if err != nil {
		return err
	}

	if r.HasSessionResult {
		return nil
	}

	credentials, err := getCredentials()
	if err != nil {
		return err
	}

	res, err := c.Auth.Login(&auth24.Login{
		Credential: credentials,
	})
	if err != nil {
		return err
	}

	sessionId := res.LoginResult

	c.SetSessionId(sessionId)

	if callback == nil {
		return nil
	}
	return callback(sessionId)
}

func (c *Client) SetSessionId(sessionId string) {
	c.headers["cookie"] = fmt.Sprintf("ASP.NET_SessionId=%s", sessionId)
	c.SessionId = sessionId
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
