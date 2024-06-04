package config

import "github.com/senyu-up/toolbox/tool/config"

type Config struct {
	App    *config.App         `yaml:"app" json:"app"`
	Gin    *config.GinConfig   `yaml:"gin" json:"gin"`
	Log    *config.LogConfig   `yaml:"Log" json:"Log"`
	Trace  *config.TraceConfig `yaml:"Trace" json:"Trace"`
	Redis  *config.RedisConfig `yaml:"redis" json:"redis"`
	Mysql  *config.MysqlConfig `yaml:"mysql" json:"mysql"`
	Health *config.HealthCheck `yaml:"health" json:"health"`
	Jwt    *JwtConfig          `yaml:"jwt" json:"jwt"`
}

type JwtConfig struct {
	TokenSecret     string `yaml:"tokenSecret" json:"tokenSecret"`
	TokenExpiration int64  `yaml:"tokenExpiration" json:"tokenExpiration"`
}
