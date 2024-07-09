package api_client_go

import (
	"context"
	"fmt"

	"github.com/air-iot/api-client-go/v4/algorithm"
	"github.com/air-iot/api-client-go/v4/api"
	"github.com/air-iot/api-client-go/v4/apicontext"
	"github.com/air-iot/api-client-go/v4/config"
	internalError "github.com/air-iot/api-client-go/v4/errors"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
	"github.com/air-iot/logger"
)

// AlgorithmRunById 算法执行
func (c *Client) AlgorithmRunById(ctx context.Context, projectId, id string, data interface{}) ([]byte, error) {

	cli, err := c.AlgorithmClient.GetAlgorithmServiceClient()
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("数据为空")
	}

	bts, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("数据为空")
	}

	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	res, err := cli.Run(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&algorithm.ClientRunByIdRequest{
			Id:   id,
			Data: bts,
		})
	if err != nil {
		return nil, errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		logger.Errorf("算法服务grpc响应错误: %+v", res)
		return nil, internalError.ParseResponse(res)
	}

	return res.GetResult(), nil
}

func (c *Client) QueryLocalAlgorithm(ctx context.Context, projectId string, query, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return err
	}
	res, err := cli.Query(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	if result != nil {
		if err := json.Unmarshal(res.GetResult(), result); err != nil {
			return errors.Wrap(err, "解析请求结果错误")
		}
	}
	return nil
}

func (c *Client) GetLocalAlgorithm(ctx context.Context, projectId, id string, result interface{}) ([]byte, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return nil, fmt.Errorf("id为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
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

func (c *Client) DeleteLocalAlgorithm(ctx context.Context, projectId, id string) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
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
	return nil
}

func (c *Client) UpdateLocalAlgorithm(ctx context.Context, projectId, id string, updateData interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}

	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return err
	}

	bts, err := json.Marshal(updateData)
	if err != nil {
		return err
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
	return nil
}

func (c *Client) ReplaceLocalAlgorithm(ctx context.Context, projectId, id string, updateData interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return err
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
	return nil
}

func (c *Client) CreateLocalAlgorithm(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return err
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

func (c *Client) CreateManyLocalAlgorithm(ctx context.Context, projectId string, createData, result interface{}) error {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return err
	}
	res, err := cli.CreateMany(
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

func (c *Client) DeleteManyLocalAlgorithm(ctx context.Context, projectId string, query interface{}) (int64, error) {
	if projectId == "" {
		projectId = config.XRequestProjectDefault
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return 0, errors.Wrap(err, "序列化查询参数错误")
	}
	cli, err := c.AlgorithmClient.GetLocalAlgorithmClient()
	if err != nil {
		return 0, err
	}
	res, err := cli.DeleteMany(
		apicontext.GetGrpcContext(ctx, map[string]string{config.XRequestProject: projectId}),
		&api.QueryRequest{Query: bts})
	if err != nil {
		return 0, errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return 0, internalError.ParseResponse(res)
	}
	return res.Count, nil
}
