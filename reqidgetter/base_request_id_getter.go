package reqidgetter

import (
	"golang.org/x/net/context"
)

const REQUEST_ID_KEY = "reqidgetter.request-id"

func BaseRequestIDGetter(ctx context.Context) string {
	return ctx.Value(REQUEST_ID_KEY).(string)
}

func BaseSetRequestId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, REQUEST_ID_KEY, reqId)
}
