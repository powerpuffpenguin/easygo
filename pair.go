package easygo

type Pair[T0 any, T1 any] struct {
	First  T0
	Second T1
}

func MakePair[T0 any, T1 any](first T0, second T1) Pair[T0, T1] {
	return Pair[T0, T1]{
		First:  first,
		Second: second,
	}
}
func NewPair[T0 any, T1 any](first T0, second T1) *Pair[T0, T1] {
	return &Pair[T0, T1]{
		First:  first,
		Second: second,
	}
}
