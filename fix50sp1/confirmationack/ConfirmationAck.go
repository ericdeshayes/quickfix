//Package confirmationack msg type = AU.
package confirmationack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ConfirmationAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AU"`
	fixt11.Header
	//ConfirmID is a required field for ConfirmationAck.
	ConfirmID string `fix:"664"`
	//TradeDate is a required field for ConfirmationAck.
	TradeDate string `fix:"75"`
	//TransactTime is a required field for ConfirmationAck.
	TransactTime time.Time `fix:"60"`
	//AffirmStatus is a required field for ConfirmationAck.
	AffirmStatus int `fix:"940"`
	//ConfirmRejReason is a non-required field for ConfirmationAck.
	ConfirmRejReason *int `fix:"774"`
	//MatchStatus is a non-required field for ConfirmationAck.
	MatchStatus *string `fix:"573"`
	//Text is a non-required field for ConfirmationAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ConfirmationAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ConfirmationAck.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetConfirmID(v string)       { m.ConfirmID = v }
func (m *Message) SetTradeDate(v string)       { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = v }
func (m *Message) SetAffirmStatus(v int)       { m.AffirmStatus = v }
func (m *Message) SetConfirmRejReason(v int)   { m.ConfirmRejReason = &v }
func (m *Message) SetMatchStatus(v string)     { m.MatchStatus = &v }
func (m *Message) SetText(v string)            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)     { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP1, "AU", r
}
