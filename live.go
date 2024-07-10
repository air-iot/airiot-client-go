package api_client_go

import (
	"context"

	"github.com/air-iot/api-client-go/v4/api"
	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
)

func (c *Client) RtspPull(ctx context.Context, projectId string, createData interface{}) (string, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return "", errors.New("插入数据为空")
	}
	cli, err := c.LiveClient.GetLiveServiceClient()
	if err != nil {
		return "", err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return "", errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	result, err := parseRes(err, res, nil)
	if err != nil {
		return "", err
	}
	if result != nil || string(result) == "" {
		return string(result), nil
	} else {
		return "", errors.New("result中path为空")
	}
}
