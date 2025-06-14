//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/rest.json

package rest24

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Rest24Client struct {
	token      *oauth2.Token
	conf       *clientcredentials.Config
	confURL    string
	httpClient *http.Client
	*ClientWithResponses
}

type Credentials struct {
	ClientId       string
	ClientSecret   string
	OrganizationId string
}

type Option func(*Rest24Client)

// WithHttpClient sets a custom http.Client. Defaults to [http.DefaultClient].
func WithHttpClient(client *http.Client) Option {
	return func(c *Rest24Client) {
		c.httpClient = client
	}
}

// Returns new [Rest24Client].
//
// You can reuse an already generated token and have it revalidated if it has
// expired, by using [Rest24Client.SetToken].
//
// You can provide options to customize the client behavior.
//
// WARNING: The client hasn't been tested.
func New(credentials *Credentials, options ...Option) *Rest24Client {
	conf := &clientcredentials.Config{
		ClientID:     credentials.ClientId,
		ClientSecret: credentials.ClientSecret,
		TokenURL:     "https://login.24sevenoffice.com/oauth/token",
		EndpointParams: map[string][]string{
			"login_organization": {credentials.OrganizationId},
			"audience":           {"https://api.24sevenoffice.com"},
		},
	}

	baseUrl := "https://rest.api.24sevenoffice.com/v1"
	client := &Rest24Client{conf: conf, httpClient: http.DefaultClient}

	for _, option := range options {
		option(client)
	}

	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(client.InterceptToken),
		WithHTTPClient(client.httpClient))
	if err != nil {
		panic(fmt.Errorf("failed to init client: %w", err))
	}
	client.ClientWithResponses = c
	return client
}
