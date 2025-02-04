package legpreallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties2"
)

//NoLegAllocs is a repeating group in LegPreAllocGrp
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegAllocSettlCurrency is a non-required field for NoLegAllocs.
	LegAllocSettlCurrency *string `fix:"1367"`
	//NestedParties2 Component
	nestedparties2.NestedParties2
}

//LegPreAllocGrp is a fix50sp2 Component
type LegPreAllocGrp struct {
	//NoLegAllocs is a non-required field for LegPreAllocGrp.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
}

func (m *LegPreAllocGrp) SetNoLegAllocs(v []NoLegAllocs) { m.NoLegAllocs = v }
