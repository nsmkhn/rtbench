package openrtb

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Parse decodes raw JSON into a BidRequest.
// It returns an error only if JSON is malformed — not if required fields are missing. Use Validate
// for spec compliance checks.
func Parse(data []byte) (*BidRequest, error) {
	var br BidRequest
	if err := json.Unmarshal(data, &br); err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return &br, nil
}

// Validate checks a parsed BidRequest against OpenRTB 2.6 required fields rules.
// Returns all violations found, not just the first
func Validate(br *BidRequest) []error {
	var errs []error

	if br.ID == "" {
		errs = append(errs, errors.New("BidRequest.id is required"))
	}

	if len(br.Imp) == 0 {
		errs = append(errs, errors.New("BidRequest.imp must contain at least one impression"))
	}

	for i, imp := range br.Imp {
		if imp.ID == "" {
			errs = append(errs, fmt.Errorf("Imp[%d].id is required", i))
		}

		if imp.Banner == nil && imp.Video == nil && imp.Native == nil {
			errs = append(errs, fmt.Errorf("Imp[%d] must have at least one of: banner, video, native", i))
		}

		if imp.Video != nil && len(imp.Video.MIMEs) == 0 {
			errs = append(errs, fmt.Errorf("Imp[%d].video.mimes is required", i))
		}
	}

	return errs
}
