# InkMath

Go module for mathematical and numerical calculations.

## Publish New Version

To publish a new version of this module, first tag it respecting [semantic versioning](https://semver.org/):

```bash
git tag vM.m.p
```

Then push the tag:

```bash
git push origin vM.m.p
```

# Packages

This module is made of four packages:

- _nums_: utilities for working with numbers
- _vec_: definition of the vector primitive
- _mat_: definition of sparse and dense matrices
- _lineq_: linear equation system solvers

## Number Utilities

Includes number comparison functions:

### `FuzzyEqualEps`

```go
func FuzzyEqualEps(a, b, epsilon float64) bool
```

Compares two numbers, `a` and `b`, and returns `true` if both numbers are closer to eachother than the third argument, `epsilon`.

### `FuzzyEqual`

```go
func FuzzyEqual(a, b float64) bool
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

## Linear Equation Solvers
