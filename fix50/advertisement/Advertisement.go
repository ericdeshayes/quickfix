//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a Advertisement FIX Message
type Message struct {
	FIXMsgType string `fix:"7"`
	fixt11.Header
	//AdvId is a required field for Advertisement.
	AdvId string `fix:"2"`
	//AdvTransType is a required field for Advertisement.
	AdvTransType string `fix:"5"`
	//AdvRefID is a non-required field for Advertisement.
	AdvRefID *string `fix:"3"`
	//Instrument Component
	instrument.Instrument
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//AdvSide is a required field for Advertisement.
	AdvSide string `fix:"4"`
	//Quantity is a required field for Advertisement.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for Advertisement.
	QtyType *int `fix:"854"`
	//Price is a non-required field for Advertisement.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for Advertisement.
	Currency *string `fix:"15"`
	//TradeDate is a non-required field for Advertisement.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for Advertisement.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for Advertisement.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Advertisement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Advertisement.
	EncodedText *string `fix:"355"`
	//URLLink is a non-required field for Advertisement.
	URLLink *string `fix:"149"`
	//LastMkt is a non-required field for Advertisement.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for Advertisement.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for Advertisement.
	TradingSessionSubID *string `fix:"625"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAdvId(v string)               { m.AdvId = v }
func (m *Message) SetAdvTransType(v string)        { m.AdvTransType = v }
func (m *Message) SetAdvRefID(v string)            { m.AdvRefID = &v }
func (m *Message) SetAdvSide(v string)             { m.AdvSide = v }
func (m *Message) SetQuantity(v float64)           { m.Quantity = v }
func (m *Message) SetQtyType(v int)                { m.QtyType = &v }
func (m *Message) SetPrice(v float64)              { m.Price = &v }
func (m *Message) SetCurrency(v string)            { m.Currency = &v }
func (m *Message) SetTradeDate(v string)           { m.TradeDate = &v }
func (m *Message) SetTransactTime(v time.Time)     { m.TransactTime = &v }
func (m *Message) SetText(v string)                { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)         { m.EncodedText = &v }
func (m *Message) SetURLLink(v string)             { m.URLLink = &v }
func (m *Message) SetLastMkt(v string)             { m.LastMkt = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

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
	return enum.BeginStringFIX50, "7", r
}
