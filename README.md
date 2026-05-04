# rtbench [![CI](https://github.com/nsmkhn/rtbench/actions/workflows/ci.yml/badge.svg)](https://github.com/nsmkhn/rtbench/actions/workflows/ci.yml)

OpenRTB 2.6 bid request parsing benchmarks in Go — profiling `encoding/json` with pprof and comparing three JSON implementations.

Companion source for the blog post: **[Go JSON Performance in Adtech: Profiling OpenRTB Bid Request Parsing](https://tokarevxvi.dev/blog/go-json-performance-openrtb/)**

## Run

**Benchmark suite:**

```bash
go test -bench=. -benchmem -count=5 ./openrtb/
```

**CPU profile (stdlib):**

```bash
go run ./cmd/profile
go tool pprof -top cpu.prof
```

**CPU profile (goccy/go-json):**

```bash
go run ./cmd/profile -impl gojson
go tool pprof -top cpu.prof
```

## Structure

```
rtbench/
├── openrtb/
│   ├── types.go              # OpenRTB 2.6 struct definitions
│   ├── parse.go              # Parse() — encoding/json wrapper
│   ├── parse_test.go         # correctness tests (table-driven)
│   └── parse_bench_test.go   # benchmark suite
├── cmd/
│   └── profile/
│       └── main.go           # pprof CPU profiler (200k iterations)
└── testdata/
    ├── valid_banner.json
    └── valid_video.json
```
