package openrtb2

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFormat_Clone(t *testing.T) {
	tests := []struct {
		name string
		f    *Format
	}{
		{
			name: "nil format",
			f:    nil,
		},
		{
			name: "empty format",
			f:    &Format{},
		},
		{
			name: "complete format",
			f: &Format{
				W:      300,
				H:      250,
				WRatio: 16,
				HRatio: 9,
				WMin:   100,
				Ext:    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.f.Clone()

			if tt.f == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Format")
				}
				return
			}

			// Verify initial equality
			if !reflect.DeepEqual(clone, tt.f) {
				t.Errorf("Clone() = %v, want %v", clone, tt.f)
			}

			// Verify deep copy by modifying the original
			if tt.f.Ext != nil {
				orig := make(json.RawMessage, len(tt.f.Ext))
				copy(orig, tt.f.Ext)
				tt.f.Ext[0] = 'X'
				if !reflect.DeepEqual(clone.Ext, orig) {
					t.Error("Clone() should create a deep copy of Ext")
				}
			}

			// Modify primitive fields
			origW := tt.f.W
			tt.f.W = 999
			if clone.W != origW {
				t.Error("Clone() should create a deep copy of primitive fields")
			}
		})
	}
}
