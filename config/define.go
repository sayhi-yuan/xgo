package config

import "time"

// 配置
// 映射项目的配置文件config.yaml

type Config struct {
	Serve       Serve            `yaml:"serve" desc:"服务信息"`
	MySQL       MySQL            `yaml:"mysql"`
	Redis       map[string]Redis `yaml:"redis"`
	Logger      Logger           `yaml:"logger"`
	FileService FileService      `yaml:"fileservice"`
}

// Serve 服务信息
type Serve struct {
	Name string `yaml:"name" desc:"服务名称"`
	Port int    `yaml:"port" desc:"服务访问端口"`
}

type MySQL struct {
	Default MySQLConfig `yaml:"default"`
}

type MySQLConfig struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	UserName    string `yaml:"username"`
	Password    string `yaml:"password"`
	DataBase    string `yaml:"database"`
	Config      string `yaml:"config"`
	MaxIdleConn int    `yaml:"maxidleconn"`
	MaxOpenConn int    `yaml:"maxopenconn"`
	Debug       bool   `yaml:"debug"`
}

type Redis struct {
	Host         []string      `yaml:"host"`
	DB           int           `yaml:"db" desc:"默认启用的数据库索引"`
	Password     string        `yaml:"password"`
	Mode         int           `yaml:"mode" desc:"redis模式；1单机;2集群"`
	ReadTimeout  time.Duration `yaml:"readtimeout" desc:"读超时时间"`
	WriteTimeout time.Duration `yaml:"writetimeout" desc:"超时时间"`
}

type Logger struct {
	Path         string `yaml:"path"`
	MaxCount     uint   `yaml:"maxcount"`
	CallerEnable bool   `json:"callerenable"`
}

type FileService struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
