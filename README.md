[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go)

# 24SevenOffice API client for Go

This package contains a generated API client based on the 24SevenOffice's SOAP
API and payroll Swagger API.

## Services currently covered in the package

- [`AccountService`](https://developer.24sevenoffice.com/docs/accountservice.html)
- [`AttachmentService`](https://developer.24sevenoffice.com/docs/attachmentservice.html)
- [`AuthService`](https://developer.24sevenoffice.com/docs/authservice.html)
- [`ClientService`](https://developer.24sevenoffice.com/docs/clientservice.html)
- [`CompanyService`](https://developer.24sevenoffice.com/docs/companyservice.html)
- [`InvoiceService`](https://developer.24sevenoffice.com/docs/invoiceservice.html)
- [`Payroll`](https://swagger.api.24sevenoffice.com/?url=https://me.24sevenoffice.com/swagger.json)
- [`PersonService`](https://developer.24sevenoffice.com/docs/personservice.html)
- [`ProductService`](https://developer.24sevenoffice.com/docs/productservice.html)
- [`ProjectService`](https://developer.24sevenoffice.com/docs/projectservice.html)
- [`TransactionService`](https://developer.24sevenoffice.com/docs/transactionservice.html)

## Usage

```bash
go get github.com/valuetechdev/24sevenoffice-go
```

Go example

```go
import (
	"github.com/valuetechdev/24sevenoffice-go/so24"
	"github.com/valuetechdev/24sevenoffice-go/so24/account24"
)

func func() {
   so24Client, err := so24.NewAuthenticatedClient(&so24.Credentials{
      ApplicationId: "your-ApplicationId",
      Username:      "your-Username",
      Password:      "your-Password",
      PayrollAPI:    "your-PayrollAPI",
   })
   if err != nil {
      return nil, err
   }

   so24RatesResult, err := so24Client.Account.GetTaxCodeList(&account24.GetTaxCodeList{})
   if err != nil {
      return nil, err
   }

   // Use the data
}
```

Usage is described in [24SevenOffice's API documentation](https://developer.24sevenoffice.com/docs/)

### Changing identity

By default the client uses your accounts default identity (the one you log in
to automatically in the UI).

```go
id := auth24.Guid("new-identity")
_, err = so24Client.Auth.SetIdentityById(&auth24.SetIdentityById{IdentityId: &id})
if err != nil {
   return nil, err
}
```

## SOAP: Adding or updating services

1. Add or replace the `.wsdl`-file in `./wsdl/so24`-directory, follow the same
   naming-convention as 24SevenOffice.
1. Add the new service in `clientsToGenerate` in `generate24.go`.
1. Add the new service in `Client` struct in `client.go`.
1. Run `make generate`.
1. Open a new PR.

## Things to know

- Only the payroll-API is using OpenAPI/Swagger.
  - It's originally based on Swagger 2.0, but it was converted to OpenAPI 3
    using `npx swagger2openapi -o <openapi>.json <swagger2>.json`
  - It has been altered to with defined schemas/models that was missing.
