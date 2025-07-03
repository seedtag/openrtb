package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestNetwork_Clone(t *testing.T) {
	tests := []struct {
		name string
		n    *Network
	}{
		{
			name: "nil network",
			n:    nil,
		},
		{
			name: "empty network",
			n:    &Network{},
		},
		{
			name: "fully populated network",
			n: &Network{
				ID:     "net1",
				Name:   "Network 1",
				Domain: "network1.com",
				Ext:    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.n.Clone()

			if tt.n == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Network")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.n.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.n.ID)
			}
			if clone.Name != tt.n.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.n.Name)
			}
			if clone.Domain != tt.n.Domain {
				t.Errorf("Clone() Domain = %v, want %v", clone.Domain, tt.n.Domain)
			}

			// Test deep copy of extension
			if tt.n.Ext != nil {
				orig := string(tt.n.Ext)
				tt.n.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
