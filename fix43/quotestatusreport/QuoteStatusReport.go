//Package quotestatusreport msg type = AI.
package quotestatusreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//Message is a QuoteStatusReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AI"`
	fix43.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusReport.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteReqID is a non-required field for QuoteStatusReport.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for QuoteStatusReport.
	QuoteID string `fix:"117"`
	//QuoteType is a non-required field for QuoteStatusReport.
	QuoteType *int `fix:"537"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for QuoteStatusReport.
	Account *string `fix:"1"`
	//AccountType is a non-required field for QuoteStatusReport.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for QuoteStatusReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteStatusReport.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	instrument.Instrument
	//BidPx is a non-required field for QuoteStatusReport.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for QuoteStatusReport.
	OfferPx *float64 `fix:"133"`
	//MktBidPx is a non-required field for QuoteStatusReport.
	MktBidPx *float64 `fix:"645"`
	//MktOfferPx is a non-required field for QuoteStatusReport.
	MktOfferPx *float64 `fix:"646"`
	//MinBidSize is a non-required field for QuoteStatusReport.
	MinBidSize *float64 `fix:"647"`
	//BidSize is a non-required field for QuoteStatusReport.
	BidSize *float64 `fix:"134"`
	//MinOfferSize is a non-required field for QuoteStatusReport.
	MinOfferSize *float64 `fix:"648"`
	//OfferSize is a non-required field for QuoteStatusReport.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for QuoteStatusReport.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for QuoteStatusReport.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for QuoteStatusReport.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for QuoteStatusReport.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for QuoteStatusReport.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for QuoteStatusReport.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for QuoteStatusReport.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for QuoteStatusReport.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for QuoteStatusReport.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for QuoteStatusReport.
	TransactTime *time.Time `fix:"60"`
	//FutSettDate is a non-required field for QuoteStatusReport.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for QuoteStatusReport.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for QuoteStatusReport.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for QuoteStatusReport.
	OrderQty2 *float64 `fix:"192"`
	//BidForwardPoints2 is a non-required field for QuoteStatusReport.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
	OfferForwardPoints2 *float64 `fix:"643"`
	//Currency is a non-required field for QuoteStatusReport.
	Currency *string `fix:"15"`
	//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
	SettlCurrFxRateCalc *string `fix:"156"`
	//Commission is a non-required field for QuoteStatusReport.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for QuoteStatusReport.
	CommType *string `fix:"13"`
	//CustOrderCapacity is a non-required field for QuoteStatusReport.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for QuoteStatusReport.
	ExDestination *string `fix:"100"`
	//QuoteStatus is a non-required field for QuoteStatusReport.
	QuoteStatus *int `fix:"297"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteStatusReqID(v string)      { m.QuoteStatusReqID = &v }
func (m *Message) SetQuoteReqID(v string)            { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)               { m.QuoteID = v }
func (m *Message) SetQuoteType(v int)                { m.QuoteType = &v }
func (m *Message) SetAccount(v string)               { m.Account = &v }
func (m *Message) SetAccountType(v int)              { m.AccountType = &v }
func (m *Message) SetTradingSessionID(v string)      { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)   { m.TradingSessionSubID = &v }
func (m *Message) SetBidPx(v float64)                { m.BidPx = &v }
func (m *Message) SetOfferPx(v float64)              { m.OfferPx = &v }
func (m *Message) SetMktBidPx(v float64)             { m.MktBidPx = &v }
func (m *Message) SetMktOfferPx(v float64)           { m.MktOfferPx = &v }
func (m *Message) SetMinBidSize(v float64)           { m.MinBidSize = &v }
func (m *Message) SetBidSize(v float64)              { m.BidSize = &v }
func (m *Message) SetMinOfferSize(v float64)         { m.MinOfferSize = &v }
func (m *Message) SetOfferSize(v float64)            { m.OfferSize = &v }
func (m *Message) SetValidUntilTime(v time.Time)     { m.ValidUntilTime = &v }
func (m *Message) SetBidSpotRate(v float64)          { m.BidSpotRate = &v }
func (m *Message) SetOfferSpotRate(v float64)        { m.OfferSpotRate = &v }
func (m *Message) SetBidForwardPoints(v float64)     { m.BidForwardPoints = &v }
func (m *Message) SetOfferForwardPoints(v float64)   { m.OfferForwardPoints = &v }
func (m *Message) SetMidPx(v float64)                { m.MidPx = &v }
func (m *Message) SetBidYield(v float64)             { m.BidYield = &v }
func (m *Message) SetMidYield(v float64)             { m.MidYield = &v }
func (m *Message) SetOfferYield(v float64)           { m.OfferYield = &v }
func (m *Message) SetTransactTime(v time.Time)       { m.TransactTime = &v }
func (m *Message) SetFutSettDate(v string)           { m.FutSettDate = &v }
func (m *Message) SetOrdType(v string)               { m.OrdType = &v }
func (m *Message) SetFutSettDate2(v string)          { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)            { m.OrderQty2 = &v }
func (m *Message) SetBidForwardPoints2(v float64)    { m.BidForwardPoints2 = &v }
func (m *Message) SetOfferForwardPoints2(v float64)  { m.OfferForwardPoints2 = &v }
func (m *Message) SetCurrency(v string)              { m.Currency = &v }
func (m *Message) SetSettlCurrBidFxRate(v float64)   { m.SettlCurrBidFxRate = &v }
func (m *Message) SetSettlCurrOfferFxRate(v float64) { m.SettlCurrOfferFxRate = &v }
func (m *Message) SetSettlCurrFxRateCalc(v string)   { m.SettlCurrFxRateCalc = &v }
func (m *Message) SetCommission(v float64)           { m.Commission = &v }
func (m *Message) SetCommType(v string)              { m.CommType = &v }
func (m *Message) SetCustOrderCapacity(v int)        { m.CustOrderCapacity = &v }
func (m *Message) SetExDestination(v string)         { m.ExDestination = &v }
func (m *Message) SetQuoteStatus(v int)              { m.QuoteStatus = &v }

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
	return enum.BeginStringFIX43, "AI", r
}
