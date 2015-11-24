package settings

import (
	"encoding/json"
	"fmt"

	"github.com/iogo-framework/logs"
)

type Elasticsearch struct {
	Host string
	Port int
}

func (config Config) Elasticsearch() (Elasticsearch, error) {
	var elastic Elasticsearch

	err := json.Unmarshal(config.Components["elasticsearch"], &elastic)
	if err != nil {
		logs.Warning("%s: %s", err.Error(), "missing or wrong 'elasticsearch' configuration, ignoring")
	}

	if elastic.Host == "" {
		elastic.Host = "elasticsearch"
		logs.Warning("missing elasticsearch 'host' configuration, assuming default value: 'elasticsearch'")
	}

	if elastic.Port == 0 {
		elastic.Port = 9200
		logs.Warning("missing elasticsearch 'port' configuration, assuming default value: 9200")
	}

	return elastic, nil
}

func (r Elasticsearch) String() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
