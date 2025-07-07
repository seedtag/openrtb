package openrtb2

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCloneRawMessage(t *testing.T) {
	tests := []struct {
		name string
		m    json.RawMessage
	}{
		{
			name: "nil message",
			m:    nil,
		},
		{
			name: "empty message",
			m:    json.RawMessage{},
		},
		{
			name: "valid message",
			m:    json.RawMessage(`{"key": "value"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := cloneRawMessage(tt.m)

			// For nil input, expect nil output
			if tt.m == nil {
				if clone != nil {
					t.Errorf("cloneRawMessage() = %v, want nil", clone)
				}
				return
			}

			// Verify initial equality
			if !reflect.DeepEqual(clone, tt.m) {
				t.Errorf("cloneRawMessage() = %v, want %v", clone, tt.m)
			}

			// Skip modification test for empty message
			if len(tt.m) == 0 {
				return
			}

			// Store original state and verify deep copy
			orig := make(json.RawMessage, len(tt.m))
			copy(orig, tt.m)
			tt.m[0] = 'X'
			if !reflect.DeepEqual(clone, orig) {
				t.Errorf("clone was modified when original changed, clone = %v, want %v", clone, orig)
			}
		})
	}
}

func TestCloneInt8Ptr(t *testing.T) {
	var (
		nilPtr *int8
		val    int8 = 42
	)

	tests := []struct {
		name string
		i    *int8
	}{
		{
			name: "nil pointer",
			i:    nilPtr,
		},
		{
			name: "valid pointer",
			i:    &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := cloneInt8Ptr(tt.i)
			if tt.i == nil {
				if clone != nil {
					t.Error("cloneInt8Ptr() should return nil for nil input")
				}
				return
			}

			if *clone != *tt.i {
				t.Errorf("cloneInt8Ptr() = %v, want %v", *clone, *tt.i)
			}

			// Verify modification of original doesn't affect clone
			orig := *tt.i
			*tt.i = 99
			if *clone != orig {
				t.Error("clone should not be affected by modifications to original")
			}
		})
	}
}

func TestCloneInt64Ptr(t *testing.T) {
	var (
		nilPtr *int64
		val    int64 = 42
	)

	tests := []struct {
		name string
		i    *int64
	}{
		{
			name: "nil pointer",
			i:    nilPtr,
		},
		{
			name: "valid pointer",
			i:    &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := cloneInt64Ptr(tt.i)
			if tt.i == nil {
				if clone != nil {
					t.Error("cloneInt64Ptr() should return nil for nil input")
				}
				return
			}

			if *clone != *tt.i {
				t.Errorf("cloneInt64Ptr() = %v, want %v", *clone, *tt.i)
			}

			// Verify modification of original doesn't affect clone
			orig := *tt.i
			*tt.i = 99
			if *clone != orig {
				t.Error("clone should not be affected by modifications to original")
			}
		})
	}
}

func TestCloneFloat64Ptr(t *testing.T) {
	var (
		nilPtr *float64
		val    float64 = 42.5
	)

	tests := []struct {
		name string
		f    *float64
	}{
		{
			name: "nil pointer",
			f:    nilPtr,
		},
		{
			name: "valid pointer",
			f:    &val,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := cloneFloat64Ptr(tt.f)
			if tt.f == nil {
				if clone != nil {
					t.Error("cloneFloat64Ptr() should return nil for nil input")
				}
				return
			}

			if *clone != *tt.f {
				t.Errorf("cloneFloat64Ptr() = %v, want %v", *clone, *tt.f)
			}

			// Verify modification of original doesn't affect clone
			orig := *tt.f
			*tt.f = 99.9
			if *clone != orig {
				t.Error("clone should not be affected by modifications to original")
			}
		})
	}
}

func TestCloneStringSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []string
	}{
		{
			name: "nil slice",
			s:    nil,
		},
		{
			name: "empty slice",
			s:    []string{},
		},
		{
			name: "slice with values",
			s:    []string{"one", "two", "three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := cloneStringSlice(tt.s)
			if !reflect.DeepEqual(clone, tt.s) {
				t.Errorf("cloneStringSlice() = %v, want %v", clone, tt.s)
			}

			if tt.s != nil {
				// Store original state
				orig := make([]string, len(tt.s))
				copy(orig, tt.s)

				// Modify original
				if len(tt.s) > 0 {
					tt.s[0] = "modified"
				}

				// Verify clone wasn't affected
				if !reflect.DeepEqual(clone, orig) {
					t.Error("clone should not be affected by modifications to original")
				}
			}
		})
	}
}
