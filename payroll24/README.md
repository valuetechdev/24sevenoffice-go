[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go/payroll24)

# 24Payroll API Client

[Official docs](https://swagger.api.24sevenoffice.com/?url=https://me.24sevenoffice.com/swagger.json)

## Usage

```go
import "github.com/valuetechdev/24sevenoffice-go/payroll24"

func func() {
   c, err := payroll24.NewAuthenticatedClient("your-api-token")
   if err != nil {
      return nil, err
   }

	 absenceRes, err := c.GetAbsenceV2WithResponse(context.TODO())
   if err != nil {
      return nil, err
   }

   // Use the data
}
```
