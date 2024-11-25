# go-simd

This repo is an ARM64 NEON architecture specific implementation of SIMD (Single Instruction, Multiple Data) operations in Go.

The SIMD instructions are written entirely with Assembly, and does not use CGO, and wrapped in a more useable "API" layer function.

## Benchmarks

To see benchmarks & tests, run go test -v -bench=. -benchmem

Here are the results from `BenchmarkInt8DotProduct` on 11/25/2024

```sh
goos: darwin
goarch: arm64
pkg: go-simd
cpu: Apple M2 Pro
BenchmarkInt8DotProduct/Scalar-16-12          172527199          6.767 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-16-12            450507735          2.729 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-100-12         33798691         35.21 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-100-12           174523876          6.981 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-1000-12         3792987        318.9 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-1000-12          21343166         56.00 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-4096-12          905314       1286 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-4096-12           5278868        229.6 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-10000-12         384264       3215 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-10000-12          2157204        564.7 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-100000-12         37902      31264 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-100000-12          181093       5595 ns/op        0 B/op        0 allocs/op
PASS
ok   go-simd 17.752s
```

Results are ~5.5x faster with SIMD for 100000 elements.
