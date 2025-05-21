package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestDevice_Clone(t *testing.T) {
	var (
		lat            float64               = 40.7128
		lon            float64               = -74.0060
		dnt            int8                  = 1
		lmt            int8                  = 1
		js             int8                  = 1
		geoFetch       int8                  = 1
		mobile         int8                  = 1
		connectionType adcom1.ConnectionType = 2
	)

	tests := []struct {
		name string
		d    *Device
	}{
		{
			name: "nil device",
			d:    nil,
		},
		{
			name: "empty device",
			d:    &Device{},
		},
		{
			name: "fully populated device",
			d: &Device{
				Geo: &Geo{
					Lat:  &lat,
					Lon:  &lon,
					Type: adcom1.LocationType(1),
					Ext:  json.RawMessage(`{"key": "value"}`),
				},
				DNT: &dnt,
				Lmt: &lmt,
				UA:  "Mozilla/5.0",
				SUA: &UserAgent{
					Browsers: []BrandVersion{
						{
							Brand:   "Chrome",
							Version: []string{"91"},
							Ext:     json.RawMessage(`{"key": "value"}`),
						},
					},
					Mobile: &mobile,
					Ext:    json.RawMessage(`{"key": "value"}`),
				},
				IP:             "192.168.1.1",
				IPv6:           "2001:db8::1",
				DeviceType:     1,
				Make:           "Apple",
				Model:          "iPhone",
				OS:             "iOS",
				OSV:            "14.0",
				HWV:            "iPhone12,1",
				H:              1920,
				W:              1080,
				PPI:            458,
				PxRatio:        2.0,
				JS:             &js,
				GeoFetch:       &geoFetch,
				FlashVer:       "1.0",
				Language:       "en",
				Carrier:        "Verizon",
				MCCMNC:         "310-012",
				ConnectionType: &connectionType,
				IFA:            "AAAAAA-BBBB-CCCC-1111",
				Ext:            json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.d.Clone()

			if tt.d == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Device")
				}
				return
			}

			// Test pointer fields
			if tt.d.DNT != nil {
				orig := *tt.d.DNT
				*tt.d.DNT = 0
				if *clone.DNT != orig {
					t.Error("Clone() should create deep copy of DNT")
				}
			}

			// Test Geo
			if tt.d.Geo != nil {
				origLat := *tt.d.Geo.Lat
				*tt.d.Geo.Lat = 999.999
				if *clone.Geo.Lat != origLat {
					t.Error("Clone() should create deep copy of Geo.Lat")
				}

				// Test Geo's extension
				if tt.d.Geo.Ext != nil {
					origGeoExt := string(tt.d.Geo.Ext)
					tt.d.Geo.Ext[0] = 'X'
					clonedGeoExt := string(clone.Geo.Ext)
					if clonedGeoExt != origGeoExt {
						t.Error("Clone() should create deep copy of Geo.Ext")
					}
				}
			}

			// Test UserAgent
			if tt.d.SUA != nil && len(tt.d.SUA.Browsers) > 0 {
				origBrand := tt.d.SUA.Browsers[0].Brand
				tt.d.SUA.Browsers[0].Brand = "modified"
				if clone.SUA.Browsers[0].Brand != origBrand {
					t.Error("Clone() should create deep copy of SUA.Browsers")
				}

				if tt.d.SUA.Mobile != nil {
					origMobile := *tt.d.SUA.Mobile
					*tt.d.SUA.Mobile = 0
					if *clone.SUA.Mobile != origMobile {
						t.Error("Clone() should create deep copy of SUA.Mobile")
					}
				}
			}

			// Test ConnectionType
			if tt.d.ConnectionType != nil {
				origConnType := *tt.d.ConnectionType
				*tt.d.ConnectionType = adcom1.ConnectionType(3)
				if *clone.ConnectionType != origConnType {
					t.Error("Clone() should create deep copy of ConnectionType")
				}
			}

			// Test extension
			if tt.d.Ext != nil {
				orig := string(tt.d.Ext)
				tt.d.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
