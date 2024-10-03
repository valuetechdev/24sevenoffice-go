# 24SevenOffice API client for Go

This package contains a generated API client based on the 24SevenOffices SOAP API.

## Services currently covered in the package:

- `AccountService`
- `AuthService`
- `ClientService`
- `CompanyService`
- `InvoiceService`
- `ProductService`
- `ProjectService`

## How to use

To use the API, you need to log in with `username`, `password` and `applicationId`. These are read from as environment variables.

The names are:

- `TWENTYFOURSEVEN_API_APPLICATIONID`
- `TWENTYFOURSEVEN_API_USERNAME`
- `TWENTYFOURSEVEN_API_PASSWORD`

The program will panic in any are missing.

Usage is described in [24SevenOffices api documentation](https://developer.24sevenoffice.com/docs/)
