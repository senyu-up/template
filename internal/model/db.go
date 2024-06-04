package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/senyu-up/toolbox/tool/logger"
	"github.com/senyu-up/toolbox/tool/redis_lock"
	"github.com/senyu-up/toolbox/tool/runtime"
	"strings"
	"template/global"
	"template/internal/enum"
)

var models = make([]interface{}, 0)

func Register(model ...interface{}) {
	if len(model) == 0 {
		return
	}
	models = append(models, model...)
}

// DbAutoMigrate 异步的去同步表结构避免启动时初始化数据库卡住
func DbAutoMigrate(stage string, appKeys ...string) {
	task := func() {
		if len(appKeys) > 0 {
			for _, appKey := range appKeys {
				migrate(stage, appKey)
			}
		} else {
			migrate(stage, "")
		}
	}

	runtime.GOSafe(context.Background(), "DbAutoMigrate", task)
	return
}

func migrate(stage string, appKey string) (bool, error) {
	dbIns := global.GetFacade().GetMysqlClient()
	if dbIns == nil {
		return false, errors.New("获取实例失败")
	}
	dbIns = dbIns.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")
	redisLock, err := redis_lock.NewDistributeLockRedis(fmt.Sprintf("gm_manager_db_migrate:%s:%s", stage, appKey), 600, "1")
	if err != nil {
		return false, errors.New("获取锁失败")
	}
	defer func() {
		_ = redisLock.Unlock()
	}()

	errorList := make([]string, 0)
	for _, model := range models {
		if dbIns.Migrator().HasTable(model) && (stage == enum.Production || stage == enum.Master) {
			continue
		}
		if stage == enum.Production || stage == enum.Master {
			err = dbIns.Migrator().AutoMigrate(model)
		} else {
			err = dbIns.AutoMigrate(model)
		}
		if err != nil {
			errorList = append(errorList, err.Error())
			logger.Warn("gorm auto migrate err:", err)
		}
	}
	if len(errorList) > 0 {
		logger.Warn("gorm auto migrate err:" + strings.Join(errorList, ";"))
		return true, errors.New(strings.Join(errorList, ";"))
	}
	return true, nil
}
func UpdateTables(appKey string) error {
	DbAutoMigrate(global.GetFacade().GetEnv().Stage, appKey)
	return nil
}
