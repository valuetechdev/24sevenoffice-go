//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../api/openapi/rest.json

package rest24

import (
	"github.com/valuetechdev/24sevenoffice-go/internal/httpclient"
)

func New(opts *TokenOpts) (*ClientWithResponses, error) {
	bt, err := newToken(opts)
	if err != nil {
		return nil, err
	}
	if bt == nil {
		return nil, nil
	}
	baseUrl := "https://rest.api.24sevenoffice.com/v1"
	c, err := NewClientWithResponses(
		baseUrl,
		WithRequestEditorFn(bt.Intercept),
		WithHTTPClient(httpclient.WithRetry()))

	if err != nil {
		return nil, err
	}

	return c, nil
}
