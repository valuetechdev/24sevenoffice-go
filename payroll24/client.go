//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/payroll.json

package payroll24

import (
	"fmt"
	"net/http"
)

type Payroll24Client struct {
	token      *Token
	httpClient *http.Client
	secret     string
	*ClientWithResponses
}

type Opts struct {
	Secret string
}

type Option func(*Payroll24Client)

// WithHttpClient sets a custom http.Client. Defaults to [http.DefaultClient].
func WithHttpClient(client *http.Client) Option {
	return func(c *Payroll24Client) {
		c.httpClient = client
	}
}

// Returns new [Payroll24Client].
//
// You can reuse an already generated token and have it revalidated if it has
// expired, by using [Payroll24Client.SetToken].
//
// You can provide options to customize the client behavior.
func New(secret string, options ...Option) *Payroll24Client {
	client := &Payroll24Client{httpClient: http.DefaultClient, secret: secret}
	for _, option := range options {
		option(client)
	}
	baseUrl := "https://payroll.24sevenoffice.com/api"
	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(client.Intercept),
		WithHTTPClient(client.httpClient))

	if err != nil {
		panic(fmt.Errorf("failed to init client: %w", err))
	}
	client.ClientWithResponses = c
	return client
}
