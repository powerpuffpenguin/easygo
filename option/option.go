package option

// Option is used to implement optional parameters
type Option[T any] interface {
	Apply(*T)
}

// Create the function f as an optional parameter
func New[T any](f func(*T)) Option[T] {
	return newFuncOptions(f)
}

type funcOptions[T any] struct {
	f func(*T)
}

func newFuncOptions[T any](f func(*T)) *funcOptions[T] {
	return &funcOptions[T]{
		f: f,
	}
}
func (o *funcOptions[T]) Apply(opts *T) {
	o.f(opts)
}
