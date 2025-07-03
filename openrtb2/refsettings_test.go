package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestRefSettings_Clone(t *testing.T) {
	tests := []struct {
		name string
		r    *RefSettings
	}{
		{
			name: "nil refsettings",
			r:    nil,
		},
		{
			name: "empty refsettings",
			r:    &RefSettings{},
		},
		{
			name: "fully populated refsettings",
			r: &RefSettings{
				RefType: adcom1.AutoRefreshTriggerTime,
				MinInt:  30,
				Ext:     json.RawMessage(`{"refsettings_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.r.Clone()

			if tt.r == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil RefSettings")
				}
				return
			}

			// Test primitive fields
			if clone.RefType != tt.r.RefType {
				t.Errorf("Clone() RefType = %v, want %v", clone.RefType, tt.r.RefType)
			}
			if clone.MinInt != tt.r.MinInt {
				t.Errorf("Clone() MinInt = %v, want %v", clone.MinInt, tt.r.MinInt)
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
