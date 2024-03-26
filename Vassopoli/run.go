package desafiogo

func Run(data []float64) *Result {
	ave := calculateAverage(data)
	med := calculateMedian(data)
	p90 := calculatePercentile(data, 90)
	p99 := calculatePercentile(data, 99)

	mod, count := float64(0), 1
	hasNumber47 := false

	var floatNumberToModCountMap map[float64]int = make(map[float64]int)

	for _, d := range data {
		if (floatNumberToModCountMap[d]) == 0 {
			floatNumberToModCountMap[d] = 1
		} else {
			floatNumberToModCountMap[d] += 1

			if floatNumberToModCountMap[d] > count {
				mod, count = d, floatNumberToModCountMap[d]
			}
		}
	}

	if floatNumberToModCountMap[47] > 0 {
		hasNumber47 = true
	}

	return &Result{
		Average:      ave,
		Median:       med,
		Percentile90: p90,
		Percentile99: p99,
		Mode:         mod,
		HasNumber47:  hasNumber47,
	}
}
