package median

import (
	"math"
)

func GetMedian(i []int) float64 {
	var medianIndex float64 = float64(len(i)-1) / 2
	aIndex := int(math.Floor(medianIndex))
	bIndex := int(math.Ceil(medianIndex))
	return float64(i[aIndex]+i[bIndex]) / 2
}
