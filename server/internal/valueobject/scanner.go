package valueobject

type Row interface {
	Scan(dest ...any) error
}
