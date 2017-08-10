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

func (qm *AeQueueManager) AddMulti(c context.Context, tasks []*Task, queueName string) ([]*Task, error) {
	aeTasks := make([]*ae.Task, len(tasks))
	for i, task := range tasks {
		aeTasks[i] = task.Task
	}
	addedTasks, err := ae.AddMulti(c, aeTasks, queueName)
	addedAeTasks := make([]*Task, len(tasks))
	for i, task := range addedTasks {
		addedAeTasks[i] = &Task{task}
	}
	return addedAeTasks, err
}

func (qm *AeQueueManager) Delete(c context.Context, task *Task, queueName string) error {
	return ae.Delete(c, task.Task, queueName)
}

func (qm *AeQueueManager) DeleteMulti(c context.Context, tasks []*Task, queueName string) error {
	aeTasks := make([]*ae.Task, len(tasks))
	for i, task := range tasks {
		aeTasks[i] = task.Task
	}
	return ae.DeleteMulti(c, aeTasks, queueName)
}
