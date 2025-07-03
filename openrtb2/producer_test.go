package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestProducer_Clone(t *testing.T) {
	tests := []struct {
		name string
		p    *Producer
	}{
		{
			name: "nil producer",
			p:    nil,
		},
		{
			name: "empty producer",
			p:    &Producer{},
		},
		{
			name: "fully populated producer",
			p: &Producer{
				ID:     "prod1",
				Name:   "Producer 1",
				CatTax: adcom1.CategoryTaxonomy(1),
				Cat:    []string{"IAB1", "IAB2"},
				Domain: "producer.com",
				Ext:    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.p.Clone()

			if tt.p == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Producer")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.p.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.p.ID)
			}
			if clone.Name != tt.p.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.p.Name)
			}
			if clone.Domain != tt.p.Domain {
				t.Errorf("Clone() Domain = %v, want %v", clone.Domain, tt.p.Domain)
			}
			if clone.CatTax != tt.p.CatTax {
				t.Errorf("Clone() CatTax = %v, want %v", clone.CatTax, tt.p.CatTax)
			}

			// Test slices
			if len(tt.p.Cat) > 0 {
				original := tt.p.Cat[0]
				tt.p.Cat[0] = "modified"
				if clone.Cat[0] != original {
					t.Error("Clone() should create deep copy of Cat slice")
				}
			}

			// Test deep copy of extension
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
