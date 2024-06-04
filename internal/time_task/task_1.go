package time_task

import (
	"context"
	"github.com/roylee0704/gron"
	"github.com/senyu-up/toolbox/tool/cron"
	"github.com/senyu-up/toolbox/tool/logger"
	"github.com/senyu-up/toolbox/tool/su_logger"
	"time"
)

// TaskTest 定时任务测试
func TaskTest() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logger.Error("task panic:", panicErr)
		}
	}()
	// 定时任务逻辑
	cron.Register(gron.Every(10*time.Second), LogInfo)

}

func LogInfo() {
	su_logger.Info(context.Background(), "定时任务执行中...")
}
