package elasticsearch

import (
	"errors"
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
)

var client *es.Client

func InitElasticsearch(addresses []string) bool {
	cfg := es.Config{
		Addresses: addresses,
	}
	log.Print("creating es client...")
	esClient, err := es.NewClient(cfg)

	client = esClient

	if err != nil {
		log.Fatalf("Error creating elasticsearch client: %v", err)
		return false
	}

	return true
}

func Info() (interface{}, error) {

	ok, err := checkClientInitialization()

	if !ok {
		return nil, err
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting client info %v", err)
		return nil, err
	}

	return res, nil
}

func checkClientInitialization() (bool, error) {
	if client == nil {
		return false, errors.New("elasticsearch client not initialized yet")
	}

	return true, nil
}
