package instrmtlegioigrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
)

//NoLegs is a repeating group in InstrmtLegIOIGrp
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegIOIQty is a non-required field for NoLegs.
	LegIOIQty *string `fix:"682"`
	//LegStipulations Component
	legstipulations.LegStipulations
}

//InstrmtLegIOIGrp is a fix50sp1 Component
type InstrmtLegIOIGrp struct {
	//NoLegs is a non-required field for InstrmtLegIOIGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegIOIGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
