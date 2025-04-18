// Code generated by gowsdl DO NOT EDIT.

package invoice24

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

type OrderSlipStateType string

const (
	OrderSlipStateTypeNone OrderSlipStateType = "None"

	OrderSlipStateTypeOffer OrderSlipStateType = "Offer"

	OrderSlipStateTypePack OrderSlipStateType = "Pack"

	OrderSlipStateTypeInvoiced OrderSlipStateType = "Invoiced"

	OrderSlipStateTypeWeb OrderSlipStateType = "Web"

	OrderSlipStateTypeRepeating OrderSlipStateType = "Repeating"

	OrderSlipStateTypeCollectingOrder OrderSlipStateType = "CollectingOrder"

	OrderSlipStateTypeReturn OrderSlipStateType = "Return"

	OrderSlipStateTypeDistribution OrderSlipStateType = "Distribution"

	OrderSlipStateTypeRest OrderSlipStateType = "Rest"

	OrderSlipStateTypeForInvoicing OrderSlipStateType = "ForInvoicing"

	OrderSlipStateTypeOvertimeRent OrderSlipStateType = "OvertimeRent"

	OrderSlipStateTypeConfirmed OrderSlipStateType = "Confirmed"

	OrderSlipStateTypeOrderState_Sent OrderSlipStateType = "OrderState_Sent"

	OrderSlipStateTypeProduction OrderSlipStateType = "Production"

	OrderSlipStateTypePaymentReminder OrderSlipStateType = "PaymentReminder"

	OrderSlipStateTypeOnAccount OrderSlipStateType = "OnAccount"

	OrderSlipStateTypeLending OrderSlipStateType = "Lending"

	OrderSlipStateTypeInvoicedCashAccount OrderSlipStateType = "InvoicedCashAccount"

	OrderSlipStateTypeInvoicedPack OrderSlipStateType = "InvoicedPack"

	OrderSlipStateTypeInactive OrderSlipStateType = "Inactive"

	OrderSlipStateTypeHours OrderSlipStateType = "Hours"

	OrderSlipStateTypeOrderState_SuperStore OrderSlipStateType = "OrderState_SuperStore"

	OrderSlipStateTypeDraft OrderSlipStateType = "Draft"
)

type Distributor string

const (
	DistributorDefault Distributor = "Default"

	DistributorManual Distributor = "Manual"

	DistributorAuto Distributor = "Auto"
)

type DistributionMethod string

const (
	DistributionMethodUnchanged DistributionMethod = "Unchanged"

	DistributionMethodPrint DistributionMethod = "Print"

	DistributionMethodEMail DistributionMethod = "EMail"

	DistributionMethodElectronicInvoice DistributionMethod = "ElectronicInvoice"
)

type ChangeState string

const (
	ChangeStateNone ChangeState = "None"

	ChangeStateAdd ChangeState = "Add"

	ChangeStateEdit ChangeState = "Edit"

	ChangeStateDelete ChangeState = "Delete"
)

type RowType string

const (
	RowTypeNormal RowType = "Normal"

	RowTypeText RowType = "Text"

	RowTypeTextBold RowType = "TextBold"
)

type GetOrderSyncList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetOrderSyncList"`

	SyncParams *InvoiceSyncParameters `xml:"syncParams,omitempty" json:"syncParams,omitempty"`
}

type GetOrderSyncListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetOrderSyncListResponse"`

	GetOrderSyncListResult *ArrayOfInt `xml:"GetOrderSyncListResult,omitempty" json:"GetOrderSyncListResult,omitempty"`
}

type GetInvoices struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoices"`

	SearchParams *InvoiceSearchParameters `xml:"searchParams,omitempty" json:"searchParams,omitempty"`

	InvoiceReturnProperties *ArrayOfString `xml:"invoiceReturnProperties,omitempty" json:"invoiceReturnProperties,omitempty"`

	RowReturnProperties *ArrayOfString `xml:"rowReturnProperties,omitempty" json:"rowReturnProperties,omitempty"`
}

type GetInvoicesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoicesResponse"`

	GetInvoicesResult *ArrayOfInvoiceOrder `xml:"GetInvoicesResult,omitempty" json:"GetInvoicesResult,omitempty"`
}

type UpdateInvoices struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices UpdateInvoices"`

	Invoices *ArrayOfInvoiceOrder `xml:"invoices,omitempty" json:"invoices,omitempty"`
}

type UpdateInvoicesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices UpdateInvoicesResponse"`

	UpdateInvoicesResult *ArrayOfInvoiceOrder `xml:"UpdateInvoicesResult,omitempty" json:"UpdateInvoicesResult,omitempty"`
}

type SaveInvoices struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveInvoices"`

	Invoices *ArrayOfInvoiceOrder `xml:"invoices,omitempty" json:"invoices,omitempty"`
}

type SaveInvoicesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveInvoicesResponse"`

	SaveInvoicesResult *ArrayOfInvoiceOrder `xml:"SaveInvoicesResult,omitempty" json:"SaveInvoicesResult,omitempty"`
}

type GetInvoiceDocument struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoiceDocument"`

	Parameters *InvoiceDocumentSearchParameters `xml:"Parameters,omitempty" json:"Parameters,omitempty"`
}

type GetInvoiceDocumentResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoiceDocumentResponse"`

	GetInvoiceDocumentResult []byte `xml:"GetInvoiceDocumentResult,omitempty" json:"GetInvoiceDocumentResult,omitempty"`
}

type GetDeliveryMethods struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDeliveryMethods"`
}

type GetDeliveryMethodsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDeliveryMethodsResponse"`

	GetDeliveryMethodsResult *ArrayOfDeliveryMethod `xml:"GetDeliveryMethodsResult,omitempty" json:"GetDeliveryMethodsResult,omitempty"`
}

type GetInvoiceTemplates struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoiceTemplates"`
}

type GetInvoiceTemplatesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetInvoiceTemplatesResponse"`

	GetInvoiceTemplatesResult *ArrayOfTemplateModel `xml:"GetInvoiceTemplatesResult,omitempty" json:"GetInvoiceTemplatesResult,omitempty"`
}

type GetPaymentMethods struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetPaymentMethods"`
}

type GetPaymentMethodsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetPaymentMethodsResponse"`

	GetPaymentMethodsResult *ArrayOfPaymentMethod `xml:"GetPaymentMethodsResult,omitempty" json:"GetPaymentMethodsResult,omitempty"`
}

type InvoiceSyncParameters struct {
	ChangedAfter soap.XSDDateTime `xml:"ChangedAfter,omitempty" json:"ChangedAfter,omitempty"`

	RegisteredAfter soap.XSDDateTime `xml:"RegisteredAfter,omitempty" json:"RegisteredAfter,omitempty"`
}

type ArrayOfInt struct {
	Int []int32 `xml:"int,omitempty" json:"int,omitempty"`
}

type InvoiceSearchParameters struct {
	CustomerIds *ArrayOfInt `xml:"CustomerIds,omitempty" json:"CustomerIds,omitempty"`

	ChangedAfter soap.XSDDateTime `xml:"ChangedAfter,omitempty" json:"ChangedAfter,omitempty"`

	OrderStates *ArrayOfOrderSlipStateType `xml:"OrderStates,omitempty" json:"OrderStates,omitempty"`

	InvoiceIds *ArrayOfInt `xml:"InvoiceIds,omitempty" json:"InvoiceIds,omitempty"`

	OrderIds *ArrayOfInt `xml:"OrderIds,omitempty" json:"OrderIds,omitempty"`

	ExternalStatus *ArrayOfInt `xml:"ExternalStatus,omitempty" json:"ExternalStatus,omitempty"`
}

type ArrayOfOrderSlipStateType struct {
	OrderSlipStateType []*OrderSlipStateType `xml:"OrderSlipStateType,omitempty" json:"OrderSlipStateType,omitempty"`
}

type ArrayOfString struct {
	Astring []*string `xml:"string,omitempty" json:"string,omitempty"`
}

type ArrayOfInvoiceOrder struct {
	InvoiceOrder []*InvoiceOrder `xml:"InvoiceOrder,omitempty" json:"InvoiceOrder,omitempty"`
}

type InvoiceOrder struct {
	OrderId int32 `xml:"OrderId,omitempty" json:"OrderId,omitempty"`

	CustomerId int32 `xml:"CustomerId,omitempty" json:"CustomerId,omitempty"`

	CustomerName string `xml:"CustomerName,omitempty" json:"CustomerName,omitempty"`

	CustomerDeliveryName string `xml:"CustomerDeliveryName,omitempty" json:"CustomerDeliveryName,omitempty"`

	CustomerDeliveryPhone string `xml:"CustomerDeliveryPhone,omitempty" json:"CustomerDeliveryPhone,omitempty"`

	Addresses *Addresses `xml:"Addresses,omitempty" json:"Addresses,omitempty"`

	OrderStatus *OrderSlipStateType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty"`

	InvoiceId int32 `xml:"InvoiceId,omitempty" json:"InvoiceId,omitempty"`

	DateOrdered soap.XSDDateTime `xml:"DateOrdered,omitempty" json:"DateOrdered,omitempty"`

	DateInvoiced soap.XSDDateTime `xml:"DateInvoiced,omitempty" json:"DateInvoiced,omitempty"`

	DateChanged soap.XSDDateTime `xml:"DateChanged,omitempty" json:"DateChanged,omitempty"`

	PaymentTime int32 `xml:"PaymentTime,omitempty" json:"PaymentTime,omitempty"`

	CustomerReferenceNo string `xml:"CustomerReferenceNo,omitempty" json:"CustomerReferenceNo,omitempty"`

	ProjectId int32 `xml:"ProjectId,omitempty" json:"ProjectId,omitempty"`

	OurReference int32 `xml:"OurReference,omitempty" json:"OurReference,omitempty"`

	IncludeVAT *bool `xml:"IncludeVAT,omitempty" json:"IncludeVAT,omitempty"`

	YourReference string `xml:"YourReference,omitempty" json:"YourReference,omitempty"`

	OrderTotalIncVat float64 `xml:"OrderTotalIncVat,omitempty" json:"OrderTotalIncVat,omitempty"`

	OrderTotalVat float64 `xml:"OrderTotalVat,omitempty" json:"OrderTotalVat,omitempty"`

	InvoiceTitle string `xml:"InvoiceTitle,omitempty" json:"InvoiceTitle,omitempty"`

	InvoiceText string `xml:"InvoiceText,omitempty" json:"InvoiceText,omitempty"`

	Paid soap.XSDDateTime `xml:"Paid,omitempty" json:"Paid,omitempty"`

	OCR string `xml:"OCR,omitempty" json:"OCR,omitempty"`

	CustomerOrgNo string `xml:"CustomerOrgNo,omitempty" json:"CustomerOrgNo,omitempty"`

	Currency *Currency `xml:"Currency,omitempty" json:"Currency,omitempty"`

	PaymentMethodId int32 `xml:"PaymentMethodId,omitempty" json:"PaymentMethodId,omitempty"`

	PaymentAmount float64 `xml:"PaymentAmount,omitempty" json:"PaymentAmount,omitempty"`

	ProductionManagerId int32 `xml:"ProductionManagerId,omitempty" json:"ProductionManagerId,omitempty"`

	SalesOpportunityId int32 `xml:"SalesOpportunityId,omitempty" json:"SalesOpportunityId,omitempty"`

	TypeOfSaleId int32 `xml:"TypeOfSaleId,omitempty" json:"TypeOfSaleId,omitempty"`

	Distributor *Distributor `xml:"Distributor,omitempty" json:"Distributor,omitempty"`

	DistributionMethod *DistributionMethod `xml:"DistributionMethod,omitempty" json:"DistributionMethod,omitempty"`

	DepartmentId int32 `xml:"DepartmentId,omitempty" json:"DepartmentId,omitempty"`

	ExternalStatus int32 `xml:"ExternalStatus,omitempty" json:"ExternalStatus,omitempty"`

	InvoiceEmailAddress string `xml:"InvoiceEmailAddress,omitempty" json:"InvoiceEmailAddress,omitempty"`

	InvoiceRows *ArrayOfInvoiceRow `xml:"InvoiceRows,omitempty" json:"InvoiceRows,omitempty"`

	APIException *APIException `xml:"APIException,omitempty" json:"APIException,omitempty"`

	ProductionNumber string `xml:"ProductionNumber,omitempty" json:"ProductionNumber,omitempty"`

	DeliveryDate soap.XSDDateTime `xml:"DeliveryDate,omitempty" json:"DeliveryDate,omitempty"`

	ReferenceInvoiceId int32 `xml:"ReferenceInvoiceId,omitempty" json:"ReferenceInvoiceId,omitempty"`

	ReferenceOrderId int32 `xml:"ReferenceOrderId,omitempty" json:"ReferenceOrderId,omitempty"`

	ReferenceNumber string `xml:"ReferenceNumber,omitempty" json:"ReferenceNumber,omitempty"`

	SkipStock bool `xml:"SkipStock,omitempty" json:"SkipStock,omitempty"`

	AccrualDate soap.XSDDateTime `xml:"AccrualDate,omitempty" json:"AccrualDate,omitempty"`

	AccrualLength int32 `xml:"AccrualLength,omitempty" json:"AccrualLength,omitempty"`

	RoundFactor float64 `xml:"RoundFactor,omitempty" json:"RoundFactor,omitempty"`

	InvoiceTemplateId *Guid `xml:"InvoiceTemplateId,omitempty" json:"InvoiceTemplateId,omitempty"`

	VippsNumber string `xml:"VippsNumber,omitempty" json:"VippsNumber,omitempty"`

	DeliveryMethod *DeliveryMethod `xml:"DeliveryMethod,omitempty" json:"DeliveryMethod,omitempty"`

	DeliveryAlternative string `xml:"DeliveryAlternative,omitempty" json:"DeliveryAlternative,omitempty"`

	SendToFactoring *bool `xml:"SendToFactoring,omitempty" json:"SendToFactoring,omitempty"`

	Commission float64 `xml:"Commission,omitempty" json:"Commission,omitempty"`

	UserDefinedDimensions *ArrayOfUserDefinedDimension `xml:"UserDefinedDimensions,omitempty" json:"UserDefinedDimensions,omitempty"`

	GLNNumber string `xml:"GLNNumber,omitempty" json:"GLNNumber,omitempty"`

	CustomerDeliveryId int32 `xml:"CustomerDeliveryId,omitempty" json:"CustomerDeliveryId,omitempty"`
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

type Currency struct {
	Rate float64 `xml:"Rate,omitempty" json:"Rate,omitempty"`

	Unit int16 `xml:"Unit,omitempty" json:"Unit,omitempty"`

	Symbol string `xml:"Symbol,omitempty" json:"Symbol,omitempty"`
}

type ArrayOfInvoiceRow struct {
	InvoiceRow []*InvoiceRow `xml:"InvoiceRow,omitempty" json:"InvoiceRow,omitempty"`
}

type InvoiceRow struct {
	ProductId int32 `xml:"ProductId,omitempty" json:"ProductId,omitempty"`

	RowId int32 `xml:"RowId,omitempty" json:"RowId,omitempty"`

	VatRate float64 `xml:"VatRate,omitempty" json:"VatRate,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	DiscountRate float64 `xml:"DiscountRate,omitempty" json:"DiscountRate,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty" json:"Quantity,omitempty"`

	QuantityDelivered float64 `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered,omitempty"`

	QuantityOrdered float64 `xml:"QuantityOrdered,omitempty" json:"QuantityOrdered,omitempty"`

	QuantityRest float64 `xml:"QuantityRest,omitempty" json:"QuantityRest,omitempty"`

	Cost float64 `xml:"Cost,omitempty" json:"Cost,omitempty"`

	ChangeState *ChangeState `xml:"ChangeState,omitempty" json:"ChangeState,omitempty"`

	TypeGroupId int32 `xml:"TypeGroupId,omitempty" json:"TypeGroupId,omitempty"`

	SequenceNumber int16 `xml:"SequenceNumber,omitempty" json:"SequenceNumber,omitempty"`

	Hidden bool `xml:"Hidden,omitempty" json:"Hidden,omitempty"`

	Type *RowType `xml:"Type,omitempty" json:"Type,omitempty"`

	InPrice float64 `xml:"InPrice,omitempty" json:"InPrice,omitempty"`

	AccrualDate soap.XSDDateTime `xml:"AccrualDate,omitempty" json:"AccrualDate,omitempty"`

	AccrualLength int32 `xml:"AccrualLength,omitempty" json:"AccrualLength,omitempty"`

	AccountProject *bool `xml:"AccountProject,omitempty" json:"AccountProject,omitempty"`

	ProjectId int32 `xml:"ProjectId,omitempty" json:"ProjectId,omitempty"`

	DepartmentId int32 `xml:"DepartmentId,omitempty" json:"DepartmentId,omitempty"`

	ProductNo string `xml:"ProductNo,omitempty" json:"ProductNo,omitempty"`

	TaxSettings *TaxSettings `xml:"TaxSettings,omitempty" json:"TaxSettings,omitempty"`

	UserDefinedDimensions *ArrayOfUserDefinedDimension `xml:"UserDefinedDimensions,omitempty" json:"UserDefinedDimensions,omitempty"`

	SerialNumber string `xml:"SerialNumber,omitempty" json:"SerialNumber,omitempty"`
}

type TaxSettings struct {
	TaxAccount int32 `xml:"TaxAccount,omitempty" json:"TaxAccount,omitempty"`

	TaxCode int32 `xml:"TaxCode,omitempty" json:"TaxCode,omitempty"`

	TaxRate float64 `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
}

type ArrayOfUserDefinedDimension struct {
	UserDefinedDimension []*UserDefinedDimension `xml:"UserDefinedDimension,omitempty" json:"UserDefinedDimension,omitempty"`
}

type UserDefinedDimension struct {
	TypeId int32 `xml:"TypeId,omitempty" json:"TypeId,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type APIException struct {
	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	Message string `xml:"Message,omitempty" json:"Message,omitempty"`

	StackTrace string `xml:"StackTrace,omitempty" json:"StackTrace,omitempty"`
}

type DeliveryMethod struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty"`
}

type InvoiceDocumentSearchParameters struct {
	InvoiceId int32 `xml:"InvoiceId,omitempty" json:"InvoiceId,omitempty"`

	TemplateId *Guid `xml:"TemplateId,omitempty" json:"TemplateId,omitempty"`
}

type ArrayOfDeliveryMethod struct {
	DeliveryMethod []*DeliveryMethod `xml:"DeliveryMethod,omitempty" json:"DeliveryMethod,omitempty"`
}

type ArrayOfTemplateModel struct {
	TemplateModel []*TemplateModel `xml:"TemplateModel,omitempty" json:"TemplateModel,omitempty"`
}

type TemplateModel struct {
	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Id *Guid `xml:"Id,omitempty" json:"Id,omitempty"`

	Language string `xml:"Language,omitempty" json:"Language,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty"`

	IsSystemTemplate bool `xml:"IsSystemTemplate,omitempty" json:"IsSystemTemplate,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty" json:"IsDefault,omitempty"`

	Modified soap.XSDDateTime `xml:"Modified,omitempty" json:"Modified,omitempty"`

	Customers *ArrayOfInt `xml:"Customers,omitempty" json:"Customers,omitempty"`
}

type ArrayOfPaymentMethod struct {
	PaymentMethod []*PaymentMethod `xml:"PaymentMethod,omitempty" json:"PaymentMethod,omitempty"`
}

type PaymentMethod struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`
}

type Guid string

type InvoiceServiceSoap interface {

	/* Gets a list of OrderIds that can be used for syncing data with the GetInvoices method */
	GetOrderSyncList(request *GetOrderSyncList) (*GetOrderSyncListResponse, error)

	GetOrderSyncListContext(ctx context.Context, request *GetOrderSyncList) (*GetOrderSyncListResponse, error)

	/* Gets a list of invoices and orders based on search parameters. You have to specify which properties you want returned. */
	GetInvoices(request *GetInvoices) (*GetInvoicesResponse, error)

	GetInvoicesContext(ctx context.Context, request *GetInvoices) (*GetInvoicesResponse, error)

	/* Updates an InvoiceOrder that has been Invoiced, valid update properties: InvoiceTitle, ReferenceNumber  */
	UpdateInvoices(request *UpdateInvoices) (*UpdateInvoicesResponse, error)

	UpdateInvoicesContext(ctx context.Context, request *UpdateInvoices) (*UpdateInvoicesResponse, error)

	/* Saves an array of invoices */
	SaveInvoices(request *SaveInvoices) (*SaveInvoicesResponse, error)

	SaveInvoicesContext(ctx context.Context, request *SaveInvoices) (*SaveInvoicesResponse, error)

	/* Gets an Invoice PDF. InvoiceDocumentSearchParameters: InvoiceId, TemplateId (optional) */
	GetInvoiceDocument(request *GetInvoiceDocument) (*GetInvoiceDocumentResponse, error)

	GetInvoiceDocumentContext(ctx context.Context, request *GetInvoiceDocument) (*GetInvoiceDocumentResponse, error)

	/* Gets all DeliveryMethods */
	GetDeliveryMethods(request *GetDeliveryMethods) (*GetDeliveryMethodsResponse, error)

	GetDeliveryMethodsContext(ctx context.Context, request *GetDeliveryMethods) (*GetDeliveryMethodsResponse, error)

	/* Gets all InvoiceTemplates */
	GetInvoiceTemplates(request *GetInvoiceTemplates) (*GetInvoiceTemplatesResponse, error)

	GetInvoiceTemplatesContext(ctx context.Context, request *GetInvoiceTemplates) (*GetInvoiceTemplatesResponse, error)

	/* Gets all PaymentMethods */
	GetPaymentMethods(request *GetPaymentMethods) (*GetPaymentMethodsResponse, error)

	GetPaymentMethodsContext(ctx context.Context, request *GetPaymentMethods) (*GetPaymentMethodsResponse, error)
}

type invoiceServiceSoap struct {
	client *soap.Client
}

func NewInvoiceServiceSoap(client *soap.Client) InvoiceServiceSoap {
	return &invoiceServiceSoap{
		client: client,
	}
}

func (service *invoiceServiceSoap) GetOrderSyncListContext(ctx context.Context, request *GetOrderSyncList) (*GetOrderSyncListResponse, error) {
	response := new(GetOrderSyncListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetOrderSyncList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetOrderSyncList(request *GetOrderSyncList) (*GetOrderSyncListResponse, error) {
	return service.GetOrderSyncListContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) GetInvoicesContext(ctx context.Context, request *GetInvoices) (*GetInvoicesResponse, error) {
	response := new(GetInvoicesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetInvoices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetInvoices(request *GetInvoices) (*GetInvoicesResponse, error) {
	return service.GetInvoicesContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) UpdateInvoicesContext(ctx context.Context, request *UpdateInvoices) (*UpdateInvoicesResponse, error) {
	response := new(UpdateInvoicesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/UpdateInvoices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) UpdateInvoices(request *UpdateInvoices) (*UpdateInvoicesResponse, error) {
	return service.UpdateInvoicesContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) SaveInvoicesContext(ctx context.Context, request *SaveInvoices) (*SaveInvoicesResponse, error) {
	response := new(SaveInvoicesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/SaveInvoices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) SaveInvoices(request *SaveInvoices) (*SaveInvoicesResponse, error) {
	return service.SaveInvoicesContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) GetInvoiceDocumentContext(ctx context.Context, request *GetInvoiceDocument) (*GetInvoiceDocumentResponse, error) {
	response := new(GetInvoiceDocumentResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetInvoiceDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetInvoiceDocument(request *GetInvoiceDocument) (*GetInvoiceDocumentResponse, error) {
	return service.GetInvoiceDocumentContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) GetDeliveryMethodsContext(ctx context.Context, request *GetDeliveryMethods) (*GetDeliveryMethodsResponse, error) {
	response := new(GetDeliveryMethodsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetDeliveryMethods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetDeliveryMethods(request *GetDeliveryMethods) (*GetDeliveryMethodsResponse, error) {
	return service.GetDeliveryMethodsContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) GetInvoiceTemplatesContext(ctx context.Context, request *GetInvoiceTemplates) (*GetInvoiceTemplatesResponse, error) {
	response := new(GetInvoiceTemplatesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetInvoiceTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetInvoiceTemplates(request *GetInvoiceTemplates) (*GetInvoiceTemplatesResponse, error) {
	return service.GetInvoiceTemplatesContext(
		context.Background(),
		request,
	)
}

func (service *invoiceServiceSoap) GetPaymentMethodsContext(ctx context.Context, request *GetPaymentMethods) (*GetPaymentMethodsResponse, error) {
	response := new(GetPaymentMethodsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetPaymentMethods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *invoiceServiceSoap) GetPaymentMethods(request *GetPaymentMethods) (*GetPaymentMethodsResponse, error) {
	return service.GetPaymentMethodsContext(
		context.Background(),
		request,
	)
}
