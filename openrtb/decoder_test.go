package openrtb

import "testing"

func TestDecodeBidRequest(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(*testing.T, *BidRequest)
	}{
		{
			name:  "id only",
			input: `{"id":"req-001"}`,
			check: func(t *testing.T, br *BidRequest) {
				if br.ID != "req-001" {
					t.Errorf(`invalid br.ID: expected "req-001", got %v`, br.ID)
				}
				if br.AT != nil {
					t.Fatalf(`invalid br.AT, expected 'nil', got %v`, br.AT)
				}
				if br.TMax != nil {
					t.Fatalf(`invalid br.TMax, expected 'nil', got %v`, br.TMax)
				}
			},
		},
		{
			name:  "id and at",
			input: `{"id":"req-001", "at":1}`,
			check: func(t *testing.T, br *BidRequest) {
				if br.ID != "req-001" {
					t.Fatalf(`invalid br.ID, expected "req-001", got %v`, br.ID)
				}
				if br.AT == nil {
					t.Fatalf(`expected br.AT to be set, got nil`)
				}
				if *br.AT != 1 {
					t.Fatalf(`invalid br.AT, expected '1', got %v`, *br.AT)
				}
				if br.TMax != nil {
					t.Fatalf(`invalid br.TMax, expected 'nil', got %v`, br.TMax)
				}
			},
		},
		{
			name:  "bcat and badv",
			input: `{"id":"req-001","bcat":["IAB1","IAB2"],"badv":["example.com"]}`,
			check: func(t *testing.T, br *BidRequest) {
				if len(br.BAdv) != 1 {
					t.Fatalf("invalid len(br.BAdv): expected 1, got %v", len(br.BAdv))
				}
				if len(br.BCat) != 2 {
					t.Fatalf("invalid len(br.BCat): expected 2, got %v", len(br.BCat))
				}

				if br.BCat[0] != "IAB1" {
					t.Fatalf(`invalid br.BCat[0]: expected "IAB1", got %v`, br.BCat[0])
				}
				if br.BCat[1] != "IAB2" {
					t.Fatalf(`invalid br.BCat[1]: expected "IAB2", got %v`, br.BCat[1])
				}
				if br.BAdv[0] != "example.com" {
					t.Fatalf(`invalid br.BAdv[0]: expected "example.com", got %v`, br.BAdv[0])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLexer([]byte(tt.input))
			br, err := decodeBidRequest(l)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.check != nil {
				tt.check(t, br)
			}
		})
	}
}
