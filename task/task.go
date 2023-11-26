package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

/*
It is the smallest unit of work in an orchestration system. It is at its basics
a process that runs on a single machine. Can run an instance like Nginx or a RESTFul API/
It specifies the following:
	1-Memory, CPU, Disk
	2-Restart Policy in case of orchestration failure
	3-Name of container image used to run the task
It has secondary definitions but those 3 were the fundamental definitions
*/

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
