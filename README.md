[![go reference badge](https://pkg.go.dev/badge/github.com/valuetechdev/24sevenoffice-go.svg)](https://pkg.go.dev/github.com/valuetechdev/24sevenoffice-go)

# 24SevenOffice API client for Go

This package contains a generated API clients for 24SevenOffice's APIs.

- [`soap24`](soap24/README.md) ([API](https://developer.24sevenoffice.com/docs/))
- [`payroll24`](payroll24/README.md) ([API](https://swagger.api.24sevenoffice.com/?url=https://me.24sevenoffice.com/swagger.json))
- [`rest24`](rest24/README.md) ([API](https://rest-api.developer.24sevenoffice.com/doc/v1/))

## Usage

```bash
go get github.com/valuetechdev/24sevenoffice-go
```

## Things to know

- The payroll-API is using OpenAPI/Swagger 2.0.
  - It has been converted to OpenAPI 3 using `npx swagger2openapi -o <openapi>.json <swagger2>.json`
  - It has been altered to with defined schemas/models that was missing.
