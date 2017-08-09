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

func (qm *AeQueueManager) Add(c context.Context, task *Task, queueName string) (*Task, error) {
	aeTask, err := ae.Add(c, task.Task, queueName)
	return &Task{aeTask}, err
}

func (qm *AeQueueManager) Delete(c context.Context, task *Task, queueName string) error {
	return ae.Delete(c, task.Task, queueName)
}
