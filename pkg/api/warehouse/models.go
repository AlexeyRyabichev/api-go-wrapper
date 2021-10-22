package warehouse

import (
	"encoding/json"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

const (
	InventoryTransferACCEPTED = "ACCEPTED"
	InventoryTransferFULLFILLED = "FULLFILLED"
)

type (
	GetWarehousesResponse struct {
		Status     sharedCommon.Status `json:"status"`
		Warehouses Warehouses          `json:"records"`
	}

	Warehouse struct {
		WarehouseID            string      `json:"warehouseID"`
		PricelistID            json.Number `json:"pricelistID"`
		PricelistID2           json.Number `json:"pricelistID2"`
		PricelistID3           json.Number `json:"pricelistID3"`
		PricelistID4           json.Number `json:"pricelistID4"`
		PricelistID5           json.Number `json:"pricelistID5"`
		Name                   string      `json:"name"`
		Code                   string      `json:"code"`
		AddressID              int         `json:"addressID"`
		Address                string      `json:"address"`
		Street                 string      `json:"street"`
		Address2               string      `json:"address2"`
		City                   string      `json:"city"`
		State                  string      `json:"state"`
		Country                string      `json:"country"`
		ZIPcode                string      `json:"ZIPcode"`
		StoreGroups            string      `json:"storeGroups"`
		CompanyName            string      `json:"companyName"`
		CompanyCode            string      `json:"companyCode"`
		CompanyVatNumber       string      `json:"companyVatNumber"`
		Phone                  string      `json:"phone"`
		Fax                    string      `json:"fax"`
		Email                  string      `json:"email"`
		Website                string      `json:"website"`
		BankName               string      `json:"bankName"`
		BankAccountNumber      string      `json:"bankAccountNumber"`
		Iban                   string      `json:"iban"`
		Swift                  string      `json:"swift"`
		UsesLocalQuickButtons  int         `json:"usesLocalQuickButtons"`
		DefaultCustomerGroupID int         `json:"defaultCustomerGroupID"`
		IsOfflineInventory     int         `json:"isOfflineInventory"`
		TimeZone               string      `json:"timeZone"`
		sharedCommon.Attributes
	}

	Warehouses []Warehouse

	GetWarehousesBulkItem struct {
		Status     sharedCommon.StatusBulk `json:"status"`
		Warehouses Warehouses              `json:"records"`
	}

	GetWarehousesResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []GetWarehousesBulkItem `json:"requests"`
	}

	SaveWarehouseResult struct {
		WarehouseID int `json:"warehouseID"`
	}

	SaveWarehouseResponse struct {
		Status  sharedCommon.Status   `json:"status"`
		Results []SaveWarehouseResult `json:"records"`
	}

	SaveWarehouseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Results []SaveWarehouseResult   `json:"records"`
	}

	SaveWarehouseResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []SaveWarehouseBulkItem `json:"requests"`
	}

	SaveInventoryRegistrationResult struct {
		InventoryRegistrationID int `json:"inventoryRegistrationID"`
	}

	SaveInventoryRegistrationResponse struct {
		Status  sharedCommon.Status               `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryWriteOffResult struct {
		InventoryWriteOffID int `json:"inventoryWriteOffID"`
	}
	SaveInventoryWriteOffResponse struct {
		Status  sharedCommon.Status           `json:"status"`
		Results []SaveInventoryWriteOffResult `json:"records"`
	}

	SaveInventoryTransferResult struct {
		InventoryTransferID int `json:"inventoryTransferID"`
	}
	SaveInventoryTransferResponse struct {
		Status  sharedCommon.Status           `json:"status"`
		Results []SaveInventoryTransferResult `json:"records"`
	}

	SaveInventoryRegistrationBulkItem struct {
		Status  sharedCommon.StatusBulk           `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryRegistrationResponseBulk struct {
		Status    sharedCommon.Status                 `json:"status"`
		BulkItems []SaveInventoryRegistrationBulkItem `json:"requests"`
	}

	ReasonCode struct {
		ReasonID                             int    `json:"reasonID"`
		Name                                 string `json:"name"`
		Added                                int    `json:"added"`
		LastModified                         int    `json:"lastModified"`
		Purpose                              string `json:"purpose"`
		Code                                 string `json:"code"`
		ManualDiscountDisablesPromotionTiers []int  `json:"manualDiscountDisablesPromotionTiers"`
	}

	GetReasonCodesResponse struct {
		Status      sharedCommon.Status `json:"status"`
		ReasonCodes []ReasonCode        `json:"records"`
	}

	GetInventoryTransferResponse struct {
		Status             sharedCommon.Status `json:"status"`
		InventoryTransfers []InventoryTransfer `json:"records"`
	}

	InventoryTransfer struct {
		InventoryTransferID         int                    `json:"inventoryTransferID"`
		InventoryTransferNo         int                    `json:"inventoryTransferNo"`
		CreatorID                   int                    `json:"creatorID"`
		WarehouseFromID             int                    `json:"warehouseFromID"`
		WarehouseToID               int                    `json:"warehouseToID"`
		DeliveryAddressID           int                    `json:"deliveryAddressID"`
		CurrencyCode                string                 `json:"currencyCode"`
		CurrencyRate                string                 `json:"currencyRate"`
		Date                        string                 `json:"date"`
		InventoryTransactionDate    string                 `json:"inventoryTransactionDate"`
		ShippingDateRequested       string                 `json:"shippingDateRequested"`
		ShippingDateActual          string                 `json:"shippingDateActual"`
		Notes                       string                 `json:"notes"`
		Status                      string                 `json:"status"`
		Confirmed                   int                    `json:"confirmed"`
		Added                       int                    `json:"added"`
		LastModified                int                    `json:"lastModified"`
		InventoryTransferOrderID    int                    `json:"inventoryTransferOrderID,omitempty"`
		Type                        string                 `json:"type"`
		Rows                        []InventoryTransferRow `json:"rows,omitempty"`
		FollowupInventoryTransferID int                    `json:"followupInventoryTransferID,omitempty"`
	}

	InventoryTransferRow struct {
		ProductID        int    `json:"productID"`
		Price            string `json:"price"`
		Amount           string `json:"amount"`
		PackageID        int    `json:"packageID"`
		AmountOfPackages string `json:"amountOfPackages"`
		AmountInPackage  int    `json:"amountInPackage"`
		PackageType      string `json:"packageType"`
		PackageTypeID    int    `json:"packageTypeID"`
	}
)
