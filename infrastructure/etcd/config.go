package etcd

type Config struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
