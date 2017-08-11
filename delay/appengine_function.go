package delay

import (
	"google.golang.org/appengine/delay"
)

func appengineFuncFactory(key string, callable Callable) Function {
	return delay.Func(key, callable)
}

