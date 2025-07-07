package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestRefresh_Clone(t *testing.T) {
	count := 2

	tests := []struct {
		name string
		r    *Refresh
	}{
		{
			name: "nil refresh",
			r:    nil,
		},
		{
			name: "empty refresh",
			r:    &Refresh{},
		},
		{
			name: "fully populated refresh",
			r: &Refresh{
				RefSettings: []RefSettings{
					{
						RefType: adcom1.AutoRefreshTriggerTime,
						MinInt:  30,
						Ext:     json.RawMessage(`{"refsettings_key":"value1"}`),
					},
					{
						RefType: adcom1.AutoRefreshTriggerUserAction,
						MinInt:  60,
						Ext:     json.RawMessage(`{"refsettings_key":"value2"}`),
					},
				},
				Count: &count,
				Ext:   json.RawMessage(`{"refresh_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.r.Clone()

			if tt.r == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Refresh")
				}
				return
			}

			// Test Count pointer
			if tt.r.Count != nil {
				if *clone.Count != *tt.r.Count {
					t.Errorf("Clone() Count = %v, want %v", *clone.Count, *tt.r.Count)
				}
				oldCount := *tt.r.Count
				*tt.r.Count = 99
				if *clone.Count != oldCount {
					t.Error("Clone() should create deep copy of Count")
				}
			}

			// Test RefSettings slice
			if len(tt.r.RefSettings) > 0 {
				if len(clone.RefSettings) != len(tt.r.RefSettings) {
					t.Errorf("Clone() RefSettings length = %v, want %v", len(clone.RefSettings), len(tt.r.RefSettings))
				}

				// Test deep copy of first RefSettings
				originalRefType := tt.r.RefSettings[0].RefType
				tt.r.RefSettings[0].RefType = adcom1.AutoRefreshTriggerUnknown
				if clone.RefSettings[0].RefType != originalRefType {
					t.Error("Clone() should create deep copy of RefSettings")
				}

				// Test RefSettings[0].Ext deep copy
				if tt.r.RefSettings[0].Ext != nil {
					orig := string(tt.r.RefSettings[0].Ext)
					tt.r.RefSettings[0].Ext[0] = 'X'
					clonedStr := string(clone.RefSettings[0].Ext)
					if clonedStr != orig {
						t.Error("Clone() should create deep copy of RefSettings[0].Ext")
					}
				}
			}

			// Test extension
			if tt.r.Ext != nil {
				orig := string(tt.r.Ext)
				tt.r.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
