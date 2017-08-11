package taskqueue

import (
	"golang.org/x/net/context"
	ae "google.golang.org/appengine/taskqueue"
)

type QueueManager interface {
	Add(c context.Context, task *ae.Task, queueName string) (*ae.Task, error)
	AddMulti(c context.Context, task []*ae.Task, queueName string) ([]*ae.Task, error)
	Delete(c context.Context, task *ae.Task, queueName string) error
	DeleteMulti(c context.Context, task []*ae.Task, queueName string) error
}
