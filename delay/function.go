package delay

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"
)

type Function interface {
	Call(ctx context.Context, args... interface{}) error
	Task(args ...interface{}) (*taskqueue.Task, error)
}
