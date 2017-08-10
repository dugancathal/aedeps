package taskqueue

import "golang.org/x/net/context"

type QueueManager interface {
	Add(c context.Context, task *Task, queueName string) (*Task, error)
	AddMulti(c context.Context, task []*Task, queueName string) ([]*Task, error)
	Delete(c context.Context, task *Task, queueName string) error
	DeleteMulti(c context.Context, task []*Task, queueName string) error
}
