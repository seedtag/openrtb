package openrtb2

import "encoding/json"

// Object: Refresh
type Refresh struct {

	// Attribute:
	//   refsettings
	// Type:
	//   object array; recommended
	// Description:
	//   A RefSettings object (see Section 3.2.34) describing the mechanics
	//   of how an ad placement automatically refreshes.
	RefSettings []RefSettings `json:"refsettings,omitempty"`

	// Attribute:
	//   count
	// Type:
	//   integer; recommended
	// Description:
	//   The number of times this ad slot had been refreshed since last page load.
	Count *int `json:"count,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Clone returns a deep copy of the Refresh object.
func (r *Refresh) Clone() *Refresh {
	if r == nil {
		return nil
	}

	clone := *r

	// Deep copy RefSettings
	if r.RefSettings != nil {
		clone.RefSettings = make([]RefSettings, len(r.RefSettings))
		for i := range r.RefSettings {
			clone.RefSettings[i] = *r.RefSettings[i].Clone()
		}
	}

	// Deep copy Count
	if r.Count != nil {
		count := *r.Count
		clone.Count = &count
	}

	// Deep copy ext
	clone.Ext = cloneRawMessage(r.Ext)

	return &clone
}
