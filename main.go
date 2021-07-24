package main

import (
	"log"
	"os"

	elasticsearch "github.com/KhalifaGad/test-elasticsearch-golang/elasticsearch"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading env variables %v", err)
		log.Fatal("process exiting...")
		return
	}

	elasticsearchURL := os.Getenv("ELASTICSEARCH_URL")

	addresses := []string{elasticsearchURL}

	if !elasticsearch.Init(addresses) {
		return
	}

}
