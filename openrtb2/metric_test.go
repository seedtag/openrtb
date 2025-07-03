package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestMetric_Clone(t *testing.T) {
	tests := []struct {
		name string
		m    *Metric
	}{
		{
			name: "nil metric",
			m:    nil,
		},
		{
			name: "empty metric",
			m:    &Metric{},
		},
		{
			name: "fully populated metric",
			m: &Metric{
				Type:   "viewability",
				Value:  0.85,
				Vendor: "EXCHANGE",
				Ext:    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.m.Clone()

			if tt.m == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Metric")
				}
				return
			}

			// Test primitive fields
			if clone.Type != tt.m.Type {
				t.Errorf("Clone() Type = %v, want %v", clone.Type, tt.m.Type)
			}
			if clone.Value != tt.m.Value {
				t.Errorf("Clone() Value = %v, want %v", clone.Value, tt.m.Value)
			}
			if clone.Vendor != tt.m.Vendor {
				t.Errorf("Clone() Vendor = %v, want %v", clone.Vendor, tt.m.Vendor)
			}

			// Test deep copy of extension
			if tt.m.Ext != nil {
				orig := string(tt.m.Ext)
				tt.m.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
