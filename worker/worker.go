package worker

import (
	"fmt"

	"github.com/arka-cell/cube/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/*
If we think of the task as the foundation of this orchestration system, then we
can think of the worker as the next layer that sits atop the foundation. Let’s
remind ourselves what the worker’s requirements are:
 1. Run tasks as Docker containers.
 2. Accept tasks to run from a manager.
 3. Provide relevant statistics to the manager for the purpose of scheduling
    tasks.
 4. Keep track of its tasks and their state
*/
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Collecing stats")
}

func (w *Worker) RunTask() {
	fmt.Println("Will run/stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("Starting a task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stopping a task")
}
