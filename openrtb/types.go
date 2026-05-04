package openrtb

import "encoding/json"

// BidRequest is the top-level object sent by an exchange to a bidder.
// Spec: OpenRTB 2.6 — 3.2.1
type BidRequest struct {
	ID     string          `json:"id"`
	Imp    []Imp           `json:"imp"`
	Site   *Site           `json:"site,omitempty"`
	App    *App            `json:"app,omitempty"`
	User   *User           `json:"user,omitempty"`
	Device *Device         `json:"device,omitempty"`
	AT     *int            `json:"at,omitempty"`
	TMax   *int            `json:"tmax,omitempty"`
	BCat   []string        `json:"bcat,omitempty"`
	BAdv   []string        `json:"badv,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// Imp represents one ad opportunity within a bid request.
// Spec: OpenRTB 2.6 — 3.2.4
type Imp struct {
	ID          string          `json:"id"`
	Banner      *Banner         `json:"banner,omitempty"`
	Video       *Video          `json:"video,omitempty"`
	Native      *Native         `json:"native,omitempty"`
	BidFloor    *float64        `json:"bidfloor,omitempty"`
	BidFloorCur string          `json:"bidfloorcur,omitempty"`
	Secure      *int            `json:"secure,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`
}

// Banner represents a banner ad slot within an impression.
// Spec: OpenRTB 2.6 — 3.2.6
type Banner struct {
	Format []Format        `json:"format,omitempty"`
	W      *int            `json:"w,omitempty"`
	H      *int            `json:"h,omitempty"`
	MIMEs  []string        `json:"mimes,omitempty"`
	Pos    *int            `json:"pos,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// Format defines a width/height pair for banner ads.
type Format struct {
	W int `json:"w"`
	H int `json:"h"`
}

// Video represents a video ad opportunity.
// Spec: OpenRTB 2.6 — 3.2.7
type Video struct {
	MIMEs          []string        `json:"mimes"`
	MinDuration    *int            `json:"minduration,omitempty"`
	MaxDuration    *int            `json:"maxduration,omitempty"`
	Protocols      []int           `json:"protocols,omitempty"`
	W              *int            `json:"w,omitempty"`
	H              *int            `json:"h,omitempty"`
	Linearity      *int            `json:"linearity,omitempty"`
	Skip           *int            `json:"skip,omitempty"`
	PlaybackMethod []int           `json:"playbackmethod,omitempty"`
	Ext            json.RawMessage `json:"ext,omitempty"`
}

// Native represents a native ad opportunity.
type Native struct {
	Request string          `json:"request"`
	Ver     *string         `json:"ver,omitempty"`
	Ext     json.RawMessage `json:"ext,omitempty"`
}

// Site describes the publisher's website.
// Spec: OpenRTB 2.6 — 3.2.13
type Site struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	Page      string          `json:"page,omitempty"`
	Publisher *Publisher      `json:"publisher,omitempty"`
	Ext       json.RawMessage `json:"ext,omitempty"`
}

// App describes a mobile application.
// Spec: OpenRTB 2.6 — 3.2.14
type App struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Bundle    string          `json:"bundle,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	Publisher *Publisher      `json:"publisher,omitempty"`
	Ext       json.RawMessage `json:"ext,omitempty"`
}

// Publisher identifies the seller of the inventory.
// Spec: OpenRTB 2.6 — 3.2.15
type Publisher struct {
	ID     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Domain string          `json:"domain,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// Device describes the end user's device.
// Spec: OpenRTB 2.6 — 3.2.18
type Device struct {
	UA         string          `json:"ua,omitempty"`
	IP         string          `json:"ip,omitempty"`
	DeviceType *int            `json:"devicetype,omitempty"`
	Make       string          `json:"make,omitempty"`
	Model      string          `json:"model,omitempty"`
	OS         string          `json:"os,omitempty"`
	OSV        string          `json:"osv,omitempty"`
	Language   string          `json:"language,omitempty"`
	Ext        json.RawMessage `json:"ext,omitempty"`
}

// User contains data about the human user of the device.
// Spec: OpenRTB 2.6 — 3.2.20
type User struct {
	ID       string          `json:"id,omitempty"`
	BuyerUID string          `json:"buyeruid,omitempty"`
	Yob      *int            `json:"yob,omitempty"`
	Gender   string          `json:"gender,omitempty"`
	Ext      json.RawMessage `json:"ext,omitempty"`
}
