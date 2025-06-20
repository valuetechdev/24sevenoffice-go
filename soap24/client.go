//go:generate go run generate24.go

package soap24

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hooklift/gowsdl/soap"
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
	accountURL     = "https://api.24sevenoffice.com/Economy/Account/V004/Accountservice.asmx"
	attachmentURL  = "https://webservices.24sevenoffice.com/Economy/Accounting/Accounting_V001/AttachmentService.asmx"
	authURL        = "https://api.24sevenoffice.com/authenticate/v001/authenticate.asmx"
	clientURL      = "https://api.24sevenoffice.com/Client/V001/ClientService.asmx"
	companyURL     = "https://api.24sevenoffice.com/CRM/Company/V001/CompanyService.asmx"
	invoiceURL     = "https://api.24sevenoffice.com/Economy/InvoiceOrder/V001/InvoiceService.asmx"
	personURL      = "https://webservices.24sevenoffice.com/CRM/Contact/PersonService.asmx"
	productURL     = "https://api.24sevenoffice.com/Logistics/Product/V001/ProductService.asmx"
	projectURl     = "https://webservices.24sevenoffice.com/Project/V001/ProjectService.asmx"
	transactionURL = "https://api.24sevenoffice.com/Economy/Accounting/V001/TransactionService.asmx"
)

type Client struct {
	sessionId   string
	headers     map[string]string
	credentials *auth24.Credential
	httpClient  *http.Client

	Account     account24.AccountServiceSoap
	Attachment  attachment24.AttachmentServiceSoap
	Auth        auth24.AuthenticateSoap
	Client      client24.ClientServiceSoap
	Company     company24.CompanyServiceSoap
	Invoice     invoice24.InvoiceServiceSoap
	Person      person24.PersonServiceSoap
	Product     product24.ProductServiceSoap
	Project     project24.ProjectServiceSoap
	Transaction transaction24.TransactionServiceSoap
}

type Option func(*Client)

// WithHttpClient sets a custom http.Client. Defaults to [http.DefaultClient].
func WithHttpClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

// authInterceptor wraps an http.RoundTripper to automatically check auth before requests
type authInterceptor struct {
	client    *Client
	transport http.RoundTripper
}

func (a *authInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	// Skip auth check for auth service URLs to avoid recursion
	if !strings.Contains(req.URL.String(), "authenticate") {
		if err := a.client.CheckAuth(); err != nil {
			return nil, err
		}
	}

	// Add current headers to the request
	for key, value := range a.client.headers {
		req.Header.Set(key, value)
	}

	return a.transport.RoundTrip(req)
}

// Returns new [Client].
//
// You can reuse an already generated sessionId by using [Client.SetSessionId]
func New(credentials auth24.Credential, options ...Option) *Client {
	headers := map[string]string{}
	client := &Client{
		credentials: &credentials,
		headers:     headers,
		httpClient:  http.DefaultClient,
	}
	for _, option := range options {
		option(client)
	}

	// Wrap the httpClient transport with auth interceptor
	interceptedClient := &http.Client{
		Transport: &authInterceptor{
			client:    client,
			transport: client.httpClient.Transport,
		},
		Timeout: client.httpClient.Timeout,
	}
	if interceptedClient.Transport.(*authInterceptor).transport == nil {
		interceptedClient.Transport.(*authInterceptor).transport = http.DefaultTransport
	}
	client.httpClient = interceptedClient

	authService := auth24.NewAuthenticateSoap(newSoapClient(authURL, client))
	invoiceService := invoice24.NewInvoiceServiceSoap(newSoapClient(invoiceURL, client))
	productService := product24.NewProductServiceSoap(newSoapClient(productURL, client))
	accountService := account24.NewAccountServiceSoap(newSoapClient(accountURL, client))
	companyService := company24.NewCompanyServiceSoap(newSoapClient(companyURL, client))
	clientService := client24.NewClientServiceSoap(newSoapClient(clientURL, client))
	personService := person24.NewPersonServiceSoap(newSoapClient(personURL, client))
	projectService := project24.NewProjectServiceSoap(newSoapClient(projectURl, client))
	transactionService := transaction24.NewTransactionServiceSoap(newSoapClient(transactionURL, client))
	attachmentService := attachment24.NewAttachmentServiceSoap(newSoapClient(attachmentURL, client))

	client.Account = accountService
	client.Attachment = attachmentService
	client.Auth = authService
	client.Client = clientService
	client.Company = companyService
	client.Invoice = invoiceService
	client.Person = personService
	client.Product = productService
	client.Project = projectService
	client.Transaction = transactionService

	return client
}

func newSoapClient(url string, client *Client) *soap.Client {
	return soap.NewClient(url, soap.WithHTTPClient(client.httpClient))
}

// Returns sessionId
func (c *Client) GetSessionId() string {
	return c.sessionId
}

// Sets sessionId
//
// And sets headers with "cookie: ASP.NET_SessionId=<sessionId>".
func (c *Client) SetSessionId(sessionId string) {
	c.headers["cookie"] = fmt.Sprintf("ASP.NET_SessionId=%s", sessionId)
	c.sessionId = sessionId
}

// Returns true if sessionId is valid
func (c *Client) IsSessionIdValid() bool {
	r, err := c.Auth.HasSession(&auth24.HasSession{})
	if err != nil {
		return false
	}

	return r.HasSessionResult
}

// Checks if auth is valid.
//
// Reauthenticates if auth is invalid.
func (c *Client) CheckAuth() error {
	if c.IsSessionIdValid() {
		return nil
	}
	res, err := c.Auth.Login(&auth24.Login{
		Credential: c.credentials,
	})
	if err != nil {
		return fmt.Errorf("soap24: failed to login: %w", err)
	}

	sessionId := res.LoginResult
	c.SetSessionId(sessionId)

	return nil
}
