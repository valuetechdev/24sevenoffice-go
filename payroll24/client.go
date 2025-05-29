//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/payroll.json

package payroll24

import (
	"github.com/valuetechdev/24sevenoffice-go/internal/httpclient"
	"github.com/valuetechdev/24sevenoffice-go/payroll24/bearer"
)

type PayrollService struct {
	Client *ClientWithResponses
}

func New(apiToken string) (*PayrollService, error) {
	bt, err := bearer.New(apiToken)
	if err != nil {
		return nil, err
	}
	if bt == nil {
		return nil, nil
	}
	baseUrl := "https://payroll.24sevenoffice.com/api"
	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(bt.Intercept),
		WithHTTPClient(httpclient.WithRetry()))

	if err != nil {
		return nil, err
	}

	return &PayrollService{c}, nil
}
