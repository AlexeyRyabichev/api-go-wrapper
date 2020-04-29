package api

const (
	baseURL = "https://%s.erply.com/api/"

	GetSalesDocumentsMethod           = "getSalesDocuments"
	GetCustomersMethod                = "getCustomers"
	getSuppliersMethod                = "getSuppliers"
	GetCountriesMethod                = "getCountries"
	GetEmployeesMethod                = "getEmployees"
	GetBusinessAreasMethod            = "getBusinessAreas"
	GetProjectsMethod                 = "getProjects"
	GetProjectStatusesMethod          = "getProjectStatuses"
	GetCurrenciesMethod               = "getCurrencies"
	GetVatRatesMethod                 = "getVatRates"
	GetPaymentsMethod                 = "getPayments"
	GetUserRightsMethod               = "getUserRights"
	GetCompanyInfoMethod              = "getCompanyInfo"
	VerifyIdentityTokenMethod         = "verifyIdentityToken"
	GetPointsOfSaleMethod             = "getPointsOfSale"
	GetIdentityToken                  = "getIdentityToken"
	GetJWTTokenMethod                 = "getJwtToken"
	GetConfParametersMethod           = "getConfParameters"
	logProcessingOfCustomerDataMethod = "logProcessingOfCustomerData"
	saveSalesDocumentMethod           = "saveSalesDocument"
	savePurchaseDocumentMethod        = "savePurchaseDocument"
	saveCustomerMethod                = "saveCustomer"
	saveAddressMethod                 = "saveAddress"
	saveSupplierMethod                = "saveSupplier"
	savePaymentMethod                 = "savePayment"
	createInstallationMethod          = "createInstallation"
	GetWarehousesMethod               = "getWarehouses"
	GetAddressesMethod                = "getAddresses"
	GetProductsMethod                 = "getProducts"
	GetProductCategoriesMethod        = "getProductCategories"
	GetProductBrandsMethod            = "getBrands"
	GetProductGroupsMethod            = "getProductGroups"
	GetProductUnitsMethod             = "getProductUnits"
	VerifyCustomerUserMethod          = "verifyCustomerUser"
	calculateShoppingCartMethod       = "calculateShoppingCart"
	validateCustomerUsernameMethod    = "validateCustomerUsername"
	clientCode                        = "clientCode"
	sessionKey                        = "sessionKey"
	responseStatusOK                  = "ok"
	//MaxIdleConns for Erply API
	MaxIdleConns = 25

	//MaxConnsPerHost for Erply API
	MaxConnsPerHost = 25
)
