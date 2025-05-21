package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestQty_Clone(t *testing.T) {
	tests := []struct {
		name string
		q    *Qty
	}{
		{
			name: "nil qty",
			q:    nil,
		},
		{
			name: "empty qty",
			q:    &Qty{},
		},
		{
			name: "fully populated qty",
			q: &Qty{
				Multiplier: 14.2,
				SourceType: adcom1.MultiplierMeasurementVendorProvided,
				Vendor:     "example.com",
				Ext:        json.RawMessage(`{"qty_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.q.Clone()

			if tt.q == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Qty")
				}
				return
			}

			// Test primitive fields
			if clone.Multiplier != tt.q.Multiplier {
				t.Errorf("Clone() Multiplier = %v, want %v", clone.Multiplier, tt.q.Multiplier)
			}
			if clone.SourceType != tt.q.SourceType {
				t.Errorf("Clone() SourceType = %v, want %v", clone.SourceType, tt.q.SourceType)
			}
			if clone.Vendor != tt.q.Vendor {
				t.Errorf("Clone() Vendor = %v, want %v", clone.Vendor, tt.q.Vendor)
			}

			// Test extension
			if tt.q.Ext != nil {
				orig := string(tt.q.Ext)
				tt.q.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
