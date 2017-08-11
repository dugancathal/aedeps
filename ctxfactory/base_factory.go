package ctxfactory

import (
	"net/http"

	"golang.org/x/net/context"
)

func BaseFactory(r *http.Request) context.Context {
	return context.Background()
}
