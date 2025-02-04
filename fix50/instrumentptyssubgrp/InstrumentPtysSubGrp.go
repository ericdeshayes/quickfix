package instrumentptyssubgrp

//NoInstrumentPartySubIDs is a repeating group in InstrumentPtysSubGrp
type NoInstrumentPartySubIDs struct {
	//InstrumentPartySubID is a non-required field for NoInstrumentPartySubIDs.
	InstrumentPartySubID *string `fix:"1053"`
	//InstrumentPartySubIDType is a non-required field for NoInstrumentPartySubIDs.
	InstrumentPartySubIDType *int `fix:"1054"`
}

//InstrumentPtysSubGrp is a fix50 Component
type InstrumentPtysSubGrp struct {
	//NoInstrumentPartySubIDs is a non-required field for InstrumentPtysSubGrp.
	NoInstrumentPartySubIDs []NoInstrumentPartySubIDs `fix:"1052,omitempty"`
}

func (m *InstrumentPtysSubGrp) SetNoInstrumentPartySubIDs(v []NoInstrumentPartySubIDs) {
	m.NoInstrumentPartySubIDs = v
}
