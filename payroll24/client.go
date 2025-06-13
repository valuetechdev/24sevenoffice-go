//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/payroll.json

package payroll24

import (
	"fmt"
	"net/http"

	"github.com/valuetechdev/24sevenoffice-go/internal/httpclient"
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

func New(opts Opts) *Payroll24Client {
	payrollClient := &Payroll24Client{httpClient: httpclient.WithRetry(), secret: opts.Secret}
	baseUrl := "https://payroll.24sevenoffice.com/api"
	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(payrollClient.Intercept),
		WithHTTPClient(payrollClient.httpClient))

	if err != nil {
		panic(fmt.Errorf("failed to init client: %w", err))

	}
	payrollClient.ClientWithResponses = c
	return payrollClient
}
