package openrtb2

import "encoding/json"

// cloneRawMessage creates a deep copy of a json.RawMessage
func cloneRawMessage(m json.RawMessage) json.RawMessage {
	if m == nil {
		return nil
	}
	if len(m) == 0 {
		return json.RawMessage{}
	}
	// Ensure we create a completely new slice with its own backing array
	clone := make(json.RawMessage, len(m))
	copy(clone, m)
	return clone
}

// cloneInt8Ptr creates a deep copy of an *int8
func cloneInt8Ptr(i *int8) *int8 {
	if i == nil {
		return nil
	}
	clone := *i
	return &clone
}

// cloneInt64Ptr creates a deep copy of an *int64
func cloneInt64Ptr(i *int64) *int64 {
	if i == nil {
		return nil
	}
	clone := *i
	return &clone
}

// cloneFloat64Ptr creates a deep copy of a *float64
func cloneFloat64Ptr(f *float64) *float64 {
	if f == nil {
		return nil
	}
	clone := *f
	return &clone
}

// cloneStringSlice creates a deep copy of a string slice
func cloneStringSlice(s []string) []string {
	if s == nil {
		return nil
	}
	clone := make([]string, len(s))
	copy(clone, s)
	return clone
}
