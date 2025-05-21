package openrtb2

import "encoding/json"

// 3.2.11 Object: Pmp
//
// This object is the private marketplace container for direct deals between buyers and sellers that may pertain to this impression.
// The actual deals are represented as a collection of Deal objects.
// Refer to Section 7.3 for more details.
type PMP struct {

	// Attribute:
	//   private_auction
	// Type:
	//   integer; default 0
	// Description:
	//   Indicator of auction eligibility to seats named in the Direct
	//   Deals object, where 0 = all bids are accepted, 1 = bids are
	//   restricted to the deals specified and the terms thereof.
	PrivateAuction int8 `json:"private_auction,omitempty"`

	// Attribute:
	//   deals
	// Type:
	//   object array
	// Description:
	//   Array of Deal (Section 3.2.12) objects that convey the specific
	//   deals applicable to this impression.
	Deals []Deal `json:"deals,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Clone returns a deep copy of the PMP object.
func (p *PMP) Clone() *PMP {
	if p == nil {
		return nil
	}
	clone := *p

	// Deep copy deals
	if p.Deals != nil {
		clone.Deals = make([]Deal, len(p.Deals))
		for i := range p.Deals {
			clone.Deals[i] = *p.Deals[i].Clone()
		}
	}

	// Deep copy ext
	if p.Ext != nil {
		clone.Ext = make(json.RawMessage, len(p.Ext))
		copy(clone.Ext, p.Ext)
	}

	return &clone
}
