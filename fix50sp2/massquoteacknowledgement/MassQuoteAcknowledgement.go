//Package massquoteacknowledgement msg type = b.
package massquoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/quotsetackgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/targetparties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MassQuoteAcknowledgement FIX Message
type Message struct {
	FIXMsgType string `fix:"b"`
	fixt11.Header
	//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for MassQuoteAcknowledgement.
	QuoteID *string `fix:"117"`
	//QuoteStatus is a required field for MassQuoteAcknowledgement.
	QuoteStatus int `fix:"297"`
	//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
	QuoteRejectReason *int `fix:"300"`
	//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
	QuoteResponseLevel *int `fix:"301"`
	//QuoteType is a non-required field for MassQuoteAcknowledgement.
	QuoteType *int `fix:"537"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for MassQuoteAcknowledgement.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for MassQuoteAcknowledgement.
	AccountType *int `fix:"581"`
	//Text is a non-required field for MassQuoteAcknowledgement.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MassQuoteAcknowledgement.
	EncodedText *string `fix:"355"`
	//QuotSetAckGrp Component
	quotsetackgrp.QuotSetAckGrp
	//QuoteCancelType is a non-required field for MassQuoteAcknowledgement.
	QuoteCancelType *int `fix:"298"`
	//TargetParties Component
	targetparties.TargetParties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)      { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)         { m.QuoteID = &v }
func (m *Message) SetQuoteStatus(v int)        { m.QuoteStatus = v }
func (m *Message) SetQuoteRejectReason(v int)  { m.QuoteRejectReason = &v }
func (m *Message) SetQuoteResponseLevel(v int) { m.QuoteResponseLevel = &v }
func (m *Message) SetQuoteType(v int)          { m.QuoteType = &v }
func (m *Message) SetAccount(v string)         { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)       { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)        { m.AccountType = &v }
func (m *Message) SetText(v string)            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)     { m.EncodedText = &v }
func (m *Message) SetQuoteCancelType(v int)    { m.QuoteCancelType = &v }

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
	return enum.ApplVerID_FIX50SP2, "b", r
}
