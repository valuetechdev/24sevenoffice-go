# 24SevenOffice api client for Go

This package contains a generatated api client based on the 24SevenOffices SOAP api.

## Services currently covered in the package:

- IvoiceService
- ProjectService
- ProductService
- AuthService
- AccountService

## How to use

To use the api, you need to log in with username, password and applicationId. These are read from as Environment variables.

The names are:

- TWENTYFOURSEVEN_API_APPLICATIONID
- TWENTYFOURSEVEN_API_USERNAME
- TWENTYFOURSEVEN_API_PASSWORD

The program will panic in any are missing.
