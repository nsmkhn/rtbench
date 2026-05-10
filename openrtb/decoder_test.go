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
