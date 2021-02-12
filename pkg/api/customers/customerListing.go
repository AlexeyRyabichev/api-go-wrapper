package customers

import (
	"context"
)

type CustomerListingDataProvider struct {
	erplyAPI Manager
}

func NewCustomerListingDataProvider(erplyClient Manager) *CustomerListingDataProvider {
	return &CustomerListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (l *CustomerListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := l.erplyAPI.GetCustomersBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (l *CustomerListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := l.erplyAPI.GetCustomersBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for i := range bulkItem.Customers {
			callback(bulkItem.Customers[i])
		}
	}

	return nil
}
