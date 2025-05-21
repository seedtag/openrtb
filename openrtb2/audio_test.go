package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestAudio_Clone(t *testing.T) {
	var (
		stitched   int8 = 1
		startDelay      = adcom1.StartDelay(0)
		nvol            = adcom1.VolumeNormalizationMode(1)
	)

	tests := []struct {
		name string
		a    *Audio
	}{
		{
			name: "nil audio",
			a:    nil,
		},
		{
			name: "empty audio",
			a:    &Audio{},
		},
		{
			name: "fully populated audio",
			a: &Audio{
				MIMEs:        []string{"audio/mp4"},
				MinDuration:  15,
				MaxDuration:  30,
				PodDur:       90,
				Protocols:    []adcom1.MediaCreativeSubtype{1, 2},
				StartDelay:   &startDelay,
				RqdDurs:      []int64{15, 30},
				PodID:        "pod123",
				PodSeq:       1,
				Sequence:     1,
				SlotInPod:    1,
				MinCPMPerSec: 1.5,
				BAttr:        []adcom1.CreativeAttribute{1, 2},
				MaxExtended:  60,
				MinBitrate:   1000,
				MaxBitrate:   2000,
				Delivery:     []adcom1.DeliveryMethod{1, 2},
				CompanionAd: []Banner{
					{
						ID:  "banner1",
						Ext: json.RawMessage(`{"key": "value"}`),
					},
				},
				API:           []adcom1.APIFramework{1, 2},
				CompanionType: []adcom1.CompanionType{1, 2},
				MaxSeq:        3,
				Feed:          1,
				Stitched:      &stitched,
				NVol:          &nvol,
				DurFloors:     []DurFloors{{MinDur: 15, BidFloor: 5.0}},
				Ext:           json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.a.Clone()

			if tt.a == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Audio")
				}
				return
			}

			// Test deep copy of extension
			if tt.a.Ext != nil {
				orig := string(tt.a.Ext)       // Store original as string
				tt.a.Ext[0] = 'X'              // Modify original
				clonedStr := string(clone.Ext) // Get clone as string
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot:  %q", orig, clonedStr)
				}
			}

			// Test pointer fields
			if tt.a.Stitched != nil {
				origStitched := *tt.a.Stitched
				*tt.a.Stitched = 0
				if *clone.Stitched != origStitched {
					t.Error("Clone() should create deep copy of Stitched pointer")
				}
			}

			// Test slices
			if len(tt.a.MIMEs) > 0 {
				original := tt.a.MIMEs[0]
				tt.a.MIMEs[0] = "modified"
				if clone.MIMEs[0] != original {
					t.Error("Clone() should create deep copy of MIMEs slice")
				}
			}

			// Test CompanionAd slice
			if len(tt.a.CompanionAd) > 0 {
				origID := tt.a.CompanionAd[0].ID
				tt.a.CompanionAd[0].ID = "modified"
				if clone.CompanionAd[0].ID != origID {
					t.Error("Clone() should create deep copy of CompanionAd")
				}
			}

			// Test DurFloors slice
			if len(tt.a.DurFloors) > 0 {
				origFloor := tt.a.DurFloors[0].BidFloor
				tt.a.DurFloors[0].BidFloor = 999.99
				if clone.DurFloors[0].BidFloor != origFloor {
					t.Error("Clone() should create deep copy of DurFloors")
				}
			}
		})
	}
}
