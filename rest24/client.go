//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/rest.json

package rest24

import (
	"fmt"
	"net/http"

	"github.com/valuetechdev/24sevenoffice-go/internal/httpclient"
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

type Opts struct {
	ClientId       string
	ClientSecret   string
	OrganizationId string
}

func New(opts *Opts) *Rest24Client {
	conf := &clientcredentials.Config{
		ClientID:     opts.ClientId,
		ClientSecret: opts.ClientSecret,
		TokenURL:     "https://login.24sevenoffice.com/oauth/token",
		EndpointParams: map[string][]string{
			"login_organization": {opts.OrganizationId},
			"audience":           {"https://api.24sevenoffice.com"},
		},
	}

	baseUrl := "https://rest.api.24sevenoffice.com/v1"
	restClient := &Rest24Client{conf: conf, httpClient: httpclient.WithRetry()}

	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(restClient.InterceptToken),
		WithHTTPClient(restClient.httpClient))
	if err != nil {
		panic(fmt.Errorf("failed to init client: %w", err))
	}
	restClient.ClientWithResponses = c
	return restClient
}
