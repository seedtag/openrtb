package openrtb2

import "encoding/json"

// 3.2.27 Object: EID
//
// Extended identifiers support in the OpenRTB specification allows buyers to use audience data in real-time bidding.
// This object can contain one or more UIDs from a single source or a technology provider.
// The exchange should ensure that business agreements allow for the sending of this data.
type EID struct {

	// Attribute:
	//   source
	// Type:
	//   string
	// Description:
	//   Source or technology provider responsible for the set of
	//   included IDs. Expressed as a top-level domain.
	Source string `json:"source,omitempty"`

	// Attribute:
	//   uids
	// Type:
	//   object array
	// Description:
	//   Array of extended ID UID objects from the given source. Refer
	//   to 3.2.28 Extended Identifier UIDs.
	UIDs []UID `json:"uids,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Clone returns a deep copy of the EID object.
func (e *EID) Clone() *EID {
	if e == nil {
		return nil
	}

	clone := *e

	// Deep copy UIDs
	if e.UIDs != nil {
		clone.UIDs = make([]UID, len(e.UIDs))
		for i := range e.UIDs {
			clone.UIDs[i] = *e.UIDs[i].Clone()
		}
	}

	// Deep copy ext
	if e.Ext != nil {
		clone.Ext = make(json.RawMessage, len(e.Ext))
		copy(clone.Ext, e.Ext)
	}

	return &clone
}
