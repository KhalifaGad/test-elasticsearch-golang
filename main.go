package main

import (
	"log"
	"os"

	elasticsearch "github.com/KhalifaGad/test-elasticsearch-golang/elasticsearch"
	server "github.com/KhalifaGad/test-elasticsearch-golang/http"
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

	es := elasticsearch.New(addresses, true)

	info, err := es.Info()

	if err != nil {
		log.Fatal("elasticsearch not initialized, exiting the process...")
		return
	}

	log.Println(info)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Start(port)
}
