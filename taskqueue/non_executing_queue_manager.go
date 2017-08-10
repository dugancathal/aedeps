package taskqueue

import (
	"errors"

	"golang.org/x/net/context"
	"fmt"
	"strings"
)

//
// A queue manager that merely stores the tasks given to it.
// This queue manager is particularly good for testing. There is an additional
// method on this manager (HasTaskInQueue) that will let you verify if a task is present by name.
//
//   // Get a QueueManager somehow
//   func ExampleNonExecutingQueueManager_TestExample() {
//      qm := taskqueue.GetQueueManager()
//   	nonExecutingQueueManager := qm.(NonExecutingQueueManager)
//   	hasItemInQueue := nonExecutingQueueManager.HasTaskInQueue("taskName", "queueName")
//   }
//
// If you're allowing dependency injection (manual or otherwise) in your application architecture,
// this implementation can also be used locally or in a staging environment when you may not want
// tasks to actually get queued up.
//
type NonExecutingQueueManager struct {
	QueuesToTasks map[string][]*Task
}

func NewNonExecutingQueueManager() *NonExecutingQueueManager {
	return &NonExecutingQueueManager{
		QueuesToTasks: make(map[string][]*Task),
	}
}

func (qm *NonExecutingQueueManager) Add(c context.Context, task *Task, queueName string) (*Task, error) {
	qm.createQueueIfNotExists(queueName)

	qm.QueuesToTasks[queueName] = append(qm.QueuesToTasks[queueName], task)
	return task, nil
}

func (qm *NonExecutingQueueManager) AddMulti(c context.Context, tasks []*Task, queueName string) ([]*Task, error) {
	qm.createQueueIfNotExists(queueName)

	qm.QueuesToTasks[queueName] = append(qm.QueuesToTasks[queueName], tasks...)
	return tasks, nil
}

func (qm *NonExecutingQueueManager) Delete(c context.Context, task *Task, queueName string) error {
	taskIndex := qm.taskIndexInQueue(task, queueName)

	if taskIndex == -1 {
		return errors.New("unable to delete task that does not exist")
	}

	qm.QueuesToTasks[queueName] = append(qm.QueuesToTasks[queueName][:taskIndex], qm.QueuesToTasks[queueName][taskIndex+1:]...)

	return nil
}

func (qm *NonExecutingQueueManager) DeleteMulti(c context.Context, tasks []*Task, queueName string) error {
	undeletableTaskNames := []string{}
	for _, task := range tasks {
		taskIndex := qm.taskIndexInQueue(task, queueName)
		if taskIndex == -1 {
			undeletableTaskNames = append(undeletableTaskNames, task.Name)
			continue
		}
		qm.QueuesToTasks[queueName] = append(qm.QueuesToTasks[queueName][:taskIndex], qm.QueuesToTasks[queueName][taskIndex+1:]...)
	}

	if len(undeletableTaskNames) > 0 {
		return fmt.Errorf("could not delete %s", strings.Join(undeletableTaskNames, ", "))
	}

	return nil
}

func (qm *NonExecutingQueueManager) HasTaskInQueue(task *Task, queueName string) bool {
	return qm.taskIndexInQueue(task, queueName) > -1
}

func (qm *NonExecutingQueueManager) createQueueIfNotExists(queueName string) {
	_, hasQueue := qm.QueuesToTasks[queueName]
	if !hasQueue {
		qm.QueuesToTasks[queueName] = []*Task{}
	}
}

func (qm *NonExecutingQueueManager) taskIndexInQueue(task *Task, queueName string) int {
	for i, t := range qm.QueuesToTasks[queueName] {
		if task.Name == t.Name {
			return i
		}
	}
	return -1
}
