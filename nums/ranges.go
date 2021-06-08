package nums

import (
	"math"
)

// IsInsideOpenRange returns true if a given number lays inside an open range (start, end).
func IsInsideOpenRange(val, start, end float64) bool {
	return val > start && val < end
}

// DoRangesOverlap returns true if two given ranges (oneStart, oneEnd) and (twoStart, twoEnd) overlap.
func DoRangesOverlap(oneStart, oneEnd, twoStart, twoEnd float64) bool {
	if FuzzyEqual(oneStart, twoStart) && FuzzyEqual(oneEnd, twoEnd) {
		return true
	}

	return IsInsideOpenRange(oneStart, twoStart, twoEnd) ||
		IsInsideOpenRange(oneEnd, twoStart, twoEnd) ||
		IsInsideOpenRange(twoStart, oneStart, oneEnd) ||
		IsInsideOpenRange(twoEnd, oneStart, oneEnd)
}

// RangesOverlap returns the overlapping of the two given ranges.
func RangesOverlap(
	oneStart, oneEnd, twoStart, twoEnd float64,
) (ok bool, start, end float64) {
	if !DoRangesOverlap(oneStart, oneEnd, twoStart, twoEnd) {
		ok = false
		return
	}

	ok = true
	start = math.Max(oneStart, twoStart)
	end = math.Min(oneEnd, twoEnd)

	return
}
