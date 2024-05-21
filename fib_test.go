package main

import (
	"fmt"
	"testing"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		fib(40)
		fib(40)
		elapsed := time.Since(start)
		fmt.Printf("Sequential: %s\n", elapsed)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		done := make(chan bool, 2)
		go func() {
			fib(40)
			done <- true
		}()
		go func() {
			fib(40)
			done <- true
		}()
		<-done
		<-done
		elapsed := time.Since(start)
		fmt.Printf("Concurrent: %s\n", elapsed)
	}
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		done := make(chan bool, 2)
		go func() {
			fib(40)
			done <- true
		}()
		go func() {
			fib(40)
			done <- true
		}()
		<-done
		<-done
		elapsed := time.Since(start)
		fmt.Printf("Parallel: %s\n", elapsed)
	}
}

func main() {
	fmt.Println("Running benchmarks...")
}
