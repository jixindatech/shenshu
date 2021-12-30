package config

import (
	"github.com/spf13/viper"
	"time"
)

type Log struct {
	Filename  string `mapstructure:"filename"`
	Level     string `mapstructure:"level"`
	MaxSize   int    `mapstructure:"maxsize"`
	MaxBackup int    `mapstructure:"maxbackup"`
	MaxAge    int    `mapstructure:"maxage"`
	Compress  bool   `mapstructure:"compress"`
}

type App struct {
	PageSize      int    `mapstructure:"page-size"`
	AdminPassword string `mapstructure:"adminpassword"`
}

type Server struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}

type DataBase struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table-prefix"`
}

type Rbac struct {
	Model  string `mapstructure:"model"`
	Policy string `mapstructure:"policy"`
	Auth   string `mapstructure:"auth"`
}

type Redis struct {
	Type      string `mapstructer:"type"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Password  string `mapstructure:"password"`
	Db        int    `mapstructure:"db"`
	KeyPrefix string `mapstructure:"key_prefix"`
}

type Memory struct {
	PurgeTime time.Duration `mapstructer:"purge_time"`
}

type Elasticsearch struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Timeout   string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	CCIndex   string `mapstructure:"cc_index" json:"cc_index" yaml:"cc_index"`
	RuleIndex string `mapstructure:"rule_index" json:"rule_index" yaml:"rule_index"`
}

type Config struct {
	RunMode       string         `mapstructure:"run_mode"`
	Log           *Log           `mapstructure:"log"`
	Server        *Server        `mapstructure:"server"`
	Database      *DataBase      `mapstructer:"database"`
	Rbac          *Rbac          `mapstructer:"rbac"`
	App           *App           `mapstructer:"app"`
	Cache         string         `mapstructer:"cache"`
	Redis         *Redis         `mapstructer:"redis"`
	Elasticsearch *Elasticsearch `mapstructer:"elasticsearch"`
	Memory        *Memory        `mapstructer:"memory"`
}

var config *Config

func ParseConfigFile(file string) (*Config, error) {
	config = new(Config)

	v := viper.New()
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := v.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
