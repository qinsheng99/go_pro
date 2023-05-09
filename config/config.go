package config

import (
	"os"

	"sigs.k8s.io/yaml"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/elastic"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/infrastructure/etcd"
	"github.com/qinsheng99/go-domain-web/infrastructure/kafka"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mongodb"
	"github.com/qinsheng99/go-domain-web/infrastructure/redis"
	"github.com/qinsheng99/go-domain-web/utils/validate"
)

// Config 整个项目的配置
type Config struct {
	Mode       string             `json:"mode"`
	Port       int                `json:"port"`
	Logger     *logger.Config     `json:"log"`
	Mysql      *mysql.Config      `json:"mysql"`
	Es         *elastic.Config    `json:"es"`
	Redis      *redis.Config      `json:"redis"`
	Mongo      *mongodb.Config    `json:"mongo"`
	Postgresql *postgresql.Config `json:"postgresql"`
	Etcd       *etcd.Config       `json:"etcd"`
	Kafka      *kafka.Config      `json:"kafka"`
	Kubernetes *kubernetes.Config `json:"kubernetes"`
}

func Init(path string) (*Config, error) {
	var cfg Config

	bys, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(os.ExpandEnv(string(bys))), &cfg)
	if err != nil {
		return nil, err
	}

	if err = validate.Vali(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
