package api_client_go

import (
	"context"

	"github.com/air-iot/api-client-go/v4/api"
	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/config"
	internalError "github.com/air-iot/api-client-go/v4/errors"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
)

// CreateFlowTask FlowTask
func (c *Client) CreateFlowTask(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

func (c *Client) GetFlowTask(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
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

func (c *Client) QueryFlowTask(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数错误")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

// BatchCreateFlowTask FlowTask
func (c *Client) BatchCreateFlowTask(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

func (c *Client) DeleteFlowTask(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

func (c *Client) UpdateFlowTask(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

func (c *Client) UpdateFlowTaskFilter(ctx context.Context, projectId string, query, updateData interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	queryBts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数错误")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
	if err != nil {
		return err
	}

	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.UpdateFilter(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.UpdateFilter{Query: queryBts, Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	return nil
}

func (c *Client) ReplaceFlowTask(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.FlowClient.GetFlowTaskServiceClient()
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

// QueryFlow Flow
func (c *Client) QueryFlow(ctx context.Context, projectId string, query, result interface{}) (int, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数错误")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
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

func (c *Client) GetFlow(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.GetOrDeleteRequest{Id: id})
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

func (c *Client) UpdateFlow(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
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

func (c *Client) BatchCreateFlow(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
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

// CreateFlowTriggerRecord FlowTriggerRecord
func (c *Client) CreateFlowTriggerRecord(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.FlowClient.GetFlowTriggerRecordServiceClient()
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

func (c *Client) DeleteFlow(ctx context.Context, projectId, id string, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
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

func (c *Client) ReplaceFlow(ctx context.Context, projectId, id string, updateData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.FlowClient.GetFlowServiceClient()
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
