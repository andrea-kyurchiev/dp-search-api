package api

//go:generate moq -out mocks.go -pkg api . ElasticSearcher QueryBuilder ResponseTransformer

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/health"
	"github.com/ONSdigital/dp-search-api/config"
	"github.com/ONSdigital/dp-search-api/elasticsearch"
	"github.com/ONSdigital/dp-search-api/service"
	elastic "github.com/ONSdigital/dp-elasticsearch/v2/elasticsearch"

	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
)

var httpServer *server.Server

//SearchAPI provides an API around elasticseach
type SearchAPI struct {
	Router        *mux.Router
	QueryBuilder  QueryBuilder
	ElasticSearch ElasticSearcher
	Transformer   ResponseTransformer
	ServiceList   *service.ExternalServiceList
}

// ElasticSearcher provides client methods for the elasticsearch package
type ElasticSearcher interface {
	Search(ctx context.Context, index string, docType string, request []byte) ([]byte, error)
	MultiSearch(ctx context.Context, index string, docType string, request []byte) ([]byte, error)
	GetStatus(ctx context.Context) ([]byte, error)
}

// QueryBuilder provides methods for the search package
type QueryBuilder interface {
	BuildSearchQuery(ctx context.Context, q, contentTypes, sort string, limit, offset int) ([]byte, error)
}

// ResponseTransformer provides methods for the transform package
type ResponseTransformer interface {
	TransformSearchResponse(ctx context.Context, responseData []byte, query string) ([]byte, error)
}

// CreateAndInitialise initiates a new Search API
func CreateAndInitialise(bindAddr string, queryBuilder QueryBuilder, elasticSearchClient ElasticSearcher, transformer ResponseTransformer, serviceList *service.ExternalServiceList, buildTime, gitCommit, version string, errorChan chan error) error {

	if elasticSearchClient == nil {
		return errors.New("CreateAndInitialise called without a valid elasticsearch client")
	}

	if queryBuilder == nil {
		return errors.New("CreateAndInitialise called without a valid query builder")
	}
	router := mux.NewRouter()

	errData := SetupData()
	if errData != nil {
		return errors.Wrap(errData, "Failed to setup data templates")
	}

	errTimeseries := SetupTimeseries()
	if errTimeseries != nil {
		return errors.Wrap(errTimeseries, "Failed to setup timeseries templates")
	}

	// Get HealthCheck
	cfg, err := config.Get()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve service configuration")
	}

	ctx := context.Background()
	hc, err := serviceList.GetHealthCheck(cfg, buildTime, gitCommit, version)
	if err != nil {
		log.Event(ctx, "could not instantiate healthcheck", log.FATAL, log.Error(err))
		return err
	}
	if err := registerCheckers(ctx, elasticSearchClient, hc); err != nil {
		return errors.Wrap(err, "unable to register checkers")
	}

	router.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)
	hc.Start(ctx)

	api := NewSearchAPI(router, elasticSearchClient, queryBuilder, transformer)

	httpServer = server.New(bindAddr, api.Router)

	// Disable this here to allow service to manage graceful shutdown of the entire app.
	httpServer.HandleOSSignals = false

	go func() {
		log.Event(nil, "search api starting", log.INFO)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Event(nil, "search api http server returned error", log.Error(err), log.ERROR)
			errorChan <- err
		}
	}()

	return nil
}

// NewSearchAPI returns a new Search API struct after registering the routes
func NewSearchAPI(router *mux.Router, elasticSearch ElasticSearcher, queryBuilder QueryBuilder, transformer ResponseTransformer) *SearchAPI {

	api := &SearchAPI{
		Router:        router,
		QueryBuilder:  queryBuilder,
		ElasticSearch: elasticSearch,
		Transformer:   transformer,
	}

	router.HandleFunc("/search", SearchHandlerFunc(queryBuilder, api.ElasticSearch, api.Transformer)).Methods("GET")
	router.HandleFunc("/timeseries/{cdid}", TimeseriesLookupHandlerFunc(api.ElasticSearch)).Methods("GET")
	router.HandleFunc("/data", DataLookupHandlerFunc(api.ElasticSearch)).Methods("GET")
	return api
}

// Close represents the graceful shutting down of the http server
func Close(ctx context.Context) error {
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Event(ctx, "graceful shutdown of http server complete", log.INFO)
	return nil
}

func registerCheckers(ctx context.Context, elasticSearchClient elasticsearch.Client, hc service.HealthChecker) (err error) {

	hasErrors := false

	if err = hc.AddCheck("Elasticsearch", elasticSearchClient.Checker); err != nil {
		log.Event(ctx, "error creating elasticsearch health check", log.ERROR, log.Error(err))
		hasErrors = true
	}

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}
