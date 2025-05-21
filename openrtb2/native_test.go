package openrtb2

import (
	"encoding/json"
	"testing"

	"github.com/prebid/openrtb/v20/adcom1"
)

func TestNative_Clone(t *testing.T) {
	tests := []struct {
		name string
		n    *Native
	}{
		{
			name: "nil native",
			n:    nil,
		},
		{
			name: "empty native",
			n:    &Native{},
		},
		{
			name: "fully populated native",
			n: &Native{
				Request: `{"native":{"layout":1,"assets":[{"id":1,"title":{"len":25}}]}}`,
				Ver:     "1.2",
				API:     []adcom1.APIFramework{1, 2},
				BAttr:   []adcom1.CreativeAttribute{1, 2},
				Ext:     json.RawMessage(`{"key": "value"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.n.Clone()

			if tt.n == nil {
				if clone != nil {
					t.Error("Clone() should return nil for nil Native")
				}
				return
			}

			// Test primitive fields
			if tt.n.Request != clone.Request {
				t.Errorf("Clone() Request = %v, want %v", clone.Request, tt.n.Request)
			}
			if tt.n.Ver != clone.Ver {
				t.Errorf("Clone() Ver = %v, want %v", clone.Ver, tt.n.Ver)
			}

			// Test slices
			if len(tt.n.API) > 0 {
				original := tt.n.API[0]
				tt.n.API[0] = 999
				if clone.API[0] != original {
					t.Error("Clone() should create deep copy of API slice")
				}
			}

			if len(tt.n.BAttr) > 0 {
				original := tt.n.BAttr[0]
				tt.n.BAttr[0] = 999
				if clone.BAttr[0] != original {
					t.Error("Clone() should create deep copy of BAttr slice")
				}
			}

			// Test deep copy of extension
			if tt.n.Ext != nil {
				orig := string(tt.n.Ext)
				tt.n.Ext[0] = 'X'
				clonedStr := string(clone.Ext)
				if clonedStr != orig {
					t.Errorf("Clone() Ext not properly copied\nwant: %q\ngot: %q", orig, clonedStr)
				}
			}
		})
	}
}
