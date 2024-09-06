package apicontext

import (
	"context"

	"github.com/air-iot/api-client-go/v4/config"
	"google.golang.org/grpc/metadata"
)

type tokenKey struct{}

type projectKey struct{}

func SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

func SetProjectID(ctx context.Context, projectId string) context.Context {
	return context.WithValue(ctx, projectKey{}, projectId)
}

func GetGrpcContext(ctx context.Context, data map[string]string) context.Context {
	md := metadata.New(data)
	// 发送 metadata
	// 创建带有meta的context
	ctx = metadata.NewOutgoingContext(ctx, md)
	token := ctx.Value(tokenKey{})
	tokenString, ok := token.(string)
	if ok {
		if tokenString != "" {
			return metadata.AppendToOutgoingContext(ctx, config.XRequestHeaderAuthorization, tokenString)
		}
	}
	project := ctx.Value(projectKey{})
	projectString, ok := project.(string)
	if ok {
		if projectString != "" {
			return metadata.AppendToOutgoingContext(ctx, config.XRequestProject, projectString)
		}
	}
	return ctx
}

func GetGrpcInContext(ctx context.Context, data map[string]string) context.Context {
	md := metadata.New(data)
	// 发送 metadata
	// 创建带有meta的context
	return metadata.NewIncomingContext(ctx, md)
}
