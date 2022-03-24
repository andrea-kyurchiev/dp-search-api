package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/ONSdigital/dp-search-api/query"
	"github.com/ONSdigital/log.go/v2/log"
)

// SearchReleasesHandlerFunc returns a http handler function handling release calendar search api requests.
func SearchReleasesHandlerFunc(validator QueryParamValidator, builder ReleaseQueryBuilder, searcher ElasticSearcher, transformer ResponseTransformer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		params := req.URL.Query()

		q, err := url.QueryUnescape(params.Get("query"))
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "query", "value": params.Get("query")})
			http.Error(w, "Bad url encoding of the query parameter", http.StatusBadRequest)
			return
		}

		limitParam := paramGet(params, "limit", "10")
		limit, err := validator.Validate(ctx, "limit", limitParam)
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "limit", "value": limitParam})
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}

		offsetParam := paramGet(params, "offset", "0")
		offset, err := validator.Validate(ctx, "offset", offsetParam)
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "offset", "value": offsetParam})
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}

		sortParam := paramGet(params, "sort", query.RelDateAsc.String())
		sort, err := validator.Validate(ctx, "sort", sortParam)
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "sort", "value": sortParam})
			http.Error(w, "Invalid sort parameter", http.StatusBadRequest)
			return
		}

		fromDateParam := paramGet(params, "fromDate", "")
		fromDate, err := validator.Validate(ctx, "date", fromDateParam)
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "fromDate", "value": fromDateParam})
			http.Error(w, "Invalid dateFrom parameter", http.StatusBadRequest)
			return
		}

		toDateParam := paramGet(params, "toDate", "")
		toDate, err := validator.Validate(ctx, "date", toDateParam)
		if err != nil {
			log.Warn(ctx, err.Error(), log.Data{"param": "toDate", "value": toDateParam})
			http.Error(w, "Invalid dateTo parameter", http.StatusBadRequest)
			return
		}

		if time.Time(fromDate.(query.Date)).After(time.Time(toDate.(query.Date))) {
			log.Warn(ctx, "fromDate after toDate", log.Data{"fromDate": fromDateParam, "toDate": toDateParam})
			http.Error(w, "Invalid date parameters", http.StatusBadRequest)
			return
		}
		upcoming := paramGetBool(params, "type-upcoming", false)
		highlight := paramGetBool(params, "highlight", true)

		formattedQuery, err := builder.BuildSearchQuery(ctx,
			query.ReleaseSearchRequest{
				Term:           q,
				From:           offset.(int),
				Size:           limit.(int),
				SortBy:         sort.(query.Sort),
				ReleasedAfter:  fromDate.(query.Date),
				ReleasedBefore: toDate.(query.Date),
				Upcoming:       upcoming,
				Published:      !upcoming,
				Highlight:      highlight,
				Now:            query.Date(time.Now()),
			})
		if err != nil {
			log.Error(ctx, "creation of search release query failed", err, log.Data{"q": q, "sort": sort, "limit": limit, "offset": offset})
			http.Error(w, "Failed to create search release query", http.StatusInternalServerError)
			return
		}

		responseData, err := searcher.Search(ctx, "ons", "", formattedQuery)
		if err != nil {
			log.Error(ctx, "elasticsearch release query failed", err)
			http.Error(w, "Failed to run search release query", http.StatusInternalServerError)
			return
		}

		if !json.Valid(responseData) {
			log.Error(ctx, "elastic search returned invalid JSON for search release query", errors.New("elastic search returned invalid JSON for search release query"))
			http.Error(w, "Failed to process search release query", http.StatusInternalServerError)
			return
		}

		if !paramGetBool(params, "raw", false) {
			responseData, err = transformer.TransformSearchResponse(ctx, responseData, q, highlight)
			if err != nil {
				log.Error(ctx, "transformation of response data failed", err)
				http.Error(w, "Failed to transform search result", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		_, err = w.Write(responseData)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}