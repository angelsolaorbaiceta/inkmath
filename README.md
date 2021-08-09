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

If the ranges `[oneStart, oneEnd]` and `[twoStart, twoEnd]` overlap, returns `ok = true` and the overlap as a new range: `[start, end]`. In case there is no overlap, `ok = false`.

## Vectors

## Matrices

## Linear Equation Solvers
