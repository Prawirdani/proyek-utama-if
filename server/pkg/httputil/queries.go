package httputil

import (
	"fmt"
	"net/http"
)

type QueryParam struct {
	// K = query param, V = db field
	FieldColumn map[string]string
	// K = db field, V = query param value
	KeyValue map[string]interface{}
	StmtArgs []interface{}
}

// Query Parameters Parser
// fc stands for field column, which is the mapping of query param to db field
// define the available query param and its corresponding db field
func NewQueryParam(fc map[string]string) *QueryParam {
	return &QueryParam{
		FieldColumn: fc,
	}
}

func (q *QueryParam) Parse(r *http.Request) error {
	q.KeyValue = make(map[string]interface{})
	q.StmtArgs = []interface{}{}

	for k, v := range q.FieldColumn {
		if r.URL.Query().Get(k) != "" {
			q.KeyValue[v] = r.URL.Query().Get(k)
		}
	}

	if len(q.KeyValue) == 0 {
		return fmt.Errorf("no query parameter found")
	}
	return nil
}

func (q *QueryParam) Build(baseQuery string) string {
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

func (q *QueryParam) Args() []interface{} {
	return q.StmtArgs
}
