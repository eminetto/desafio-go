package desafiogo

import (
	"sync"
)

const LEN_PER_WORKER = 1300
const MAX_WORKERS = 60

func calculeMode(data []float64) float64 {
	countMap := make(map[float64]int)
	var moda float64
	var modaCount int

	for _, d := range data {
		countMap[d]++
		if countMap[d] > modaCount {
			moda = d
			modaCount = countMap[d]
		}
	}

	if modaCount == 1 {
		return 0
	}

	return moda
}

func hasNumber(data []float64, number float64) bool {
	for _, d := range data {
		if d == number {
			return true
		}
	}
	return false

}

func calculateStatistics(data []float64) *Result {
	var wg sync.WaitGroup
	wg.Add(6)

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
		mod = calculeMode(data)
		wg.Done()
	}()

	go func() {
		hasNumber47 = hasNumber(data, 47)
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

func Run(data []float64) *Result {

	var ave, med, p90, p99, mod float64
	var hasNumber47 bool

	N_WORKERS := len(data) / LEN_PER_WORKER

	if N_WORKERS == 0 {
		N_WORKERS = 1
	}

	if N_WORKERS > MAX_WORKERS {
		N_WORKERS = MAX_WORKERS
	}

	partSize := len(data) / N_WORKERS

	resultsChan := make(chan *Result, N_WORKERS)

	for i := 0; i < len(data); i += partSize {
		endIndex := i + partSize

		if endIndex > len(data) {
			endIndex = len(data)
		}

		go func(slice []float64) {
			resultsChan <- calculateStatistics(slice)
		}(data[i:endIndex])
	}

	var aveTotal, medTotal, p90Total, p99Total, modTotal float64

	for i := 0; i < N_WORKERS; i++ {
		result := <-resultsChan

		aveTotal += result.Average
		medTotal += result.Median
		p90Total += result.Percentile90
		p99Total += result.Percentile99
		modTotal += result.Mode

		if result.HasNumber47 {
			hasNumber47 = true
		}
	}

	ave = aveTotal / float64(N_WORKERS)
	med = medTotal / float64(N_WORKERS)
	p90 = p90Total / float64(N_WORKERS)
	p99 = p99Total / float64(N_WORKERS)
	mod = modTotal / float64(N_WORKERS)

	return &Result{
		Average:      ave,
		Median:       med,
		Percentile90: p90,
		Percentile99: p99,
		Mode:         mod,
		HasNumber47:  hasNumber47,
	}

}
