[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go/payroll24)

# 24Payroll API Client

[Official docs](https://swagger.api.24sevenoffice.com/?url=https://me.24sevenoffice.com/swagger.json)

## Usage

```go
import "github.com/valuetechdev/24sevenoffice-go/payroll24"

func yourFunc() (any, error) {
   c := payroll24.New("your-secret")
   if err := c.Authenticate(); err != nil {
      return nil, err
   }

	 absenceRes, err := c.GetAbsenceV2WithResponse(context.TODO())
   if err != nil {
      return nil, err
   }

   // Use the data
}
```

## Things to know

- The API is using OpenAPI/Swagger 2.0.
  - It has been converted to OpenAPI 3 using `npx swagger2openapi -o <openapi>.json <swagger2>.json`
  - It has been altered to with defined schemas/models that was missing.
