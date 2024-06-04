package model

type Query interface {
	Build(baseQuery string) (string, []interface{})
}
