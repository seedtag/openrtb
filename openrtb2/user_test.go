package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestUser_Clone(t *testing.T) {
	var (
		yob int64   = 1990
		lat float64 = 42.3601
		lon float64 = -71.0589
	)

	tests := []struct {
		name string
		u    *User
	}{
		{
			name: "nil user",
			u:    nil,
		},
		{
			name: "empty user",
			u:    &User{},
		},
		{
			name: "fully populated user",
			u: &User{
				ID:       "user1",
				BuyerUID: "buyer123",
				Yob:      yob,
				Gender:   "M",
				Keywords: "sports,news",
				Geo: &Geo{
					Lat:     &lat,
					Lon:     &lon,
					Type:    1,
					City:    "Boston",
					Region:  "MA",
					Country: "USA",
					Ext:     json.RawMessage(`{"geo_key":"value"}`),
				},
				EIDs: []EID{
					{
						Source: "example.com",
						UIDs: []UID{
							{
								ID:    "uid123",
								AType: 1,
								Ext:   json.RawMessage(`{"uid_key":"value"}`),
							},
						},
						Ext: json.RawMessage(`{"eid_key":"value"}`),
					},
				},
				Data: []Data{
					{
						ID:   "data1",
						Name: "Example Data",
						Ext:  json.RawMessage(`{"data_key":"value"}`),
					},
				},
				Consent: "GDPR-CONSENT",
				Ext:     json.RawMessage(`{"user_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.u.Clone()

			if tt.u == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil User")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.u.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.u.ID)
			}
			if clone.BuyerUID != tt.u.BuyerUID {
				t.Errorf("Clone() BuyerUID = %v, want %v", clone.BuyerUID, tt.u.BuyerUID)
			}
			if clone.Yob != tt.u.Yob {
				t.Errorf("Clone() Yob = %v, want %v", clone.Yob, tt.u.Yob)
			}
			if clone.Gender != tt.u.Gender {
				t.Errorf("Clone() Gender = %v, want %v", clone.Gender, tt.u.Gender)
			}
			if clone.Keywords != tt.u.Keywords {
				t.Errorf("Clone() Keywords = %v, want %v", clone.Keywords, tt.u.Keywords)
			}
			if clone.Consent != tt.u.Consent {
				t.Errorf("Clone() Consent = %v, want %v", clone.Consent, tt.u.Consent)
			}

			// Test Geo
			if tt.u.Geo != nil {
				if clone.Geo == tt.u.Geo {
					t.Error("Clone() should create a new pointer for Geo")
				}
				originalLat := *tt.u.Geo.Lat
				newLat := float64(0)
				tt.u.Geo.Lat = &newLat
				if *clone.Geo.Lat != originalLat {
					t.Error("Clone() should create deep copy of Geo")
				}
			}

			// Test EIDs slice
			if len(tt.u.EIDs) > 0 {
				if len(clone.EIDs) != len(tt.u.EIDs) {
					t.Errorf("Clone() EIDs length = %v, want %v", len(clone.EIDs), len(tt.u.EIDs))
				}
				originalSource := tt.u.EIDs[0].Source
				tt.u.EIDs[0].Source = "modified"
				if clone.EIDs[0].Source != originalSource {
					t.Error("Clone() should create deep copy of EIDs")
				}
			}

			// Test Data slice
			if len(tt.u.Data) > 0 {
				if len(clone.Data) != len(tt.u.Data) {
					t.Errorf("Clone() Data length = %v, want %v", len(clone.Data), len(tt.u.Data))
				}
				originalID := tt.u.Data[0].ID
				tt.u.Data[0].ID = "modified"
				if clone.Data[0].ID != originalID {
					t.Error("Clone() should create deep copy of Data")
				}
			}

			// Test extension
			if tt.u.Ext != nil {
				orig := string(tt.u.Ext)
				tt.u.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
