package openrtb

import (
	"os"
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
