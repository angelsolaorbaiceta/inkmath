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

package lineq

import (
	"fmt"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
Solution is the solution data for a linear equation system solver.
*/
type Solution struct {
	ReachedMaxIter bool
	MinError       float64
	IterCount      int
	Solution       *vec.Vector
}

/* <-- Construction --> */

func makeSolution(iterCount int, minError float64, solution *vec.Vector) *Solution {
	return &Solution{
		ReachedMaxIter: false,
		MinError:       minError,
		IterCount:      iterCount,
		Solution:       solution,
	}
}

func makeErrorSolution(
	iterCount int,
	minError float64,
	partialSolution *vec.Vector,
) *Solution {
	return &Solution{
		ReachedMaxIter: true,
		MinError:       minError,
		IterCount:      iterCount,
		Solution:       partialSolution,
	}
}

/* <-- Methods --> */
func (sol Solution) String() string {
	if sol.ReachedMaxIter {
		return fmt.Sprintf(
			"[KO] -> Min Error: %f, Iter Count: %d",
			sol.MinError, sol.IterCount,
		)
	}

	return fmt.Sprintf(
		"[OK] -> Iter Count: %d, Min Error: %f, Solution: %v",
		sol.IterCount, sol.MinError, sol.Solution,
	)
}
