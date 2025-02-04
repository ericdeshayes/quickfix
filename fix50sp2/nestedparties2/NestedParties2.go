package nestedparties2

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nstdptys2subgrp"
)

//NoNested2PartyIDs is a repeating group in NestedParties2
type NoNested2PartyIDs struct {
	//Nested2PartyID is a non-required field for NoNested2PartyIDs.
	Nested2PartyID *string `fix:"757"`
	//Nested2PartyIDSource is a non-required field for NoNested2PartyIDs.
	Nested2PartyIDSource *string `fix:"758"`
	//Nested2PartyRole is a non-required field for NoNested2PartyIDs.
	Nested2PartyRole *int `fix:"759"`
	//NstdPtys2SubGrp Component
	nstdptys2subgrp.NstdPtys2SubGrp
}

//NestedParties2 is a fix50sp2 Component
type NestedParties2 struct {
	//NoNested2PartyIDs is a non-required field for NestedParties2.
	NoNested2PartyIDs []NoNested2PartyIDs `fix:"756,omitempty"`
}

func (m *NestedParties2) SetNoNested2PartyIDs(v []NoNested2PartyIDs) { m.NoNested2PartyIDs = v }
