// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package api

import (
	"context"
	"github.com/ONSdigital/dp-authorisation/auth"
	"github.com/ONSdigital/dp-elasticsearch/v3/client"
	"github.com/ONSdigital/dp-search-api/query"
	"net/http"
	"sync"
)

// Ensure, that ElasticSearcherMock does implement ElasticSearcher.
// If this is not the case, regenerate this file with moq.
var _ ElasticSearcher = &ElasticSearcherMock{}

// ElasticSearcherMock is a mock implementation of ElasticSearcher.
//
//	func TestSomethingThatUsesElasticSearcher(t *testing.T) {
//
//		// make and configure a mocked ElasticSearcher
//		mockedElasticSearcher := &ElasticSearcherMock{
//			MultiSearchFunc: func(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
//				panic("mock out the MultiSearch method")
//			},
//			SearchFunc: func(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
//				panic("mock out the Search method")
//			},
//		}
//
//		// use mockedElasticSearcher in code that requires ElasticSearcher
//		// and then make assertions.
//
//	}
type ElasticSearcherMock struct {
	// MultiSearchFunc mocks the MultiSearch method.
	MultiSearchFunc func(ctx context.Context, index string, docType string, request []byte) ([]byte, error)

	// SearchFunc mocks the Search method.
	SearchFunc func(ctx context.Context, index string, docType string, request []byte) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// MultiSearch holds details about calls to the MultiSearch method.
		MultiSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Index is the index argument value.
			Index string
			// DocType is the docType argument value.
			DocType string
			// Request is the request argument value.
			Request []byte
		}
		// Search holds details about calls to the Search method.
		Search []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Index is the index argument value.
			Index string
			// DocType is the docType argument value.
			DocType string
			// Request is the request argument value.
			Request []byte
		}
	}
	lockMultiSearch sync.RWMutex
	lockSearch      sync.RWMutex
}

// MultiSearch calls MultiSearchFunc.
func (mock *ElasticSearcherMock) MultiSearch(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
	if mock.MultiSearchFunc == nil {
		panic("ElasticSearcherMock.MultiSearchFunc: method is nil but ElasticSearcher.MultiSearch was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}{
		Ctx:     ctx,
		Index:   index,
		DocType: docType,
		Request: request,
	}
	mock.lockMultiSearch.Lock()
	mock.calls.MultiSearch = append(mock.calls.MultiSearch, callInfo)
	mock.lockMultiSearch.Unlock()
	return mock.MultiSearchFunc(ctx, index, docType, request)
}

// MultiSearchCalls gets all the calls that were made to MultiSearch.
// Check the length with:
//
//	len(mockedElasticSearcher.MultiSearchCalls())
func (mock *ElasticSearcherMock) MultiSearchCalls() []struct {
	Ctx     context.Context
	Index   string
	DocType string
	Request []byte
} {
	var calls []struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}
	mock.lockMultiSearch.RLock()
	calls = mock.calls.MultiSearch
	mock.lockMultiSearch.RUnlock()
	return calls
}

// Search calls SearchFunc.
func (mock *ElasticSearcherMock) Search(ctx context.Context, index string, docType string, request []byte) ([]byte, error) {
	if mock.SearchFunc == nil {
		panic("ElasticSearcherMock.SearchFunc: method is nil but ElasticSearcher.Search was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}{
		Ctx:     ctx,
		Index:   index,
		DocType: docType,
		Request: request,
	}
	mock.lockSearch.Lock()
	mock.calls.Search = append(mock.calls.Search, callInfo)
	mock.lockSearch.Unlock()
	return mock.SearchFunc(ctx, index, docType, request)
}

// SearchCalls gets all the calls that were made to Search.
// Check the length with:
//
//	len(mockedElasticSearcher.SearchCalls())
func (mock *ElasticSearcherMock) SearchCalls() []struct {
	Ctx     context.Context
	Index   string
	DocType string
	Request []byte
} {
	var calls []struct {
		Ctx     context.Context
		Index   string
		DocType string
		Request []byte
	}
	mock.lockSearch.RLock()
	calls = mock.calls.Search
	mock.lockSearch.RUnlock()
	return calls
}

// Ensure, that DpElasticSearcherMock does implement DpElasticSearcher.
// If this is not the case, regenerate this file with moq.
var _ DpElasticSearcher = &DpElasticSearcherMock{}

// DpElasticSearcherMock is a mock implementation of DpElasticSearcher.
//
//	func TestSomethingThatUsesDpElasticSearcher(t *testing.T) {
//
//		// make and configure a mocked DpElasticSearcher
//		mockedDpElasticSearcher := &DpElasticSearcherMock{
//			CountFunc: func(ctx context.Context, count client.Count) ([]byte, error) {
//				panic("mock out the Count method")
//			},
//			CreateIndexFunc: func(ctx context.Context, indexName string, indexSettings []byte) error {
//				panic("mock out the CreateIndex method")
//			},
//			MultiSearchFunc: func(ctx context.Context, searches []client.Search, params *client.QueryParams) ([]byte, error) {
//				panic("mock out the MultiSearch method")
//			},
//		}
//
//		// use mockedDpElasticSearcher in code that requires DpElasticSearcher
//		// and then make assertions.
//
//	}
type DpElasticSearcherMock struct {
	// CountFunc mocks the Count method.
	CountFunc func(ctx context.Context, count client.Count) ([]byte, error)

	// CreateIndexFunc mocks the CreateIndex method.
	CreateIndexFunc func(ctx context.Context, indexName string, indexSettings []byte) error

	// MultiSearchFunc mocks the MultiSearch method.
	MultiSearchFunc func(ctx context.Context, searches []client.Search, params *client.QueryParams) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Count holds details about calls to the Count method.
		Count []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Count is the count argument value.
			Count client.Count
		}
		// CreateIndex holds details about calls to the CreateIndex method.
		CreateIndex []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// IndexName is the indexName argument value.
			IndexName string
			// IndexSettings is the indexSettings argument value.
			IndexSettings []byte
		}
		// MultiSearch holds details about calls to the MultiSearch method.
		MultiSearch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Searches is the searches argument value.
			Searches []client.Search
			// Params is the params argument value.
			Params *client.QueryParams
		}
	}
	lockCount       sync.RWMutex
	lockCreateIndex sync.RWMutex
	lockMultiSearch sync.RWMutex
}

// Count calls CountFunc.
func (mock *DpElasticSearcherMock) Count(ctx context.Context, count client.Count) ([]byte, error) {
	if mock.CountFunc == nil {
		panic("DpElasticSearcherMock.CountFunc: method is nil but DpElasticSearcher.Count was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Count client.Count
	}{
		Ctx:   ctx,
		Count: count,
	}
	mock.lockCount.Lock()
	mock.calls.Count = append(mock.calls.Count, callInfo)
	mock.lockCount.Unlock()
	return mock.CountFunc(ctx, count)
}

// CountCalls gets all the calls that were made to Count.
// Check the length with:
//
//	len(mockedDpElasticSearcher.CountCalls())
func (mock *DpElasticSearcherMock) CountCalls() []struct {
	Ctx   context.Context
	Count client.Count
} {
	var calls []struct {
		Ctx   context.Context
		Count client.Count
	}
	mock.lockCount.RLock()
	calls = mock.calls.Count
	mock.lockCount.RUnlock()
	return calls
}

// CreateIndex calls CreateIndexFunc.
func (mock *DpElasticSearcherMock) CreateIndex(ctx context.Context, indexName string, indexSettings []byte) error {
	if mock.CreateIndexFunc == nil {
		panic("DpElasticSearcherMock.CreateIndexFunc: method is nil but DpElasticSearcher.CreateIndex was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		IndexName     string
		IndexSettings []byte
	}{
		Ctx:           ctx,
		IndexName:     indexName,
		IndexSettings: indexSettings,
	}
	mock.lockCreateIndex.Lock()
	mock.calls.CreateIndex = append(mock.calls.CreateIndex, callInfo)
	mock.lockCreateIndex.Unlock()
	return mock.CreateIndexFunc(ctx, indexName, indexSettings)
}

// CreateIndexCalls gets all the calls that were made to CreateIndex.
// Check the length with:
//
//	len(mockedDpElasticSearcher.CreateIndexCalls())
func (mock *DpElasticSearcherMock) CreateIndexCalls() []struct {
	Ctx           context.Context
	IndexName     string
	IndexSettings []byte
} {
	var calls []struct {
		Ctx           context.Context
		IndexName     string
		IndexSettings []byte
	}
	mock.lockCreateIndex.RLock()
	calls = mock.calls.CreateIndex
	mock.lockCreateIndex.RUnlock()
	return calls
}

// MultiSearch calls MultiSearchFunc.
func (mock *DpElasticSearcherMock) MultiSearch(ctx context.Context, searches []client.Search, params *client.QueryParams) ([]byte, error) {
	if mock.MultiSearchFunc == nil {
		panic("DpElasticSearcherMock.MultiSearchFunc: method is nil but DpElasticSearcher.MultiSearch was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Searches []client.Search
		Params   *client.QueryParams
	}{
		Ctx:      ctx,
		Searches: searches,
		Params:   params,
	}
	mock.lockMultiSearch.Lock()
	mock.calls.MultiSearch = append(mock.calls.MultiSearch, callInfo)
	mock.lockMultiSearch.Unlock()
	return mock.MultiSearchFunc(ctx, searches, params)
}

// MultiSearchCalls gets all the calls that were made to MultiSearch.
// Check the length with:
//
//	len(mockedDpElasticSearcher.MultiSearchCalls())
func (mock *DpElasticSearcherMock) MultiSearchCalls() []struct {
	Ctx      context.Context
	Searches []client.Search
	Params   *client.QueryParams
} {
	var calls []struct {
		Ctx      context.Context
		Searches []client.Search
		Params   *client.QueryParams
	}
	mock.lockMultiSearch.RLock()
	calls = mock.calls.MultiSearch
	mock.lockMultiSearch.RUnlock()
	return calls
}

// Ensure, that QueryParamValidatorMock does implement QueryParamValidator.
// If this is not the case, regenerate this file with moq.
var _ QueryParamValidator = &QueryParamValidatorMock{}

// QueryParamValidatorMock is a mock implementation of QueryParamValidator.
//
//	func TestSomethingThatUsesQueryParamValidator(t *testing.T) {
//
//		// make and configure a mocked QueryParamValidator
//		mockedQueryParamValidator := &QueryParamValidatorMock{
//			ValidateFunc: func(ctx context.Context, name string, value string) (interface{}, error) {
//				panic("mock out the Validate method")
//			},
//		}
//
//		// use mockedQueryParamValidator in code that requires QueryParamValidator
//		// and then make assertions.
//
//	}
type QueryParamValidatorMock struct {
	// ValidateFunc mocks the Validate method.
	ValidateFunc func(ctx context.Context, name string, value string) (interface{}, error)

	// calls tracks calls to the methods.
	calls struct {
		// Validate holds details about calls to the Validate method.
		Validate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Value is the value argument value.
			Value string
		}
	}
	lockValidate sync.RWMutex
}

// Validate calls ValidateFunc.
func (mock *QueryParamValidatorMock) Validate(ctx context.Context, name string, value string) (interface{}, error) {
	if mock.ValidateFunc == nil {
		panic("QueryParamValidatorMock.ValidateFunc: method is nil but QueryParamValidator.Validate was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Name  string
		Value string
	}{
		Ctx:   ctx,
		Name:  name,
		Value: value,
	}
	mock.lockValidate.Lock()
	mock.calls.Validate = append(mock.calls.Validate, callInfo)
	mock.lockValidate.Unlock()
	return mock.ValidateFunc(ctx, name, value)
}

// ValidateCalls gets all the calls that were made to Validate.
// Check the length with:
//
//	len(mockedQueryParamValidator.ValidateCalls())
func (mock *QueryParamValidatorMock) ValidateCalls() []struct {
	Ctx   context.Context
	Name  string
	Value string
} {
	var calls []struct {
		Ctx   context.Context
		Name  string
		Value string
	}
	mock.lockValidate.RLock()
	calls = mock.calls.Validate
	mock.lockValidate.RUnlock()
	return calls
}

// Ensure, that QueryBuilderMock does implement QueryBuilder.
// If this is not the case, regenerate this file with moq.
var _ QueryBuilder = &QueryBuilderMock{}

// QueryBuilderMock is a mock implementation of QueryBuilder.
//
//	func TestSomethingThatUsesQueryBuilder(t *testing.T) {
//
//		// make and configure a mocked QueryBuilder
//		mockedQueryBuilder := &QueryBuilderMock{
//			BuildCountQueryFunc: func(ctx context.Context, req *query.CountRequest) ([]byte, error) {
//				panic("mock out the BuildCountQuery method")
//			},
//			BuildSearchQueryFunc: func(ctx context.Context, req *query.SearchRequest, esVersion710 bool) ([]byte, error) {
//				panic("mock out the BuildSearchQuery method")
//			},
//		}
//
//		// use mockedQueryBuilder in code that requires QueryBuilder
//		// and then make assertions.
//
//	}
type QueryBuilderMock struct {
	// BuildCountQueryFunc mocks the BuildCountQuery method.
	BuildCountQueryFunc func(ctx context.Context, req *query.CountRequest) ([]byte, error)

	// BuildSearchQueryFunc mocks the BuildSearchQuery method.
	BuildSearchQueryFunc func(ctx context.Context, req *query.SearchRequest, esVersion710 bool) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// BuildCountQuery holds details about calls to the BuildCountQuery method.
		BuildCountQuery []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req *query.CountRequest
		}
		// BuildSearchQuery holds details about calls to the BuildSearchQuery method.
		BuildSearchQuery []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req *query.SearchRequest
			// EsVersion710 is the esVersion710 argument value.
			EsVersion710 bool
		}
	}
	lockBuildCountQuery  sync.RWMutex
	lockBuildSearchQuery sync.RWMutex
}

// BuildCountQuery calls BuildCountQueryFunc.
func (mock *QueryBuilderMock) BuildCountQuery(ctx context.Context, req *query.CountRequest) ([]byte, error) {
	if mock.BuildCountQueryFunc == nil {
		panic("QueryBuilderMock.BuildCountQueryFunc: method is nil but QueryBuilder.BuildCountQuery was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req *query.CountRequest
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockBuildCountQuery.Lock()
	mock.calls.BuildCountQuery = append(mock.calls.BuildCountQuery, callInfo)
	mock.lockBuildCountQuery.Unlock()
	return mock.BuildCountQueryFunc(ctx, req)
}

// BuildCountQueryCalls gets all the calls that were made to BuildCountQuery.
// Check the length with:
//
//	len(mockedQueryBuilder.BuildCountQueryCalls())
func (mock *QueryBuilderMock) BuildCountQueryCalls() []struct {
	Ctx context.Context
	Req *query.CountRequest
} {
	var calls []struct {
		Ctx context.Context
		Req *query.CountRequest
	}
	mock.lockBuildCountQuery.RLock()
	calls = mock.calls.BuildCountQuery
	mock.lockBuildCountQuery.RUnlock()
	return calls
}

// BuildSearchQuery calls BuildSearchQueryFunc.
func (mock *QueryBuilderMock) BuildSearchQuery(ctx context.Context, req *query.SearchRequest, esVersion710 bool) ([]byte, error) {
	if mock.BuildSearchQueryFunc == nil {
		panic("QueryBuilderMock.BuildSearchQueryFunc: method is nil but QueryBuilder.BuildSearchQuery was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Req          *query.SearchRequest
		EsVersion710 bool
	}{
		Ctx:          ctx,
		Req:          req,
		EsVersion710: esVersion710,
	}
	mock.lockBuildSearchQuery.Lock()
	mock.calls.BuildSearchQuery = append(mock.calls.BuildSearchQuery, callInfo)
	mock.lockBuildSearchQuery.Unlock()
	return mock.BuildSearchQueryFunc(ctx, req, esVersion710)
}

// BuildSearchQueryCalls gets all the calls that were made to BuildSearchQuery.
// Check the length with:
//
//	len(mockedQueryBuilder.BuildSearchQueryCalls())
func (mock *QueryBuilderMock) BuildSearchQueryCalls() []struct {
	Ctx          context.Context
	Req          *query.SearchRequest
	EsVersion710 bool
} {
	var calls []struct {
		Ctx          context.Context
		Req          *query.SearchRequest
		EsVersion710 bool
	}
	mock.lockBuildSearchQuery.RLock()
	calls = mock.calls.BuildSearchQuery
	mock.lockBuildSearchQuery.RUnlock()
	return calls
}

// Ensure, that ReleaseQueryBuilderMock does implement ReleaseQueryBuilder.
// If this is not the case, regenerate this file with moq.
var _ ReleaseQueryBuilder = &ReleaseQueryBuilderMock{}

// ReleaseQueryBuilderMock is a mock implementation of ReleaseQueryBuilder.
//
//	func TestSomethingThatUsesReleaseQueryBuilder(t *testing.T) {
//
//		// make and configure a mocked ReleaseQueryBuilder
//		mockedReleaseQueryBuilder := &ReleaseQueryBuilderMock{
//			BuildSearchQueryFunc: func(ctx context.Context, request interface{}) ([]client.Search, error) {
//				panic("mock out the BuildSearchQuery method")
//			},
//		}
//
//		// use mockedReleaseQueryBuilder in code that requires ReleaseQueryBuilder
//		// and then make assertions.
//
//	}
type ReleaseQueryBuilderMock struct {
	// BuildSearchQueryFunc mocks the BuildSearchQuery method.
	BuildSearchQueryFunc func(ctx context.Context, request interface{}) ([]client.Search, error)

	// calls tracks calls to the methods.
	calls struct {
		// BuildSearchQuery holds details about calls to the BuildSearchQuery method.
		BuildSearchQuery []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Request is the request argument value.
			Request interface{}
		}
	}
	lockBuildSearchQuery sync.RWMutex
}

// BuildSearchQuery calls BuildSearchQueryFunc.
func (mock *ReleaseQueryBuilderMock) BuildSearchQuery(ctx context.Context, request interface{}) ([]client.Search, error) {
	if mock.BuildSearchQueryFunc == nil {
		panic("ReleaseQueryBuilderMock.BuildSearchQueryFunc: method is nil but ReleaseQueryBuilder.BuildSearchQuery was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Request interface{}
	}{
		Ctx:     ctx,
		Request: request,
	}
	mock.lockBuildSearchQuery.Lock()
	mock.calls.BuildSearchQuery = append(mock.calls.BuildSearchQuery, callInfo)
	mock.lockBuildSearchQuery.Unlock()
	return mock.BuildSearchQueryFunc(ctx, request)
}

// BuildSearchQueryCalls gets all the calls that were made to BuildSearchQuery.
// Check the length with:
//
//	len(mockedReleaseQueryBuilder.BuildSearchQueryCalls())
func (mock *ReleaseQueryBuilderMock) BuildSearchQueryCalls() []struct {
	Ctx     context.Context
	Request interface{}
} {
	var calls []struct {
		Ctx     context.Context
		Request interface{}
	}
	mock.lockBuildSearchQuery.RLock()
	calls = mock.calls.BuildSearchQuery
	mock.lockBuildSearchQuery.RUnlock()
	return calls
}

// Ensure, that ResponseTransformerMock does implement ResponseTransformer.
// If this is not the case, regenerate this file with moq.
var _ ResponseTransformer = &ResponseTransformerMock{}

// ResponseTransformerMock is a mock implementation of ResponseTransformer.
//
//	func TestSomethingThatUsesResponseTransformer(t *testing.T) {
//
//		// make and configure a mocked ResponseTransformer
//		mockedResponseTransformer := &ResponseTransformerMock{
//			TransformCountResponseFunc: func(ctx context.Context, responseData []byte) (int, error) {
//				panic("mock out the TransformCountResponse method")
//			},
//			TransformSearchResponseFunc: func(ctx context.Context, responseData []byte, queryMoqParam string, highlight bool) ([]byte, error) {
//				panic("mock out the TransformSearchResponse method")
//			},
//		}
//
//		// use mockedResponseTransformer in code that requires ResponseTransformer
//		// and then make assertions.
//
//	}
type ResponseTransformerMock struct {
	// TransformCountResponseFunc mocks the TransformCountResponse method.
	TransformCountResponseFunc func(ctx context.Context, responseData []byte) (int, error)

	// TransformSearchResponseFunc mocks the TransformSearchResponse method.
	TransformSearchResponseFunc func(ctx context.Context, responseData []byte, queryMoqParam string, highlight bool) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// TransformCountResponse holds details about calls to the TransformCountResponse method.
		TransformCountResponse []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResponseData is the responseData argument value.
			ResponseData []byte
		}
		// TransformSearchResponse holds details about calls to the TransformSearchResponse method.
		TransformSearchResponse []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResponseData is the responseData argument value.
			ResponseData []byte
			// QueryMoqParam is the queryMoqParam argument value.
			QueryMoqParam string
			// Highlight is the highlight argument value.
			Highlight bool
		}
	}
	lockTransformCountResponse  sync.RWMutex
	lockTransformSearchResponse sync.RWMutex
}

// TransformCountResponse calls TransformCountResponseFunc.
func (mock *ResponseTransformerMock) TransformCountResponse(ctx context.Context, responseData []byte) (int, error) {
	if mock.TransformCountResponseFunc == nil {
		panic("ResponseTransformerMock.TransformCountResponseFunc: method is nil but ResponseTransformer.TransformCountResponse was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ResponseData []byte
	}{
		Ctx:          ctx,
		ResponseData: responseData,
	}
	mock.lockTransformCountResponse.Lock()
	mock.calls.TransformCountResponse = append(mock.calls.TransformCountResponse, callInfo)
	mock.lockTransformCountResponse.Unlock()
	return mock.TransformCountResponseFunc(ctx, responseData)
}

// TransformCountResponseCalls gets all the calls that were made to TransformCountResponse.
// Check the length with:
//
//	len(mockedResponseTransformer.TransformCountResponseCalls())
func (mock *ResponseTransformerMock) TransformCountResponseCalls() []struct {
	Ctx          context.Context
	ResponseData []byte
} {
	var calls []struct {
		Ctx          context.Context
		ResponseData []byte
	}
	mock.lockTransformCountResponse.RLock()
	calls = mock.calls.TransformCountResponse
	mock.lockTransformCountResponse.RUnlock()
	return calls
}

// TransformSearchResponse calls TransformSearchResponseFunc.
func (mock *ResponseTransformerMock) TransformSearchResponse(ctx context.Context, responseData []byte, queryMoqParam string, highlight bool) ([]byte, error) {
	if mock.TransformSearchResponseFunc == nil {
		panic("ResponseTransformerMock.TransformSearchResponseFunc: method is nil but ResponseTransformer.TransformSearchResponse was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ResponseData  []byte
		QueryMoqParam string
		Highlight     bool
	}{
		Ctx:           ctx,
		ResponseData:  responseData,
		QueryMoqParam: queryMoqParam,
		Highlight:     highlight,
	}
	mock.lockTransformSearchResponse.Lock()
	mock.calls.TransformSearchResponse = append(mock.calls.TransformSearchResponse, callInfo)
	mock.lockTransformSearchResponse.Unlock()
	return mock.TransformSearchResponseFunc(ctx, responseData, queryMoqParam, highlight)
}

// TransformSearchResponseCalls gets all the calls that were made to TransformSearchResponse.
// Check the length with:
//
//	len(mockedResponseTransformer.TransformSearchResponseCalls())
func (mock *ResponseTransformerMock) TransformSearchResponseCalls() []struct {
	Ctx           context.Context
	ResponseData  []byte
	QueryMoqParam string
	Highlight     bool
} {
	var calls []struct {
		Ctx           context.Context
		ResponseData  []byte
		QueryMoqParam string
		Highlight     bool
	}
	mock.lockTransformSearchResponse.RLock()
	calls = mock.calls.TransformSearchResponse
	mock.lockTransformSearchResponse.RUnlock()
	return calls
}

// Ensure, that AuthHandlerMock does implement AuthHandler.
// If this is not the case, regenerate this file with moq.
var _ AuthHandler = &AuthHandlerMock{}

// AuthHandlerMock is a mock implementation of AuthHandler.
//
//	func TestSomethingThatUsesAuthHandler(t *testing.T) {
//
//		// make and configure a mocked AuthHandler
//		mockedAuthHandler := &AuthHandlerMock{
//			RequireFunc: func(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc {
//				panic("mock out the Require method")
//			},
//		}
//
//		// use mockedAuthHandler in code that requires AuthHandler
//		// and then make assertions.
//
//	}
type AuthHandlerMock struct {
	// RequireFunc mocks the Require method.
	RequireFunc func(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc

	// calls tracks calls to the methods.
	calls struct {
		// Require holds details about calls to the Require method.
		Require []struct {
			// Required is the required argument value.
			Required auth.Permissions
			// Handler is the handler argument value.
			Handler http.HandlerFunc
		}
	}
	lockRequire sync.RWMutex
}

// Require calls RequireFunc.
func (mock *AuthHandlerMock) Require(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc {
	if mock.RequireFunc == nil {
		panic("AuthHandlerMock.RequireFunc: method is nil but AuthHandler.Require was just called")
	}
	callInfo := struct {
		Required auth.Permissions
		Handler  http.HandlerFunc
	}{
		Required: required,
		Handler:  handler,
	}
	mock.lockRequire.Lock()
	mock.calls.Require = append(mock.calls.Require, callInfo)
	mock.lockRequire.Unlock()
	return mock.RequireFunc(required, handler)
}

// RequireCalls gets all the calls that were made to Require.
// Check the length with:
//
//	len(mockedAuthHandler.RequireCalls())
func (mock *AuthHandlerMock) RequireCalls() []struct {
	Required auth.Permissions
	Handler  http.HandlerFunc
} {
	var calls []struct {
		Required auth.Permissions
		Handler  http.HandlerFunc
	}
	mock.lockRequire.RLock()
	calls = mock.calls.Require
	mock.lockRequire.RUnlock()
	return calls
}
