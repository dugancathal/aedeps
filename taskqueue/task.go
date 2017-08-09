package taskqueue

import (
	"google.golang.org/appengine/taskqueue"
)

type Task struct {
	*taskqueue.Task
}
