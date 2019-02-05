package app

import (
	"strings"
	"github.com/saromanov/selitra/backend/internal/storage"
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
	if len()
}