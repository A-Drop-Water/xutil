package xutil

func SliceAToBFunc[T any, V any](a []T, f func(T) V) []V {
	if a == nil {
		return nil
	}
	b := make([]V, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}

// SliceToMap 列表转map
func SliceToMap[T any, K comparable](arr []T, fun func(v T) K) map[K]T {
	res := make(map[K]T, len(arr))
	for _, v := range arr {
		res[fun(v)] = v
	}
	return res
}
