package delay

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"
)

// BaseFunction is merely a named function that matches
// the interface of an appengine "Function". Any calls to
// its Call method will execute the underlying "delayed"
// function immediately.
type BaseFunction struct {
	key string
	actual Callable
}

func baseFuncFactory(key string, callable Callable) Function {
	return &BaseFunction{
		key: key,
		actual: callable,
	}
}

func (f *BaseFunction) Call(ctx context.Context, args... interface{}) error {
	return f.actual(ctx, args...)
}

// Task returns a pointer to a zero-value taskqueue.Task.
func (f *BaseFunction) Task(args... interface{}) (*taskqueue.Task, error) {
	return &taskqueue.Task{}, nil
}
