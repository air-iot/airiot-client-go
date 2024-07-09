package api_client_go

import (
	"context"
	netHttp "net/http"
	"net/url"

	"github.com/air-iot/api-client-go/v4/api"
	internalError "github.com/air-iot/api-client-go/v4/errors"
	"github.com/air-iot/errors"
	"github.com/air-iot/json"
)

func (c *Client) QueryProject(ctx context.Context, query, result interface{}) error {
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数错误")
	}
	res, err := cli.Query(ctx, &api.QueryRequest{Query: bts})
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

func (c *Client) QueryProjectAvailable(ctx context.Context, result interface{}) error {
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.QueryAvailable(ctx, &api.EmptyRequest{})
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

func (c *Client) RestQueryProject(ctx context.Context, query, result interface{}) error {
	u := url.URL{Path: "/spm/project"}
	if query != nil {
		bts, err := json.Marshal(query)
		if err != nil {
			return errors.Wrap(err, "序列化查询参数错误")
		}
		params := url.Values{}
		params.Set("query", string(bts))
		u.RawQuery = params.Encode()
	}
	cli, err := c.SpmClient.GetRestClient()
	if err != nil {
		return err
	}
	if err := cli.Invoke(ctx, netHttp.MethodGet, u.RequestURI(), map[string]interface{}{}, result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetProject(ctx context.Context, id string, result interface{}) ([]byte, error) {
	if id == "" {
		return nil, errors.New("id为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return nil, err
	}
	res, err := cli.Get(ctx, &api.GetOrDeleteRequest{Id: id})
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

func (c *Client) DeleteProject(ctx context.Context, id string, result interface{}) error {
	if id == "" {
		return errors.New("id为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	res, err := cli.Delete(ctx, &api.GetOrDeleteRequest{Id: id})
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

func (c *Client) UpdateProject(ctx context.Context, id string, updateData, result interface{}) error {
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.Update(ctx, &api.UpdateRequest{Id: id, Data: bts})
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

func (c *Client) UpdateProjectLicense(ctx context.Context, id string, updateData, _ interface{}) error {
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.UpdateLicense(ctx, &api.UpdateRequest{Id: id, Data: bts})
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}
	if !res.GetStatus() {
		return internalError.ParseResponse(res)
	}
	//if err := json.Unmarshal(res.GetResult(), result); err != nil {
	//	return errors.Wrap(err, "解析请求结果错误")
	//}
	return nil
}

func (c *Client) ReplaceProject(ctx context.Context, id string, updateData, result interface{}) error {
	if id == "" {
		return errors.New("id为空")
	}
	if updateData == nil {
		return errors.New("更新数据为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(updateData)
	if err != nil {
		return errors.Wrap(err, "序列化更新数据错误")
	}
	res, err := cli.Replace(ctx, &api.UpdateRequest{Id: id, Data: bts})
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

func (c *Client) CreateProject(ctx context.Context, createData, result interface{}) error {
	if createData == nil {
		return errors.New("插入数据为空")
	}
	cli, err := c.SpmClient.GetProjectServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(createData)
	if err != nil {
		return errors.Wrap(err, "序列化插入数据错误")
	}
	res, err := cli.Create(ctx, &api.CreateRequest{Data: bts})
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

func (c *Client) QueryPmSetting(ctx context.Context, query, result interface{}) error {
	cli, err := c.SpmClient.GetSettingServiceClient()
	if err != nil {
		return err
	}
	bts, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "序列化查询参数错误")
	}
	res, err := cli.Query(ctx, &api.QueryRequest{Query: bts})
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
