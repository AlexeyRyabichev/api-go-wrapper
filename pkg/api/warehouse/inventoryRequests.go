package warehouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
	"net/http"
)

func (cli *Client) SaveInventoryRegistration(ctx context.Context, filters map[string]string) (inventoryRegistrationID int, err error) {
	resp, err := cli.SendRequest(ctx, "saveInventoryRegistration", filters)
	if err != nil {
		return 0, sharedCommon.NewFromError("saveInventoryRegistration: error sending POST request", err, 0)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, sharedCommon.NewFromError(fmt.Sprintf("saveInventoryRegistration: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	respData := SaveInventoryRegistrationResponse{}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, sharedCommon.NewFromError("saveInventoryRegistration: error decoding JSON response body", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, sharedCommon.NewFromError(fmt.Sprintf("saveInventoryRegistration: API error %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
	}
	if len(respData.Results) < 1 {
		return 0, sharedCommon.NewFromError("saveInventoryRegistration: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Results[0].InventoryRegistrationID, nil
}

func (cli *Client) SaveInventoryRegistrationBulk(
	ctx context.Context,
	bulkRequest []map[string]interface{},
	baseFilters map[string]string) (
	SaveInventoryRegistrationResponseBulk,
	error,
) {
	var bulkResp SaveInventoryRegistrationResponseBulk

	if len(bulkRequest) > sharedCommon.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d inventory registrations in one bulk request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveInventoryRegistration",
			Filters:    bulkInput,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveInventoryRegistrationResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
				bulkResp.Status.ErrorCode,
			)
		}
	}

	return bulkResp, nil
}
