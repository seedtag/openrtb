package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestData_Clone(t *testing.T) {
	tests := []struct {
		name string
		d    *Data
	}{
		{
			name: "nil data",
			d:    nil,
		},
		{
			name: "empty data",
			d:    &Data{},
		},
		{
			name: "fully populated data",
			d: &Data{
				ID:   "data1",
				Name: "Data Provider 1",
				Segment: []Segment{
					{
						ID:    "seg1",
						Name:  "Segment 1",
						Value: "value1",
						Ext:   json.RawMessage(`{"key1": "value1"}`),
					},
					{
						ID:    "seg2",
						Name:  "Segment 2",
						Value: "value2",
						Ext:   json.RawMessage(`{"key2": "value2"}`),
					},
				},
				Ext: json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.d.Clone()

			if tt.d == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Data")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.d.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.d.ID)
			}
			if clone.Name != tt.d.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.d.Name)
			}

			// Test segments
			if len(tt.d.Segment) > 0 {
				origSegName := tt.d.Segment[0].Name
				tt.d.Segment[0].Name = "modified"
				if clone.Segment[0].Name != origSegName {
					t.Error("Clone() should create deep copy of Segment slice")
				}

				// Test segment's extension
				if tt.d.Segment[0].Ext != nil {
					origSegExt := string(tt.d.Segment[0].Ext)
					tt.d.Segment[0].Ext[0] = 'X'
					clonedSegExt := string(clone.Segment[0].Ext)
					if clonedSegExt != origSegExt {
						t.Error("Clone() should create deep copy of Segment's Ext")
					}
				}
			}

			// Test deep copy of extension
			if tt.d.Ext != nil {
				orig := string(tt.d.Ext)
				tt.d.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
