package easygo

func Map[T0 any, T1 any](vals []T0, to func(v T0) T1) []T1 {
	arrs := make([]T1, len(vals))
	for i, val := range vals {
		arrs[i] = to(val)
	}
	return arrs
}
