//Package crossordercancelrequest msg type = u.
package crossordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/rootparties"
	"github.com/quickfixgo/quickfix/fix50/sidecrossordcxlgrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CrossOrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"u"`
	fixt11.Header
	//OrderID is a non-required field for CrossOrderCancelRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelRequest.
	CrossPrioritization int `fix:"550"`
	//SideCrossOrdCxlGrp Component
	sidecrossordcxlgrp.SideCrossOrdCxlGrp
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//TransactTime is a required field for CrossOrderCancelRequest.
	TransactTime time.Time `fix:"60"`
	//HostCrossID is a non-required field for CrossOrderCancelRequest.
	HostCrossID *string `fix:"961"`
	//RootParties Component
	rootparties.RootParties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = &v }
func (m *Message) SetCrossID(v string)          { m.CrossID = v }
func (m *Message) SetOrigCrossID(v string)      { m.OrigCrossID = v }
func (m *Message) SetCrossType(v int)           { m.CrossType = v }
func (m *Message) SetCrossPrioritization(v int) { m.CrossPrioritization = v }
func (m *Message) SetTransactTime(v time.Time)  { m.TransactTime = v }
func (m *Message) SetHostCrossID(v string)      { m.HostCrossID = &v }

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
	return enum.BeginStringFIX50, "u", r
}
