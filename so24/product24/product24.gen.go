// Code generated by gowsdl DO NOT EDIT.

package product24

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type SetStockQuantity struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SetStockQuantity"`

	ProductId int32 `xml:"productId,omitempty" json:"productId,omitempty"`

	StockQuantity float64 `xml:"stockQuantity,omitempty" json:"stockQuantity,omitempty"`
}

type SetStockQuantityResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SetStockQuantityResponse"`

	SetStockQuantityResult bool `xml:"SetStockQuantityResult,omitempty" json:"SetStockQuantityResult,omitempty"`
}

type GetProducts struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetProducts"`

	SearchParams *ProductSearchParameters `xml:"searchParams,omitempty" json:"searchParams,omitempty"`

	ReturnProperties *ArrayOfString `xml:"returnProperties,omitempty" json:"returnProperties,omitempty"`
}

type GetProductsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetProductsResponse"`

	GetProductsResult *ArrayOfProduct `xml:"GetProductsResult,omitempty" json:"GetProductsResult,omitempty"`
}

type SaveProducts struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveProducts"`

	Products *ArrayOfProduct `xml:"products,omitempty" json:"products,omitempty"`
}

type SaveProductsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveProductsResponse"`

	SaveProductsResult *ArrayOfProduct `xml:"SaveProductsResult,omitempty" json:"SaveProductsResult,omitempty"`
}

type DeleteProducts struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices DeleteProducts"`

	Products *ArrayOfProduct `xml:"products,omitempty" json:"products,omitempty"`
}

type DeleteProductsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices DeleteProductsResponse"`

	DeleteProductsResult *ArrayOfProduct `xml:"DeleteProductsResult,omitempty" json:"DeleteProductsResult,omitempty"`
}

type GetPriceList struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetPriceList"`

	PriceListId int32 `xml:"priceListId,omitempty" json:"priceListId,omitempty"`
}

type GetPriceListResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetPriceListResponse"`

	GetPriceListResult *ArrayOfKeyValuePair `xml:"GetPriceListResult,omitempty" json:"GetPriceListResult,omitempty"`
}

type GetAllPriceLists struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetAllPriceLists"`
}

type GetAllPriceListsResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetAllPriceListsResponse"`

	GetAllPriceListsResult *ArrayOfProductPriceList `xml:"GetAllPriceListsResult,omitempty" json:"GetAllPriceListsResult,omitempty"`
}

type GetCategories struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetCategories"`

	ReturnProperties *ArrayOfString `xml:"returnProperties,omitempty" json:"returnProperties,omitempty"`
}

type GetCategoriesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetCategoriesResponse"`

	GetCategoriesResult *ArrayOfCategory `xml:"GetCategoriesResult,omitempty" json:"GetCategoriesResult,omitempty"`
}

type SaveCategories struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveCategories"`

	Categories *ArrayOfCategory `xml:"categories,omitempty" json:"categories,omitempty"`
}

type SaveCategoriesResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices SaveCategoriesResponse"`

	SaveCategoriesResult *ArrayOfCategory `xml:"SaveCategoriesResult,omitempty" json:"SaveCategoriesResult,omitempty"`
}

type GetDiscountMatrixPriceGroup struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDiscountMatrixPriceGroup"`
}

type GetDiscountMatrixPriceGroupResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDiscountMatrixPriceGroupResponse"`

	GetDiscountMatrixPriceGroupResult *ArrayOfDiscount `xml:"GetDiscountMatrixPriceGroupResult,omitempty" json:"GetDiscountMatrixPriceGroupResult,omitempty"`
}

type GetDiscountMatrixCategory struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDiscountMatrixCategory"`
}

type GetDiscountMatrixCategoryResponse struct {
	XMLName xml.Name `xml:"http://24sevenOffice.com/webservices GetDiscountMatrixCategoryResponse"`

	GetDiscountMatrixCategoryResult *ArrayOfDiscount `xml:"GetDiscountMatrixCategoryResult,omitempty" json:"GetDiscountMatrixCategoryResult,omitempty"`
}

type ProductSearchParameters struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	CategoryId int32 `xml:"CategoryId,omitempty" json:"CategoryId,omitempty"`

	No string `xml:"No,omitempty" json:"No,omitempty"`

	EAN1 string `xml:"EAN1,omitempty" json:"EAN1,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty"`

	DateChanged soap.XSDDateTime `xml:"DateChanged,omitempty" json:"DateChanged,omitempty"`

	ProductIds *ArrayOfInt `xml:"ProductIds,omitempty" json:"ProductIds,omitempty"`
}

type ArrayOfInt struct {
	Int []int32 `xml:"int,omitempty" json:"int,omitempty"`
}

type ArrayOfString struct {
	Astring []*string `xml:"string,omitempty" json:"string,omitempty"`
}

type ArrayOfProduct struct {
	Product []*Product `xml:"Product,omitempty" json:"Product,omitempty"`
}

type Product struct {
	CashPriceIncTax float64 `xml:"CashPriceIncTax,omitempty" json:"CashPriceIncTax,omitempty"`

	WebPrice float64 `xml:"WebPrice,omitempty" json:"WebPrice,omitempty"`

	SupplierId int32 `xml:"SupplierId,omitempty" json:"SupplierId,omitempty"`

	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	TaxRate float64 `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`

	Stock float64 `xml:"Stock,omitempty" json:"Stock,omitempty"`

	StatusId int32 `xml:"StatusId,omitempty" json:"StatusId,omitempty"`

	CategoryId int32 `xml:"CategoryId,omitempty" json:"CategoryId,omitempty"`

	PriceGroupID int32 `xml:"PriceGroupID,omitempty" json:"PriceGroupID,omitempty"`

	InPrice float64 `xml:"InPrice,omitempty" json:"InPrice,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	DescriptionLong string `xml:"DescriptionLong,omitempty" json:"DescriptionLong,omitempty"`

	Cost float64 `xml:"Cost,omitempty" json:"Cost,omitempty"`

	EAN1 string `xml:"EAN1,omitempty" json:"EAN1,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty"`

	No string `xml:"No,omitempty" json:"No,omitempty"`

	DateChanged soap.XSDDateTime `xml:"DateChanged,omitempty" json:"DateChanged,omitempty"`

	APIException *APIException `xml:"APIException,omitempty" json:"APIException,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty"`

	MinimumStock float64 `xml:"MinimumStock,omitempty" json:"MinimumStock,omitempty"`

	OrderProposal float64 `xml:"OrderProposal,omitempty" json:"OrderProposal,omitempty"`

	StockLocation string `xml:"StockLocation,omitempty" json:"StockLocation,omitempty"`

	SupplierProductCode string `xml:"SupplierProductCode,omitempty" json:"SupplierProductCode,omitempty"`

	SupplierProductName string `xml:"SupplierProductName,omitempty" json:"SupplierProductName,omitempty"`

	Web *bool `xml:"Web,omitempty" json:"Web,omitempty"`

	ProductCode string `xml:"ProductCode,omitempty" json:"ProductCode,omitempty"`

	StockControlled *bool `xml:"StockControlled,omitempty" json:"StockControlled,omitempty"`

	NewOffer *bool `xml:"NewOffer,omitempty" json:"NewOffer,omitempty"`

	NewProduct *bool `xml:"NewProduct,omitempty" json:"NewProduct,omitempty"`
}

type APIException struct {
	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	Message string `xml:"Message,omitempty" json:"Message,omitempty"`

	StackTrace string `xml:"StackTrace,omitempty" json:"StackTrace,omitempty"`
}

type ArrayOfKeyValuePair struct {
	KeyValuePair []*KeyValuePair `xml:"KeyValuePair,omitempty" json:"KeyValuePair,omitempty"`
}

type KeyValuePair struct {
	Key string `xml:"Key,omitempty" json:"Key,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ArrayOfProductPriceList struct {
	ProductPriceList []*ProductPriceList `xml:"ProductPriceList,omitempty" json:"ProductPriceList,omitempty"`
}

type ProductPriceList struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	IncVat bool `xml:"IncVat,omitempty" json:"IncVat,omitempty"`

	Currency string `xml:"Currency,omitempty" json:"Currency,omitempty"`
}

type ArrayOfCategory struct {
	Category []*Category `xml:"Category,omitempty" json:"Category,omitempty"`
}

type Category struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	No string `xml:"No,omitempty" json:"No,omitempty"`

	ParentId int32 `xml:"ParentId,omitempty" json:"ParentId,omitempty"`

	APIException *APIException `xml:"APIException,omitempty" json:"APIException,omitempty"`
}

type ArrayOfDiscount struct {
	Discount []*Discount `xml:"Discount,omitempty" json:"Discount,omitempty"`
}

type Discount struct {
	PriceGroupId int32 `xml:"PriceGroupId,omitempty" json:"PriceGroupId,omitempty"`

	CategoryId int32 `xml:"CategoryId,omitempty" json:"CategoryId,omitempty"`

	CustomerId int32 `xml:"CustomerId,omitempty" json:"CustomerId,omitempty"`

	Value float64 `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ProductServiceSoap interface {

	/* Sets the stock quantity and returns true if no error occured. */
	SetStockQuantity(request *SetStockQuantity) (*SetStockQuantityResponse, error)

	SetStockQuantityContext(ctx context.Context, request *SetStockQuantity) (*SetStockQuantityResponse, error)

	GetProducts(request *GetProducts) (*GetProductsResponse, error)

	GetProductsContext(ctx context.Context, request *GetProducts) (*GetProductsResponse, error)

	SaveProducts(request *SaveProducts) (*SaveProductsResponse, error)

	SaveProductsContext(ctx context.Context, request *SaveProducts) (*SaveProductsResponse, error)

	DeleteProducts(request *DeleteProducts) (*DeleteProductsResponse, error)

	DeleteProductsContext(ctx context.Context, request *DeleteProducts) (*DeleteProductsResponse, error)

	/* Returns the product prices for a given price list Id */
	GetPriceList(request *GetPriceList) (*GetPriceListResponse, error)

	GetPriceListContext(ctx context.Context, request *GetPriceList) (*GetPriceListResponse, error)

	/* Returns all product price lists */
	GetAllPriceLists(request *GetAllPriceLists) (*GetAllPriceListsResponse, error)

	GetAllPriceListsContext(ctx context.Context, request *GetAllPriceLists) (*GetAllPriceListsResponse, error)

	GetCategories(request *GetCategories) (*GetCategoriesResponse, error)

	GetCategoriesContext(ctx context.Context, request *GetCategories) (*GetCategoriesResponse, error)

	SaveCategories(request *SaveCategories) (*SaveCategoriesResponse, error)

	SaveCategoriesContext(ctx context.Context, request *SaveCategories) (*SaveCategoriesResponse, error)

	GetDiscountMatrixPriceGroup(request *GetDiscountMatrixPriceGroup) (*GetDiscountMatrixPriceGroupResponse, error)

	GetDiscountMatrixPriceGroupContext(ctx context.Context, request *GetDiscountMatrixPriceGroup) (*GetDiscountMatrixPriceGroupResponse, error)

	GetDiscountMatrixCategory(request *GetDiscountMatrixCategory) (*GetDiscountMatrixCategoryResponse, error)

	GetDiscountMatrixCategoryContext(ctx context.Context, request *GetDiscountMatrixCategory) (*GetDiscountMatrixCategoryResponse, error)
}

type productServiceSoap struct {
	client *soap.Client
}

func NewProductServiceSoap(client *soap.Client) ProductServiceSoap {
	return &productServiceSoap{
		client: client,
	}
}

func (service *productServiceSoap) SetStockQuantityContext(ctx context.Context, request *SetStockQuantity) (*SetStockQuantityResponse, error) {
	response := new(SetStockQuantityResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/SetStockQuantity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) SetStockQuantity(request *SetStockQuantity) (*SetStockQuantityResponse, error) {
	return service.SetStockQuantityContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetProductsContext(ctx context.Context, request *GetProducts) (*GetProductsResponse, error) {
	response := new(GetProductsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetProducts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetProducts(request *GetProducts) (*GetProductsResponse, error) {
	return service.GetProductsContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) SaveProductsContext(ctx context.Context, request *SaveProducts) (*SaveProductsResponse, error) {
	response := new(SaveProductsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/SaveProducts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) SaveProducts(request *SaveProducts) (*SaveProductsResponse, error) {
	return service.SaveProductsContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) DeleteProductsContext(ctx context.Context, request *DeleteProducts) (*DeleteProductsResponse, error) {
	response := new(DeleteProductsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/DeleteProducts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) DeleteProducts(request *DeleteProducts) (*DeleteProductsResponse, error) {
	return service.DeleteProductsContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetPriceListContext(ctx context.Context, request *GetPriceList) (*GetPriceListResponse, error) {
	response := new(GetPriceListResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetPriceList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetPriceList(request *GetPriceList) (*GetPriceListResponse, error) {
	return service.GetPriceListContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetAllPriceListsContext(ctx context.Context, request *GetAllPriceLists) (*GetAllPriceListsResponse, error) {
	response := new(GetAllPriceListsResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetAllPriceLists", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetAllPriceLists(request *GetAllPriceLists) (*GetAllPriceListsResponse, error) {
	return service.GetAllPriceListsContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetCategoriesContext(ctx context.Context, request *GetCategories) (*GetCategoriesResponse, error) {
	response := new(GetCategoriesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetCategories", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetCategories(request *GetCategories) (*GetCategoriesResponse, error) {
	return service.GetCategoriesContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) SaveCategoriesContext(ctx context.Context, request *SaveCategories) (*SaveCategoriesResponse, error) {
	response := new(SaveCategoriesResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/SaveCategories", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) SaveCategories(request *SaveCategories) (*SaveCategoriesResponse, error) {
	return service.SaveCategoriesContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetDiscountMatrixPriceGroupContext(ctx context.Context, request *GetDiscountMatrixPriceGroup) (*GetDiscountMatrixPriceGroupResponse, error) {
	response := new(GetDiscountMatrixPriceGroupResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetDiscountMatrixPriceGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetDiscountMatrixPriceGroup(request *GetDiscountMatrixPriceGroup) (*GetDiscountMatrixPriceGroupResponse, error) {
	return service.GetDiscountMatrixPriceGroupContext(
		context.Background(),
		request,
	)
}

func (service *productServiceSoap) GetDiscountMatrixCategoryContext(ctx context.Context, request *GetDiscountMatrixCategory) (*GetDiscountMatrixCategoryResponse, error) {
	response := new(GetDiscountMatrixCategoryResponse)
	err := service.client.CallContext(ctx, "http://24sevenOffice.com/webservices/GetDiscountMatrixCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *productServiceSoap) GetDiscountMatrixCategory(request *GetDiscountMatrixCategory) (*GetDiscountMatrixCategoryResponse, error) {
	return service.GetDiscountMatrixCategoryContext(
		context.Background(),
		request,
	)
}
