package desafiogo

import (
	"sort"
	"sync"
)

func Run(data []float64) *Result {
	var ave, med, p90, p99, mod float64
	var hasNumber47 bool
	var wg sync.WaitGroup
	wg.Add(6)

	// Enhanced map-based approach for mode calculation
	modeMap := make(map[float64]int)
	var modeVal float64
	var maxCount int

	go func() {
		defer wg.Done()
		for _, v := range data {
			modeMap[v]++
			if modeMap[v] > maxCount {
				modeVal = v
				maxCount = modeMap[v]
			}
		}
		mod = modeVal
	}()

	go func() {
		defer wg.Done()
		ave = calculateAverage(data)
	}()

	go func() {
		defer wg.Done()
		med = calculateMedian(data)
	}()

	go func() {
		defer wg.Done()
		p90 = calculatePercentile(data, 90)
	}()

	go func() {
		defer wg.Done()
		p99 = calculatePercentile(data, 99)
	}()

	// Parallel processing for the hasNumber47 check
	go func() {
		defer wg.Done()
		sort.Float64s(data) // Ensure data is sorted for binary search
		hasNumber47 = sort.SearchFloat64s(data, 47) < len(data) && data[sort.SearchFloat64s(data, 47)] == 47
	}()

	wg.Wait()

	return &Result{
		Average:      ave,
		Median:       med,
		Percentile90: p90,
		Percentile99: p99,
		Mode:         mod,
		HasNumber47:  hasNumber47,
	}
}
