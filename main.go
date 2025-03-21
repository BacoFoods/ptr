package ptr

// Ptr returns a pointer to the value passed in.
func Ptr[T any](v T) *T {
	return &v
}

// Val returns the value of the pointer passed in, or the default value if the pointer is nil.
func Val[T any](p *T, d T) T {
	if p == nil {
		return d
	}

	return *p
}

// EqualVals returns true if the values of the pointers passed in are equal.
// If either pointer is nil, it returns false.
func EqualVals[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return false
	}

	return *a == *b
}
