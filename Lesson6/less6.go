/*
	Напишите программу, которая запускает n потоков и дожидается завершения их всех
*/
package main

import (
	"fmt"
	"github.com/pkg/profile"
	"sync"
)

type Mass struct {
	mut   sync.Mutex
	count int
}

func (m *Mass) F1(wg *sync.WaitGroup) {
	defer wg.Done()

	//m.mut.Lock()
	//defer m.mut.Unlock()

	m.count++
}

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	m := Mass{}
	n := 1000

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go m.F1(&wg)
	}
	wg.Wait()

	fmt.Println("Count=", m.count)
}
