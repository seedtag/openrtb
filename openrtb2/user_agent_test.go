package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestUserAgent_Clone(t *testing.T) {
	mobile := int8(1)

	tests := []struct {
		name string
		ua   *UserAgent
	}{
		{
			name: "nil user agent",
			ua:   nil,
		},
		{
			name: "empty user agent",
			ua:   &UserAgent{},
		},
		{
			name: "fully populated user agent",
			ua: &UserAgent{
				Browsers: []BrandVersion{
					{
						Brand:   "Chrome",
						Version: []string{"91", "0", "4472"},
						Ext:     json.RawMessage(`{"browser_key":"value"}`),
					},
					{
						Brand:   "Firefox",
						Version: []string{"89", "0"},
						Ext:     json.RawMessage(`{"browser_key":"value2"}`),
					},
				},
				Platform: &BrandVersion{
					Brand:   "Windows",
					Version: []string{"10", "0"},
					Ext:     json.RawMessage(`{"platform_key":"value"}`),
				},
				Mobile:       &mobile,
				Architecture: "x86",
				Bitness:      "64",
				Model:        "PC",
				Source:       adcom1.UASourceHighEntropy,
				Ext:          json.RawMessage(`{"ua_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.ua.Clone()

			if tt.ua == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil UserAgent")
				}
				return
			}

			// Test primitive fields
			if clone.Architecture != tt.ua.Architecture {
				t.Errorf("Clone() Architecture = %v, want %v", clone.Architecture, tt.ua.Architecture)
			}
			if clone.Bitness != tt.ua.Bitness {
				t.Errorf("Clone() Bitness = %v, want %v", clone.Bitness, tt.ua.Bitness)
			}
			if clone.Model != tt.ua.Model {
				t.Errorf("Clone() Model = %v, want %v", clone.Model, tt.ua.Model)
			}
			if clone.Source != tt.ua.Source {
				t.Errorf("Clone() Source = %v, want %v", clone.Source, tt.ua.Source)
			}

			// Test Browsers slice
			if len(tt.ua.Browsers) > 0 {
				if len(clone.Browsers) != len(tt.ua.Browsers) {
					t.Errorf("Clone() Browsers length = %v, want %v", len(clone.Browsers), len(tt.ua.Browsers))
				}
				originalBrand := tt.ua.Browsers[0].Brand
				tt.ua.Browsers[0].Brand = "modified"
				if clone.Browsers[0].Brand != originalBrand {
					t.Error("Clone() should create deep copy of Browsers")
				}

				// Test Browsers[0].Version slice
				if len(tt.ua.Browsers[0].Version) > 0 {
					originalVersion := tt.ua.Browsers[0].Version[0]
					tt.ua.Browsers[0].Version[0] = "modified"
					if clone.Browsers[0].Version[0] != originalVersion {
						t.Error("Clone() should create deep copy of Browser Version")
					}
				}
			}

			// Test Platform
			if tt.ua.Platform != nil {
				if clone.Platform == tt.ua.Platform {
					t.Error("Clone() should create a new pointer for Platform")
				}
				originalBrand := tt.ua.Platform.Brand
				tt.ua.Platform.Brand = "modified"
				if clone.Platform.Brand != originalBrand {
					t.Error("Clone() should create deep copy of Platform")
				}

				// Test Platform Version slice
				if len(tt.ua.Platform.Version) > 0 {
					originalVersion := tt.ua.Platform.Version[0]
					tt.ua.Platform.Version[0] = "modified"
					if clone.Platform.Version[0] != originalVersion {
						t.Error("Clone() should create deep copy of Platform Version")
					}
				}
			}

			// Test Mobile pointer
			if tt.ua.Mobile != nil {
				if *clone.Mobile != *tt.ua.Mobile {
					t.Errorf("Clone() Mobile = %v, want %v", *clone.Mobile, *tt.ua.Mobile)
				}
				oldMobile := *tt.ua.Mobile
				*tt.ua.Mobile = 0
				if *clone.Mobile != oldMobile {
					t.Error("Clone() should create deep copy of Mobile")
				}
			}

			// Test extension
			if tt.ua.Ext != nil {
				orig := string(tt.ua.Ext)
				tt.ua.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
