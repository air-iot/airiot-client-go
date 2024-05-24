package api_client_go

import (
	"context"
	"fmt"

	"github.com/air-iot/api-client-go/v4/errors"
	"github.com/air-iot/api-client-go/v4/jsserver"
	"github.com/air-iot/json"
)

func (c *Client) RunJsScript(ctx context.Context, variables interface{}, script string) ([]byte, error) {
	cli, err := c.JsServerClient.GetScriptClient()
	if err != nil {
		return nil, errors.NewMsg("获取客户端错误,%s", err)
	}
	b, err := json.Marshal(variables)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error: %w", err)
	}
	res, err := cli.Run(ctx, &jsserver.Request{
		Content: script,
		Params:  b,
	})
	if err != nil {
		return nil, fmt.Errorf("client.Run error: %w", err)
	}
	if res.Status {
		return res.Result, nil
	} else {
		return nil, fmt.Errorf("返回错误: info=%s,detail=%s", res.Info, res.Detail)
	}
}
