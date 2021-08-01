package main

import (
	"math"
	"sync"
	"testing"
)

func BenchmarkF1_10_RW(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F1(&wg, 100, 900)
	}
}

func BenchmarkF1_50_RW(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F1(&wg, 500, 500)
	}
}

func BenchmarkF1_90_RW(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F1(&wg, 900, 100)
	}
}
func BenchmarkF1_10(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F2(&wg, 100, 900)
	}
}

func BenchmarkF1_50(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F2(&wg, 500, 500)
	}
}

func BenchmarkF1_90(b *testing.B) {
	m := Mass{}

	// Make test data
	m.sm = make(map[int]float64)
	for i := 0; i < 1000; i++ {
		m.sm[i] = math.Sqrt(float64(i))
	}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go m.F2(&wg, 900, 100)
	}
}
