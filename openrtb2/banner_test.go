package openrtb2

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestBanner_Clone(t *testing.T) {
	var (
		w   int64 = 300
		h   int64 = 250
		pos       = adcom1.PlacementPosition(1)
		vcm int8  = 1
	)

	tests := []struct {
		name string
		b    *Banner
	}{
		{
			name: "nil banner",
			b:    nil,
		},
		{
			name: "empty banner",
			b:    &Banner{},
		},
		{
			name: "fully populated banner",
			b: &Banner{
				Format: []Format{
					{
						W:      300,
						H:      250,
						WRatio: 16,
						HRatio: 9,
						WMin:   100,
						Ext:    json.RawMessage(`{"key": "value"}`),
					},
				},
				W:        &w,
				H:        &h,
				WMax:     1000,
				HMax:     1000,
				WMin:     0,
				HMin:     0,
				BType:    []BannerAdType{1, 2},
				BAttr:    []adcom1.CreativeAttribute{1, 2},
				Pos:      &pos,
				MIMEs:    []string{"image/jpeg", "image/png"},
				TopFrame: 1,
				ExpDir:   []adcom1.ExpandableDirection{1, 2},
				API:      []adcom1.APIFramework{1, 2},
				ID:       "test",
				Vcm:      &vcm,
				Ext:      json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.b.Clone()

			if tt.b == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Banner")
				}
				return
			}

			// Verify initial equality
			if !reflect.DeepEqual(clone, tt.b) {
				t.Errorf("Clone() = %v, want %v", clone, tt.b)
			}

			// Test deep copy of Format slice
			if tt.b.Format != nil {
				origFormat := make([]Format, len(tt.b.Format))
				copy(origFormat, tt.b.Format)
				tt.b.Format[0].W = 999
				if !reflect.DeepEqual(clone.Format, origFormat) {
					t.Error("Clone() should create a deep copy of Format slice")
				}
			}

			// Test deep copy of pointer fields
			if tt.b.W != nil {
				origW := *tt.b.W
				*tt.b.W = 999
				if *clone.W != origW {
					t.Error("Clone() should create a deep copy of W pointer")
				}
			}

			// Test deep copy of extension
			if tt.b.Ext != nil {
				orig := string(tt.b.Ext)       // Store original as string
				tt.b.Ext[0] = 'X'              // Modify original
				clonedStr := string(clone.Ext) // Get clone as string

				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot:  %q", orig, clonedStr)
				}
			}

			// Test deep copy of slices
			if len(tt.b.MIMEs) > 0 {
				origMIMEs := make([]string, len(tt.b.MIMEs))
				copy(origMIMEs, tt.b.MIMEs)
				tt.b.MIMEs[0] = "modified"
				if !reflect.DeepEqual(clone.MIMEs, origMIMEs) {
					t.Error("Clone() should create a deep copy of MIMEs slice")
				}
			}
		})
	}
}
