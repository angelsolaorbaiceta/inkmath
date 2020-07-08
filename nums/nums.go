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

import "math"

const defaultEpsilon = 1E-10

/*
FuzzyEqualsEps compares two float64 values and returns true if the difference between
the two is smaller than a given epsilon.
*/
func FuzzyEqualEps(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

/*
FuzzyEqual compares two float64 values and returns true if the difference between the
two is smaller than a default epsilon value (1E-10).
*/
func FuzzyEqual(a, b float64) bool {
	return FuzzyEqualEps(a, b, defaultEpsilon)
}

/*
IsCloseToZero returns true if the given number is close enough to zero.
Comparation made with the default epsilon value.
*/
func IsCloseToZero(a float64) bool {
	return FuzzyEqual(a, 0.0)
}

/*
LinInterpol computes the linear interpolation for a given position given two points on
the desired line: (startPos, startVal) and (endPos, endVal).
*/
func LinInterpol(startPos, startVal, endPos, endVal, posToInterpolate float64) float64 {
	return startVal + (posToInterpolate-startPos)*(endVal-startVal)/(endPos-startPos)
}
