package dataservice

import context "context"

type deleteAllKey struct{}
type deleteAllHeader struct{}

func Ctx4DeleteAll(ctx context.Context, header, pw string) context.Context {
	return context.WithValue(context.WithValue(ctx, deleteAllHeader{}, header), deleteAllKey{}, pw)
}

func GetDeleteAllPw(ctx context.Context) string {
	v := ctx.Value(deleteAllKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func GetDeleteAllHeader(ctx context.Context) string {
	v := ctx.Value(deleteAllHeader{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func ContextDataWPw(ctx context.Context) map[string]string {
	data := make(map[string]string)
	data[GetDeleteAllHeader(ctx)] = GetDeleteAllPw(ctx)
	return data
}
