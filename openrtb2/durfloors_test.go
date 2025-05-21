package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestDurFloors_Clone(t *testing.T) {
	tests := []struct {
		name string
		d    *DurFloors
	}{
		{
			name: "nil durfloors",
			d:    nil,
		},
		{
			name: "empty durfloors",
			d:    &DurFloors{},
		},
		{
			name: "fully populated durfloors",
			d: &DurFloors{
				MinDur:   15,
				MaxDur:   30,
				BidFloor: 10.5,
				Ext:      json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.d.Clone()

			if tt.d == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil DurFloors")
				}
				return
			}

			// Test primitive fields
			if clone.MinDur != tt.d.MinDur {
				t.Errorf("Clone() MinDur = %v, want %v", clone.MinDur, tt.d.MinDur)
			}
			if clone.MaxDur != tt.d.MaxDur {
				t.Errorf("Clone() MaxDur = %v, want %v", clone.MaxDur, tt.d.MaxDur)
			}
			if clone.BidFloor != tt.d.BidFloor {
				t.Errorf("Clone() BidFloor = %v, want %v", clone.BidFloor, tt.d.BidFloor)
			}

			// Test deep copy of extension
			if tt.d.Ext != nil {
				orig := string(tt.d.Ext)       // Store original as string
				tt.d.Ext[0] = 'X'              // Modify original
				clonedStr := string(clone.Ext) // Get clone as string

				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot:  %q", orig, clonedStr)
				}
			}
		})
	}
}
