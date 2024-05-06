package desafiogo

import (
	"sync"
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    float64 // The value of the item; arbitrary.
	priority int     // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func Run(data []float64) *Result {
	var ave, med, p90, p99, mod float64
	var hasNumber47 bool
	var wg sync.WaitGroup
	wg.Add(6)

	// Use a priority queue to manage mode calculation efficiently
	pq := make(PriorityQueue, 0)
	itemMap := make(map[float64]*Item)
	heap.Init(&pq)

	go func() {
		defer wg.Done()
		for _, v := range data {
			if item, ok := itemMap[v]; ok {
				item.priority++
				heap.Fix(&pq, item.index)
			} else {
				item := &Item{
					value:    v,
					priority: 1,
				}
				heap.Push(&pq, item)
				itemMap[v] = item
			}
		}
		if pq.Len() > 0 {
			mod = heap.Pop(&pq).(*Item).value
		}
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

	go func() {
		defer wg.Done()
		for _, v := range data {
			if v == 47 {
				hasNumber47 = true
				break
			}
		}
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
