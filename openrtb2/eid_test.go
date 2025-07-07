package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestEID_Clone(t *testing.T) {
	tests := []struct {
		name string
		e    *EID
	}{
		{
			name: "nil eid",
			e:    nil,
		},
		{
			name: "empty eid",
			e:    &EID{},
		},
		{
			name: "fully populated eid",
			e: &EID{
				Source: "example.com",
				UIDs: []UID{
					{
						ID:    "uid1",
						AType: adcom1.AgentTypeWeb,
						Ext:   json.RawMessage(`{"uid_key":"value1"}`),
					},
					{
						ID:    "uid2",
						AType: adcom1.AgentTypeApp,
						Ext:   json.RawMessage(`{"uid_key":"value2"}`),
					},
				},
				Ext: json.RawMessage(`{"eid_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.e.Clone()

			if tt.e == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil EID")
				}
				return
			}

			// Test primitive fields
			if clone.Source != tt.e.Source {
				t.Errorf("Clone() Source = %v, want %v", clone.Source, tt.e.Source)
			}

			// Test UIDs array
			if len(tt.e.UIDs) > 0 {
				if len(clone.UIDs) != len(tt.e.UIDs) {
					t.Errorf("Clone() UIDs length = %v, want %v", len(clone.UIDs), len(tt.e.UIDs))
				}

				// Test deep copy of first UID
				originalID := tt.e.UIDs[0].ID
				tt.e.UIDs[0].ID = "modified"
				if clone.UIDs[0].ID != originalID {
					t.Error("Clone() should create deep copy of UIDs")
				}

				if tt.e.UIDs[0].Ext != nil {
					orig := string(tt.e.UIDs[0].Ext)
					tt.e.UIDs[0].Ext[0] = 'X'
					clonedStr := string(clone.UIDs[0].Ext)
					if clonedStr != orig {
						t.Error("Clone() should create deep copy of UIDs[0].Ext")
					}
				}
			}

			// Test extension
			if tt.e.Ext != nil {
				orig := string(tt.e.Ext)
				tt.e.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
