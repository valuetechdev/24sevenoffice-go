//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config payroll.yml ../../api/openapi/payroll.json

package payroll24

import (
	"net"
	"net/http"
	"time"

	"github.com/valuetechdev/24sevenoffice-go/so24/payroll24/bearer"
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
		WithHTTPClient(&http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   60 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 0,
			}}))

	if err != nil {
		return nil, err
	}

	return &PayrollService{c}, nil
}
