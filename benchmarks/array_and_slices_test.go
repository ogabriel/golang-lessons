package main

import (
	"testing"
)

type Foobar struct {
	foo string
	bar string
}

func BenchmarkArrayAndSlicesDeclarationSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var foobar []Foobar
		for i := 0; i < 3; i++ {
			foobar = append(foobar, Foobar{foo: "foo", bar: "bar"})
		}
	}

	b.StopTimer()
}

func BenchmarkArrayAndSlicesDeclarationArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var foobar [10]Foobar
		for i := 0; i < 3; i++ {
			foobar[i] = Foobar{foo: "foo", bar: "bar"}
		}

		_ = foobar[:]
	}

	b.StopTimer()
}

func BenchmarkArrayAndSlicesDeclarationSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		foobar := make([]Foobar, 0, 10)
		for i := 0; i < 3; i++ {
			foobar = append(foobar, Foobar{foo: "foo", bar: "bar"})
		}
	}

	b.StopTimer()
}

func BenchmarkArrayAndSlicesDeclarationArrayPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var foobar [3]*Foobar
		for i := 0; i < 3; i++ {
			foobar[i] = &Foobar{foo: "foo", bar: "bar"}
		}
	}

	b.StopTimer()
}
