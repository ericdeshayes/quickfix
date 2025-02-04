//Package ioi msg type = 6.
package ioi

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtlegioigrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/ioiqualgrp"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/routinggrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a IOI FIX Message
type Message struct {
	FIXMsgType string `fix:"6"`
	fixt11.Header
	//IOIID is a required field for IOI.
	IOIID string `fix:"23"`
	//IOITransType is a required field for IOI.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IOI.
	IOIRefID *string `fix:"26"`
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//Side is a required field for IOI.
	Side string `fix:"54"`
	//QtyType is a non-required field for IOI.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//IOIQty is a required field for IOI.
	IOIQty string `fix:"27"`
	//Currency is a non-required field for IOI.
	Currency *string `fix:"15"`
	//Stipulations Component
	stipulations.Stipulations
	//InstrmtLegIOIGrp Component
	instrmtlegioigrp.InstrmtLegIOIGrp
	//PriceType is a non-required field for IOI.
	PriceType *int `fix:"423"`
	//Price is a non-required field for IOI.
	Price *float64 `fix:"44"`
	//ValidUntilTime is a non-required field for IOI.
	ValidUntilTime *time.Time `fix:"62"`
	//IOIQltyInd is a non-required field for IOI.
	IOIQltyInd *string `fix:"25"`
	//IOINaturalFlag is a non-required field for IOI.
	IOINaturalFlag *bool `fix:"130"`
	//IOIQualGrp Component
	ioiqualgrp.IOIQualGrp
	//Text is a non-required field for IOI.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for IOI.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for IOI.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for IOI.
	TransactTime *time.Time `fix:"60"`
	//URLLink is a non-required field for IOI.
	URLLink *string `fix:"149"`
	//RoutingGrp Component
	routinggrp.RoutingGrp
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//Parties Component
	parties.Parties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetIOIID(v string)             { m.IOIID = v }
func (m *Message) SetIOITransType(v string)      { m.IOITransType = v }
func (m *Message) SetIOIRefID(v string)          { m.IOIRefID = &v }
func (m *Message) SetSide(v string)              { m.Side = v }
func (m *Message) SetQtyType(v int)              { m.QtyType = &v }
func (m *Message) SetIOIQty(v string)            { m.IOIQty = v }
func (m *Message) SetCurrency(v string)          { m.Currency = &v }
func (m *Message) SetPriceType(v int)            { m.PriceType = &v }
func (m *Message) SetPrice(v float64)            { m.Price = &v }
func (m *Message) SetValidUntilTime(v time.Time) { m.ValidUntilTime = &v }
func (m *Message) SetIOIQltyInd(v string)        { m.IOIQltyInd = &v }
func (m *Message) SetIOINaturalFlag(v bool)      { m.IOINaturalFlag = &v }
func (m *Message) SetText(v string)              { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)       { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)       { m.EncodedText = &v }
func (m *Message) SetTransactTime(v time.Time)   { m.TransactTime = &v }
func (m *Message) SetURLLink(v string)           { m.URLLink = &v }

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
	return enum.BeginStringFIX50, "6", r
}
