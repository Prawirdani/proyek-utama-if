package model

import (
	"fmt"
	"net/http"
)

// available query options
// K = query param, V = db field
var pesananQueries = map[string]string{
	"id":         "id",
	"status":     "p.status_pesanan",
	"mejaID":     "m.id",
	"statusMeja": "m.status",
}

type Query struct {
	KeyValue map[string]interface{}
	StmtArgs []interface{}
}

// K = db field, V = query param value
func ParsePesananQuery(r *http.Request) (*Query, error) {
	q := r.URL.Query()
	kv := make(map[string]interface{})

	for k, v := range pesananQueries {
		if q.Get(k) != "" {
			kv[v] = q.Get(k)
		}
	}

	if len(kv) == 0 {
		return nil, fmt.Errorf("no query parameter found")
	}

	return &Query{
		KeyValue: kv,
	}, nil
}

func (q *Query) Build(baseQuery string) string {
	query := baseQuery
	count := 1
	clause := " WHERE"

	for k, v := range q.KeyValue {
		query += fmt.Sprintf("%s %s=$%v", clause, k, count)
		q.StmtArgs = append(q.StmtArgs, v)
		clause = " AND"
		count++
	}

	fmt.Println(query)
	return query
}
