package app

import (
	"errors"
	"strings"
	"github.com/jinzhu/now"
	"github.com/saromanov/selitra/backend/internal/storage"
)

var (
	errDateParse = errors.New("unable to parse date")
	errDateRange = errors.New("unknown format of the date. It should be today, week, month, year")
)

// parseQuery provides parsing of the query and convert
// it to the search request
func parseQuery(query string)(*storage.SearchRequest, error) {
	exprs := strings.Split(query, ";")
	for i := 0;i < len(exprs);i++ {
		switch exp
	}
}

// getDate returns from and to timestamps
// it should be in format 'date=today'
func getDate(expr string) (int64, int64, error) {
	values := strings.Split(expr, "=")
	if len(values) <= 1 {
		return 0, 0, errDateParse
	}

	switch values[1] {
	case "today":
		return now.BeginningOfDay().UnixNano(), now.EndOfDay().UnixNano(), nil
	}

	return 0, 0, errDateRange
}