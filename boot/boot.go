package boot

import (
	"fmt"
	"github.com/senyu-up/toolbox/combz/facade"
	"github.com/senyu-up/toolbox/tool/config/loader"
	"github.com/senyu-up/toolbox/tool/file"
	"template/config"
	"template/global"
)

// 传入指定路径，初始化配置
func bootConfig(path string) (conf *config.Config, err error) {
	confLoader, err := loader.InitConf(&loader.File{},
		loader.ConfOptWithPath(file.ScanConfigPath(path)), // ./config.yaml
		loader.ConfOptWithType("yaml"))
	if err != nil {
		return conf, err
	}

	conf = &config.Config{}
	err = confLoader.Unmarshal(&conf)
	global.SetConfig(conf)
	return
}

var tbf = &facade.ToolFacade{}

func Boot(confPath string) (err error) {
	conf, err := bootConfig(confPath)
	if err != nil {
		return err
	}

	if conf == nil {
		return fmt.Errorf("config is nil")
	}

	tbf, err = facade.InitApp(
		facade.ConfigOptionWithApp(conf.App),
		facade.ConfigOptionWithGin(conf.Gin),
		facade.ConfigOptionWithLog(conf.Log),
		facade.ConfigOptionWithRedis(conf.Redis),
		facade.ConfigOptionWithMysql(conf.Mysql),
		facade.ConfigOptionWithTrace(conf.Trace),
		facade.ConfigOptionWithHealth(conf.Health),
	)
	if err != nil {
		fmt.Printf("init app facade failed, err: %v", err)
		return
	}
	global.SetFacade(tbf) // 把门面设置到全局对象
	return err
}
