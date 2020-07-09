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

package vec

import "github.com/angelsolaorbaiceta/inkmath/nums"

/*
VectorContainsData tests whether a given vector contains exactly the same data as the
slice of float64 numbers.

Both need to be of the same size in order for this test to be true.
*/
func VectorContainsData(vector *Vector, data []float64) bool {
	if vector.length != len(data) {
		return false
	}

	for i := 0; i < vector.length; i++ {
		if !nums.FuzzyEqual(vector.Value(i), data[i]) {
			return false
		}
	}

	return true
}
