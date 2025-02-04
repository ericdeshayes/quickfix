//Package orderstatusrequest msg type = H.
package orderstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
)

//Message is a OrderStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"H"`
	fix43.Header
	//OrderID is a non-required field for OrderStatusRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderStatusRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderStatusRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for OrderStatusRequest.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for OrderStatusRequest.
	Account *string `fix:"1"`
	//Instrument Component
	instrument.Instrument
	//Side is a required field for OrderStatusRequest.
	Side string `fix:"54"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)          { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)      { m.ClOrdLinkID = &v }
func (m *Message) SetAccount(v string)          { m.Account = &v }
func (m *Message) SetSide(v string)             { m.Side = v }

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
	return enum.BeginStringFIX43, "H", r
}
