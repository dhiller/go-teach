package median

import (
	"sort"
)

func medianFunc(values []float64) float64 {
	sort.Float64s(values)
	i := len(values)
	return values[i/ 2]
}
