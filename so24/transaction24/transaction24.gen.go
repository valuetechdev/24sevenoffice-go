// Code generated by gowsdl DO NOT EDIT.

package transaction24

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type DimensionType string

const (
	DimensionTypeNone DimensionType = "None"

	DimensionTypeProject DimensionType = "Project"

	DimensionTypeDepartment DimensionType = "Department"

	DimensionTypeEmployee DimensionType = "Employee"

	DimensionTypeProduct DimensionType = "Product"

	DimensionTypeCustomer DimensionType = "Customer"

	DimensionTypeCustomerOrderSlip DimensionType = "CustomerOrderSlip"

	DimensionTypeSupplierOrderSlip DimensionType = "SupplierOrderSlip"

	DimensionTypeUserDefined DimensionType = "UserDefined"
)

type DateSearchParameters string

const (
	DateSearchParametersEntryDate DateSearchParameters = "EntryDate"

	DateSearchParametersDateChangedUTC DateSearchParameters = "DateChangedUTC"
)

type TransactionSystemType string

const (
	TransactionSystemTypeInvoiceCustomer TransactionSystemType = "InvoiceCustomer"

	TransactionSystemTypeCreditnoteCustomer TransactionSystemType = "CreditnoteCustomer"

	TransactionSystemTypeDifferenceOutbound TransactionSystemType = "DifferenceOutbound"

	TransactionSystemTypeDisbursment TransactionSystemType = "Disbursment"

	TransactionSystemTypeInvoiceSupplier TransactionSystemType = "InvoiceSupplier"

	TransactionSystemTypeCreditnoteSupplier TransactionSystemType = "CreditnoteSupplier"

	TransactionSystemTypePayment TransactionSystemType = "Payment"

	TransactionSystemTypeDifferenceInbound TransactionSystemType = "DifferenceInbound"

	TransactionSystemTypeMiscellaneous TransactionSystemType = "Miscellaneous"

	TransactionSystemTypeCashSale TransactionSystemType = "CashSale"

	TransactionSystemTypeReminderFee TransactionSystemType = "ReminderFee"

	TransactionSystemTypeReminderNote TransactionSystemType = "ReminderNote"
)

type PeriodType string

const (
	PeriodTypeNone PeriodType = "None"

	PeriodTypeMonth PeriodType = "Month"
)

type GetTransactionTypes struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetTransactionTypes"`
}

type GetTransactionTypesResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetTransactionTypesResponse"`

	GetTransactionTypesResult *ArrayOfTransactionType `xml:"GetTransactionTypesResult,omitempty" json:"GetTransactionTypesResult,omitempty"`
}

type GetUserDefinedDimensions struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetUserDefinedDimensions"`
}

type GetUserDefinedDimensionsResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetUserDefinedDimensionsResponse"`

	GetUserDefinedDimensionsResult *ArrayOfDimension `xml:"GetUserDefinedDimensionsResult,omitempty" json:"GetUserDefinedDimensionsResult,omitempty"`
}

type GetUserDefinedDimensionValues struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetUserDefinedDimensionValues"`

	DimensionType int32 `xml:"DimensionType,omitempty" json:"DimensionType,omitempty"`
}

type GetUserDefinedDimensionValuesResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetUserDefinedDimensionValuesResponse"`

	GetUserDefinedDimensionValuesResult *ArrayOfDimension `xml:"GetUserDefinedDimensionValuesResult,omitempty" json:"GetUserDefinedDimensionValuesResult,omitempty"`
}

type GetFiscalYears struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetFiscalYears"`
}

type GetFiscalYearsResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetFiscalYearsResponse"`

	GetFiscalYearsResult *ArrayOfFiscalYear `xml:"GetFiscalYearsResult,omitempty" json:"GetFiscalYearsResult,omitempty"`
}

type GetTransactions struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetTransactions"`

	SearchParams *TransactionSearchParameters `xml:"searchParams,omitempty" json:"searchParams,omitempty"`
}

type GetTransactionsResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetTransactionsResponse"`

	GetTransactionsResult *ArrayOfTransaction `xml:"GetTransactionsResult,omitempty" json:"GetTransactionsResult,omitempty"`
}

type GetAggregated struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetAggregated"`

	StartDate soap.XSDDateTime `xml:"startDate,omitempty" json:"startDate,omitempty"`

	Dimensions *ArrayOfDimension `xml:"dimensions,omitempty" json:"dimensions,omitempty"`

	AccountNos *ArrayOfShort `xml:"accountNos,omitempty" json:"accountNos,omitempty"`

	PeriodType *PeriodType `xml:"periodType,omitempty" json:"periodType,omitempty"`

	ReturnIB bool `xml:"returnIB,omitempty" json:"returnIB,omitempty"`
}

type GetAggregatedResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetAggregatedResponse"`

	GetAggregatedResult *ArrayOfAggregatedData `xml:"GetAggregatedResult,omitempty" json:"GetAggregatedResult,omitempty"`
}

type ArrayOfTransactionType struct {
	TransactionType []*TransactionType `xml:"TransactionType,omitempty" json:"TransactionType,omitempty"`
}

type TransactionType struct {
	Turnover bool `xml:"Turnover,omitempty" json:"Turnover,omitempty"`

	Title string `xml:"Title,omitempty" json:"Title,omitempty"`

	No int32 `xml:"No,omitempty" json:"No,omitempty"`

	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`
}

type ArrayOfDimension struct {
	Dimension []*Dimension `xml:"Dimension,omitempty" json:"Dimension,omitempty"`
}

type Dimension struct {
	Type *DimensionType `xml:"Type,omitempty" json:"Type,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`

	Percent float64 `xml:"Percent,omitempty" json:"Percent,omitempty"`

	TypeId int32 `xml:"TypeId,omitempty" json:"TypeId,omitempty"`
}

type ArrayOfFiscalYear struct {
	FiscalYear []*FiscalYear `xml:"FiscalYear,omitempty" json:"FiscalYear,omitempty"`
}

type FiscalYear struct {
	Id *Guid `xml:"Id,omitempty" json:"Id,omitempty"`

	StartingDate soap.XSDDateTime `xml:"StartingDate,omitempty" json:"StartingDate,omitempty"`

	EndingDate soap.XSDDateTime `xml:"EndingDate,omitempty" json:"EndingDate,omitempty"`

	IsLocked bool `xml:"IsLocked,omitempty" json:"IsLocked,omitempty"`
}

type TransactionSearchParameters struct {
	DateSearchParameters *DateSearchParameters `xml:"DateSearchParameters,omitempty" json:"DateSearchParameters,omitempty"`

	ShowOpenEntries *bool `xml:"ShowOpenEntries,omitempty" json:"ShowOpenEntries,omitempty"`

	LinkId int32 `xml:"LinkId,omitempty" json:"LinkId,omitempty"`

	DateStart soap.XSDDateTime `xml:"DateStart,omitempty" json:"DateStart,omitempty"`

	DateEnd soap.XSDDateTime `xml:"DateEnd,omitempty" json:"DateEnd,omitempty"`

	ProjectId int32 `xml:"ProjectId,omitempty" json:"ProjectId,omitempty"`

	CustomerId int32 `xml:"CustomerId,omitempty" json:"CustomerId,omitempty"`

	DepartmentId int32 `xml:"DepartmentId,omitempty" json:"DepartmentId,omitempty"`

	AccountNoStart int16 `xml:"AccountNoStart,omitempty" json:"AccountNoStart,omitempty"`

	AccountNoEnd int16 `xml:"AccountNoEnd,omitempty" json:"AccountNoEnd,omitempty"`

	TransactionNoStart int32 `xml:"TransactionNoStart,omitempty" json:"TransactionNoStart,omitempty"`

	TransactionNoEnd int32 `xml:"TransactionNoEnd,omitempty" json:"TransactionNoEnd,omitempty"`

	TransactionTypeId int32 `xml:"TransactionTypeId,omitempty" json:"TransactionTypeId,omitempty"`

	SystemType *TransactionSystemType `xml:"SystemType,omitempty" json:"SystemType,omitempty"`

	RegisteredAfter soap.XSDDateTime `xml:"RegisteredAfter,omitempty" json:"RegisteredAfter,omitempty"`

	HasInvoiceId *bool `xml:"HasInvoiceId,omitempty" json:"HasInvoiceId,omitempty"`

	InvoiceNo string `xml:"InvoiceNo,omitempty" json:"InvoiceNo,omitempty"`

	VatCode *int16 `xml:"VatCode,omitempty" json:"VatCode,omitempty"`
}

type ArrayOfTransaction struct {
	Transaction []*Transaction `xml:"Transaction,omitempty" json:"Transaction,omitempty"`
}

type Transaction struct {
	Date soap.XSDDateTime `xml:"Date,omitempty" json:"Date,omitempty"`

	AccountNo int16 `xml:"AccountNo,omitempty" json:"AccountNo,omitempty"`

	Currency string `xml:"Currency,omitempty" json:"Currency,omitempty"`

	CurrencyRate float64 `xml:"CurrencyRate,omitempty" json:"CurrencyRate,omitempty"`

	CurrencyUnit int16 `xml:"CurrencyUnit,omitempty" json:"CurrencyUnit,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty"`

	StampNo int32 `xml:"StampNo,omitempty" json:"StampNo,omitempty"`

	Period byte `xml:"Period,omitempty" json:"Period,omitempty"`

	TransactionTypeId int32 `xml:"TransactionTypeId,omitempty" json:"TransactionTypeId,omitempty"`

	Comment string `xml:"Comment,omitempty" json:"Comment,omitempty"`

	TransactionNo int32 `xml:"TransactionNo,omitempty" json:"TransactionNo,omitempty"`

	VatCode int16 `xml:"VatCode,omitempty" json:"VatCode,omitempty"`

	Id *Guid `xml:"Id,omitempty" json:"Id,omitempty"`

	LinkId int32 `xml:"LinkId,omitempty" json:"LinkId,omitempty"`

	InvoiceNo string `xml:"InvoiceNo,omitempty" json:"InvoiceNo,omitempty"`

	SequenceNo int16 `xml:"SequenceNo,omitempty" json:"SequenceNo,omitempty"`

	SystemType *TransactionSystemType `xml:"SystemType,omitempty" json:"SystemType,omitempty"`

	DueDate soap.XSDDateTime `xml:"DueDate,omitempty" json:"DueDate,omitempty"`

	Dimensions *ArrayOfDimension `xml:"Dimensions,omitempty" json:"Dimensions,omitempty"`

	RegistrationDate soap.XSDDateTime `xml:"RegistrationDate,omitempty" json:"RegistrationDate,omitempty"`

	DateChanged soap.XSDDateTime `xml:"DateChanged,omitempty" json:"DateChanged,omitempty"`

	Hidden bool `xml:"Hidden,omitempty" json:"Hidden,omitempty"`

	Open bool `xml:"Open,omitempty" json:"Open,omitempty"`

	OCR string `xml:"OCR,omitempty" json:"OCR,omitempty"`

	VatDividend float64 `xml:"VatDividend,omitempty" json:"VatDividend,omitempty"`

	HasVatDividend bool `xml:"HasVatDividend,omitempty" json:"HasVatDividend,omitempty"`

	PeriodDate soap.XSDDate `xml:"PeriodDate,omitempty" json:"PeriodDate,omitempty"`
}

type ArrayOfShort struct {
	Short []int16 `xml:"short,omitempty" json:"short,omitempty"`
}

type ArrayOfAggregatedData struct {
	AggregatedData []*AggregatedData `xml:"AggregatedData,omitempty" json:"AggregatedData,omitempty"`
}

type AggregatedData struct {
	IncomingBalance float64 `xml:"IncomingBalance,omitempty" json:"IncomingBalance,omitempty"`

	TotalBalance float64 `xml:"TotalBalance,omitempty" json:"TotalBalance,omitempty"`

	Periods *ArrayOfKeyValuePair `xml:"Periods,omitempty" json:"Periods,omitempty"`

	Dimension *Dimension `xml:"Dimension,omitempty" json:"Dimension,omitempty"`

	AccountNo int16 `xml:"AccountNo,omitempty" json:"AccountNo,omitempty"`
}

type ArrayOfKeyValuePair struct {
	KeyValuePair []*KeyValuePair `xml:"KeyValuePair,omitempty" json:"KeyValuePair,omitempty"`
}

type KeyValuePair struct {
	Key string `xml:"Key,omitempty" json:"Key,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type Guid string

type TransactionServiceSoap interface {
	GetTransactionTypes(request *GetTransactionTypes) (*GetTransactionTypesResponse, error)

	GetTransactionTypesContext(ctx context.Context, request *GetTransactionTypes) (*GetTransactionTypesResponse, error)

	GetUserDefinedDimensions(request *GetUserDefinedDimensions) (*GetUserDefinedDimensionsResponse, error)

	GetUserDefinedDimensionsContext(ctx context.Context, request *GetUserDefinedDimensions) (*GetUserDefinedDimensionsResponse, error)

	GetUserDefinedDimensionValues(request *GetUserDefinedDimensionValues) (*GetUserDefinedDimensionValuesResponse, error)

	GetUserDefinedDimensionValuesContext(ctx context.Context, request *GetUserDefinedDimensionValues) (*GetUserDefinedDimensionValuesResponse, error)

	/* A list of fiscal years and their boundaries. A fiscal year is starting at 'StartingDate' up to 'EndingDate', do not include the ending date (eg.: StartingDate >= [FiscalYear] < EndingDate) */
	GetFiscalYears(request *GetFiscalYears) (*GetFiscalYearsResponse, error)

	GetFiscalYearsContext(ctx context.Context, request *GetFiscalYears) (*GetFiscalYearsResponse, error)

	GetTransactions(request *GetTransactions) (*GetTransactionsResponse, error)

	GetTransactionsContext(ctx context.Context, request *GetTransactions) (*GetTransactionsResponse, error)

	GetAggregated(request *GetAggregated) (*GetAggregatedResponse, error)

	GetAggregatedContext(ctx context.Context, request *GetAggregated) (*GetAggregatedResponse, error)
}

type transactionServiceSoap struct {
	client *soap.Client
}

func NewTransactionServiceSoap(client *soap.Client) TransactionServiceSoap {
	return &transactionServiceSoap{
		client: client,
	}
}

func (service *transactionServiceSoap) GetTransactionTypesContext(ctx context.Context, request *GetTransactionTypes) (*GetTransactionTypesResponse, error) {
	response := new(GetTransactionTypesResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetTransactionTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetTransactionTypes(request *GetTransactionTypes) (*GetTransactionTypesResponse, error) {
	return service.GetTransactionTypesContext(
		context.Background(),
		request,
	)
}

func (service *transactionServiceSoap) GetUserDefinedDimensionsContext(ctx context.Context, request *GetUserDefinedDimensions) (*GetUserDefinedDimensionsResponse, error) {
	response := new(GetUserDefinedDimensionsResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetUserDefinedDimensions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetUserDefinedDimensions(request *GetUserDefinedDimensions) (*GetUserDefinedDimensionsResponse, error) {
	return service.GetUserDefinedDimensionsContext(
		context.Background(),
		request,
	)
}

func (service *transactionServiceSoap) GetUserDefinedDimensionValuesContext(ctx context.Context, request *GetUserDefinedDimensionValues) (*GetUserDefinedDimensionValuesResponse, error) {
	response := new(GetUserDefinedDimensionValuesResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetUserDefinedDimensionValues", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetUserDefinedDimensionValues(request *GetUserDefinedDimensionValues) (*GetUserDefinedDimensionValuesResponse, error) {
	return service.GetUserDefinedDimensionValuesContext(
		context.Background(),
		request,
	)
}

func (service *transactionServiceSoap) GetFiscalYearsContext(ctx context.Context, request *GetFiscalYears) (*GetFiscalYearsResponse, error) {
	response := new(GetFiscalYearsResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetFiscalYears", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetFiscalYears(request *GetFiscalYears) (*GetFiscalYearsResponse, error) {
	return service.GetFiscalYearsContext(
		context.Background(),
		request,
	)
}

func (service *transactionServiceSoap) GetTransactionsContext(ctx context.Context, request *GetTransactions) (*GetTransactionsResponse, error) {
	response := new(GetTransactionsResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetTransactions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetTransactions(request *GetTransactions) (*GetTransactionsResponse, error) {
	return service.GetTransactionsContext(
		context.Background(),
		request,
	)
}

func (service *transactionServiceSoap) GetAggregatedContext(ctx context.Context, request *GetAggregated) (*GetAggregatedResponse, error) {
	response := new(GetAggregatedResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetAggregated", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *transactionServiceSoap) GetAggregated(request *GetAggregated) (*GetAggregatedResponse, error) {
	return service.GetAggregatedContext(
		context.Background(),
		request,
	)
}
