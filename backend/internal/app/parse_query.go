package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/now"
	"github.com/saromanov/selitra/backend/internal/storage"
)

var (
	errQueryParse = errors.New("unable to parse query")
	errDateRange  = errors.New("unknown format of the date. It should be today, week, month, year")
)

// parseQuery provides parsing of the query and convert
// it to the search request
func parseQuery(query string) (*storage.SearchRequest, error) {
	exprs := strings.Split(query, " ")
	fmt.Println("QUERY: ", exprs)
	result := &storage.SearchRequest{}
	for i := 0; i < len(exprs); i++ {
		if strings.HasPrefix(exprs[i], "date") {
			fts, ets, err := getDate(exprs[i])
			if err != nil {
				return nil, err
			}
			result.FromTimestamp = fts
			result.ToTimestamp = ets
		}
		if strings.HasPrefix(exprs[i], "service") {
			service, err := getService(exprs[i])
			if err != nil {
				return nil, err
			}
			result.Service = service
		}
	}

	return result, nil
}

// getDate returns from and to timestamps
// it should be in format 'date=today'
func getDate(expr string) (int64, int64, error) {
	values := strings.Split(expr, "=")
	if len(values) <= 1 {
		return 0, 0, errQueryParse
	}

	switch values[1] {
	case "today":
		return now.BeginningOfDay().UnixNano(), now.EndOfDay().UnixNano(), nil
	case "week":
		return now.BeginningOfWeek().UnixNano(), now.EndOfDay().UnixNano(), nil
	case "month":
		return now.BeginningOfWeek().UnixNano(), now.EndOfDay().UnixNano(), nil
	}

	return 0, 0, errDateRange
}

// getService returns service name in the case "service=name"
func getService(line string) (string, error) {
	values := strings.Split(line, "=")
	if len(values) <= 1 {
		return "", errQueryParse
	}
	return values[1], nil
}
