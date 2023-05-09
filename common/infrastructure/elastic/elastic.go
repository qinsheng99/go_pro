package elastic

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

var es *elastic.Client

func Init(cfg *Config) (err error) {
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
		elastic.SetSniff(false),
	}

	//if c.Password != "" {
	//	options = append(options, elastic.SetBasicAuth(c.Username, c.Password))
	//}

	es, err = elastic.NewClient(options...)
	if err != nil {
		return
	}

	return nil
}
