package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	gojson "github.com/goccy/go-json"
	"github.com/nsmkhn/rtbench/openrtb"
)

const iterations = 200_000

func main() {
	impl := flag.String("impl", "stdlib", "json implementation: stdlib, gojson, or handwritten")
	flag.Parse()

	data, err := os.ReadFile("testdata/valid_banner.json")
	if err != nil {
		log.Fatalf("could not read testdata: %v", err)
	}

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("could not create cpu.prof: %v", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("could not start CPU profile: %v", err)
	}
	defer pprof.StopCPUProfile()

	for range iterations {
		switch *impl {
		case "gojson":
			var br openrtb.BidRequest
			if err := gojson.Unmarshal(data, &br); err != nil {
				log.Fatalf("gojson parse error: %v", err)
			}
		case "handwritten":
			_, err := openrtb.ParseFast(data)
			if err != nil {
				log.Fatalf("parse error: %v", err)
			}
		default:
			_, err := openrtb.Parse(data)
			if err != nil {
				log.Fatalf("parse error: %v", err)
			}
		}
	}

	fmt.Printf("profiled %d iterations, wrote cpu.prof\n", iterations)
}
