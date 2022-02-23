package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	dpelasticsearch "github.com/ONSdigital/dp-elasticsearch/v2/elasticsearch"
	"github.com/ONSdigital/dp-search-api/config"
	"github.com/ONSdigital/dp-search-api/elasticsearch"

	"log"
	"sync"
)

var (
	maxConcurrentExtractions = 20
	maxConcurrentIndexings   = 20

	ctx = context.Background()
)

type zebedeeClient interface {
	GetPublishedIndex(ctx context.Context) (zebedee.PublishedIndex, error)
	GetPublishedData(ctx context.Context, uriString string) ([]byte, error)
}

type elasticSearchClient interface {
	CreateIndex(ctx context.Context, indexName string, indexSettings []byte) (int, error)
	AddDocument(ctx context.Context, indexName, documentType, documentID string, document []byte) (int, error)
}

type Document struct {
	URI  string
	Body []byte
}

//type TransformedDoc struct {
//	ID  string
//	Body []byte
//}

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("error retrieving config: %s", err)
	}

	zebClient := zebedee.New(cfg.ZebedeeURL)
	esClient := dpelasticsearch.NewClient(cfg.ElasticSearchAPIURL, cfg.SignElasticsearchRequests, 5)

	// get URI from zebedee
	urisChan := uriProducer(ctx, zebClient)

	// get publishedContents corresponding to URI
	extractedChan, extractionFailuresChan := docExtractor(ctx, zebClient, urisChan, maxConcurrentExtractions)

	// Transformer
	transformedChan := docTransformer(extractedChan)

	// reindex ONS Index
	indexedChan := docIndexer(ctx, esClient, transformedChan, maxConcurrentIndexings)

	// Summarize the re-indexing the elastic search
	summarize(indexedChan, extractionFailuresChan)
}

func uriProducer(ctx context.Context, z zebedeeClient) chan string {
	uriChan := make(chan string)
	go func() {
		defer close(uriChan)
		for _, uri := range getPublishedURIs(ctx, z) {
			for i := 0; i < 1; i++ {
				uriChan <- uri
			}
			//fmt.Printf("Sending %s to channel\n", uri)
		}
		fmt.Println("Finished listing uris")
	}()
	return uriChan
}

func getPublishedURIs(ctx context.Context, z zebedeeClient) []string {
	index, err := z.GetPublishedIndex(ctx)
	if err != nil {
		//TODO error handling
		log.Fatalf("Fatal error getting index from zebedee: %s", err)
	}
	fmt.Printf("Fetched %d uris from zebedee\n", len(index.URIs))
	return index.URIs
}

func docExtractor(ctx context.Context, z zebedeeClient, uriChan chan string, maxExtractions int) (chan Document, chan string) {
	extractedChan := make(chan Document)
	extractionFailuresChan := make(chan string)
	go func() {
		defer close(extractedChan)
		defer close(extractionFailuresChan)

		var wg sync.WaitGroup

		for w := 0; w < maxExtractions; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				extractDoc(ctx, z, uriChan, extractedChan, extractionFailuresChan)
			}()
		}
		wg.Wait()
		fmt.Println("Finished extracting docs")
	}()
	return extractedChan, extractionFailuresChan
}

func extractDoc(ctx context.Context, z zebedeeClient, uriChan chan string, extractedChan chan Document, extractionFailuresChan chan string) {
	for uri := range uriChan {
		body, err := z.GetPublishedData(ctx, uri)
		//time.Sleep(time.Second)
		if err != nil {
			extractionFailuresChan <- uri
		}

		extractedDoc := Document{
			URI:  uri,
			Body: body,
		}
		extractedChan <- extractedDoc
	}
}

func docTransformer(extractedChan chan Document) chan Document {
	transformedChan := make(chan Document)
	go func() {
		defer close(transformedChan)
		var wg sync.WaitGroup

		for extractedDoc := range extractedChan {
			wg.Add(1)
			go func(doc Document) {
				defer wg.Done()
				transformDoc(doc, transformedChan)
			}(extractedDoc)
		}
		wg.Wait()
		fmt.Println("Finished transforming docs")
	}()
	return transformedChan
}

func transformDoc(extractedDoc Document, transformedChan chan Document) {
	body := []byte(fmt.Sprintf("{\"decription\": {\"title\":\"Transform some data from '%s'\"}}", extractedDoc.URI)) //TODO implement actual transform
	//time.Sleep(time.Second)
	transformedDoc := Document{
		URI:  extractedDoc.URI,
		Body: body,
	}
	transformedChan <- transformedDoc
}

func docIndexer(ctx context.Context, es elasticSearchClient, transformedChan chan Document, maxIndexings int) chan bool {
	indexedChan := make(chan bool)
	go func() {
		defer close(indexedChan)

		indexName := createIndexName("ons")
		fmt.Printf("Index created: %s\n", indexName)
		status, err := es.CreateIndex(ctx, indexName, elasticsearch.GetSearchIndexSettings())
		if err != nil {
			log.Fatal(ctx, "error creating index", err)
		}
		if status != http.StatusOK {
			log.Fatal(ctx, "error creating index http status - ", status)
		}

		var wg sync.WaitGroup

		for w := 0; w < maxIndexings; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				indexDoc(ctx, es, transformedChan, indexedChan, indexName)
			}()
		}
		wg.Wait()
		fmt.Println("Finished indexing docs")
	}()
	return indexedChan
}

func createIndexName(s string) string {
	now := time.Now()
	return fmt.Sprintf("%s%d", s, now.UnixMicro())
}

func indexDoc(ctx context.Context, es elasticSearchClient, transformedChan chan Document, indexedChan chan bool, indexName string) {
	for transformedDoc := range transformedChan {

		id := url.PathEscape(transformedDoc.URI) //TODO this isn't right, the client should url-escape the id
		indexed := true
		status, err := es.AddDocument(ctx, indexName, "_create", id, transformedDoc.Body)
		if err != nil || status != http.StatusCreated {
			indexed = false
		}

		indexedChan <- indexed
	}
}

func summarize(indexedChan chan bool, extractionFailuresChan chan string) {
	totalIndexed, totalFailed := 0, 0
	for range extractionFailuresChan {
		totalFailed++
	}
	for indexed := range indexedChan {
		if indexed {
			totalIndexed++
		} else {
			totalFailed++
		}
	}

	fmt.Printf("Indexed: %d, Failed: %d\n", totalIndexed, totalFailed)
}
