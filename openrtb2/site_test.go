package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestSite_Clone(t *testing.T) {
	var (
		mobile        int8 = 1
		privacyPolicy int8 = 1
	)

	tests := []struct {
		name string
		s    *Site
	}{
		{
			name: "nil site",
			s:    nil,
		},
		{
			name: "empty site",
			s:    &Site{},
		},
		{
			name: "fully populated site",
			s: &Site{
				ID:            "site1",
				Name:          "Test Site",
				Domain:        "test.com",
				CatTax:        adcom1.CategoryTaxonomy(1),
				Cat:           []string{"IAB1", "IAB2"},
				SectionCat:    []string{"IAB3", "IAB4"},
				PageCat:       []string{"IAB5", "IAB6"},
				Page:          "https://test.com/page",
				Ref:           "https://referrer.com",
				Search:        "test search",
				Mobile:        &mobile,
				PrivacyPolicy: &privacyPolicy,
				Publisher: &Publisher{
					ID:   "pub1",
					Name: "Publisher 1",
					Ext:  json.RawMessage(`{"key": "value"}`),
				},
				Content: &Content{
					ID:    "content1",
					Title: "Content Title",
					Ext:   json.RawMessage(`{"key": "value"}`),
				},
				Keywords:               "test,site",
				KwArray:                []string{"test", "site"},
				InventoryPartnerDomain: "partner.com",
				Ext:                    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.s.Clone()

			if tt.s == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Site")
				}
				return
			}

			// Test pointer fields
			if tt.s.Mobile != nil {
				orig := *tt.s.Mobile
				*tt.s.Mobile = 0
				if *clone.Mobile != orig {
					t.Error("Clone() should create deep copy of Mobile")
				}
			}

			// Test string slices
			if len(tt.s.Cat) > 0 {
				original := tt.s.Cat[0]
				tt.s.Cat[0] = "modified"
				if clone.Cat[0] != original {
					t.Error("Clone() should create deep copy of Cat slice")
				}
			}

			// Test Publisher
			if tt.s.Publisher != nil {
				origName := tt.s.Publisher.Name
				tt.s.Publisher.Name = "modified"
				if clone.Publisher.Name != origName {
					t.Error("Clone() should create deep copy of Publisher")
				}

				if tt.s.Publisher.Ext != nil {
					origExt := string(tt.s.Publisher.Ext)
					tt.s.Publisher.Ext[0] = 'X'
					clonedExt := string(clone.Publisher.Ext)
					if clonedExt != origExt {
						t.Error("Clone() should create deep copy of Publisher.Ext")
					}
				}
			}

			// Test Content
			if tt.s.Content != nil {
				origTitle := tt.s.Content.Title
				tt.s.Content.Title = "modified"
				if clone.Content.Title != origTitle {
					t.Error("Clone() should create deep copy of Content")
				}

				if tt.s.Content.Ext != nil {
					origExt := string(tt.s.Content.Ext)
					tt.s.Content.Ext[0] = 'X'
					clonedExt := string(clone.Content.Ext)
					if clonedExt != origExt {
						t.Error("Clone() should create deep copy of Content.Ext")
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
