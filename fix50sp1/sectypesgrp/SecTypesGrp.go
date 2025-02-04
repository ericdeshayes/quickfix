package sectypesgrp

//NoSecurityTypes is a repeating group in SecTypesGrp
type NoSecurityTypes struct {
	//SecurityType is a non-required field for NoSecurityTypes.
	SecurityType *string `fix:"167"`
	//SecuritySubType is a non-required field for NoSecurityTypes.
	SecuritySubType *string `fix:"762"`
	//Product is a non-required field for NoSecurityTypes.
	Product *int `fix:"460"`
	//CFICode is a non-required field for NoSecurityTypes.
	CFICode *string `fix:"461"`
}

//SecTypesGrp is a fix50sp1 Component
type SecTypesGrp struct {
	//NoSecurityTypes is a non-required field for SecTypesGrp.
	NoSecurityTypes []NoSecurityTypes `fix:"558,omitempty"`
}

func (m *SecTypesGrp) SetNoSecurityTypes(v []NoSecurityTypes) { m.NoSecurityTypes = v }
