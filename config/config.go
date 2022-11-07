package config

import (
	"flag"
	"io/ioutil"
	"os"
	"sigs.k8s.io/yaml"
)

// Config 整个项目的配置
type Config struct {
	Mode              string `json:"mode"`
	Port              int    `json:"port"`
	*LogConfig        `json:"log"`
	*MysqlConfig      `json:"mysql"`
	*EsConfig         `json:"es"`
	*RedisConfig      `json:"redis"`
	*MongoConfig      `json:"mongo"`
	*PostgresqlConfig `json:"postgresql"`
	*EtcdConfig       `json:"etcd"`
	*KubernetesConfig `json:"kubernetes"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type MysqlConfig struct {
	DbHost    string `json:"db_host"`
	DbPort    int64  `json:"db_port"`
	DbUser    string `json:"db_user"`
	DbPwd     string `json:"db_pwd"`
	DbName    string `json:"db_name"`
	DbMaxConn int    `json:"db_max_conn"`
	DbMaxidle int    `json:"db_maxidle"`
}

type PostgresqlConfig struct {
	DbHost    string `json:"db_host"`
	DbPort    int64  `json:"db_port"`
	DbUser    string `json:"db_user"`
	DbPwd     string `json:"db_pwd"`
	DbName    string `json:"db_name"`
	DbMaxConn int    `json:"db_max_conn"`
	DbMaxidle int    `json:"db_maxidle"`
}

type EsConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type MongoConfig struct {
	Host       string `json:"host"`
	Port       int64  `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type EtcdConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type KubernetesConfig struct {
	NameSpace string `json:"namespace"`
	Pod       struct {
		Image  string `json:"image"`
		Secret string `json:"secret"`
	} `json:"pod"`
	ConfigMap struct {
		ConfigMapName string `json:"config_map_name"`
		ConfigName    string `json:"config_name"`
		MounthPath    string `json:"mounth_path"`
	} `json:"config_map"`
	Crd struct {
		Group   string `json:"group"`
		Kind    string `json:"kind"`
		Version string `json:"version"`
	} `json:"crd"`
}

var path = flag.String("config-file", "", "")
var Conf = new(Config)

func Init() error {
	var file = "config/config.yaml"
	if *path != "" {
		file = *path
	}
	bys, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(os.ExpandEnv(string(bys))), Conf)
	if err != nil {
		return err
	}
	return nil
}
