package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestPMP_Clone(t *testing.T) {
	tests := []struct {
		name string
		p    *PMP
	}{
		{
			name: "nil pmp",
			p:    nil,
		},
		{
			name: "empty pmp",
			p:    &PMP{},
		},
		{
			name: "fully populated pmp",
			p: &PMP{
				PrivateAuction: 1,
				Deals: []Deal{
					{
						ID:          "deal1",
						BidFloor:    1.23,
						BidFloorCur: "USD",
						AT:          1,
						WSeat:       []string{"seat1", "seat2"},
						WADomain:    []string{"domain1.com"},
						Ext:         json.RawMessage(`{"deal_key": "value"}`),
					},
					{
						ID:          "deal2",
						BidFloor:    2.34,
						BidFloorCur: "EUR",
					},
				},
				Ext: json.RawMessage(`{"pmp_key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.p.Clone()

			if tt.p == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil PMP")
				}
				return
			}

			// Test primitive fields
			if clone.PrivateAuction != tt.p.PrivateAuction {
				t.Errorf("Clone() PrivateAuction = %v, want %v", clone.PrivateAuction, tt.p.PrivateAuction)
			}

			// Test deals slice
			if len(tt.p.Deals) > 0 {
				if len(clone.Deals) != len(tt.p.Deals) {
					t.Errorf("Clone() Deals length = %v, want %v", len(clone.Deals), len(tt.p.Deals))
				}

				// Test deep copy by modifying original
				originalDealID := tt.p.Deals[0].ID
				tt.p.Deals[0].ID = "modified"
				if clone.Deals[0].ID != originalDealID {
					t.Error("Clone() should create deep copy of Deals")
				}

				// Test nested slice deep copy
				if len(tt.p.Deals[0].WSeat) > 0 {
					originalSeat := tt.p.Deals[0].WSeat[0]
					tt.p.Deals[0].WSeat[0] = "modified"
					if clone.Deals[0].WSeat[0] != originalSeat {
						t.Error("Clone() should create deep copy of Deal.WSeat")
					}
				}
			}

			// Test extension
			if tt.p.Ext != nil {
				orig := string(tt.p.Ext)
				tt.p.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
