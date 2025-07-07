package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestContent_Clone(t *testing.T) {
	var (
		prodQ           = adcom1.ProductionQuality(1)
		liveStream int8 = 1
		sourceRel  int8 = 1
		embeddable int8 = 1
	)

	tests := []struct {
		name string
		c    *Content
	}{
		{
			name: "nil content",
			c:    nil,
		},
		{
			name: "empty content",
			c:    &Content{},
		},
		{
			name: "fully populated content",
			c: &Content{
				ID:                 "content1",
				Episode:            1,
				Title:              "Test Content",
				Series:             "Test Series",
				Season:             "Season 1",
				Artist:             "Test Artist",
				Genre:              "Test Genre",
				Album:              "Test Album",
				ISRC:               "TEST123",
				Producer:           &Producer{ID: "prod1", Ext: json.RawMessage(`{"key": "value"}`)},
				URL:                "http://test.com",
				CatTax:             adcom1.CategoryTaxonomy(1),
				Cat:                []string{"IAB1", "IAB2"},
				ProdQ:              &prodQ,
				VideoQuality:       &prodQ,
				Context:            1,
				ContentRating:      "PG",
				UserRating:         "4.5",
				QAGMediaRating:     1,
				Keywords:           "test,content",
				KwArray:            []string{"test", "content"},
				LiveStream:         &liveStream,
				SourceRelationship: &sourceRel,
				Len:                300,
				Language:           "en",
				LangB:              "en-US",
				Embeddable:         &embeddable,
				Data: []Data{
					{
						ID:   "data1",
						Name: "Data 1",
						Segment: []Segment{
							{ID: "seg1", Value: "val1"},
						},
						Ext: json.RawMessage(`{"key": "value"}`),
					},
				},
				Network: &Network{
					ID:   "net1",
					Name: "Network 1",
					Ext:  json.RawMessage(`{"key": "value"}`),
				},
				Channel: &Channel{
					ID:   "ch1",
					Name: "Channel 1",
					Ext:  json.RawMessage(`{"key": "value"}`),
				},
				Ext: json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.c.Clone()

			if tt.c == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Content")
				}
				return
			}

			// Test pointer fields
			if tt.c.ProdQ != nil {
				orig := *tt.c.ProdQ
				*tt.c.ProdQ = adcom1.ProductionQuality(2)
				if *clone.ProdQ != orig {
					t.Error("Clone() should create deep copy of ProdQ")
				}
			}

			// Test Producer
			if tt.c.Producer != nil {
				origID := tt.c.Producer.ID
				tt.c.Producer.ID = "modified"
				if clone.Producer.ID != origID {
					t.Error("Clone() should create deep copy of Producer")
				}
			}

			// Test Data slice
			if len(tt.c.Data) > 0 {
				origName := tt.c.Data[0].Name
				tt.c.Data[0].Name = "modified"
				if clone.Data[0].Name != origName {
					t.Error("Clone() should create deep copy of Data")
				}
			}

			// Test Network
			if tt.c.Network != nil {
				origID := tt.c.Network.ID
				tt.c.Network.ID = "modified"
				if clone.Network.ID != origID {
					t.Error("Clone() should create deep copy of Network")
				}
			}

			// Test Channel
			if tt.c.Channel != nil {
				origID := tt.c.Channel.ID
				tt.c.Channel.ID = "modified"
				if clone.Channel.ID != origID {
					t.Error("Clone() should create deep copy of Channel")
				}
			}

			// Test extension
			if tt.c.Ext != nil {
				orig := string(tt.c.Ext)
				tt.c.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
