package ctxfactory

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func AppEngineFactory(r *http.Request) context.Context {
	ctx := appengine.NewContext(r)
	return ctx
}
