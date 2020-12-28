package company

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"testing"
)

//works
func TestErplyClient_GetConfParameters(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	resp, err := cli.GetConfParameters(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
