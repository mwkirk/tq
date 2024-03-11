package container

type Filterable[T any] interface {
	Filter(func(T) bool) []T
}
