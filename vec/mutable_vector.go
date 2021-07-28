package vec

type MutableVector interface {
	ReadOnlyVector

	/* Methods */
	SetValue(i int, value float64)
	SetZero(i int)
}

// SetValue sets the given value at the given index.
func (v *Vector) SetValue(i int, value float64) {
	v.data[i] = value
}

// SetZero sets a zero value in the given index.
func (v *Vector) SetZero(i int) {
	v.data[i] = 0.0
}
