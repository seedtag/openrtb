package openrtb2

import (
	"encoding/json"
	"testing"
)

func TestChannel_Clone(t *testing.T) {
	tests := []struct {
		name string
		c    *Channel
	}{
		{
			name: "nil channel",
			c:    nil,
		},
		{
			name: "empty channel",
			c:    &Channel{},
		},
		{
			name: "fully populated channel",
			c: &Channel{
				ID:     "ch1",
				Name:   "Channel 1",
				Domain: "channel1.com",
				Ext:    json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.c.Clone()

			if tt.c == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Channel")
				}
				return
			}

			// Test primitive fields
			if clone.ID != tt.c.ID {
				t.Errorf("Clone() ID = %v, want %v", clone.ID, tt.c.ID)
			}
			if clone.Name != tt.c.Name {
				t.Errorf("Clone() Name = %v, want %v", clone.Name, tt.c.Name)
			}
			if clone.Domain != tt.c.Domain {
				t.Errorf("Clone() Domain = %v, want %v", clone.Domain, tt.c.Domain)
			}

			// Test deep copy of extension
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
