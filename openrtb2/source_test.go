package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestSource_Clone(t *testing.T) {
	fd := int8(1)
	hp := int8(1)

	tests := []struct {
		name string
		s    *Source
	}{
		{
			name: "nil source",
			s:    nil,
		},
		{
			name: "empty source",
			s:    &Source{},
		},
		{
			name: "fully populated source",
			s: &Source{
				FD:     &fd,
				TID:    "transaction123",
				PChain: "pchain123",
				SChain: &SupplyChain{
					Complete: 1,
					Nodes: []SupplyChainNode{
						{
							ASI:    "ssp.example.com",
							SID:    "seller123",
							RID:    "request123",
							Name:   "Example SSP",
							Domain: "example.com",
							HP:     &hp,
							Ext:    json.RawMessage(`{"node_key":"value"}`),
						},
					},
					Ver: "1.0",
					Ext: json.RawMessage(`{"chain_key":"value"}`),
				},
				Ext: json.RawMessage(`{"source_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.s.Clone()

			if tt.s == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Source")
				}
				return
			}

			// Test primitive fields
			if clone.TID != tt.s.TID {
				t.Errorf("Clone() TID = %v, want %v", clone.TID, tt.s.TID)
			}
			if clone.PChain != tt.s.PChain {
				t.Errorf("Clone() PChain = %v, want %v", clone.PChain, tt.s.PChain)
			}

			// Test FD pointer
			if tt.s.FD != nil {
				if *clone.FD != *tt.s.FD {
					t.Errorf("Clone() FD = %v, want %v", *clone.FD, *tt.s.FD)
				}
				oldFD := *tt.s.FD
				*tt.s.FD = 0
				if *clone.FD != oldFD {
					t.Error("Clone() should create deep copy of FD")
				}
			}

			// Test SChain
			if tt.s.SChain != nil {
				if clone.SChain == tt.s.SChain {
					t.Error("Clone() should create a new pointer for SChain")
				}
				if tt.s.SChain.Complete != clone.SChain.Complete {
					t.Errorf("Clone() SChain.Complete = %v, want %v", clone.SChain.Complete, tt.s.SChain.Complete)
				}

				// Test deep copy of chain's node
				if len(tt.s.SChain.Nodes) > 0 {
					originalASI := tt.s.SChain.Nodes[0].ASI
					tt.s.SChain.Nodes[0].ASI = "modified"
					if clone.SChain.Nodes[0].ASI != originalASI {
						t.Error("Clone() should create deep copy of SChain.Nodes")
					}

					// Test node's extension
					if tt.s.SChain.Nodes[0].Ext != nil {
						orig := string(tt.s.SChain.Nodes[0].Ext)
						tt.s.SChain.Nodes[0].Ext[0] = 'X'
						clonedStr := string(clone.SChain.Nodes[0].Ext)
						if clonedStr != orig {
							t.Error("Clone() should create deep copy of SChain.Nodes[0].Ext")
						}
					}
				}

				// Test chain's extension
				if tt.s.SChain.Ext != nil {
					orig := string(tt.s.SChain.Ext)
					tt.s.SChain.Ext[0] = 'X'
					clonedStr := string(clone.SChain.Ext)
					if clonedStr != orig {
						t.Error("Clone() should create deep copy of SChain.Ext")
					}
				}
			}

			// Test extension
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
