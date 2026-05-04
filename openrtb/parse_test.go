package openrtb

import (
	"os"
	"testing"
)

func readTestData(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := os.ReadFile("../testdata/" + filename)
	if err != nil {
		t.Fatalf("could not read testdata/%s: %v", filename, err)
	}
	return data
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{"valid banner request", "valid_banner.json", false},
		{"valid video request", "valid_video.json", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readTestData(t, tt.file)
			_, err := Parse(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestParse_MalformedJSON(t *testing.T) {
	_, err := Parse([]byte(`{"id": "bad}`))
	if err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name           string
		file           string
		wantErrorCount int
	}{
		{"valid banner", "valid_banner.json", 0},
		{"valid video", "valid_video.json", 0},
		{"missing imp", "invalid_no_imp.json", 1},
		{"video missing mimes", "invalid_video_no_mimes.json", 1},
		{"imp missing format", "invalid_imp_no_format.json", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readTestData(t, tt.file)
			br, err := Parse(data)
			if err != nil {
				t.Fatalf("Parse() failed: %v", err)
			}
			errs := Validate(br)
			if len(errs) != tt.wantErrorCount {
				t.Errorf("Validate() returned %d errors, want %d: %v", len(errs), tt.wantErrorCount, errs)
			}
		})
	}
}
