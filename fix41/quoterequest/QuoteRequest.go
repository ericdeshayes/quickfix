//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a QuoteRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"R"`
	fix41.Header
	//QuoteReqID is a required field for QuoteRequest.
	QuoteReqID string `fix:"131"`
	//Symbol is a required field for QuoteRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for QuoteRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for QuoteRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for QuoteRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for QuoteRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for QuoteRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for QuoteRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for QuoteRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for QuoteRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for QuoteRequest.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for QuoteRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for QuoteRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for QuoteRequest.
	SecurityDesc *string `fix:"107"`
	//PrevClosePx is a non-required field for QuoteRequest.
	PrevClosePx *float64 `fix:"140"`
	//Side is a non-required field for QuoteRequest.
	Side *string `fix:"54"`
	//OrderQty is a non-required field for QuoteRequest.
	OrderQty *int `fix:"38"`
	//FutSettDate is a non-required field for QuoteRequest.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for QuoteRequest.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for QuoteRequest.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for QuoteRequest.
	OrderQty2 *float64 `fix:"192"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)        { m.QuoteReqID = v }
func (m *Message) SetSymbol(v string)            { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)         { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)        { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)          { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)      { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string) { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)          { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)            { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)      { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)      { m.OptAttribute = &v }
func (m *Message) SetSecurityExchange(v string)  { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)            { m.Issuer = &v }
func (m *Message) SetSecurityDesc(v string)      { m.SecurityDesc = &v }
func (m *Message) SetPrevClosePx(v float64)      { m.PrevClosePx = &v }
func (m *Message) SetSide(v string)              { m.Side = &v }
func (m *Message) SetOrderQty(v int)             { m.OrderQty = &v }
func (m *Message) SetFutSettDate(v string)       { m.FutSettDate = &v }
func (m *Message) SetOrdType(v string)           { m.OrdType = &v }
func (m *Message) SetFutSettDate2(v string)      { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)        { m.OrderQty2 = &v }

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
	return enum.BeginStringFIX41, "R", r
}
