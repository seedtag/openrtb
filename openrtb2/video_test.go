package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestVideo_Clone(t *testing.T) {
	var (
		w             int64 = 640
		h             int64 = 480
		skip          int8  = 1
		boxingAllowed int8  = 1
		startDelay          = adcom1.StartDelay(0)
		pos                 = adcom1.PlacementPosition(1)
	)

	tests := []struct {
		name string
		v    *Video
	}{
		{
			name: "nil video",
			v:    nil,
		},
		{
			name: "empty video",
			v:    &Video{},
		},
		{
			name: "fully populated video",
			v: &Video{
				MIMEs:          []string{"video/mp4", "video/webm"},
				MinDuration:    15,
				MaxDuration:    30,
				StartDelay:     &startDelay,
				MaxSeq:         3,
				PodDur:         90,
				Protocols:      []adcom1.MediaCreativeSubtype{1, 2},
				Protocol:       1,
				W:              &w,
				H:              &h,
				PodID:          "pod123",
				PodSeq:         1,
				RqdDurs:        []int64{15, 30},
				Placement:      1,
				Plcmt:          1,
				Linearity:      1,
				Skip:           &skip,
				SkipMin:        5,
				SkipAfter:      5,
				Sequence:       1,
				SlotInPod:      1,
				MinCPMPerSec:   1.5,
				BAttr:          []adcom1.CreativeAttribute{1, 2},
				MaxExtended:    60,
				MinBitRate:     1000,
				MaxBitRate:     2000,
				BoxingAllowed:  &boxingAllowed,
				PlaybackMethod: []adcom1.PlaybackMethod{1, 2},
				PlaybackEnd:    1,
				Delivery:       []adcom1.DeliveryMethod{1, 2},
				Pos:            &pos,
				CompanionAd: []Banner{
					{
						W:   &w,
						H:   &h,
						ID:  "banner1",
						Ext: json.RawMessage(`{"key": "value"}`),
					},
				},
				API:           []adcom1.APIFramework{1, 2},
				CompanionType: []adcom1.CompanionType{1, 2},
				PodDedupe:     []adcom1.PodDedupe{1, 2},
				DurFloors:     []DurFloors{{MinDur: 15, BidFloor: 5.0}},
				Ext:           json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.v.Clone()

			if tt.v == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Video")
				}
				return
			}

			// Test deep copy of extension
			if tt.v.Ext != nil {
				orig := string(tt.v.Ext)       // Store original as string
				tt.v.Ext[0] = 'X'              // Modify original
				clonedStr := string(clone.Ext) // Get clone as string
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot:  %q", orig, clonedStr)
				}
			}

			// Test pointer fields
			if tt.v.W != nil {
				origW := *tt.v.W
				*tt.v.W = 999
				if *clone.W != origW {
					t.Error("Clone() should create deep copy of W pointer")
				}
			}

			// Test slices
			if len(tt.v.MIMEs) > 0 {
				original := tt.v.MIMEs[0]
				tt.v.MIMEs[0] = "modified"
				if clone.MIMEs[0] != original {
					t.Error("Clone() should create deep copy of MIMEs slice")
				}
			}

			// Test CompanionAd slice
			if len(tt.v.CompanionAd) > 0 {
				origID := tt.v.CompanionAd[0].ID
				tt.v.CompanionAd[0].ID = "modified"
				if clone.CompanionAd[0].ID != origID {
					t.Error("Clone() should create deep copy of CompanionAd")
				}
			}

			// Test DurFloors slice
			if len(tt.v.DurFloors) > 0 {
				origFloor := tt.v.DurFloors[0].BidFloor
				tt.v.DurFloors[0].BidFloor = 999.99
				if clone.DurFloors[0].BidFloor != origFloor {
					t.Error("Clone() should create deep copy of DurFloors")
				}
			}
		})
	}
}
