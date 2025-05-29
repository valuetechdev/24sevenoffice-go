[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go/soap24)

# 24SevenOffice SOAP API Client

[Official docs](https://developer.24sevenoffice.com/docs/)

## Services currently covered in the package

- [`AccountService`](https://developer.24sevenoffice.com/docs/accountservice.html)
- [`AttachmentService`](https://developer.24sevenoffice.com/docs/attachmentservice.html)
- [`AuthService`](https://developer.24sevenoffice.com/docs/authservice.html)
- [`ClientService`](https://developer.24sevenoffice.com/docs/clientservice.html)
- [`CompanyService`](https://developer.24sevenoffice.com/docs/companyservice.html)
- [`InvoiceService`](https://developer.24sevenoffice.com/docs/invoiceservice.html)
- [`PersonService`](https://developer.24sevenoffice.com/docs/personservice.html)
- [`ProductService`](https://developer.24sevenoffice.com/docs/productservice.html)
- [`ProjectService`](https://developer.24sevenoffice.com/docs/projectservice.html)
- [`TransactionService`](https://developer.24sevenoffice.com/docs/transactionservice.html)

## Usage

```go
import (
	"github.com/valuetechdev/24sevenoffice-go/soap24"
	"github.com/valuetechdev/24sevenoffice-go/soap24/account24"
)

func func() {
   c, err := soap24.NewAuthenticatedClient(&soap24.Credentials{
      ApplicationId: "your-ApplicationId",
      Username:      "your-Username",
      Password:      "your-Password",
   })
   if err != nil {
      return nil, err
   }

   ratesResult, err := c.Account.GetTaxCodeList(&account24.GetTaxCodeList{})
   if err != nil {
      return nil, err
   }

   // Use the data
}
```

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
