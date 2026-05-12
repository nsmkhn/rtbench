package openrtb

import (
	"os"
	"sync"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkParse_StdlibJSON(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	for b.Loop() {
		_, err := Parse(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkParse_JsonIterator(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	for b.Loop() {
		var br BidRequest
		if err := jsonIterator.Unmarshal(data, &br); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParse_GoJson(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	for b.Loop() {
		var br BidRequest
		if err := gojson.Unmarshal(data, &br); err != nil {
			b.Fatal(err)
		}
	}
}

var bidRequestPool = sync.Pool{
	New: func() any { return new(BidRequest) },
}

func BenchmarkParse_GoJsonPool(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	for b.Loop() {
		br := bidRequestPool.Get().(*BidRequest)
		*br = BidRequest{}
		if err := gojson.Unmarshal(data, br); err != nil {
			b.Fatal(err)
		}
		bidRequestPool.Put(br)
	}
}

func BenchmarkParse_GoJsonParallel(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var br BidRequest
			if err := gojson.Unmarshal(data, &br); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkParse_GoJsonPoolParallel(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			br := bidRequestPool.Get().(*BidRequest)
			*br = BidRequest{}
			if err := gojson.Unmarshal(data, br); err != nil {
				b.Fatal(err)
			}
			bidRequestPool.Put(br)
		}
	})
}

func BenchmarkParse_HandWritten(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	for b.Loop() {
		_, err := ParseFast(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParse_HandWrittenParallel(b *testing.B) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		b.Fatalf("could not read testdata: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := ParseFast(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
