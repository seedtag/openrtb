package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestUID_Clone(t *testing.T) {
	tests := []struct {
		name string
		u    *UID
	}{
		{
			name: "nil uid",
			u:    nil,
		},
		{
			name: "empty uid",
			u:    &UID{},
		},
		{
			name: "fully populated uid",
			u: &UID{
				ID:    "uid123",
				AType: adcom1.AgentTypeWeb,
				Ext:   json.RawMessage(`{"uid_key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.u.Clone()

			if tt.u == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil UID")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.u.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.u.ID)
			}
			if clone.AType != tt.u.AType {
				t.Errorf("Clone() AType = %v, want %v", clone.AType, tt.u.AType)
			}

			// Test extension
			if tt.u.Ext != nil {
				orig := string(tt.u.Ext)
				tt.u.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
