package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestRegs_Clone(t *testing.T) {
	gdpr := int8(1)

	tests := []struct {
		name string
		r    *Regs
	}{
		{
			name: "nil regs",
			r:    nil,
		},
		{
			name: "empty regs",
			r:    &Regs{},
		},
		{
			name: "fully populated regs",
			r: &Regs{
				COPPA:     1,
				GDPR:      &gdpr,
				USPrivacy: "1YNY",
				GPP:       "DBACNYA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA",
				GPPSID:    []int8{1, 2, 6},
				Ext:       json.RawMessage(`{"regs_key":"value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.r.Clone()

			if tt.r == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Regs")
				}
				return
			}

			// Test primitive fields
			if clone.COPPA != tt.r.COPPA {
				t.Errorf("Clone() COPPA = %v, want %v", clone.COPPA, tt.r.COPPA)
			}
			if clone.USPrivacy != tt.r.USPrivacy {
				t.Errorf("Clone() USPrivacy = %v, want %v", clone.USPrivacy, tt.r.USPrivacy)
			}
			if clone.GPP != tt.r.GPP {
				t.Errorf("Clone() GPP = %v, want %v", clone.GPP, tt.r.GPP)
			}

			// Test GDPR pointer
			if tt.r.GDPR != nil {
				if *clone.GDPR != *tt.r.GDPR {
					t.Errorf("Clone() GDPR = %v, want %v", *clone.GDPR, *tt.r.GDPR)
				}
				oldGDPR := *tt.r.GDPR
				*tt.r.GDPR = 0
				if *clone.GDPR != oldGDPR {
					t.Error("Clone() should create deep copy of GDPR")
				}
			}

			// Test GPPSID slice
			if len(tt.r.GPPSID) > 0 {
				if len(clone.GPPSID) != len(tt.r.GPPSID) {
					t.Errorf("Clone() GPPSID length = %v, want %v", len(clone.GPPSID), len(tt.r.GPPSID))
				}
				original := tt.r.GPPSID[0]
				tt.r.GPPSID[0] = 99
				if clone.GPPSID[0] != original {
					t.Error("Clone() should create deep copy of GPPSID")
				}
			}

			// Test extension
			if tt.r.Ext != nil {
				orig := string(tt.r.Ext)
				tt.r.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
