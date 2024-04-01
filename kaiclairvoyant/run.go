package desafiogo

import "sync"

func Run(data []float64) *Result {
	var ave float64
	var med float64
	var p90 float64
	var p99 float64
	var mod float64
	var hasNumber47 bool
	var mostAppearances int
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		repetitions := make(map[float64]int)
		for _, v := range data {
			repetitions[v]++
			if repetitions[v] > mostAppearances {
				mod = v
				mostAppearances = repetitions[v]
			}
		}
		if mostAppearances < 2 {
			mod = 0
		}
		if repetitions[47] != 0 {
			hasNumber47 = true
		}
		wg.Done()
	}()
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
