package elasticsearch

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/qinsheng99/go-domain-web/config"
)

type (
	Config struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ConfLoader func(v interface{}) error
)

var es *elastic.Client

func Init(cfg *config.EsConfig) (err error) {
	var c = Config{
		URL:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		Username: "",
	}
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(c.URL),
		elastic.SetSniff(false),
	}

	if c.Password != "" {
		options = append(options, elastic.SetBasicAuth(c.Username, c.Password))
	}

	es, err = elastic.NewClient(options...)
	if err != nil {
		return
	}

	return nil
}

func GetElasticsearch() *elastic.Client {
	return es
}
