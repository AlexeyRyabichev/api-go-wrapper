package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/addresses"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	common.Die(err)

	addrs, err := GetAddressesBulk(apiClient)
	common.Die(err)

	fmt.Printf("%+v\n", addrs)

	err = SaveAddressesBulk(apiClient)
	common.Die(err)

	DeleteAddress(apiClient)

	DeleteAddressBulk(apiClient)

	CheckAddressListing(apiClient)
}

func GetAddressesBulk(cl *api.Client) (addresses []sharedCommon.Address, err error) {
	addressCli := cl.AddressProvider

	bulkFilters := []map[string]interface{}{
		{
			"recordsOnPage": 2,
			"pageNo":        1,
		},
		{
			"recordsOnPage": 2,
			"pageNo":        2,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := addressCli.GetAddressesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, supplier := range bulkItem.Addresses {
			addresses = append(addresses, supplier)
		}
	}

	return
}

func SaveAddressesBulk(cl *api.Client) (err error) {
	addressProvider := cl.AddressProvider

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	dataToUpdate := []map[string]interface{}{
		{
			"street":     "Elmstr",
			"city":       "Rome",
			"postalCode": "123456",
			"country":    "Italy",
			"ownerID":    12355,
			"typeID":     1,
		},
	}
	bulkResponse, err := addressProvider.SaveAddressesBulk(ctx, dataToUpdate, map[string]string{})
	if err != nil {
		return
	}

	fmt.Printf("%+v", bulkResponse)

	return
}

func DeleteAddress(cl *api.Client) {
	cli := cl.AddressProvider

	req := map[string]string{
		"addressID": "6320",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := cli.DeleteAddress(ctx, req)
	common.Die(err)
}

func DeleteAddressBulk(cl *api.Client) {
	addressProvider := cl.AddressProvider

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	addressesToDelete := []map[string]interface{}{
		{
			"addressID":     7238,
		},
		{
			"addressID":     7239,
		},
	}
	bulkResponse, err := addressProvider.DeleteAddressBulk(ctx, addressesToDelete, map[string]string{})
	common.Die(err)

	fmt.Printf("%s", common.ConvertSourceToJsonStrIfPossible(bulkResponse))

	return
}

func CheckAddressListing(cl *api.Client) {
	addressesDataProvider := addresses.NewAddressListingDataProvider(cl.AddressProvider)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        300,
			MaxFetchersCount:          10,
		},
		addressesDataProvider,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	addressesChan := lister.Get(ctx, map[string]interface{}{
		"recordsOnPage": 1,
	})

	addrs := make([]sharedCommon.Address, 0)
	for addr := range addressesChan {
		common.Die(addr.Err)
		addrs = append(addrs, addr.Payload.(sharedCommon.Address))
	}

	fmt.Println(common.ConvertSourceToJsonStrIfPossible(addrs))
}
