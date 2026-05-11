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
		{
			name:  "site with publisher",
			input: `{"id":"req-001","site":{"page":"https://example.com","publisher":{"id":"pub-123"}}}`,
			check: func(t *testing.T, br *BidRequest) {
				if br.Site == nil {
					t.Fatalf("expected br.Site to be set, got nil")
				}
				if br.Site.Page != "https://example.com" {
					t.Fatalf(`invalid br.Site.Page: expected "https://example.com", got %v`, br.Site.Page)
				}
				if br.Site.Publisher == nil {
					t.Fatalf("expected br.Site.Publisher to be set, got nil")
				}
				if br.Site.Publisher.ID != "pub-123" {
					t.Fatalf(`invalid br.Site.Publisher.ID: expected "pub-123", got %v`, br.Site.Publisher.ID)
				}
			},
		},
		{
			name:  "device and user",
			input: `{"id":"req-001","device":{"ua":"Mozilla/5.0","ip":"192.168.1.1","os":"Windows","devicetype":1},"user":{"id":"user-abc","buyeruid":"buyer-xyz","yob":1990}}`,
			check: func(t *testing.T, br *BidRequest) {
				if br.Device == nil {
					t.Fatalf("expected br.Device to be set, got nil")
				}
				if br.Device.UA != "Mozilla/5.0" {
					t.Fatalf(`invalid br.Device.UA: expected "Mozilla/5.0", got %v`, br.Device.UA)
				}
				if br.Device.OS != "Windows" {
					t.Fatalf(`invalid br.Device.OS: expected "Windows", got %v`, br.Device.OS)
				}
				if br.Device.DeviceType == nil {
					t.Fatalf("expected br.Device.DeviceType to be set, got nil")
				}
				if *br.Device.DeviceType != 1 {
					t.Fatalf("invalid br.Device.DeviceType: expected 1, got %v", *br.Device.DeviceType)
				}
				if br.User == nil {
					t.Fatalf("expected br.User to be set, got nil")
				}
				if br.User.ID != "user-abc" {
					t.Fatalf(`invalid br.User.ID: expected "user-abc", got %v`, br.User.ID)
				}
				if br.User.Yob == nil {
					t.Fatalf("expected br.User.Yob to be set, got nil")
				}
				if *br.User.Yob != 1990 {
					t.Fatalf("invalid br.User.Yob: expected 1990, got %v", *br.User.Yob)
				}
			},
		},
		{
			name:  "imp with banner and bidfloor",
			input: `{"id":"req-001","imp":[{"id":"imp-001","banner":{"format":[{"w":300,"h":250},{"w":728,"h":90}]},"bidfloor":0.5,"bidfloorcur":"USD"}]}`,
			check: func(t *testing.T, br *BidRequest) {
				if len(br.Imp) != 1 {
					t.Fatalf("expected 1 imp, got %d", len(br.Imp))
				}
				if br.Imp[0].ID != "imp-001" {
					t.Fatalf(`invalid Imp[0].ID: expected "imp-001", got %v`, br.Imp[0].ID)
				}
				if br.Imp[0].Banner == nil {
					t.Fatalf("expected Imp[0].Banner to be set, got nil")
				}
				if len(br.Imp[0].Banner.Format) != 2 {
					t.Fatalf("expected 2 formats, got %d", len(br.Imp[0].Banner.Format))
				}
				if br.Imp[0].Banner.Format[0].W != 300 || br.Imp[0].Banner.Format[0].H != 250 {
					t.Fatalf("invalid Format[0]: expected 300x250, got %dx%d", br.Imp[0].Banner.Format[0].W, br.Imp[0].Banner.Format[0].H)
				}
				if br.Imp[0].BidFloor == nil {
					t.Fatalf("expected Imp[0].BidFloor to be set, got nil")
				}
				if *br.Imp[0].BidFloor != 0.5 {
					t.Fatalf("invalid Imp[0].BidFloor: expected 0.5, got %v", *br.Imp[0].BidFloor)
				}
				if br.Imp[0].BidFloorCur != "USD" {
					t.Fatalf(`invalid Imp[0].BidFloorCur: expected "USD", got %v`, br.Imp[0].BidFloorCur)
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
