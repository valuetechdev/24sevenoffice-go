# 24SevenOffice API client for Go

This package contains a generated API client based on the 24SevenOffice's SOAP API.

## Services currently covered in the package:

- `AccountService`
- `AuthService`
- `ClientService`
- `CompanyService`
- `InvoiceService`
- `PersonService`
- `ProductService`
- `ProjectService`

## Usage

```bash
go get github.com/valuetechdev/api-client-24so
```

Go example

```go
import (
	"github.com/valuetechdev/api-client-24so/so24"
	"github.com/valuetechdev/api-client-24so/so24/account24"
)

func func() {
   so24Client, err := so24.NewAuthenticatedClient()
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

To use the API, you need to log in with `username`, `password` and
`applicationId`. These are read from as environment variables.

The names are:

- `TWENTYFOURSEVEN_API_APPLICATIONID`
- `TWENTYFOURSEVEN_API_USERNAME`
- `TWENTYFOURSEVEN_API_PASSWORD`

The program will panic in any are missing.

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

## Adding or updating services

1. Add or replace the `.wsdl`-file in `./wsdl/so24`-directory, follow the same
   naming-convention as 24SevenOffice.
1. Add the new service in `clientsToGenerate` in `generate24.go`.
1. Add the new service in `Client` struct in `client.go`.
1. Run `make generate`.
1. Open a new PR.
