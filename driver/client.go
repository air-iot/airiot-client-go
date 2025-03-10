package driver

import (
	"github.com/air-iot/errors"
	"sync"

	"github.com/air-iot/api-client-go/v4/config"
	"github.com/air-iot/api-client-go/v4/conn"
	"github.com/air-iot/logger"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
)

const serviceName = "driver"

type Client struct {
	lock        sync.RWMutex
	registry    *etcd.Registry
	config      config.Config
	conn        *grpc.ClientConn
	restClient  *http.Client
	opts        []grpc.DialOption
	middlewares []middleware.Middleware

	driverClient                    DriverServiceClient
	driverInstanceServiceClient     DriverInstanceServiceClient
	driverEventCronServiceClient    DriverEventCronServiceClient
	driverInstructCronServiceClient DriverInstructCronServiceClient
	driverInstructServiceClient     DriverInstructServiceClient
}

func NewClient(cfg config.Config, registry *etcd.Registry, cred grpc.DialOption, httpCred middleware.Middleware) (*Client, func(), error) {
	c := &Client{
		registry:    registry,
		config:      cfg,
		opts:        []grpc.DialOption{cred},
		middlewares: []middleware.Middleware{httpCred},
	}
	//if err := c.createRestConn(); err != nil {
	//	return nil, nil, err
	//}
	//if err := c.createConn(); err != nil {
	//	return nil, nil, err
	//}
	cleanFunc := func() {
		if c.restClient != nil {
			if err := c.restClient.Close(); err != nil {
				logger.Errorf("rest close error: %s", err.Error())
			}
		}
		if c.conn != nil {
			if err := c.conn.Close(); err != nil {
				logger.Errorf("grpc close error: %s", err.Error())
			}
		}
	}
	return c, cleanFunc, nil
}

func (c *Client) createConn() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.conn != nil {
		return nil
	}
	logger.Infof("%s grpc client cc, %+v", serviceName, c.config)
	cc, err := conn.CreateConn(serviceName, c.config, c.registry, c.opts...)
	if err != nil {
		return err
	}
	c.driverClient = NewDriverServiceClient(cc)
	c.driverInstanceServiceClient = NewDriverInstanceServiceClient(cc)
	c.driverEventCronServiceClient = NewDriverEventCronServiceClient(cc)
	c.driverInstructCronServiceClient = NewDriverInstructCronServiceClient(cc)
	c.driverInstructServiceClient = NewDriverInstructServiceClient(cc)
	c.conn = cc
	return nil
}

func (c *Client) createRestConn() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.restClient != nil {
		return nil
	}
	logger.Infof("%s http client createConn, %+v", serviceName, c.config)
	cc, err := conn.CreateRestConn(serviceName, c.config, c.registry, c.middlewares...)
	if err != nil {
		return err
	}
	c.restClient = cc
	return nil
}

func (c *Client) GetRestClient() (*http.Client, error) {
	if c.restClient == nil {
		if err := c.createRestConn(); err != nil {
			return nil, err
		}
	}
	return c.restClient, nil
}

func (c *Client) GetDriverServiceClient() (DriverServiceClient, error) {
	if c.conn == nil {
		if err := c.createConn(); err != nil {
			return nil, err
		}
	}
	if c.driverClient == nil {
		return nil, errors.New("客户端是空")
	}
	return c.driverClient, nil
}

func (c *Client) GetDriverInstanceServiceClient() (DriverInstanceServiceClient, error) {
	if c.conn == nil {
		if err := c.createConn(); err != nil {
			return nil, err
		}
	}
	if c.driverInstanceServiceClient == nil {
		return nil, errors.New("客户端是空")
	}
	return c.driverInstanceServiceClient, nil
}

func (c *Client) GetDriverEventCronServiceClient() (DriverEventCronServiceClient, error) {
	if c.conn == nil {
		if err := c.createConn(); err != nil {
			return nil, err
		}
	}
	if c.driverEventCronServiceClient == nil {
		return nil, errors.New("客户端是空")
	}
	return c.driverEventCronServiceClient, nil
}

func (c *Client) GetDriverInstructCronServiceClient() (DriverEventCronServiceClient, error) {
	if c.conn == nil {
		if err := c.createConn(); err != nil {
			return nil, err
		}
	}
	if c.driverInstructCronServiceClient == nil {
		return nil, errors.New("客户端是空")
	}
	return c.driverInstructCronServiceClient, nil
}

func (c *Client) GetDriverInstructServiceClient() (DriverInstructServiceClient, error) {
	if c.conn == nil {
		if err := c.createConn(); err != nil {
			return nil, err
		}
	}
	if c.driverInstructServiceClient == nil {
		return nil, errors.New("客户端是空")
	}
	return c.driverInstructServiceClient, nil
}
