package httputil

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type Pagination struct {
	// Page query field name ?PageQuery=
	PageQuery string
	// Page query value
	Page int
	// Limit query field name ?LimitQuery=
	LimitQuery string
	// Limit query value
	Limit int
}

type Filter struct {
	// K = query param field, V = db field
	// TODO: Should use a struct for this
	Fields map[string]string
	// K = db field, V = query param value
	Values map[string]interface{}
}

type Sort struct {
	// Sort and Order param query (sort=name&order=ASC/DESC)
	// A Map for storing available sort fields, e.g. {"name": "db_field_name", "created_at": "db_created_at"}
	// K = query param field, V = db field
	// It will ignore the sort query if the order by query is not provider.
	// Otherwise if the sort query is provided, but no order query is provided, it will return error.
	Fields map[string]string

	// K = SORT/ORDER, V = DB FIELD or ASC/DESC
	Values map[string]interface{}
}

type queryOption func(*QueryProcessor)

type QueryProcessor struct {
	Pagination *Pagination
	Filter     *Filter
	Sort       *Sort
}

// Enable filtering with query param. Catch all defined query param and its value, then map it to db field.
// It will ignored if no query param found.
// fc stands for field column, which is the mapping of query param to db field.
// define the available query param and its corresponding db field.
func WithFilter(fc map[string]string) queryOption {
	return func(q *QueryProcessor) {
		q.Filter = &Filter{Fields: fc}
	}
}

// Enable pagination with query param.
// If this option is enabled, it will require both page and limit query param.
// It will return error if no query param found.
func WithPagination(pageQ, limitQ string) queryOption {
	return func(q *QueryProcessor) {
		q.Pagination = &Pagination{
			PageQuery:  pageQ,
			LimitQuery: limitQ,
		}
	}
}

// Enable sorting with query param.
// It will ignored if no sort & order query param found.
func WithSort(fc map[string]string) queryOption {
	return func(q *QueryProcessor) {
		q.Sort = &Sort{Fields: fc}
	}
}

func NewQueryProcessor(opts ...queryOption) *QueryProcessor {
	q := &QueryProcessor{}

	for _, opt := range opts {
		opt(q)
	}
	return q
}

const MAX_LIMIT = 100

func (q *QueryProcessor) Parse(r *http.Request) error {
	querier := r.URL.Query()
	if q.Filter != nil {
		q.Filter.Values = make(map[string]interface{})
		// q.Filter.StmtArgs = []interface{}{}

		for k, v := range q.Filter.Fields {
			if querier.Get(k) != "" {
				q.Filter.Values[v] = querier.Get(k)
			}
		}

		if len(q.Filter.Values) == 0 {
			return fmt.Errorf("no query parameter found")
		}
	}

	if q.Pagination != nil {
		pageQuery := querier.Get(q.Pagination.PageQuery)
		limitQuery := querier.Get(q.Pagination.LimitQuery)

		if pageQuery == "" && limitQuery == "" {
			return ErrBadRequest("missing page or limit query.")
		}

		pageInt, err := strconv.Atoi(pageQuery)
		if err != nil {
			return ErrBadRequest("invalid page query")
		}
		q.Pagination.Page = pageInt

		limitInt, err := strconv.Atoi(limitQuery)
		if err != nil {
			return ErrBadRequest("invalid limit query")
		}

		if limitInt > MAX_LIMIT {
			return ErrBadRequest("limit query exceeds maximum limit of 100.")
		}
		q.Pagination.Limit = limitInt
	}

	if q.Sort != nil {
		q.Sort.Values = make(map[string]interface{})
		sortQuery := querier.Get("sort")
		orderQuery := strings.ToUpper(querier.Get("order"))
		for k, v := range q.Sort.Fields {
			if sortQuery == k {
				if orderQuery != "ASC" && orderQuery != "DESC" {
					return ErrBadRequest("invalid or missing order query")
				}
				q.Sort.Values[v] = orderQuery
			}
		}
	}

	return nil
}

// Build query string with safety prepared statement arguments placeholder
func (q *QueryProcessor) Build(baseQuery string) (query string, stmtArgs []interface{}) {
	query = baseQuery
	stmtCount := 1
	if q.Filter != nil && q.Filter.Values != nil {
		clause := " WHERE"
		for k, v := range q.Filter.Values {
			query += fmt.Sprintf("%s %s=$%v", clause, k, stmtCount)
			stmtArgs = append(stmtArgs, v)
			clause = " AND"
			stmtCount++
		}
	}

	if q.Sort != nil && q.Sort.Values != nil {
		clause := " ORDER BY"
		for k, v := range q.Sort.Values {
			query += fmt.Sprintf("%s %s %s", clause, k, v)
		}
	}

	if q.Pagination != nil {
		query += fmt.Sprintf(" LIMIT $%v OFFSET $%v", stmtCount, stmtCount+1)
		stmtArgs = append(stmtArgs, q.Pagination.Limit, (q.Pagination.Page-1)*q.Pagination.Limit)
	}

	slog.Debug("QueryProcessor.Build", slog.Any("query", query), slog.Any("args", stmtArgs))
	return query, stmtArgs
}
