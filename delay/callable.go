package delay

import "golang.org/x/net/context"

type Callable func(ctx context.Context, args... interface{}) error
