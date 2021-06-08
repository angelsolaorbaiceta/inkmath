package lineq

import (
	"fmt"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

// Solution is the solution data for a linear equation system solver.
type Solution struct {
	ReachedMaxIter bool
	MinError       float64
	IterCount      int
	Solution       *vec.Vector
}

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
