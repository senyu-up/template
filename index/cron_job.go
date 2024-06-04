package index

import "template/internal/time_task"

func RegisterCronJob() (err error) {
	time_task.StartAllTask()
	return nil
}
