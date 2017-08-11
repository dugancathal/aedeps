package ctxfactory

import (
	"net/http"

	"golang.org/x/net/context"
)

// A function type that returns NEW contexts for an HTTP request.
// Use this type to inject a context creator rather than using the
// built in context.Background or AppEngine's NewContext.
type ContextFactory func(r *http.Request) context.Context
