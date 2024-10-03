// Code generated by gowsdl DO NOT EDIT.

package client24

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

type FaxNumberType string

const (
	FaxNumberTypeUnknown FaxNumberType = "Unknown"

	FaxNumberTypeWork FaxNumberType = "Work"
)

type TypeGroupModule string

const (
	TypeGroupModuleAll TypeGroupModule = "All"

	TypeGroupModuleSale TypeGroupModule = "Sale"

	TypeGroupModulePurchase TypeGroupModule = "Purchase"
)

type GetUsers struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetUsers"`
}

type GetUsersResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetUsersResponse"`

	GetUsersResult *ArrayOfUser `xml:"GetUsersResult,omitempty" json:"GetUsersResult,omitempty"`
}

type GetClientInformation struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetClientInformation"`
}

type GetClientInformationResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetClientInformationResponse"`

	GetClientInformationResult *Client `xml:"GetClientInformationResult,omitempty" json:"GetClientInformationResult,omitempty"`
}

type GetDepartmentList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDepartmentList"`
}

type GetDepartmentListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDepartmentListResponse"`

	GetDepartmentListResult *ArrayOfDepartment `xml:"GetDepartmentListResult,omitempty" json:"GetDepartmentListResult,omitempty"`
}

type GetCurrencyList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetCurrencyList"`
}

type GetCurrencyListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetCurrencyListResponse"`

	GetCurrencyListResult *ArrayOfCurrency `xml:"GetCurrencyListResult,omitempty" json:"GetCurrencyListResult,omitempty"`
}

type GetTypeGroupList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetTypeGroupList"`

	Module *TypeGroupModule `xml:"module,omitempty" json:"module,omitempty"`
}

type GetTypeGroupListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetTypeGroupListResponse"`

	GetTypeGroupListResult *ArrayOfAccountingGroup `xml:"GetTypeGroupListResult,omitempty" json:"GetTypeGroupListResult,omitempty"`
}

type GetVatTypeList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetVatTypeList"`
}

type GetVatTypeListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetVatTypeListResponse"`

	GetVatTypeListResult *ArrayOfVatType `xml:"GetVatTypeListResult,omitempty" json:"GetVatTypeListResult,omitempty"`
}

type ArrayOfUser struct {
	User []*User `xml:"User,omitempty" json:"User,omitempty"`
}

type User struct {
	ContactId int32 `xml:"ContactId,omitempty" json:"ContactId,omitempty"`

	Id *Guid `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	EmployeeId int32 `xml:"EmployeeId,omitempty" json:"EmployeeId,omitempty"`
}

type Client struct {
	AddressList *Addresses `xml:"AddressList,omitempty" json:"AddressList,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty"`

	EmailAddressList *EmailAddresses `xml:"EmailAddressList,omitempty" json:"EmailAddressList,omitempty"`

	FactoringClientNo string `xml:"FactoringClientNo,omitempty" json:"FactoringClientNo,omitempty"`

	FactoringText string `xml:"FactoringText,omitempty" json:"FactoringText,omitempty"`

	FaxNumberList *ArrayOfFaxNumber `xml:"FaxNumberList,omitempty" json:"FaxNumberList,omitempty"`

	IsUsingFactoring bool `xml:"IsUsingFactoring,omitempty" json:"IsUsingFactoring,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	PhoneNumberList *PhoneNumbers `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`

	ReminderDays int32 `xml:"ReminderDays,omitempty" json:"ReminderDays,omitempty"`

	ReminderDays2 int32 `xml:"ReminderDays2,omitempty" json:"ReminderDays2,omitempty"`

	ReminderDays3 int32 `xml:"ReminderDays3,omitempty" json:"ReminderDays3,omitempty"`

	UserId int64 `xml:"UserId,omitempty" json:"UserId,omitempty"`

	BankAccount string `xml:"BankAccount,omitempty" json:"BankAccount,omitempty"`

	IBAN string `xml:"IBAN,omitempty" json:"IBAN,omitempty"`

	OrganizationNumber string `xml:"OrganizationNumber,omitempty" json:"OrganizationNumber,omitempty"`

	Swift string `xml:"Swift,omitempty" json:"Swift,omitempty"`

	ResellerId int64 `xml:"ResellerId,omitempty" json:"ResellerId,omitempty"`

	ResellerName string `xml:"ResellerName,omitempty" json:"ResellerName,omitempty"`

	DefaultCurrency string `xml:"DefaultCurrency,omitempty" json:"DefaultCurrency,omitempty"`
}

type Addresses struct {
	Post *Address `xml:"Post,omitempty" json:"Post,omitempty"`

	Delivery *Address `xml:"Delivery,omitempty" json:"Delivery,omitempty"`

	Visit *Address `xml:"Visit,omitempty" json:"Visit,omitempty"`

	Invoice *Address `xml:"Invoice,omitempty" json:"Invoice,omitempty"`
}

type Address struct {
	Street string `xml:"Street,omitempty" json:"Street,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty" json:"PostalCode,omitempty"`

	PostalArea string `xml:"PostalArea,omitempty" json:"PostalArea,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty"`
}

type EmailAddresses struct {
	Home *EmailAddress `xml:"Home,omitempty" json:"Home,omitempty"`

	Invoice *EmailAddress `xml:"Invoice,omitempty" json:"Invoice,omitempty"`

	Primary *EmailAddress `xml:"Primary,omitempty" json:"Primary,omitempty"`

	Work *EmailAddress `xml:"Work,omitempty" json:"Work,omitempty"`

	Alternative *EmailAddress `xml:"Alternative,omitempty" json:"Alternative,omitempty"`
}

type EmailAddress struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ArrayOfFaxNumber struct {
	FaxNumber []*FaxNumber `xml:"FaxNumber,omitempty" json:"FaxNumber,omitempty"`
}

type FaxNumber struct {
	Type *FaxNumberType `xml:"Type,omitempty" json:"Type,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type PhoneNumbers struct {
	Home *PhoneNumber `xml:"Home,omitempty" json:"Home,omitempty"`

	Fax *PhoneNumber `xml:"Fax,omitempty" json:"Fax,omitempty"`

	Mobile *PhoneNumber `xml:"Mobile,omitempty" json:"Mobile,omitempty"`

	Primary *PhoneNumber `xml:"Primary,omitempty" json:"Primary,omitempty"`

	Work *PhoneNumber `xml:"Work,omitempty" json:"Work,omitempty"`
}

type PhoneNumber struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ArrayOfDepartment struct {
	Department []*Department `xml:"Department,omitempty" json:"Department,omitempty"`
}

type Department struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Addresses *Addresses `xml:"Addresses,omitempty" json:"Addresses,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty"`
}

type ArrayOfCurrency struct {
	Currency []*Currency `xml:"Currency,omitempty" json:"Currency,omitempty"`
}

type Currency struct {
	Symbol string `xml:"Symbol,omitempty" json:"Symbol,omitempty"`

	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Rate float64 `xml:"Rate,omitempty" json:"Rate,omitempty"`

	Unit int32 `xml:"Unit,omitempty" json:"Unit,omitempty"`
}

type ArrayOfAccountingGroup struct {
	AccountingGroup []*AccountingGroup `xml:"AccountingGroup,omitempty" json:"AccountingGroup,omitempty"`
}

type AccountingGroup struct {
	AccountId int32 `xml:"AccountId,omitempty" json:"AccountId,omitempty"`

	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`
}

type ArrayOfVatType struct {
	VatType []*VatType `xml:"VatType,omitempty" json:"VatType,omitempty"`
}

type VatType struct {
	No int32 `xml:"No,omitempty" json:"No,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Rate float64 `xml:"Rate,omitempty" json:"Rate,omitempty"`
}

type Guid string

type ClientServiceSoap interface {

	/* Returns a list of users of the currently logged in client. */
	GetUsers(request *GetUsers) (*GetUsersResponse, error)

	GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error)

	/* Get 24SevenOffice's client data for the client logged in */
	GetClientInformation(request *GetClientInformation) (*GetClientInformationResponse, error)

	GetClientInformationContext(ctx context.Context, request *GetClientInformation) (*GetClientInformationResponse, error)

	/* Returns the list of departments */
	GetDepartmentList(request *GetDepartmentList) (*GetDepartmentListResponse, error)

	GetDepartmentListContext(ctx context.Context, request *GetDepartmentList) (*GetDepartmentListResponse, error)

	/* Gets a list of the client's currencies */
	GetCurrencyList(request *GetCurrencyList) (*GetCurrencyListResponse, error)

	GetCurrencyListContext(ctx context.Context, request *GetCurrencyList) (*GetCurrencyListResponse, error)

	/* Gets a list of the client's typegroups */
	GetTypeGroupList(request *GetTypeGroupList) (*GetTypeGroupListResponse, error)

	GetTypeGroupListContext(ctx context.Context, request *GetTypeGroupList) (*GetTypeGroupListResponse, error)

	/* Gets a list of the client's VAT Types */
	GetVatTypeList(request *GetVatTypeList) (*GetVatTypeListResponse, error)

	GetVatTypeListContext(ctx context.Context, request *GetVatTypeList) (*GetVatTypeListResponse, error)
}

type clientServiceSoap struct {
	client *soap.Client
}

func NewClientServiceSoap(client *soap.Client) ClientServiceSoap {
	return &clientServiceSoap{
		client: client,
	}
}

func (service *clientServiceSoap) GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error) {
	response := new(GetUsersResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetUsers(request *GetUsers) (*GetUsersResponse, error) {
	return service.GetUsersContext(
		context.Background(),
		request,
	)
}

func (service *clientServiceSoap) GetClientInformationContext(ctx context.Context, request *GetClientInformation) (*GetClientInformationResponse, error) {
	response := new(GetClientInformationResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetClientInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetClientInformation(request *GetClientInformation) (*GetClientInformationResponse, error) {
	return service.GetClientInformationContext(
		context.Background(),
		request,
	)
}

func (service *clientServiceSoap) GetDepartmentListContext(ctx context.Context, request *GetDepartmentList) (*GetDepartmentListResponse, error) {
	response := new(GetDepartmentListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetDepartmentList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetDepartmentList(request *GetDepartmentList) (*GetDepartmentListResponse, error) {
	return service.GetDepartmentListContext(
		context.Background(),
		request,
	)
}

func (service *clientServiceSoap) GetCurrencyListContext(ctx context.Context, request *GetCurrencyList) (*GetCurrencyListResponse, error) {
	response := new(GetCurrencyListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetCurrencyList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetCurrencyList(request *GetCurrencyList) (*GetCurrencyListResponse, error) {
	return service.GetCurrencyListContext(
		context.Background(),
		request,
	)
}

func (service *clientServiceSoap) GetTypeGroupListContext(ctx context.Context, request *GetTypeGroupList) (*GetTypeGroupListResponse, error) {
	response := new(GetTypeGroupListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetTypeGroupList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetTypeGroupList(request *GetTypeGroupList) (*GetTypeGroupListResponse, error) {
	return service.GetTypeGroupListContext(
		context.Background(),
		request,
	)
}

func (service *clientServiceSoap) GetVatTypeListContext(ctx context.Context, request *GetVatTypeList) (*GetVatTypeListResponse, error) {
	response := new(GetVatTypeListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetVatTypeList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *clientServiceSoap) GetVatTypeList(request *GetVatTypeList) (*GetVatTypeListResponse, error) {
	return service.GetVatTypeListContext(
		context.Background(),
		request,
	)
}
