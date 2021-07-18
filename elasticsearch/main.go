package elasticsearch

import (
	"errors"
	"log"

	goelasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func GetClient(addresses []string) *goelasticsearch.Client {
	cfg := goelasticsearch.Config{
		Addresses: addresses,
	}
	log.Print("creating es client...")
	client, err := goelasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Err creating elasticsearch client: %v", err)
		return nil
	}

	return client
}

type Elasticsearch struct {
	client goelasticsearch.Client
}

func (elasticsearch *Elasticsearch) Info() (interface{}, error) {

	if elasticsearch == nil {
		log.Fatalf("Elasticsearch client not initialized yet!.")
		return nil, errors.New("Elasticsearch client not initialized yet!")
	}

	res, err := elasticsearch.client.Info()
	if err != nil {
		log.Fatalf("Error getting client info %v", err)
		return nil, err
	}

	defer res.Body.Close()

	return res, nil
}
