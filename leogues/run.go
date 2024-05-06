package desafiogo

import (
	"sort"
	"sync"
)

// Improved sorting algorithm for better efficiency
func calculateModeAndHasNumber47(data []float64) (float64, bool) {
	if !sort.Float64sAreSorted(data) {
		sort.Float64s(data)
	}

	var mode float64
	var maxFrequency, frequencyCount int
	hasNumber47 := false

	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			frequencyCount++
			if frequencyCount > maxFrequency {
				maxFrequency = frequencyCount
				mode = data[i]
			}
		} else {
			frequencyCount = 1
		}
		if data[i] == 47 {
			hasNumber47 = true
		}
	}
	if maxFrequency == 1 {
		return 0, hasNumber47
	}

	return mode, hasNumber47
}

func Run(data []float64) *Result {
	sort.Float64s(data) // Ensuring data is sorted once for all operations
	var wg sync.WaitGroup
	wg.Add(5)

	var (
		ave         float64
		med         float64
		p90         float64
		p99         float64
		mod         float64
		hasNumber47 bool
	)

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
		mod, hasNumber47 = calculateModeAndHasNumber47(data)
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
