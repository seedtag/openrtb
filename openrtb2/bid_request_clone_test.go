package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestBidRequest_Clone(t *testing.T) {
	secure := int8(1)
	width := int64(300)
	height := int64(250)

	tests := []struct {
		name string
		br   *BidRequest
	}{
		{
			name: "nil bid request",
			br:   nil,
		},
		{
			name: "empty bid request",
			br:   &BidRequest{},
		},
		{
			name: "fully populated bid request",
			br: &BidRequest{
				ID: "bid1",
				Imp: []Imp{
					{
						ID: "imp1",
						Banner: &Banner{
							W:   &width,
							H:   &height,
							Ext: json.RawMessage(`{"banner_key":"value"}`),
						},
						Secure: &secure,
						Ext:    json.RawMessage(`{"imp_key":"value"}`),
					},
				},
				Site: &Site{
					ID:   "site1",
					Name: "example.com",
					Publisher: &Publisher{
						ID:   "pub1",
						Name: "Example Publisher",
						Ext:  json.RawMessage(`{"publisher_key":"value"}`),
					},
					Ext: json.RawMessage(`{"site_key":"value"}`),
				},
				App: &App{
					ID:   "app1",
					Name: "Example App",
					Publisher: &Publisher{
						ID:   "pub2",
						Name: "Example App Publisher",
						Ext:  json.RawMessage(`{"publisher_key":"value"}`),
					},
					Ext: json.RawMessage(`{"app_key":"value"}`),
				},
				DOOH: &DOOH{
					ID:   "dooh1",
					Name: "Times Square Display",
					Ext:  json.RawMessage(`{"dooh_key":"value"}`),
				},
				Device: &Device{
					UA:    "test-user-agent",
					IP:    "192.168.1.1",
					Model: "iPhone",
					OS:    "iOS",
					OSV:   "14.0",
					Make:  "Apple",
					Ext:   json.RawMessage(`{"device_key":"value"}`),
				},
				User: &User{
					ID:     "user1",
					Yob:    1990,
					Gender: "M",
					Ext:    json.RawMessage(`{"user_key":"value"}`),
				},
				Test:    1,
				AT:      2,
				TMax:    500,
				WSeat:   []string{"seat1", "seat2"},
				BSeat:   []string{"seat3", "seat4"},
				AllImps: 1,
				Cur:     []string{"USD", "EUR"},
				WLang:   []string{"en", "es"},
				WLangB:  []string{"en-US", "es-ES"},
				ACat:    []string{"IAB1", "IAB2"},
				BCat:    []string{"IAB3", "IAB4"},
				CatTax:  adcom1.CatTaxIABContent10,
				BAdv:    []string{"adv1.com", "adv2.com"},
				BApp:    []string{"com.example.app1", "com.example.app2"},
				Source: &Source{
					FD:     &secure,
					TID:    "tid1",
					PChain: "pchain1",
					Ext:    json.RawMessage(`{"source_key":"value"}`),
				},
				Regs: &Regs{
					COPPA: 1,
					Ext:   json.RawMessage(`{"regs_key":"value"}`),
				},
				Ext: json.RawMessage(`{"request_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.br.Clone()

			if tt.br == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil BidRequest")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.br.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.br.ID)
			}
			if clone.Test != tt.br.Test {
				t.Errorf("Clone() Test = %v, want %v", clone.Test, tt.br.Test)
			}
			if clone.AT != tt.br.AT {
				t.Errorf("Clone() AT = %v, want %v", clone.AT, tt.br.AT)
			}
			if clone.TMax != tt.br.TMax {
				t.Errorf("Clone() TMax = %v, want %v", clone.TMax, tt.br.TMax)
			}
			if clone.AllImps != tt.br.AllImps {
				t.Errorf("Clone() AllImps = %v, want %v", clone.AllImps, tt.br.AllImps)
			}
			if clone.CatTax != tt.br.CatTax {
				t.Errorf("Clone() CatTax = %v, want %v", clone.CatTax, tt.br.CatTax)
			}

			// Test Imp slice
			if len(tt.br.Imp) > 0 {
				if len(clone.Imp) != len(tt.br.Imp) {
					t.Errorf("Clone() Imp length = %v, want %v", len(clone.Imp), len(tt.br.Imp))
				}
				originalID := tt.br.Imp[0].ID
				tt.br.Imp[0].ID = "modified"
				if clone.Imp[0].ID != originalID {
					t.Error("Clone() should create deep copy of Imp")
				}
			}

			// Test Site
			if tt.br.Site != nil {
				if clone.Site == tt.br.Site {
					t.Error("Clone() should create a new pointer for Site")
				}
				originalID := tt.br.Site.ID
				tt.br.Site.ID = "modified"
				if clone.Site.ID != originalID {
					t.Error("Clone() should create deep copy of Site")
				}
			}

			// Test App
			if tt.br.App != nil {
				if clone.App == tt.br.App {
					t.Error("Clone() should create a new pointer for App")
				}
				originalID := tt.br.App.ID
				tt.br.App.ID = "modified"
				if clone.App.ID != originalID {
					t.Error("Clone() should create deep copy of App")
				}
			}

			// Test DOOH
			if tt.br.DOOH != nil {
				if clone.DOOH == tt.br.DOOH {
					t.Error("Clone() should create a new pointer for DOOH")
				}
				originalID := tt.br.DOOH.ID
				tt.br.DOOH.ID = "modified"
				if clone.DOOH.ID != originalID {
					t.Error("Clone() should create deep copy of DOOH")
				}
			}

			// Test Device
			if tt.br.Device != nil {
				if clone.Device == tt.br.Device {
					t.Error("Clone() should create a new pointer for Device")
				}
				originalUA := tt.br.Device.UA
				tt.br.Device.UA = "modified"
				if clone.Device.UA != originalUA {
					t.Error("Clone() should create deep copy of Device")
				}
			}

			// Test User
			if tt.br.User != nil {
				if clone.User == tt.br.User {
					t.Error("Clone() should create a new pointer for User")
				}
				originalID := tt.br.User.ID
				tt.br.User.ID = "modified"
				if clone.User.ID != originalID {
					t.Error("Clone() should create deep copy of User")
				}
			}

			// Test string slices
			testStringSlice := func(name string, original, cloned []string) {
				if len(original) > 0 {
					if len(cloned) != len(original) {
						t.Errorf("Clone() %s length = %v, want %v", name, len(cloned), len(original))
					}
					originalValue := original[0]
					original[0] = "modified"
					if cloned[0] != originalValue {
						t.Errorf("Clone() should create deep copy of %s", name)
					}
				}
			}

			testStringSlice("WSeat", tt.br.WSeat, clone.WSeat)
			testStringSlice("BSeat", tt.br.BSeat, clone.BSeat)
			testStringSlice("Cur", tt.br.Cur, clone.Cur)
			testStringSlice("WLang", tt.br.WLang, clone.WLang)
			testStringSlice("WLangB", tt.br.WLangB, clone.WLangB)
			testStringSlice("ACat", tt.br.ACat, clone.ACat)
			testStringSlice("BCat", tt.br.BCat, clone.BCat)
			testStringSlice("BAdv", tt.br.BAdv, clone.BAdv)
			testStringSlice("BApp", tt.br.BApp, clone.BApp)

			// Test Source
			if tt.br.Source != nil {
				if clone.Source == tt.br.Source {
					t.Error("Clone() should create a new pointer for Source")
				}
				originalTID := tt.br.Source.TID
				tt.br.Source.TID = "modified"
				if clone.Source.TID != originalTID {
					t.Error("Clone() should create deep copy of Source")
				}
			}

			// Test Regs
			if tt.br.Regs != nil {
				if clone.Regs == tt.br.Regs {
					t.Error("Clone() should create a new pointer for Regs")
				}
				originalCOPPA := tt.br.Regs.COPPA
				tt.br.Regs.COPPA = 0
				if clone.Regs.COPPA != originalCOPPA {
					t.Error("Clone() should create deep copy of Regs")
				}
			}

			// Test extension
			if tt.br.Ext != nil {
				orig := string(tt.br.Ext)
				tt.br.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
