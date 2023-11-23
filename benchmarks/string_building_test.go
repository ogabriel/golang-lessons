package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringBuildingStringBuilder(b *testing.B) {
	var builder strings.Builder

	for n := 0; n < b.N; n++ {
		builder.WriteString("x")
	}

	_ = builder.String()

	b.StopTimer()
}

func BenchmarkStringBuildingStringBuilderGrow(b *testing.B) {
	var builder strings.Builder
	builder.Grow(b.N)

	for n := 0; n < b.N; n++ {
		builder.WriteString("x")
	}

	_ = builder.String()

	b.StopTimer()
}

// func BenchmarkStringBuildingInterpolation(b *testing.B) {
// 	var str string
//
// 	for n := 0; n < b.N; n++ {
// 		str = fmt.Sprintf("%s%s", str, "x")
// 	}
//
// 	b.StopTimer()
// }

// func BenchmarkStringBuildingConcat(b *testing.B) {
// 	var str string
//
// 	for n := 0; n < b.N; n++ {
// 		str += "x"
// 	}
//
// 	b.StopTimer()
// }

func BenchmarkStringBuildingBuffer(b *testing.B) {
	var buffer bytes.Buffer

	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}

	b.StopTimer()
}
