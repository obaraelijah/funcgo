package funcgo

func Foreach[T any](f func(elem T)) func(iterable []T) {
	return func(iterable []T) {
		for i := 0; i < len(iterable); i++ {
			f(iterable[i])
		}
	}
}

func Map[S []E, T []F, E any, F any](f func(elem E) F) func(iterable S) T {
	return func(iterable S) T {
		r := make([]F, 0)
		for _, elem := range iterable {
			r = append(r, f(elem))
		}
		return r
	}
}

func Filter[T []E, E any](f func(elem E) bool) func(iterable T) T {
	return func(iterable T) T {
		r := make([]E, 0)
		for _, elem := range iterable {
			if f(elem) {
				r = append(r, elem)
			}
		}
		return r
	}
}

func Exclude[T []E, E any](f func(elem E) bool) func(iterable T) T {
	return func(iterable T) T {
		r := make([]E, 0)
		for _, elem := range iterable {
			if !f(elem) {
				r = append(r, elem)
			}
		}
		return r
	}
}

func Any[T []E, E any](f func(elem E) bool) func(iterable T) bool {
	return func(iterable T) bool {
		for i := 0; i < len(iterable); i++ {
			if f(iterable[i]) {
				return true
			}
		}
		return false
	}
}

func All[T []E, E any](f func(elem E) bool) func(iterable T) bool {
	return func(iterable T) bool {
		for i := 0; i < len(iterable); i++ {
			if !f(iterable[i]) {
				return false
			}
		}
		return true
	}
}

func None[T []E, E any](f func(elem E) bool) func(iterable T) bool {
	return func(iterable T) bool {
		for i := 0; i < len(iterable); i++ {
			if f(iterable[i]) {
				return false
			}
		}
		return true
	}
}

func CountMatch[T []E, E any](f func(elem E) bool) func(iterable T) int {
	return func(iterable T) int {
		count := 0
		for i := 0; i < len(iterable); i++ {
			if f(iterable[i]) {
				count++
			}
		}
		return count
	}
}

func Reduce[T []E, E any](f func(prev, next E) E) func(iterable T) E {
	return func(iterable T) E {
		var res E
		for idx, elem := range iterable {
			if idx == 0 {
				res = elem
				continue
			}
			res = f(res, elem)
		}
		return res
	}
}
