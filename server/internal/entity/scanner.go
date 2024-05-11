package entity

type Row interface {
	Scan(dest ...any) error
}
