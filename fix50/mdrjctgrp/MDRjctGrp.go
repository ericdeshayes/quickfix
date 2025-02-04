package mdrjctgrp

//NoAltMDSource is a repeating group in MDRjctGrp
type NoAltMDSource struct {
	//AltMDSourceID is a non-required field for NoAltMDSource.
	AltMDSourceID *string `fix:"817"`
}

//MDRjctGrp is a fix50 Component
type MDRjctGrp struct {
	//NoAltMDSource is a non-required field for MDRjctGrp.
	NoAltMDSource []NoAltMDSource `fix:"816,omitempty"`
}

func (m *MDRjctGrp) SetNoAltMDSource(v []NoAltMDSource) { m.NoAltMDSource = v }
