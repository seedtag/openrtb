package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestDeal_Clone(t *testing.T) {
	tests := []struct {
		name string
		d    *Deal
	}{
		{
			name: "nil deal",
			d:    nil,
		},
		{
			name: "empty deal",
			d:    &Deal{},
		},
		{
			name: "fully populated deal",
			d: &Deal{
				ID:           "deal1",
				BidFloor:     1.23,
				BidFloorCur:  "USD",
				AT:           1,
				WSeat:        []string{"seat1", "seat2"},
				WADomain:     []string{"domain1.com", "domain2.com"},
				Guar:         1,
				MinCPMPerSec: 0.5,
				DurFloors: []DurFloors{
					{
						MinDur:   15,
						BidFloor: 2.50,
						Ext:      json.RawMessage(`{"key1": "value1"}`),
					},
					{
						MinDur:   30,
						MaxDur:   30,
						BidFloor: 4.50,
						Ext:      json.RawMessage(`{"key2": "value2"}`),
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
					t.Error("Clone() should return nil for nil Deal")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.d.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.d.ID)
			}
			if clone.BidFloor != tt.d.BidFloor {
				t.Errorf("Clone() BidFloor = %v, want %v", clone.BidFloor, tt.d.BidFloor)
			}
			if clone.MinCPMPerSec != tt.d.MinCPMPerSec {
				t.Errorf("Clone() MinCPMPerSec = %v, want %v", clone.MinCPMPerSec, tt.d.MinCPMPerSec)
			}

			// Test string slices
			if len(tt.d.WSeat) > 0 {
				original := tt.d.WSeat[0]
				tt.d.WSeat[0] = "modified"
				if clone.WSeat[0] != original {
					t.Error("Clone() should create deep copy of WSeat slice")
				}
			}

			// Test DurFloors
			if len(tt.d.DurFloors) > 0 {
				original := tt.d.DurFloors[0].BidFloor
				tt.d.DurFloors[0].BidFloor = 999.99
				if clone.DurFloors[0].BidFloor != original {
					t.Error("Clone() should create deep copy of DurFloors")
				}

				if tt.d.DurFloors[0].Ext != nil {
					origExt := string(tt.d.DurFloors[0].Ext)
					tt.d.DurFloors[0].Ext[0] = 'X'
					clonedExt := string(clone.DurFloors[0].Ext)
					if clonedExt != origExt {
						t.Error("Clone() should create deep copy of DurFloors[0].Ext")
					}
				}
			}

			// Test extension
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
