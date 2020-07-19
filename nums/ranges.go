/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nums

import (
	"math"
)

/*
IsInsideOpenRange returns true if a given number lays inside an open range (start, end).
*/
func IsInsideOpenRange(val, start, end float64) bool {
	return val > start && val < end
}

/*
DoRangesOverlap returns true if two given ranges (oneStart, oneEnd) and
(twoStart, twoEnd) overlap.
*/
func DoRangesOverlap(oneStart, oneEnd, twoStart, twoEnd float64) bool {
	if FuzzyEqual(oneStart, twoStart) && FuzzyEqual(oneEnd, twoEnd) {
		return true
	}

	return IsInsideOpenRange(oneStart, twoStart, twoEnd) ||
		IsInsideOpenRange(oneEnd, twoStart, twoEnd) ||
		IsInsideOpenRange(twoStart, oneStart, oneEnd) ||
		IsInsideOpenRange(twoEnd, oneStart, oneEnd)
}

/*
RangesOverlap returns the overlapping of the two given ranges.
*/
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
