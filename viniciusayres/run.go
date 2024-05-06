package desafiogo

import (
	"sort"
	"sync"
)

// Optimized binary search implementation for the hasNumber47 check
func binarySearch(data []float64, seek float64) bool {
	begin, end := 0, len(data)-1

	for begin <= end {
		mid := (begin + end) / 2

		if data[mid] == seek {
			return true
		} else if data[mid] < seek {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

// Utilize more efficient data structures for mode calculation
func calculateMode(data []float64) float64 {
	modeMap := make(map[float64]int)
	var mode float64
	maxCount := 0

	for _, value := range data {
		modeMap[value]++
		if modeMap[value] > maxCount {
			maxCount = modeMap[value]
			mode = value
		}
	}

	if maxCount <= 1 {
		return 0
	}

	return mode
}

func Run(data []float64) *Result {
	if !sort.IsSorted(sort.Float64Slice(data)) {
		sort.Float64s(data)
	}

	var ave float64
	var med float64
	var p90 float64
	var p99 float64
	var mod float64
	hasNumber47 := false
	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		ave = calculateAverage(data)
		wg.Done()
	}()
	go func() {
		med = calculateMedian(data)
		wg.Done()
	}()
	go func() {
		p90 = calculatePercentile(data, 90)
		wg.Done()
	}()
	go func() {
		p99 = calculatePercentile(data, 99)
		wg.Done()
	}()

	go func() {
		mod = calculateMode(data)
		wg.Done()
	}()

	go func() {
		hasNumber47 = binarySearch(data, 47)
		wg.Done()
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
