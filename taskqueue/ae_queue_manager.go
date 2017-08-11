package taskqueue

import (
	"golang.org/x/net/context"
	ae "google.golang.org/appengine/taskqueue"
)

type AeQueueManager struct {
}

func NewAeQueueManager() *AeQueueManager {
	return &AeQueueManager{}
}

func (qm *AeQueueManager) Add(c context.Context, task *ae.Task, queueName string) (*ae.Task, error) {
	return ae.Add(c, task, queueName)
}

func (qm *AeQueueManager) AddMulti(c context.Context, tasks []*ae.Task, queueName string) ([]*ae.Task, error) {
	return ae.AddMulti(c, tasks, queueName)
}

func (qm *AeQueueManager) Delete(c context.Context, task *ae.Task, queueName string) error {
	return ae.Delete(c, task, queueName)
}

func (qm *AeQueueManager) DeleteMulti(c context.Context, tasks []*ae.Task, queueName string) error {
	return ae.DeleteMulti(c, tasks, queueName)
}
