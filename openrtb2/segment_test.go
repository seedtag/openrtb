package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestSegment_Clone(t *testing.T) {
	tests := []struct {
		name string
		s    *Segment
	}{
		{
			name: "nil segment",
			s:    nil,
		},
		{
			name: "empty segment",
			s:    &Segment{},
		},
		{
			name: "fully populated segment",
			s: &Segment{
				ID:    "seg1",
				Name:  "Segment 1",
				Value: "value1",
				Ext:   json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.s.Clone()

			if tt.s == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Segment")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.s.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.s.ID)
			}
			if clone.Name != tt.s.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.s.Name)
			}
			if clone.Value != tt.s.Value {
				t.Errorf("Clone() Value = %v, want %v", clone.Value, tt.s.Value)
			}

			// Test deep copy of extension
			if tt.s.Ext != nil {
				orig := string(tt.s.Ext)
				tt.s.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
