package util

import goog_query "github.com/google/go-querystring/query"

type FilterSort string

type FilterParams struct {
	Period string     `url:"period"`
	Page   int        `url:"page"`
	SortBy FilterSort `url:"sort_by"`
}
type CommonParams struct {
	Query string `url:"q"`
}

type Params struct {
	FilterParams
	CommonParams
}

func ParseFilterParams[Q any](query Q) string {
	vals, _ := goog_query.Values(query)

	return vals.Encode()
}
