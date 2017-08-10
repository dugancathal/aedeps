package taskqueue_test

import (
	"github.com/dugancathal/aedeps/taskqueue"
	ae "google.golang.org/appengine/taskqueue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("NonExecutingQueueManager", func() {
	var c context.Context

	BeforeEach(func() {
		c = context.Background()
	})

	It("tracks items being added to a queue", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		qm.Add(c, task, "queue1")

		Expect(qm.HasTaskInQueue(task, "queue1")).To(BeTrue())
	})

	It("can distinguish between multiple queues", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		qm.Add(c, task, "queue1")

		Expect(qm.HasTaskInQueue(task, "queue2")).To(BeFalse())
	})

	It("can delete tasks from a queue", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		qm.Add(c, task, "queue1")

		qm.Delete(c, task, "queue1")

		Expect(qm.HasTaskInQueue(task, "queue1")).To(BeFalse())
	})

	It("only deletes items from the specified queue", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		qm.Add(c, task, "queue1")
		qm.Add(c, task, "queue2")

		err := qm.Delete(c, task, "queue2")
		Expect(err).To(BeNil())

		Expect(qm.HasTaskInQueue(task, "queue1")).To(BeTrue())
		Expect(qm.HasTaskInQueue(task, "queue2")).To(BeFalse())
	})

	It("returns an error when the task does not exist", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}

		err := qm.Delete(c, task, "queue2")

		Expect(err).ToNot(BeNil())
	})

	It("can add multiple items to the same queue at once", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task1 := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		task2 := &taskqueue.Task{
			&ae.Task{Name: "barTask"},
		}
		qm.AddMulti(c, []*taskqueue.Task{task1, task2}, "queue1")
		Expect(qm.HasTaskInQueue(task1, "queue1")).To(BeTrue())
		Expect(qm.HasTaskInQueue(task2, "queue1")).To(BeTrue())
	})

	It("can delete multiple items from the same queue at once", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task1 := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		task2 := &taskqueue.Task{
			&ae.Task{Name: "barTask"},
		}
		qm.AddMulti(c, []*taskqueue.Task{task1, task2}, "queue1")
		qm.DeleteMulti(c, []*taskqueue.Task{task1, task2}, "queue1")

		Expect(qm.QueuesToTasks["queue1"]).To(HaveLen(0))
		Expect(qm.HasTaskInQueue(task1, "queue1")).To(BeFalse())
		Expect(qm.HasTaskInQueue(task2, "queue1")).To(BeFalse())
	})

	It("returns an error that says the tasks that it could not remove", func() {
		qm := taskqueue.NewNonExecutingQueueManager()
		task1 := &taskqueue.Task{
			&ae.Task{Name: "fooTask"},
		}
		task2 := &taskqueue.Task{
			&ae.Task{Name: "barTask"},
		}
		err := qm.DeleteMulti(c, []*taskqueue.Task{task1, task2}, "queue1")
		Expect(err.Error()).To(ContainSubstring("could not delete fooTask, barTask"))
	})
})
