package api

import "context"

//IClient interface for cached and simple client
type IClient interface {
	GetConfParameters() (*ConfParameter, error)
	GetWarehouses() (Warehouses, error)
	GetSalesDocumentByID(id string) ([]SaleDocument, error)
	GetSalesDocumentsByIDs(id []string) ([]SaleDocument, error)
	GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
	GetCustomersByIDs(customerID []string) (Customers, error)
	GetCustomerByRegNumber(regNumber string) (*Customer, error)
	GetCustomerByGLN(gln string) (*Customer, error)
	GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
	GetSupplierByName(name string) (*Customer, error)
	GetVatRatesByID(vatRateID string) (VatRates, error)
	GetCompanyInfo() (*CompanyInfo, error)
	GetProductUnits() ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductsByIDs(ids []string) ([]Product, error)
	GetProductsByCode3(code3 string) (*Product, error)
	GetAddresses() (*Address, error)
	GetCountries(ctx context.Context, filters map[string]string) ([]Country, error)
	GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error)
	GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error)
	GetProjects(ctx context.Context, filters map[string]string) ([]Project, error)
	GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error)
	GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error)
	PostPurchaseDocument(in *PurchaseDocumentConstructor, provider string) (PurchaseDocImportReports, error)
	PostSalesDocumentFromWoocomm(in *SaleDocumentConstructor, shopOrderID string) (SaleDocImportReports, error)
	PostSalesDocument(in *SaleDocumentConstructor, provider string) (SaleDocImportReports, error)
	PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error)
	PostSupplier(in *CustomerConstructor) (*CustomerImportReport, error)
	DeleteDocumentsByID(id string) error
	GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error)
	GetPointsOfSaleByID(posID string) (*PointOfSale, error)
	VerifyIdentityToken(jwt string) (*SessionInfo, error)
	GetIdentityToken() (*IdentityToken, error)
	Close()
}
