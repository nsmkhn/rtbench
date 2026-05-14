# rtbench [![CI](https://github.com/nsmkhn/rtbench/actions/workflows/ci.yml/badge.svg)](https://github.com/nsmkhn/rtbench/actions/workflows/ci.yml)

OpenRTB 2.6 bid request parsing benchmarks in Go — comparing JSON library performance, profiling with pprof, and building a hand-written decoder with arm64 NEON and arena allocation.

Companion source for the blog series:
- **[Go JSON Performance in AdTech: Profiling OpenRTB Bid Request Parsing](https://tokarevxvi.dev/blog/go-json-performance-openrtb/)** — profiling `encoding/json`, comparing `json-iterator` and `goccy/go-json`
- **[Hand-Writing an OpenRTB JSON Decoder in Go](https://tokarevxvi.dev/blog/go-json-handwritten-openrtb/)** — lexer, `parseIntBytes`, NEON string scanning, arena allocation

## Results

Apple M2 Pro, Go 1.26.2, `-bench=. -benchmem -count=5`, 804-byte banner bid request:

| Implementation         | ns/op | B/op  | allocs/op |
| ---------------------- | ----- | ----- | --------- |
| `encoding/json`        | 7,685 | 1,840 | 49        |
| `json-iterator`        | 2,185 | 1,456 | 48        |
| `goccy/go-json`        | 1,590 | 1,963 | 23        |
| `goccy/go-json` + pool | 1,514 | 1,804 | 22        |
| `ParseFast`            | 1,360 | 968   | 19        |
| `ParseFastArena`       | 1,245 | 392   | 14        |

Parallel throughput (10 goroutines):

| Implementation         | ns/op |
| ---------------------- | ----- |
| `goccy/go-json` + pool | 876   |
| `ParseFastArena`       | 348   |

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
│   ├── types.go               # OpenRTB 2.6 struct definitions
│   ├── parse.go               # Parse(), ParseFast(), ParseFastArena(), ReleaseArena()
│   ├── lexer.go               # JSON lexer (zero-copy string scanning)
│   ├── decoder.go             # Hand-written field decoder
│   ├── arena.go               # Arena struct and pool management
│   ├── indexstop_arm64.go     # indexStopByte stub (arm64)
│   ├── indexstop_arm64.s      # NEON implementation
│   ├── indexstop_generic.go   # Scalar fallback (!arm64)
│   ├── parse_test.go          # Correctness tests
│   ├── parse_bench_test.go    # Benchmark suite
│   ├── decoder_test.go        # Decoder unit tests
│   └── lexer_test.go          # Lexer unit tests
├── cmd/
│   └── profile/
│       └── main.go            # pprof CPU profiler (200k iterations)
└── testdata/
    ├── valid_banner.json
    └── valid_video.json
```
