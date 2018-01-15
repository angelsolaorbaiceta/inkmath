package vec

// Vector is an array of values.
type Vector struct {
	length int
	data   []float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

// Make returns a vector with the given size all filled with zeroes.
func Make(size int) *Vector {
	return &Vector{size, make([]float64, size)}
}
