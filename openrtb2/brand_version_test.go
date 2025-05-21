package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestBrandVersion_Clone(t *testing.T) {
	tests := []struct {
		name string
		b    *BrandVersion
	}{
		{
			name: "nil brand version",
			b:    nil,
		},
		{
			name: "empty brand version",
			b:    &BrandVersion{},
		},
		{
			name: "fully populated brand version",
			b: &BrandVersion{
				Brand:   "Chrome",
				Version: []string{"91", "0", "4472", "77"},
				Ext:     json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.b.Clone()

			if tt.b == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil BrandVersion")
				}
				return
			}

			// Test primitive fields
			if clone.Brand != tt.b.Brand {
				t.Errorf("Clone() Brand = %v, want %v", clone.Brand, tt.b.Brand)
			}

			// Test slice
			if len(tt.b.Version) > 0 {
				original := tt.b.Version[0]
				tt.b.Version[0] = "modified"
				if clone.Version[0] != original {
					t.Error("Clone() should create deep copy of Version slice")
				}
			}

			// Test deep copy of extension
			if tt.b.Ext != nil {
				orig := string(tt.b.Ext)
				tt.b.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
