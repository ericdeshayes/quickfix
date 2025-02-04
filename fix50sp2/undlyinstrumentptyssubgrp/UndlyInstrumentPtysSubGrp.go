package undlyinstrumentptyssubgrp

//NoUndlyInstrumentPartySubIDs is a repeating group in UndlyInstrumentPtysSubGrp
type NoUndlyInstrumentPartySubIDs struct {
	//UnderlyingInstrumentPartySubID is a non-required field for NoUndlyInstrumentPartySubIDs.
	UnderlyingInstrumentPartySubID *string `fix:"1063"`
	//UnderlyingInstrumentPartySubIDType is a non-required field for NoUndlyInstrumentPartySubIDs.
	UnderlyingInstrumentPartySubIDType *int `fix:"1064"`
}

//UndlyInstrumentPtysSubGrp is a fix50sp2 Component
type UndlyInstrumentPtysSubGrp struct {
	//NoUndlyInstrumentPartySubIDs is a non-required field for UndlyInstrumentPtysSubGrp.
	NoUndlyInstrumentPartySubIDs []NoUndlyInstrumentPartySubIDs `fix:"1062,omitempty"`
}

func (m *UndlyInstrumentPtysSubGrp) SetNoUndlyInstrumentPartySubIDs(v []NoUndlyInstrumentPartySubIDs) {
	m.NoUndlyInstrumentPartySubIDs = v
}
