package elasticsearch

import (
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
)

/*
	Hey!, combining interface with struct to make client field of elasticsearch struct private 🤙
*/

type Elasticsearch interface {
	Info() (interface{}, error)
}

type elasticsearch struct {
	client *es.Client
}

func New(addresses []string) Elasticsearch {

	cfg := es.Config{
		Addresses: addresses,
	}

	log.Print("creating es client...")
	esClient, err := es.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating elasticsearch client: %v", err)
		return nil
	}

	return &elasticsearch{client: esClient}
}

func (elasticsearchObj *elasticsearch) Info() (interface{}, error) {

	res, err := elasticsearchObj.client.Info()
	if err != nil {
		log.Fatalf("Error getting client info %v", err)
		return nil, err
	}

	return res, nil
}
