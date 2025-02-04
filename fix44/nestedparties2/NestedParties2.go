package nestedparties2

//NoNested2PartyIDs is a repeating group in NestedParties2
type NoNested2PartyIDs struct {
	//Nested2PartyID is a non-required field for NoNested2PartyIDs.
	Nested2PartyID *string `fix:"757"`
	//Nested2PartyIDSource is a non-required field for NoNested2PartyIDs.
	Nested2PartyIDSource *string `fix:"758"`
	//Nested2PartyRole is a non-required field for NoNested2PartyIDs.
	Nested2PartyRole *int `fix:"759"`
	//NoNested2PartySubIDs is a non-required field for NoNested2PartyIDs.
	NoNested2PartySubIDs []NoNested2PartySubIDs `fix:"806,omitempty"`
}

//NoNested2PartySubIDs is a repeating group in NoNested2PartyIDs
type NoNested2PartySubIDs struct {
	//Nested2PartySubID is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubID *string `fix:"760"`
	//Nested2PartySubIDType is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubIDType *int `fix:"807"`
}

//NestedParties2 is a fix44 Component
type NestedParties2 struct {
	//NoNested2PartyIDs is a non-required field for NestedParties2.
	NoNested2PartyIDs []NoNested2PartyIDs `fix:"756,omitempty"`
}

func (m *NestedParties2) SetNoNested2PartyIDs(v []NoNested2PartyIDs) { m.NoNested2PartyIDs = v }
