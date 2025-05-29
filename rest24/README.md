[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go/rest24)

# 24SevenOffice REST API Client

[Official docs](https://rest-api.developer.24sevenoffice.com/doc/v1/)

## Usage

```go
import "github.com/valuetechdev/24sevenoffice-go/rest24"

func func() {
	c, err := New(&TokenOpts{
		ClientId:       "your-client-id",
		ClientSecret:   "your-client-secret",
		OrganizationId: "your-org-id",
	})
	if err != nil {
	  return nil, err
	}

	res, err := c.GetMeWithResponse(context.Background(), nil)
	if err != nil {
	  return nil, err
	}

	// Use the data
}
```
