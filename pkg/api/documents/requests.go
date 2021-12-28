package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
	"math"
)

func (cli *Client) GetPurchaseDocuments(ctx context.Context, filters map[string]string) ([]PurchaseDocument, error) {
	docs := make([]PurchaseDocument, 0)

	resp, err := cli.SendRequest(ctx, "getPurchaseDocuments", filters)
	if err != nil {
		return nil, err
	}
	var res GetPurchaseDocumentsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []PurchaseDocument{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetPurchaseDocumentsResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	docs = res.PurchaseDocuments
	if res.Status.RecordsInResponse < res.Status.RecordsTotal {
		totalPages := int(math.Ceil(float64(res.Status.RecordsTotal) / float64(res.Status.RecordsInResponse)))
		newFilters := make([]map[string]interface{}, 0)

		bulkNewFilters := make([][]map[string]interface{}, 0)

		for i := 2; i <= totalPages; i++ {
			cpyFilter := make(map[string]interface{})
			for k, v := range filters {
				cpyFilter[k] = v
			}
			cpyFilter["pageNo"] = i
			newFilters = append(newFilters, cpyFilter)

			if i%20 == 0 {
				bulkNewFilters = append(bulkNewFilters, newFilters)
				newFilters = make([]map[string]interface{}, 0)
			}
		}

		for _, newFilters := range bulkNewFilters {
			secondaryResp, err := cli.GetPurchaseDocumentsBulk(ctx, newFilters, map[string]string{})
			if err != nil {
				return []PurchaseDocument{}, err
			}

			for _, bulkItem := range secondaryResp.BulkItems {
				docs = append(docs, bulkItem.PurchaseDocuments...)
			}
		}
	}

	return docs, nil
}

func (cli *Client) GetPurchaseDocumentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPurchaseDocumentResponseBulk, error) {
	var bulkResp GetPurchaseDocumentResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getPurchaseDocuments",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetPurchaseDocumentResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, bulkResp.Status.ErrorCode)
		}
	}

	return bulkResp, nil
}
