package lineq

import (
	"fmt"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
LineqSolution is the solution data for a linear equation system solver.
*/
type LineqSolution struct {
	ReachedMaxIter bool
	MinError       float64
	IterCount      int
	Solution       *vec.Vector
}

/* ::::::::::::::: Construction ::::::::::::::: */

func makeSolution(iterCount int, minError float64, solution *vec.Vector) *LineqSolution {
	return &LineqSolution{
		ReachedMaxIter: false,
		MinError:       minError,
		IterCount:      iterCount,
		Solution:       solution,
	}
}

func makeErrorSolution(iterCount int, minError float64, partialSolution *vec.Vector) *LineqSolution {
	return &LineqSolution{
		ReachedMaxIter: true,
		MinError:       minError,
		IterCount:      iterCount,
		Solution:       partialSolution,
	}
}

/* ::::::::::::::: Methods ::::::::::::::: */
func (sol LineqSolution) String() string {
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
