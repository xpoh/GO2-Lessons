/*
	Напишите программу, которая запускает n потоков и дожидается завершения их всех
*/
package main

import (
	"fmt"
	"math"
	"sync"
)

type Mass struct {
	sm   map[int]float64
	mut  sync.RWMutex
	mut2 sync.Mutex
	tmp  float64
}

func (m *Mass) F1(wg *sync.WaitGroup, nrd int, nwr int) {
	defer wg.Done()

	m.mut.RLock()
	for i := 0; i < nrd; i++ {
		m.tmp = m.sm[i]
	}
	m.mut.RUnlock()
	m.mut.Lock()
	defer m.mut.Unlock()

	for i := 0; i < nwr; i++ {
		m.sm[999-i] = math.Sqrt(float64(999 - i))
	}
}

func (m *Mass) F2(wg *sync.WaitGroup, nrd int, nwr int) {
	defer wg.Done()

	m.mut2.Lock()
	for i := 0; i < nrd; i++ {
		m.tmp = m.sm[i]
	}
	m.mut2.Unlock()
	m.mut2.Lock()
	defer m.mut2.Unlock()

	for i := 0; i < nwr; i++ {
		m.sm[999-i] = math.Sqrt(float64(999 - i))
	}
}

func main() {
	m := Mass{}
	n := 10

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go m.F1(&wg, 100, 900)
	}

	wg.Wait()

	fmt.Println("All work done")
}
