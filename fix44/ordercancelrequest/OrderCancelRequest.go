//Package ordercancelrequest msg type = F.
package ordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoUnderlyings is a repeating group in OrderCancelRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//Message is a OrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"F"`
	fix44.Header
	//OrigClOrdID is a required field for OrderCancelRequest.
	OrigClOrdID string `fix:"41"`
	//OrderID is a non-required field for OrderCancelRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderCancelRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderCancelRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for OrderCancelRequest.
	ClOrdLinkID *string `fix:"583"`
	//ListID is a non-required field for OrderCancelRequest.
	ListID *string `fix:"66"`
	//OrigOrdModTime is a non-required field for OrderCancelRequest.
	OrigOrdModTime *time.Time `fix:"586"`
	//Account is a non-required field for OrderCancelRequest.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for OrderCancelRequest.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for OrderCancelRequest.
	AccountType *int `fix:"581"`
	//Parties Component
	parties.Parties
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for OrderCancelRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Side is a required field for OrderCancelRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelRequest.
	TransactTime time.Time `fix:"60"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//ComplianceID is a non-required field for OrderCancelRequest.
	ComplianceID *string `fix:"376"`
	//Text is a non-required field for OrderCancelRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelRequest.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrigClOrdID(v string)            { m.OrigClOrdID = v }
func (m *Message) SetOrderID(v string)                { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)                { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)       { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)            { m.ClOrdLinkID = &v }
func (m *Message) SetListID(v string)                 { m.ListID = &v }
func (m *Message) SetOrigOrdModTime(v time.Time)      { m.OrigOrdModTime = &v }
func (m *Message) SetAccount(v string)                { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)              { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)               { m.AccountType = &v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *Message) SetSide(v string)                   { m.Side = v }
func (m *Message) SetTransactTime(v time.Time)        { m.TransactTime = v }
func (m *Message) SetComplianceID(v string)           { m.ComplianceID = &v }
func (m *Message) SetText(v string)                   { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)            { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "F", r
}
