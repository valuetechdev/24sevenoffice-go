//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config payroll.yml ../../api/openapi/payroll.json

package payroll24

import "github.com/valuetechdev/api-client-24so/so24/payroll24/bearer"

type PayrollService struct {
	Client *ClientWithResponses
}

func New() (*PayrollService, error) {
	bt, err := bearer.New()
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
	)

	if err != nil {
		return nil, err
	}

	return &PayrollService{c}, nil
}
