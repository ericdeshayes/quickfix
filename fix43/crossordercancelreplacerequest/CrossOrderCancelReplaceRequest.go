//Package crossordercancelreplacerequest msg type = t.
package crossordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoSides is a repeating group in CrossOrderCancelReplaceRequest
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrigClOrdID is a required field for NoSides.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for NoSides.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NoSides.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for NoSides.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties Component
	parties.Parties
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoSides.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoSides.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NoSides.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//QuantityType is a non-required field for NoSides.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//CommissionData Component
	commissiondata.CommissionData
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoSides.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoSides.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoSides.
	CoveredOrUncovered *int `fix:"203"`
	//CashMargin is a non-required field for NoSides.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoSides.
	ClearingFeeIndicator *string `fix:"635"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//SideComplianceID is a non-required field for NoSides.
	SideComplianceID *string `fix:"659"`
}

func (m *NoSides) SetSide(v string)                 { m.Side = v }
func (m *NoSides) SetOrigClOrdID(v string)          { m.OrigClOrdID = v }
func (m *NoSides) SetClOrdID(v string)              { m.ClOrdID = v }
func (m *NoSides) SetSecondaryClOrdID(v string)     { m.SecondaryClOrdID = &v }
func (m *NoSides) SetClOrdLinkID(v string)          { m.ClOrdLinkID = &v }
func (m *NoSides) SetOrigOrdModTime(v time.Time)    { m.OrigOrdModTime = &v }
func (m *NoSides) SetTradeOriginationDate(v string) { m.TradeOriginationDate = &v }
func (m *NoSides) SetAccount(v string)              { m.Account = &v }
func (m *NoSides) SetAccountType(v int)             { m.AccountType = &v }
func (m *NoSides) SetDayBookingInst(v string)       { m.DayBookingInst = &v }
func (m *NoSides) SetBookingUnit(v string)          { m.BookingUnit = &v }
func (m *NoSides) SetPreallocMethod(v string)       { m.PreallocMethod = &v }
func (m *NoSides) SetNoAllocs(v []NoAllocs)         { m.NoAllocs = v }
func (m *NoSides) SetQuantityType(v int)            { m.QuantityType = &v }
func (m *NoSides) SetOrderCapacity(v string)        { m.OrderCapacity = &v }
func (m *NoSides) SetOrderRestrictions(v string)    { m.OrderRestrictions = &v }
func (m *NoSides) SetCustOrderCapacity(v int)       { m.CustOrderCapacity = &v }
func (m *NoSides) SetForexReq(v bool)               { m.ForexReq = &v }
func (m *NoSides) SetSettlCurrency(v string)        { m.SettlCurrency = &v }
func (m *NoSides) SetText(v string)                 { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *NoSides) SetPositionEffect(v string)       { m.PositionEffect = &v }
func (m *NoSides) SetCoveredOrUncovered(v int)      { m.CoveredOrUncovered = &v }
func (m *NoSides) SetCashMargin(v string)           { m.CashMargin = &v }
func (m *NoSides) SetClearingFeeIndicator(v string) { m.ClearingFeeIndicator = &v }
func (m *NoSides) SetSolicitedFlag(v bool)          { m.SolicitedFlag = &v }
func (m *NoSides) SetSideComplianceID(v string)     { m.SideComplianceID = &v }

//NoAllocs is a repeating group in NoSides
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	nestedparties.NestedParties
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

func (m *NoAllocs) SetAllocAccount(v string)      { m.AllocAccount = &v }
func (m *NoAllocs) SetIndividualAllocID(v string) { m.IndividualAllocID = &v }
func (m *NoAllocs) SetAllocQty(v float64)         { m.AllocQty = &v }

//NoTradingSessions is a repeating group in CrossOrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a CrossOrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"t"`
	fix43.Header
	//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelReplaceRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelReplaceRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
	CrossPrioritization int `fix:"550"`
	//NoSides is a required field for CrossOrderCancelReplaceRequest.
	NoSides []NoSides `fix:"552"`
	//Instrument Component
	instrument.Instrument
	//SettlmntTyp is a non-required field for CrossOrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for CrossOrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for CrossOrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for CrossOrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	stipulations.Stipulations
	//OrdType is a required field for CrossOrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for CrossOrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//Currency is a non-required field for CrossOrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//IOIid is a non-required field for CrossOrderCancelReplaceRequest.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for CrossOrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for CrossOrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for CrossOrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for CrossOrderCancelReplaceRequest.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for CrossOrderCancelReplaceRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for CrossOrderCancelReplaceRequest.
	NetMoney *float64 `fix:"118"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)                        { m.OrderID = &v }
func (m *Message) SetCrossID(v string)                        { m.CrossID = v }
func (m *Message) SetOrigCrossID(v string)                    { m.OrigCrossID = v }
func (m *Message) SetCrossType(v int)                         { m.CrossType = v }
func (m *Message) SetCrossPrioritization(v int)               { m.CrossPrioritization = v }
func (m *Message) SetNoSides(v []NoSides)                     { m.NoSides = v }
func (m *Message) SetSettlmntTyp(v string)                    { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)                    { m.FutSettDate = &v }
func (m *Message) SetHandlInst(v string)                      { m.HandlInst = v }
func (m *Message) SetExecInst(v string)                       { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                        { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                      { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                  { m.ExDestination = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *Message) SetProcessCode(v string)                    { m.ProcessCode = &v }
func (m *Message) SetPrevClosePx(v float64)                   { m.PrevClosePx = &v }
func (m *Message) SetLocateReqd(v bool)                       { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetOrdType(v string)                        { m.OrdType = v }
func (m *Message) SetPriceType(v int)                         { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                         { m.Price = &v }
func (m *Message) SetStopPx(v float64)                        { m.StopPx = &v }
func (m *Message) SetCurrency(v string)                       { m.Currency = &v }
func (m *Message) SetComplianceID(v string)                   { m.ComplianceID = &v }
func (m *Message) SetIOIid(v string)                          { m.IOIid = &v }
func (m *Message) SetQuoteID(v string)                        { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)                    { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)               { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                     { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                  { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                     { m.GTBookingInst = &v }
func (m *Message) SetMaxShow(v float64)                       { m.MaxShow = &v }
func (m *Message) SetPegDifference(v float64)                 { m.PegDifference = &v }
func (m *Message) SetDiscretionInst(v string)                 { m.DiscretionInst = &v }
func (m *Message) SetDiscretionOffset(v float64)              { m.DiscretionOffset = &v }
func (m *Message) SetCancellationRights(v string)             { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)          { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                       { m.RegistID = &v }
func (m *Message) SetDesignation(v string)                    { m.Designation = &v }
func (m *Message) SetAccruedInterestRate(v float64)           { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)            { m.AccruedInterestAmt = &v }
func (m *Message) SetNetMoney(v float64)                      { m.NetMoney = &v }

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
	return enum.BeginStringFIX43, "t", r
}
