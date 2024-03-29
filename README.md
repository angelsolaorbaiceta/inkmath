# InkMath

Go module implementing matrices, vectors and linear equation system solvers.

# Packages

This module is made of four packages:

- _nums_: utilities for working with numbers
- _vec_: definition of the vector primitive
- _mat_: definition of sparse and dense matrices
- _lineq_: linear equation system solvers

## Number Utilities

Includes number comparison functions:

### `FloatsEqualEps`

```go
func FloatsEqualEps(a, b, epsilon float64) bool
```

Compares two numbers, `a` and `b`, and returns `true` if both numbers are closer to eachother than the third argument, `epsilon`.

### `FloatsEqual`

```go
func FloatsEqual(a, b float64) bool
```

Compares two numbers, `a` and `b`, and returns `true` if both numbers are closer to eachother than a default `epsilon` of $1 \times 10^{10}$.

### `IsCloseToZero`

```go
func IsCloseToZero(a float64) bool
```

Returns `true` whether the number `a` is closer to zero than the default `epsilon` of $1 \times 10^{10}$.

---

A function to do linear interpolations:

### `LinInterpol`

```go
func LinInterpol(startPos, startVal, endPos, endVal, posToInterpolate float64) float64
```

Computes the y-value of a linear function at x = `posToInterpolate`, where the line is defined by the two points: (`startPos`, `startVal`) and (`endPos`, `endVal`).

---

And function to operate with open ranges:

### `IsInsideOpenRange`

```go
func IsInsideOpenRange(val, start, end float64) bool
```

Returns `true` if `start < val` and `val < end`.

### `DoRangesOverlap`

```go
func DoRangesOverlap(oneStart, oneEnd, twoStart, twoEnd float64) bool
```

Returns `true` if the ranges `[oneStart, oneEnd]` and `[twoStart, twoEnd]` overlap.

### `RangesOverlap`

```go
func RangesOverlap(oneStart, oneEnd, twoStart, twoEnd float64,) (ok bool, start, end float64)
```

If the ranges `[oneStart, oneEnd]` and `[twoStart, twoEnd]` overlap, returns `ok = true` and the overlap as a new range: `[start, end]`.
In case there is no overlap, `ok = false`.

## Vectors

A vector is a linear array of floating point numbers.
There is a single `Vector` struct implementation, but there are two interfaces to use vectors:

### `ReadOnlyVector`

Represents a vector which operations don't allow the mutation of its data.
The interface is defined as follows:

```go
type ReadOnlyVector interface {
	Length() int
	Norm() float64

	Value(i int) float64
	Opposite() ReadOnlyVector
	Scaled(factor float64) ReadOnlyVector
	Plus(other ReadOnlyVector) ReadOnlyVector
	Minus(other ReadOnlyVector) ReadOnlyVector
	Times(other ReadOnlyVector) float64

	Clone() ReadOnlyVector
	Equals(other ReadOnlyVector) bool
	AsMutable() MutableVector
}
```

Use a `ReadOnlyVector` as the type of those vector instances which shouldn't mutate their data.

### `MutableVector`

Represents a vector which, appart from the methods present in a `ReadOnlyVector`, allows modifying the vector's data.
The interface is defined as follows:

```go
type MutableVector interface {
	ReadOnlyVector

	SetValue(i int, value float64)
	SetZero(i int)
}
```

Use a `MutableVector` as the type of those vector instances which need to mutate their data.

### Vector Factories

You can create vectors using any of the following functions:

#### `Make`

```go
func Make(size int) MutableVector
```

Creates a `MutableVector` instance with the given size and filled with zeroes.

#### `MakeReadOnly`

```go
func MakeReadOnly(size int) ReadOnlyVector
```

Creates a `ReadOnlyVector` instance with the given size and filled with zeroes.

#### `MakeWithValues`

```go
func MakeWithValues(vals []float64) MutableVector
```

Creates a `MutableVector` with the passed in values (`vals`).
The length of the vector will be `len(vals)`.

## Matrices

A matrix is a bi-dimensional array of floating point numbers.
There are two matrix implementations (sparse and dense), and two interfaces to use matrices.
The two interfaces, following the same approach as with the vectors are:

### `ReadOnlyMatrix`

Represents a matrix which operations don't allow the mutation of its data.
The interface is defined as follows:

```go
type ReadOnlyMatrix interface {
	Rows() int
	Cols() int
	NonZeroIndicesAtRow(int) []int

	Value(int, int) float64

	RowTimesVector(row int, v vec.ReadOnlyVector) float64
	TimesVector(v vec.ReadOnlyVector) vec.ReadOnlyVector
	TimesMatrix(other ReadOnlyMatrix) ReadOnlyMatrix
}
```

### `MutableMatrix`

Represents a matrix that, appart from the methods in a read-only matrix, has operations to mutate its data.
The interface is defined as follows:

```go
type MutableMatrix interface {
	ReadOnlyMatrix

	SetValue(int, int, float64)
	AddToValue(int, int, float64)

	SetZeroCol(int)
	SetIdentityRow(int)
}
```

---

As mentioned above, there are two different implementations for a matrix: dense and sparse.
The sparse matrix only stores those values that are different from zero, thus reducing the memory footprint in case of large matrices containing lots of zeros.

### `SparseMat`

Represents a sparse matrix, where only non-zero values are stored.
The `SparseMat` struct implementation satisfies both the `ReadOnlyMatrix` and `MutableMatrix` interfaces.

### `DenseMat`

Represents a dense matrix, where every value (including zeroes) are stored.
The `DenseMat` struct implementation satisfies both the `ReadOnlyMatrix` and `MutableMatrix` interfaces.

## Linear Equation Solvers

`InkMath` contains one interface defining the contract for all the linear equation solver implementations: `Solver`:

```go
type Solver interface {
	CanSolve(coefficients mat.ReadOnlyMatrix, freeTerms vec.ReadOnlyVector) bool
	Solve(coefficients mat.ReadOnlyMatrix, freeTerms vec.ReadOnlyVector) *Solution
}
```

This interface has two methods:

- `CanSolve` to test whether a given set of linear equations can be solved by the current method
- `Solve` to solve the system and produce a solution

The returned `Solution` instance is defined as follows:

```go
type Solution struct {
	ReachedMaxIter bool
	MinError       float64
	IterCount      int
	Solution       vec.ReadOnlyVector
}
```

where:

- `ReachedMaxIter` is a flag that indicates, in the case of iterative methods, whether the maximum number of iterations was reached before a good enough solution could be found
- `MinError` a bound of the estimated error of the solution, which can be made as small as required
- `IterCount` the number of iterations necessary to find a solution
- `Solution` the solution vector
