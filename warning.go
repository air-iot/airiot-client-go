package api_client_go

import (
	"context"
	internalError "github.com/air-iot/api-client-go/v4/errors"

	"github.com/air-iot/api-client-go/v4/api"
	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/api-client-go/v4/warning"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
)

// QueryWarn 查询
func (c *Client) QueryWarn(ctx context.Context, projectId, token, archive string, query interface{}, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数错误")
	}
	cli, err := c.WarningClient.GetWarnServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId, config.XRequestHeaderAuthorization: token}),
		&warning.QueryWarningRequest{Query: bts, Archive: archive})
	if err != nil {
		return 0, errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return 0, internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return 0, errors.Wrap(err, "解析请求结果错误")
	}
	return int(res.Count), nil
}

func (c *Client) GetWarn(ctx context.Context, projectId, archive, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.WarningClient.GetWarnServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&warning.GetOrDeleteWarningRequest{Id: id, Archive: archive})
	if err != nil {
		return nil, errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return nil, internalError.ParseResponse(res)
	}
	if result == nil {
		return res.GetResult(), nil
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return nil, errors.Wrap(err, "解析请求结果错误")
	}
	return res.GetResult(), nil
}

func (c *Client) BatchCreateWarn(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.WarningClient.GetWarnServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.BatchCreate(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}

// QueryRule 查询
func (c *Client) QueryRule(ctx context.Context, projectId string, query interface{}, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数错误")
	}
	cli, err := c.WarningClient.GetRuleServiceClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if err != nil {
		return 0, errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return 0, internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return 0, errors.Wrap(err, "解析请求结果错误")
	}
	return int(res.Count), nil
}

func (c *Client) CreateWarn(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.WarningClient.GetWarnServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}

func (c *Client) BatchCreateRule(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.WarningClient.GetRuleServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.BatchCreate(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.CreateRequest{Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}

func (c *Client) DeleteRule(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.WarningClient.GetRuleServiceClient()
	if err != nil {
		return err
	}

	res, err := cli.Delete(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}

func (c *Client) UpdateRule(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.WarningClient.GetRuleServiceClient()
	if err != nil {
		return err
	}

	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.Update(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}

func (c *Client) ReplaceRule(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.WarningClient.GetRuleServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.Replace(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateRequest{Id: id, Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if err := json.Unmarshal(res.GetResult(), result); err != nil {
		return errors.Wrap(err, "解析请求结果错误")
	}
	return nil
}
