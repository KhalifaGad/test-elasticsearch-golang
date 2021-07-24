package elasticsearch

import (
	"errors"
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
)

/*
	Hey!, combining interface with struct to make client field of elasticsearch struct private 🤙
*/

var client Elasticsearch

type Elasticsearch interface {
	Info() (interface{}, error)
	Store()
}

type elasticsearch struct {
	client *es.Client
}

func New(addresses []string, shouldStore bool) Elasticsearch {

	cfg := es.Config{
		Addresses: addresses,
	}

	log.Print("creating es client...")
	esClient, err := es.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating elasticsearch client: %v", err)
		return nil
	}

	elasticsearchObj := &elasticsearch{client: esClient}

	if shouldStore {
		client = elasticsearchObj
	}
	return elasticsearchObj
}

func (elasticsearchObj *elasticsearch) Info() (interface{}, error) {

	res, err := elasticsearchObj.client.Info()
	if err != nil {
		log.Fatalf("Error getting client info %v", err)
		return nil, err
	}

	return res, nil
}

func (elasticsearchObj *elasticsearch) Store() {
	client = elasticsearchObj
}

func GetClient() (Elasticsearch, error) {
	if client == nil {
		return nil, errors.New("there are no stored client yet")
	}

	return client, nil
}
