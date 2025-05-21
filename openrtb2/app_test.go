package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestApp_Clone(t *testing.T) {
	var (
		privacyPolicy int8 = 1
		paid          int8 = 1
	)

	tests := []struct {
		name string
		a    *App
	}{
		{
			name: "nil app",
			a:    nil,
		},
		{
			name: "empty app",
			a:    &App{},
		},
		{
			name: "fully populated app",
			a: &App{
				ID:            "app1",
				Name:          "Test App",
				Bundle:        "com.test.app",
				Domain:        "test.com",
				StoreURL:      "https://play.google.com/store/apps/details?id=com.test.app",
				CatTax:        adcom1.CategoryTaxonomy(1),
				Cat:           []string{"IAB1", "IAB2"},
				SectionCat:    []string{"IAB3", "IAB4"},
				PageCat:       []string{"IAB5", "IAB6"},
				Ver:           "1.0",
				PrivacyPolicy: &privacyPolicy,
				Paid:          &paid,
				Publisher: &Publisher{
					ID:   "pub1",
					Name: "Publisher 1",
					Ext:  json.RawMessage(`{"key": "value"}`),
				},
				Content: &Content{
					ID:    "content1",
					Title: "Test Content",
					Ext:   json.RawMessage(`{"key": "value"}`),
				},
				Keywords:               "test,app",
				KwArray:                []string{"test", "app"},
				InventoryPartnerDomain: "partner.com",
				Ext:                    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.a.Clone()

			if tt.a == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil App")
				}
				return
			}

			// Test pointer fields
			if tt.a.PrivacyPolicy != nil {
				orig := *tt.a.PrivacyPolicy
				*tt.a.PrivacyPolicy = 0
				if *clone.PrivacyPolicy != orig {
					t.Error("Clone() should create deep copy of PrivacyPolicy")
				}
			}

			// Test string slices
			if len(tt.a.Cat) > 0 {
				original := tt.a.Cat[0]
				tt.a.Cat[0] = "modified"
				if clone.Cat[0] != original {
					t.Error("Clone() should create deep copy of Cat slice")
				}
			}

			// Test Publisher
			if tt.a.Publisher != nil {
				origName := tt.a.Publisher.Name
				tt.a.Publisher.Name = "modified"
				if clone.Publisher.Name != origName {
					t.Error("Clone() should create deep copy of Publisher")
				}
			}

			// Test Content
			if tt.a.Content != nil {
				origTitle := tt.a.Content.Title
				tt.a.Content.Title = "modified"
				if clone.Content.Title != origTitle {
					t.Error("Clone() should create deep copy of Content")
				}
			}

			// Test deep copy of extension
			if tt.a.Ext != nil {
				orig := string(tt.a.Ext)
				tt.a.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
