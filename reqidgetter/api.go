// A package for retrieving the request ID from a context.
//
// The "base" implementation uses a publicly exported key
// (hey, the whole idea is to minimize lock-in), and a function
// to set the request ID to some value if you choose.
//
// Setting the request ID might be done in some form of middleware,
// while retrieving might be done by something like a logger.
package reqidgetter

import "golang.org/x/net/context"

var requestIDGetter RequestIDGetter

// Factory() returns the currently set RequestIDGetter
// All calls to get the RequestId should go through this function
func RequestID(ctx context.Context) string {
	return requestIDGetter(ctx)
}

func UseAeRequestIDGetter() {
	requestIDGetter = AppEngineRequestIDGetter
}

func UseBaseRequestIDGetter() {
	requestIDGetter = BaseRequestIDGetter
}