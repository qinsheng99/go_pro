package config

import (
	"os"

	"sigs.k8s.io/yaml"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/elastic"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/common/logger"
	"github.com/qinsheng99/go-domain-web/infrastructure/etcd"
	"github.com/qinsheng99/go-domain-web/infrastructure/kafka"
	"github.com/qinsheng99/go-domain-web/infrastructure/kubernetes"
	"github.com/qinsheng99/go-domain-web/infrastructure/mongodb"
	"github.com/qinsheng99/go-domain-web/infrastructure/redis"
	"github.com/qinsheng99/go-domain-web/task"
	"github.com/qinsheng99/go-domain-web/utils/validate"
)

// Config 整个项目的配置
type Config struct {
	Port       int                `json:"port"          required:"true"`
	Es         *elastic.Config    `json:"es"            required:"true"`
	Etcd       *etcd.Config       `json:"etcd"          required:"true"`
	Task       *task.Config       `json:"task"          required:"true"`
	Mysql      *mysql.Config      `json:"mysql"         required:"true"`
	Redis      *redis.Config      `json:"redis"         required:"true"`
	Mongo      *mongodb.Config    `json:"mongo"         required:"true"`
	Kafka      *kafka.Config      `json:"kafka"         required:"true"`
	Logger     *logger.Config     `json:"log"           required:"true"`
	Postgres   *postgres.Config   `json:"postgres"      required:"true"`
	Kubernetes *kubernetes.Config `json:"kubernetes"    required:"true"`
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
