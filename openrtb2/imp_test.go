package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestImp_Clone(t *testing.T) {
	clickBrowser := int8(1)
	secure := int8(1)
	count := 2

	tests := []struct {
		name string
		i    *Imp
	}{
		{
			name: "nil imp",
			i:    nil,
		},
		{
			name: "empty imp",
			i:    &Imp{},
		},
		{
			name: "fully populated imp",
			i: &Imp{
				ID: "imp1",
				Metric: []Metric{
					{
						Type:   "viewability",
						Value:  0.85,
						Vendor: "vendor1",
						Ext:    json.RawMessage(`{"metric_key":"value"}`),
					},
				},
				Banner: &Banner{
					Format: []Format{
						{W: 300, H: 250, WRatio: 1, HRatio: 1},
					},
					Ext: json.RawMessage(`{"banner_key":"value"}`),
				},
				Video: &Video{
					MIMEs: []string{"video/mp4"},
					Ext:   json.RawMessage(`{"video_key":"value"}`),
				},
				Audio: &Audio{
					MIMEs: []string{"audio/mp3"},
					Ext:   json.RawMessage(`{"audio_key":"value"}`),
				},
				Native: &Native{
					Request: "native request string",
					Ext:     json.RawMessage(`{"native_key":"value"}`),
				},
				PMP: &PMP{
					PrivateAuction: 1,
					Deals: []Deal{
						{
							ID:          "deal1",
							BidFloor:    1.23,
							BidFloorCur: "USD",
							Ext:         json.RawMessage(`{"deal_key":"value"}`),
						},
					},
					Ext: json.RawMessage(`{"pmp_key":"value"}`),
				},
				DisplayManager:    "dm1",
				DisplayManagerVer: "1.0",
				Instl:             1,
				TagID:             "tag1",
				BidFloor:          1.23,
				BidFloorCur:       "USD",
				ClickBrowser:      &clickBrowser,
				Secure:            &secure,
				IframeBuster:      []string{"buster1", "buster2"},
				Rwdd:              1,
				Exp:               30,
				Qty: &Qty{
					Multiplier: 14.2,
					Vendor:     "example.com",
					Ext:        json.RawMessage(`{"qty_key":"value"}`),
				},
				DT: 1621234567890,
				Refresh: &Refresh{
					RefSettings: []RefSettings{
						{
							RefType: 1,
							MinInt:  30,
							Ext:     json.RawMessage(`{"refsettings_key":"value"}`),
						},
					},
					Count: &count,
					Ext:   json.RawMessage(`{"refresh_key":"value"}`),
				},
				Ext: json.RawMessage(`{"imp_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.i.Clone()

			if tt.i == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Imp")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.i.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.i.ID)
			}
			if clone.DisplayManager != tt.i.DisplayManager {
				t.Errorf("Clone() DisplayManager = %v, want %v", clone.DisplayManager, tt.i.DisplayManager)
			}
			if clone.DisplayManagerVer != tt.i.DisplayManagerVer {
				t.Errorf("Clone() DisplayManagerVer = %v, want %v", clone.DisplayManagerVer, tt.i.DisplayManagerVer)
			}
			if clone.Instl != tt.i.Instl {
				t.Errorf("Clone() Instl = %v, want %v", clone.Instl, tt.i.Instl)
			}
			if clone.TagID != tt.i.TagID {
				t.Errorf("Clone() TagID = %v, want %v", clone.TagID, tt.i.TagID)
			}
			if clone.BidFloor != tt.i.BidFloor {
				t.Errorf("Clone() BidFloor = %v, want %v", clone.BidFloor, tt.i.BidFloor)
			}
			if clone.BidFloorCur != tt.i.BidFloorCur {
				t.Errorf("Clone() BidFloorCur = %v, want %v", clone.BidFloorCur, tt.i.BidFloorCur)
			}
			if clone.Rwdd != tt.i.Rwdd {
				t.Errorf("Clone() Rwdd = %v, want %v", clone.Rwdd, tt.i.Rwdd)
			}
			if clone.Exp != tt.i.Exp {
				t.Errorf("Clone() Exp = %v, want %v", clone.Exp, tt.i.Exp)
			}
			if clone.DT != tt.i.DT {
				t.Errorf("Clone() DT = %v, want %v", clone.DT, tt.i.DT)
			}

			// Test Metric slice
			if len(tt.i.Metric) > 0 {
				if len(clone.Metric) != len(tt.i.Metric) {
					t.Errorf("Clone() Metric length = %v, want %v", len(clone.Metric), len(tt.i.Metric))
				}
				original := tt.i.Metric[0].Type
				tt.i.Metric[0].Type = "modified"
				if clone.Metric[0].Type != original {
					t.Error("Clone() should create deep copy of Metric")
				}
			}

			// Test Banner
			if tt.i.Banner != nil {
				if clone.Banner == tt.i.Banner {
					t.Error("Clone() should create a new pointer for Banner")
				}
				if len(tt.i.Banner.Format) > 0 {
					original := tt.i.Banner.Format[0].W
					tt.i.Banner.Format[0].W = 999
					if clone.Banner.Format[0].W != original {
						t.Error("Clone() should create deep copy of Banner.Format")
					}
				}
			}

			// Test Video
			if tt.i.Video != nil {
				if clone.Video == tt.i.Video {
					t.Error("Clone() should create a new pointer for Video")
				}
				if len(tt.i.Video.MIMEs) > 0 {
					original := tt.i.Video.MIMEs[0]
					tt.i.Video.MIMEs[0] = "modified"
					if clone.Video.MIMEs[0] != original {
						t.Error("Clone() should create deep copy of Video.MIMEs")
					}
				}
			}

			// Test Audio
			if tt.i.Audio != nil {
				if clone.Audio == tt.i.Audio {
					t.Error("Clone() should create a new pointer for Audio")
				}
				if len(tt.i.Audio.MIMEs) > 0 {
					original := tt.i.Audio.MIMEs[0]
					tt.i.Audio.MIMEs[0] = "modified"
					if clone.Audio.MIMEs[0] != original {
						t.Error("Clone() should create deep copy of Audio.MIMEs")
					}
				}
			}

			// Test Native
			if tt.i.Native != nil {
				if clone.Native == tt.i.Native {
					t.Error("Clone() should create a new pointer for Native")
				}
				original := tt.i.Native.Request
				tt.i.Native.Request = "modified"
				if clone.Native.Request != original {
					t.Error("Clone() should create deep copy of Native fields")
				}
			}

			// Test PMP
			if tt.i.PMP != nil {
				if clone.PMP == tt.i.PMP {
					t.Error("Clone() should create a new pointer for PMP")
				}
				if len(tt.i.PMP.Deals) > 0 {
					original := tt.i.PMP.Deals[0].ID
					tt.i.PMP.Deals[0].ID = "modified"
					if clone.PMP.Deals[0].ID != original {
						t.Error("Clone() should create deep copy of PMP.Deals")
					}
				}
			}

			// Test ClickBrowser pointer
			if tt.i.ClickBrowser != nil {
				if *clone.ClickBrowser != *tt.i.ClickBrowser {
					t.Errorf("Clone() ClickBrowser = %v, want %v", *clone.ClickBrowser, *tt.i.ClickBrowser)
				}
				oldCB := *tt.i.ClickBrowser
				*tt.i.ClickBrowser = 0
				if *clone.ClickBrowser != oldCB {
					t.Error("Clone() should create deep copy of ClickBrowser")
				}
			}

			// Test Secure pointer
			if tt.i.Secure != nil {
				if *clone.Secure != *tt.i.Secure {
					t.Errorf("Clone() Secure = %v, want %v", *clone.Secure, *tt.i.Secure)
				}
				oldSecure := *tt.i.Secure
				*tt.i.Secure = 0
				if *clone.Secure != oldSecure {
					t.Error("Clone() should create deep copy of Secure")
				}
			}

			// Test IframeBuster slice
			if len(tt.i.IframeBuster) > 0 {
				if len(clone.IframeBuster) != len(tt.i.IframeBuster) {
					t.Errorf("Clone() IframeBuster length = %v, want %v", len(clone.IframeBuster), len(tt.i.IframeBuster))
				}
				original := tt.i.IframeBuster[0]
				tt.i.IframeBuster[0] = "modified"
				if clone.IframeBuster[0] != original {
					t.Error("Clone() should create deep copy of IframeBuster")
				}
			}

			// Test Qty
			if tt.i.Qty != nil {
				if clone.Qty == tt.i.Qty {
					t.Error("Clone() should create a new pointer for Qty")
				}
				original := tt.i.Qty.Multiplier
				tt.i.Qty.Multiplier = 999.99
				if clone.Qty.Multiplier != original {
					t.Error("Clone() should create deep copy of Qty")
				}
			}

			// Test Refresh
			if tt.i.Refresh != nil {
				if clone.Refresh == tt.i.Refresh {
					t.Error("Clone() should create a new pointer for Refresh")
				}
				if tt.i.Refresh.Count != nil {
					original := *tt.i.Refresh.Count
					*tt.i.Refresh.Count = 999
					if *clone.Refresh.Count != original {
						t.Error("Clone() should create deep copy of Refresh.Count")
					}
				}
			}

			// Test extension
			if tt.i.Ext != nil {
				orig := string(tt.i.Ext)
				tt.i.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
