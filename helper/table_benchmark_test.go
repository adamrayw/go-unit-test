package helper

import "testing"

type DataBench struct {
	name    string
	request string
}

// It runs a benchmark for each test in the tests slice
func BenchmarkTableBenchmark(b *testing.B) {
	tests := []DataBench{
		{
			name:    "HelloWorld(Adam)",
			request: "Adam",
		},
		{
			name:    "HelloWorld(Ray)",
			request: "Ray",
		},
		{
			name:    "HelloWorld(Wibowo)",
			request: "Wibowo",
		},
	}

	for _, benchmark := range tests {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}
