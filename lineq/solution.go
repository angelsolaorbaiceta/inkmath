package lineq

import "github.com/angelsolaorbaiceta/inkmath/vec"

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

func makeSolution(iterCount int, solution *vec.Vector) *LineqSolution {
	return &LineqSolution{
		ReachedMaxIter: false,
		MinError:       0.0,
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
