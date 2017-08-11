package reqidgetter

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func AppEngineRequestIDGetter(ctx context.Context) string {
	return appengine.RequestID(ctx)
}

