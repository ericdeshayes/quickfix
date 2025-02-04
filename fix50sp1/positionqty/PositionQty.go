package positionqty

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

//NoPositions is a repeating group in PositionQty
type NoPositions struct {
	//PosType is a non-required field for NoPositions.
	PosType *string `fix:"703"`
	//LongQty is a non-required field for NoPositions.
	LongQty *float64 `fix:"704"`
	//ShortQty is a non-required field for NoPositions.
	ShortQty *float64 `fix:"705"`
	//PosQtyStatus is a non-required field for NoPositions.
	PosQtyStatus *int `fix:"706"`
	//NestedParties Component
	nestedparties.NestedParties
	//QuantityDate is a non-required field for NoPositions.
	QuantityDate *string `fix:"976"`
}

//PositionQty is a fix50sp1 Component
type PositionQty struct {
	//NoPositions is a non-required field for PositionQty.
	NoPositions []NoPositions `fix:"702,omitempty"`
}

func (m *PositionQty) SetNoPositions(v []NoPositions) { m.NoPositions = v }
