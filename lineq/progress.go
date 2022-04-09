package lineq

import "math"

type IterativeSolverProgress struct {
	ProgressPercentage int
	Error              float64
	IterCount          int
}

// computeProgressPercentage returns the percentage of the progress or the current error.
// An error equal to, or smaller than, the maxError will return 100%.
// Every order of magnitude larger than the maxError will return 10% of the maxError, thus
// the progress is logarithmic.
func computeProgressPercentage(maxError, currentError float64) int {
	diff := math.Log10(currentError) - math.Log10(maxError)
	capDiff := math.Max(math.Min(diff, 10), 0)

	return int(10 * (10 - capDiff))
}
