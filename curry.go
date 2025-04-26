package funcgo

// Curry2 curries a function with two arguments.
func Curry2[A, B, R any](f func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return f(a, b)
		}
	}
}

// Curry3 curries a function with three arguments.
func Curry3[A, B, C, R any](f func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return f(a, b, c)
			}
		}
	}
}

// Uncurry3 converts a curried function back to a function with three arguments.
func Uncurry3[A, B, C, R any](f func(A) func(B) func(C) R) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a)(b)(c)
	}
}
