package internal

func SafeDeref[T comparable](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}

func Ref[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}
