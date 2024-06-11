## Notes

1. the method must start with "Benchmark", so like BenchmarkSomething
2. use the b.T to represent time on a for
3. use `go test -bench=Something` to run only a set of tests
4. use `go test -bench=Something -benchmem ` to also measure memory
