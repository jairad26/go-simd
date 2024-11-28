# go-simd

This repo is an ARM64 NEON architecture specific implementation of SIMD (Single Instruction, Multiple Data) operations in Go.

The SIMD instructions are written entirely with Assembly, and does not use CGO, and wrapped in a more useable "API" layer.

## Benchmarks

To see benchmarks & tests, run go test -v -bench=. -benchmem

Here are the results from `BenchmarkInt8DotProduct` on 11/28/2024

```sh
goos: darwin
goarch: arm64
pkg: go-simd
cpu: Apple M2 Pro
BenchmarkInt8DotProduct/Scalar-16-12          175903202          6.641 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-16-12            506505428          2.384 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-100-12         33686653         35.22 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-100-12           78084122         15.39 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-1000-12         3858519        309.3 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-1000-12          43131205         28.08 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-4096-12          940941       1268 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-4096-12          20645560         58.13 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-10000-12         387094       3101 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-10000-12          8187171        146.4 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/Scalar-100000-12         34800      30903 ns/op        0 B/op        0 allocs/op
BenchmarkInt8DotProduct/SIMD-100000-12          542487       2214 ns/op        0 B/op        0 allocs/op
PASS
ok   go-simd 16.531s
```

Results are ~14x faster with SIMD for 100000 elements.
