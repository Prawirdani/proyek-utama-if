package model

type Query interface {
	Build(baseQuery string) (query string, countQuery string, stmtArgs []interface{})
	SetCount(count int)
}
