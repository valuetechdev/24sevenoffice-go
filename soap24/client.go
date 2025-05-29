//go:generate go run generate24.go

package soap24

import (
	"fmt"

	"github.com/hooklift/gowsdl/soap"
	"github.com/valuetechdev/24sevenoffice-go/payroll24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/account24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/attachment24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/auth24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/client24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/company24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/invoice24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/person24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/product24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/project24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/transaction24"
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
	credentials *Credentials
}

type Credentials struct {
	ApplicationId string
	Username      string
	Password      string
	PayrollAPI    string
}

// panic if missing credentials
func NewAuthenticatedClient(creds *Credentials) (*Client, error) {
	if creds == nil {
		return nil, fmt.Errorf("creds cannot be nil")
	}
	c, err := NewClient(creds)
	if err != nil {
		return nil, err
	}

	err = c.CheckAuth(nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewClient(creds *Credentials) (*Client, error) {
	if creds == nil {
		return nil, fmt.Errorf("creds cannot be nil")
	}
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
	payrollService, err := payroll24.New(creds.PayrollAPI)
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
		credentials: creds,
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
	withTLSHandshakeTimeout := soap.WithTLSHandshakeTimeout(0)
	withRequestTimeout := soap.WithRequestTimeout(0)

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

	credentials, err := getCredentials(c.credentials)
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

func getCredentials(creds *Credentials) (*auth24.Credential, error) {
	so24appGuid := auth24.Guid(creds.ApplicationId)

	return &auth24.Credential{
		ApplicationId: &so24appGuid,
		Password:      creds.Password,
		Username:      creds.Username,
	}, nil
}
