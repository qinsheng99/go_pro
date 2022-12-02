package etcd

import (
	"fmt"
	"time"

	"github.com/qinsheng99/go-domain-web/config"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client *clientv3.Client

func Init(cfg *config.EtcdConfig) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		return err
	}
	return nil
}

func GetEtcd() *clientv3.Client {
	return client
}
