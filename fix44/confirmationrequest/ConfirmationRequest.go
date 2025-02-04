//Package confirmationrequest msg type = BH.
package confirmationrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"time"
)

//NoOrders is a repeating group in ConfirmationRequest
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//NestedParties2 Component
	nestedparties2.NestedParties2
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//OrderAvgPx is a non-required field for NoOrders.
	OrderAvgPx *float64 `fix:"799"`
	//OrderBookingQty is a non-required field for NoOrders.
	OrderBookingQty *float64 `fix:"800"`
}

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }
func (m *NoOrders) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListID(v string)           { m.ListID = &v }
func (m *NoOrders) SetOrderQty(v float64)        { m.OrderQty = &v }
func (m *NoOrders) SetOrderAvgPx(v float64)      { m.OrderAvgPx = &v }
func (m *NoOrders) SetOrderBookingQty(v float64) { m.OrderBookingQty = &v }

//Message is a ConfirmationRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BH"`
	fix44.Header
	//ConfirmReqID is a required field for ConfirmationRequest.
	ConfirmReqID string `fix:"859"`
	//ConfirmType is a required field for ConfirmationRequest.
	ConfirmType int `fix:"773"`
	//NoOrders is a non-required field for ConfirmationRequest.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//AllocID is a non-required field for ConfirmationRequest.
	AllocID *string `fix:"70"`
	//SecondaryAllocID is a non-required field for ConfirmationRequest.
	SecondaryAllocID *string `fix:"793"`
	//IndividualAllocID is a non-required field for ConfirmationRequest.
	IndividualAllocID *string `fix:"467"`
	//TransactTime is a required field for ConfirmationRequest.
	TransactTime time.Time `fix:"60"`
	//AllocAccount is a non-required field for ConfirmationRequest.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for ConfirmationRequest.
	AllocAcctIDSource *int `fix:"661"`
	//AllocAccountType is a non-required field for ConfirmationRequest.
	AllocAccountType *int `fix:"798"`
	//Text is a non-required field for ConfirmationRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ConfirmationRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ConfirmationRequest.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetConfirmReqID(v string)      { m.ConfirmReqID = v }
func (m *Message) SetConfirmType(v int)          { m.ConfirmType = v }
func (m *Message) SetNoOrders(v []NoOrders)      { m.NoOrders = v }
func (m *Message) SetAllocID(v string)           { m.AllocID = &v }
func (m *Message) SetSecondaryAllocID(v string)  { m.SecondaryAllocID = &v }
func (m *Message) SetIndividualAllocID(v string) { m.IndividualAllocID = &v }
func (m *Message) SetTransactTime(v time.Time)   { m.TransactTime = v }
func (m *Message) SetAllocAccount(v string)      { m.AllocAccount = &v }
func (m *Message) SetAllocAcctIDSource(v int)    { m.AllocAcctIDSource = &v }
func (m *Message) SetAllocAccountType(v int)     { m.AllocAccountType = &v }
func (m *Message) SetText(v string)              { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)       { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)       { m.EncodedText = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX44, "BH", r
}
