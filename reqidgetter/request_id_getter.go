package reqidgetter

import "golang.org/x/net/context"

// A function type that returns the request ID from a context.
// Use this type to inject a request ID getter.
type RequestIDGetter func(c context.Context) string
