package derivativeeventsgrp

import (
	"time"
)

//NoDerivativeEvents is a repeating group in DerivativeEventsGrp
type NoDerivativeEvents struct {
	//DerivativeEventType is a non-required field for NoDerivativeEvents.
	DerivativeEventType *int `fix:"1287"`
	//DerivativeEventDate is a non-required field for NoDerivativeEvents.
	DerivativeEventDate *string `fix:"1288"`
	//DerivativeEventTime is a non-required field for NoDerivativeEvents.
	DerivativeEventTime *time.Time `fix:"1289"`
	//DerivativeEventPx is a non-required field for NoDerivativeEvents.
	DerivativeEventPx *float64 `fix:"1290"`
	//DerivativeEventText is a non-required field for NoDerivativeEvents.
	DerivativeEventText *string `fix:"1291"`
}

//DerivativeEventsGrp is a fix50sp2 Component
type DerivativeEventsGrp struct {
	//NoDerivativeEvents is a non-required field for DerivativeEventsGrp.
	NoDerivativeEvents []NoDerivativeEvents `fix:"1286,omitempty"`
}

func (m *DerivativeEventsGrp) SetNoDerivativeEvents(v []NoDerivativeEvents) { m.NoDerivativeEvents = v }
