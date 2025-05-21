package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestSupplyChainNode_Clone(t *testing.T) {
	hp := int8(1)

	tests := []struct {
		name string
		n    *SupplyChainNode
	}{
		{
			name: "nil node",
			n:    nil,
		},
		{
			name: "empty node",
			n:    &SupplyChainNode{},
		},
		{
			name: "fully populated node",
			n: &SupplyChainNode{
				ASI:    "ssp.example.com",
				SID:    "seller123",
				RID:    "request123",
				Name:   "Example SSP",
				Domain: "example.com",
				HP:     &hp,
				Ext:    json.RawMessage(`{"node_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.n.Clone()

			if tt.n == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil SupplyChainNode")
				}
				return
			}

			// Test primitive fields
			if clone.ASI != tt.n.ASI {
				t.Errorf("Clone() ASI = %v, want %v", clone.ASI, tt.n.ASI)
			}
			if clone.SID != tt.n.SID {
				t.Errorf("Clone() SID = %v, want %v", clone.SID, tt.n.SID)
			}
			if clone.RID != tt.n.RID {
				t.Errorf("Clone() RID = %v, want %v", clone.RID, tt.n.RID)
			}
			if clone.Name != tt.n.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.n.Name)
			}
			if clone.Domain != tt.n.Domain {
				t.Errorf("Clone() Domain = %v, want %v", clone.Domain, tt.n.Domain)
			}

			// Test HP pointer
			if tt.n.HP != nil {
				if *clone.HP != *tt.n.HP {
					t.Errorf("Clone() HP = %v, want %v", *clone.HP, *tt.n.HP)
				}
				oldHP := *tt.n.HP
				*tt.n.HP = 0
				if *clone.HP != oldHP {
					t.Error("Clone() should create deep copy of HP")
				}
			}

			// Test extension
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
