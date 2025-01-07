package main

import (
	"testing"
	"time"
)

var format string = "2006-01-02T15:04:05.999999Z"

func BenchmarkStringAllocationGlobal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now().Format(format)
	}

	b.StopTimer()
}

func BenchmarkStringAllocationLocal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now().Format("2006-01-02T15:04:05.999999Z")
	}

	b.StopTimer()
}
