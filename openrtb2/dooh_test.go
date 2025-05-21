package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestDOOH_Clone(t *testing.T) {
	venueTypeTax := adcom1.VenueTaxonomyOpenOOH10

	tests := []struct {
		name string
		d    *DOOH
	}{
		{
			name: "nil dooh",
			d:    nil,
		},
		{
			name: "empty dooh",
			d:    &DOOH{},
		},
		{
			name: "fully populated dooh",
			d: &DOOH{
				ID:           "dooh123",
				Name:         "Times Square Display",
				VenueType:    []string{"street", "billboard"},
				VenueTypeTax: &venueTypeTax,
				Publisher: &Publisher{
					ID:     "pub123",
					Name:   "DOOH Media Corp",
					Domain: "doohmedia.com",
					Ext:    json.RawMessage(`{"publisher_key": "value"}`),
				},
				Domain:   "display1.doohmedia.com",
				Keywords: "times square,billboard,nyc",
				Content: &Content{
					ID:    "content123",
					Title: "Times Square Feed",
					Ext:   json.RawMessage(`{"content_key": "value"}`),
				},
				Ext: json.RawMessage(`{"dooh_key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.d.Clone()

			if tt.d == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil DOOH")
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
			if clone.Domain != tt.d.Domain {
				t.Errorf("Clone() Domain = %v, want %v", clone.Domain, tt.d.Domain)
			}
			if clone.Keywords != tt.d.Keywords {
				t.Errorf("Clone() Keywords = %v, want %v", clone.Keywords, tt.d.Keywords)
			}

			// Test VenueType slice
			if len(tt.d.VenueType) > 0 {
				if len(clone.VenueType) != len(tt.d.VenueType) {
					t.Errorf("Clone() VenueType length = %v, want %v", len(clone.VenueType), len(tt.d.VenueType))
				}
				original := tt.d.VenueType[0]
				tt.d.VenueType[0] = "modified"
				if clone.VenueType[0] != original {
					t.Error("Clone() should create deep copy of VenueType")
				}
			}

			// Test VenueTypeTax
			if tt.d.VenueTypeTax != nil {
				if clone.VenueTypeTax == tt.d.VenueTypeTax {
					t.Error("Clone() should create a new pointer for VenueTypeTax")
				}
				if *clone.VenueTypeTax != *tt.d.VenueTypeTax {
					t.Error("Clone() VenueTypeTax value should match original")
				}
			}

			// Test Publisher
			if tt.d.Publisher != nil {
				if clone.Publisher == tt.d.Publisher {
					t.Error("Clone() should create a new pointer for Publisher")
				}
				origPubID := tt.d.Publisher.ID
				tt.d.Publisher.ID = "modified"
				if clone.Publisher.ID != origPubID {
					t.Error("Clone() should create deep copy of Publisher")
				}
			}

			// Test Content
			if tt.d.Content != nil {
				if clone.Content == tt.d.Content {
					t.Error("Clone() should create a new pointer for Content")
				}
				origContentID := tt.d.Content.ID
				tt.d.Content.ID = "modified"
				if clone.Content.ID != origContentID {
					t.Error("Clone() should create deep copy of Content")
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
