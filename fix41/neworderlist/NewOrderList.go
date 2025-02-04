//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	fix41.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//WaveNo is a non-required field for NewOrderList.
	WaveNo *string `fix:"105"`
	//ListSeqNo is a required field for NewOrderList.
	ListSeqNo int `fix:"67"`
	//ListNoOrds is a required field for NewOrderList.
	ListNoOrds int `fix:"68"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//ClOrdID is a required field for NewOrderList.
	ClOrdID string `fix:"11"`
	//ClientID is a non-required field for NewOrderList.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for NewOrderList.
	ExecBroker *string `fix:"76"`
	//Account is a non-required field for NewOrderList.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for NewOrderList.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NewOrderList.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for NewOrderList.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for NewOrderList.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderList.
	MinQty *int `fix:"110"`
	//MaxFloor is a non-required field for NewOrderList.
	MaxFloor *int `fix:"111"`
	//ExDestination is a non-required field for NewOrderList.
	ExDestination *string `fix:"100"`
	//ProcessCode is a non-required field for NewOrderList.
	ProcessCode *string `fix:"81"`
	//Symbol is a required field for NewOrderList.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NewOrderList.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NewOrderList.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NewOrderList.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NewOrderList.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NewOrderList.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NewOrderList.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NewOrderList.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NewOrderList.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NewOrderList.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for NewOrderList.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NewOrderList.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for NewOrderList.
	SecurityDesc *string `fix:"107"`
	//PrevClosePx is a non-required field for NewOrderList.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderList.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderList.
	LocateReqd *string `fix:"114"`
	//OrderQty is a required field for NewOrderList.
	OrderQty int `fix:"38"`
	//OrdType is a required field for NewOrderList.
	OrdType string `fix:"40"`
	//Price is a non-required field for NewOrderList.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderList.
	StopPx *float64 `fix:"99"`
	//PegDifference is a non-required field for NewOrderList.
	PegDifference *float64 `fix:"211"`
	//Currency is a non-required field for NewOrderList.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for NewOrderList.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for NewOrderList.
	ExpireTime *time.Time `fix:"126"`
	//Commission is a non-required field for NewOrderList.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NewOrderList.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for NewOrderList.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NewOrderList.
	ForexReq *string `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderList.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NewOrderList.
	Text *string `fix:"58"`
	//FutSettDate2 is a non-required field for NewOrderList.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NewOrderList.
	OrderQty2 *float64 `fix:"192"`
	//OpenClose is a non-required field for NewOrderList.
	OpenClose *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderList.
	CoveredOrUncovered *int `fix:"203"`
	//CustomerOrFirm is a non-required field for NewOrderList.
	CustomerOrFirm *int `fix:"204"`
	//MaxShow is a non-required field for NewOrderList.
	MaxShow *int `fix:"210"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)            { m.ListID = v }
func (m *Message) SetWaveNo(v string)            { m.WaveNo = &v }
func (m *Message) SetListSeqNo(v int)            { m.ListSeqNo = v }
func (m *Message) SetListNoOrds(v int)           { m.ListNoOrds = v }
func (m *Message) SetListExecInst(v string)      { m.ListExecInst = &v }
func (m *Message) SetClOrdID(v string)           { m.ClOrdID = v }
func (m *Message) SetClientID(v string)          { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)        { m.ExecBroker = &v }
func (m *Message) SetAccount(v string)           { m.Account = &v }
func (m *Message) SetSettlmntTyp(v string)       { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)       { m.FutSettDate = &v }
func (m *Message) SetHandlInst(v string)         { m.HandlInst = v }
func (m *Message) SetExecInst(v string)          { m.ExecInst = &v }
func (m *Message) SetMinQty(v int)               { m.MinQty = &v }
func (m *Message) SetMaxFloor(v int)             { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)     { m.ExDestination = &v }
func (m *Message) SetProcessCode(v string)       { m.ProcessCode = &v }
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
func (m *Message) SetSide(v string)              { m.Side = v }
func (m *Message) SetLocateReqd(v string)        { m.LocateReqd = &v }
func (m *Message) SetOrderQty(v int)             { m.OrderQty = v }
func (m *Message) SetOrdType(v string)           { m.OrdType = v }
func (m *Message) SetPrice(v float64)            { m.Price = &v }
func (m *Message) SetStopPx(v float64)           { m.StopPx = &v }
func (m *Message) SetPegDifference(v float64)    { m.PegDifference = &v }
func (m *Message) SetCurrency(v string)          { m.Currency = &v }
func (m *Message) SetTimeInForce(v string)       { m.TimeInForce = &v }
func (m *Message) SetExpireTime(v time.Time)     { m.ExpireTime = &v }
func (m *Message) SetCommission(v float64)       { m.Commission = &v }
func (m *Message) SetCommType(v string)          { m.CommType = &v }
func (m *Message) SetRule80A(v string)           { m.Rule80A = &v }
func (m *Message) SetForexReq(v string)          { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)     { m.SettlCurrency = &v }
func (m *Message) SetText(v string)              { m.Text = &v }
func (m *Message) SetFutSettDate2(v string)      { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)        { m.OrderQty2 = &v }
func (m *Message) SetOpenClose(v string)         { m.OpenClose = &v }
func (m *Message) SetCoveredOrUncovered(v int)   { m.CoveredOrUncovered = &v }
func (m *Message) SetCustomerOrFirm(v int)       { m.CustomerOrFirm = &v }
func (m *Message) SetMaxShow(v int)              { m.MaxShow = &v }

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
	return enum.BeginStringFIX41, "E", r
}
