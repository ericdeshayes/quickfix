//Package allocationinstruction msg type = J.
package allocationinstruction

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoOrders is a repeating group in AllocationInstruction
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//NestedParties2 Component
	nestedparties2.NestedParties2
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//OrderAvgPx is a non-required field for NoOrders.
	OrderAvgPx *float64 `fix:"799"`
	//OrderBookingQty is a non-required field for NoOrders.
	OrderBookingQty *float64 `fix:"800"`
}

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }
func (m *NoOrders) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListID(v string)           { m.ListID = &v }
func (m *NoOrders) SetOrderQty(v float64)        { m.OrderQty = &v }
func (m *NoOrders) SetOrderAvgPx(v float64)      { m.OrderAvgPx = &v }
func (m *NoOrders) SetOrderBookingQty(v float64) { m.OrderBookingQty = &v }

//NoExecs is a repeating group in AllocationInstruction
type NoExecs struct {
	//LastQty is a non-required field for NoExecs.
	LastQty *float64 `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for NoExecs.
	SecondaryExecID *string `fix:"527"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastParPx is a non-required field for NoExecs.
	LastParPx *float64 `fix:"669"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
}

func (m *NoExecs) SetLastQty(v float64)        { m.LastQty = &v }
func (m *NoExecs) SetExecID(v string)          { m.ExecID = &v }
func (m *NoExecs) SetSecondaryExecID(v string) { m.SecondaryExecID = &v }
func (m *NoExecs) SetLastPx(v float64)         { m.LastPx = &v }
func (m *NoExecs) SetLastParPx(v float64)      { m.LastParPx = &v }
func (m *NoExecs) SetLastCapacity(v string)    { m.LastCapacity = &v }

//NoUnderlyings is a repeating group in AllocationInstruction
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in AllocationInstruction
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//NoAllocs is a repeating group in AllocationInstruction
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//MatchStatus is a non-required field for NoAllocs.
	MatchStatus *string `fix:"573"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//NestedParties Component
	nestedparties.NestedParties
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *bool `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//CommissionData Component
	commissiondata.CommissionData
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//AllocSettlCurrAmt is a non-required field for NoAllocs.
	AllocSettlCurrAmt *float64 `fix:"737"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AllocAccruedInterestAmt is a non-required field for NoAllocs.
	AllocAccruedInterestAmt *float64 `fix:"742"`
	//AllocInterestAtMaturity is a non-required field for NoAllocs.
	AllocInterestAtMaturity *float64 `fix:"741"`
	//NoMiscFees is a non-required field for NoAllocs.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//NoClearingInstructions is a non-required field for NoAllocs.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
	//ClearingFeeIndicator is a non-required field for NoAllocs.
	ClearingFeeIndicator *string `fix:"635"`
	//AllocSettlInstType is a non-required field for NoAllocs.
	AllocSettlInstType *int `fix:"780"`
	//SettlInstructionsData Component
	settlinstructionsdata.SettlInstructionsData
}

func (m *NoAllocs) SetAllocAccount(v string)                             { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                           { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetMatchStatus(v string)                              { m.MatchStatus = &v }
func (m *NoAllocs) SetAllocPrice(v float64)                              { m.AllocPrice = &v }
func (m *NoAllocs) SetAllocQty(v float64)                                { m.AllocQty = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                        { m.IndividualAllocID = &v }
func (m *NoAllocs) SetProcessCode(v string)                              { m.ProcessCode = &v }
func (m *NoAllocs) SetNotifyBrokerOfCredit(v bool)                       { m.NotifyBrokerOfCredit = &v }
func (m *NoAllocs) SetAllocHandlInst(v int)                              { m.AllocHandlInst = &v }
func (m *NoAllocs) SetAllocText(v string)                                { m.AllocText = &v }
func (m *NoAllocs) SetEncodedAllocTextLen(v int)                         { m.EncodedAllocTextLen = &v }
func (m *NoAllocs) SetEncodedAllocText(v string)                         { m.EncodedAllocText = &v }
func (m *NoAllocs) SetAllocAvgPx(v float64)                              { m.AllocAvgPx = &v }
func (m *NoAllocs) SetAllocNetMoney(v float64)                           { m.AllocNetMoney = &v }
func (m *NoAllocs) SetSettlCurrAmt(v float64)                            { m.SettlCurrAmt = &v }
func (m *NoAllocs) SetAllocSettlCurrAmt(v float64)                       { m.AllocSettlCurrAmt = &v }
func (m *NoAllocs) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                       { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetSettlCurrFxRate(v float64)                         { m.SettlCurrFxRate = &v }
func (m *NoAllocs) SetSettlCurrFxRateCalc(v string)                      { m.SettlCurrFxRateCalc = &v }
func (m *NoAllocs) SetAllocAccruedInterestAmt(v float64)                 { m.AllocAccruedInterestAmt = &v }
func (m *NoAllocs) SetAllocInterestAtMaturity(v float64)                 { m.AllocInterestAtMaturity = &v }
func (m *NoAllocs) SetNoMiscFees(v []NoMiscFees)                         { m.NoMiscFees = v }
func (m *NoAllocs) SetNoClearingInstructions(v []NoClearingInstructions) { m.NoClearingInstructions = v }
func (m *NoAllocs) SetClearingFeeIndicator(v string)                     { m.ClearingFeeIndicator = &v }
func (m *NoAllocs) SetAllocSettlInstType(v int)                          { m.AllocSettlInstType = &v }

//NoMiscFees is a repeating group in NoAllocs
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
	//MiscFeeBasis is a non-required field for NoMiscFees.
	MiscFeeBasis *int `fix:"891"`
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }
func (m *NoMiscFees) SetMiscFeeBasis(v int)   { m.MiscFeeBasis = &v }

//NoClearingInstructions is a repeating group in NoAllocs
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

func (m *NoClearingInstructions) SetClearingInstruction(v int) { m.ClearingInstruction = &v }

//Message is a AllocationInstruction FIX Message
type Message struct {
	FIXMsgType string `fix:"J"`
	fix44.Header
	//AllocID is a required field for AllocationInstruction.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for AllocationInstruction.
	AllocTransType string `fix:"71"`
	//AllocType is a required field for AllocationInstruction.
	AllocType int `fix:"626"`
	//SecondaryAllocID is a non-required field for AllocationInstruction.
	SecondaryAllocID *string `fix:"793"`
	//RefAllocID is a non-required field for AllocationInstruction.
	RefAllocID *string `fix:"72"`
	//AllocCancReplaceReason is a non-required field for AllocationInstruction.
	AllocCancReplaceReason *int `fix:"796"`
	//AllocIntermedReqType is a non-required field for AllocationInstruction.
	AllocIntermedReqType *int `fix:"808"`
	//AllocLinkID is a non-required field for AllocationInstruction.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for AllocationInstruction.
	AllocLinkType *int `fix:"197"`
	//BookingRefID is a non-required field for AllocationInstruction.
	BookingRefID *string `fix:"466"`
	//AllocNoOrdersType is a required field for AllocationInstruction.
	AllocNoOrdersType int `fix:"857"`
	//NoOrders is a non-required field for AllocationInstruction.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//NoExecs is a non-required field for AllocationInstruction.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//PreviouslyReported is a non-required field for AllocationInstruction.
	PreviouslyReported *bool `fix:"570"`
	//ReversalIndicator is a non-required field for AllocationInstruction.
	ReversalIndicator *bool `fix:"700"`
	//MatchType is a non-required field for AllocationInstruction.
	MatchType *string `fix:"574"`
	//Side is a required field for AllocationInstruction.
	Side string `fix:"54"`
	//Instrument Component
	instrument.Instrument
	//InstrumentExtension Component
	instrumentextension.InstrumentExtension
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for AllocationInstruction.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for AllocationInstruction.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//Quantity is a required field for AllocationInstruction.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for AllocationInstruction.
	QtyType *int `fix:"854"`
	//LastMkt is a non-required field for AllocationInstruction.
	LastMkt *string `fix:"30"`
	//TradeOriginationDate is a non-required field for AllocationInstruction.
	TradeOriginationDate *string `fix:"229"`
	//TradingSessionID is a non-required field for AllocationInstruction.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for AllocationInstruction.
	TradingSessionSubID *string `fix:"625"`
	//PriceType is a non-required field for AllocationInstruction.
	PriceType *int `fix:"423"`
	//AvgPx is a required field for AllocationInstruction.
	AvgPx float64 `fix:"6"`
	//AvgParPx is a non-required field for AllocationInstruction.
	AvgParPx *float64 `fix:"860"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Currency is a non-required field for AllocationInstruction.
	Currency *string `fix:"15"`
	//AvgPxPrecision is a non-required field for AllocationInstruction.
	AvgPxPrecision *int `fix:"74"`
	//Parties Component
	parties.Parties
	//TradeDate is a required field for AllocationInstruction.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationInstruction.
	TransactTime *time.Time `fix:"60"`
	//SettlType is a non-required field for AllocationInstruction.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for AllocationInstruction.
	SettlDate *string `fix:"64"`
	//BookingType is a non-required field for AllocationInstruction.
	BookingType *int `fix:"775"`
	//GrossTradeAmt is a non-required field for AllocationInstruction.
	GrossTradeAmt *float64 `fix:"381"`
	//Concession is a non-required field for AllocationInstruction.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for AllocationInstruction.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for AllocationInstruction.
	NetMoney *float64 `fix:"118"`
	//PositionEffect is a non-required field for AllocationInstruction.
	PositionEffect *string `fix:"77"`
	//AutoAcceptIndicator is a non-required field for AllocationInstruction.
	AutoAcceptIndicator *bool `fix:"754"`
	//Text is a non-required field for AllocationInstruction.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationInstruction.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationInstruction.
	EncodedText *string `fix:"355"`
	//NumDaysInterest is a non-required field for AllocationInstruction.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for AllocationInstruction.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for AllocationInstruction.
	AccruedInterestAmt *float64 `fix:"159"`
	//TotalAccruedInterestAmt is a non-required field for AllocationInstruction.
	TotalAccruedInterestAmt *float64 `fix:"540"`
	//InterestAtMaturity is a non-required field for AllocationInstruction.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for AllocationInstruction.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for AllocationInstruction.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for AllocationInstruction.
	EndCash *float64 `fix:"922"`
	//LegalConfirm is a non-required field for AllocationInstruction.
	LegalConfirm *bool `fix:"650"`
	//Stipulations Component
	stipulations.Stipulations
	//YieldData Component
	yielddata.YieldData
	//TotNoAllocs is a non-required field for AllocationInstruction.
	TotNoAllocs *int `fix:"892"`
	//LastFragment is a non-required field for AllocationInstruction.
	LastFragment *bool `fix:"893"`
	//NoAllocs is a non-required field for AllocationInstruction.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAllocID(v string)                  { m.AllocID = v }
func (m *Message) SetAllocTransType(v string)           { m.AllocTransType = v }
func (m *Message) SetAllocType(v int)                   { m.AllocType = v }
func (m *Message) SetSecondaryAllocID(v string)         { m.SecondaryAllocID = &v }
func (m *Message) SetRefAllocID(v string)               { m.RefAllocID = &v }
func (m *Message) SetAllocCancReplaceReason(v int)      { m.AllocCancReplaceReason = &v }
func (m *Message) SetAllocIntermedReqType(v int)        { m.AllocIntermedReqType = &v }
func (m *Message) SetAllocLinkID(v string)              { m.AllocLinkID = &v }
func (m *Message) SetAllocLinkType(v int)               { m.AllocLinkType = &v }
func (m *Message) SetBookingRefID(v string)             { m.BookingRefID = &v }
func (m *Message) SetAllocNoOrdersType(v int)           { m.AllocNoOrdersType = v }
func (m *Message) SetNoOrders(v []NoOrders)             { m.NoOrders = v }
func (m *Message) SetNoExecs(v []NoExecs)               { m.NoExecs = v }
func (m *Message) SetPreviouslyReported(v bool)         { m.PreviouslyReported = &v }
func (m *Message) SetReversalIndicator(v bool)          { m.ReversalIndicator = &v }
func (m *Message) SetMatchType(v string)                { m.MatchType = &v }
func (m *Message) SetSide(v string)                     { m.Side = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)   { m.NoUnderlyings = v }
func (m *Message) SetNoLegs(v []NoLegs)                 { m.NoLegs = v }
func (m *Message) SetQuantity(v float64)                { m.Quantity = v }
func (m *Message) SetQtyType(v int)                     { m.QtyType = &v }
func (m *Message) SetLastMkt(v string)                  { m.LastMkt = &v }
func (m *Message) SetTradeOriginationDate(v string)     { m.TradeOriginationDate = &v }
func (m *Message) SetTradingSessionID(v string)         { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)      { m.TradingSessionSubID = &v }
func (m *Message) SetPriceType(v int)                   { m.PriceType = &v }
func (m *Message) SetAvgPx(v float64)                   { m.AvgPx = v }
func (m *Message) SetAvgParPx(v float64)                { m.AvgParPx = &v }
func (m *Message) SetCurrency(v string)                 { m.Currency = &v }
func (m *Message) SetAvgPxPrecision(v int)              { m.AvgPxPrecision = &v }
func (m *Message) SetTradeDate(v string)                { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)          { m.TransactTime = &v }
func (m *Message) SetSettlType(v string)                { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                { m.SettlDate = &v }
func (m *Message) SetBookingType(v int)                 { m.BookingType = &v }
func (m *Message) SetGrossTradeAmt(v float64)           { m.GrossTradeAmt = &v }
func (m *Message) SetConcession(v float64)              { m.Concession = &v }
func (m *Message) SetTotalTakedown(v float64)           { m.TotalTakedown = &v }
func (m *Message) SetNetMoney(v float64)                { m.NetMoney = &v }
func (m *Message) SetPositionEffect(v string)           { m.PositionEffect = &v }
func (m *Message) SetAutoAcceptIndicator(v bool)        { m.AutoAcceptIndicator = &v }
func (m *Message) SetText(v string)                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)              { m.EncodedText = &v }
func (m *Message) SetNumDaysInterest(v int)             { m.NumDaysInterest = &v }
func (m *Message) SetAccruedInterestRate(v float64)     { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)      { m.AccruedInterestAmt = &v }
func (m *Message) SetTotalAccruedInterestAmt(v float64) { m.TotalAccruedInterestAmt = &v }
func (m *Message) SetInterestAtMaturity(v float64)      { m.InterestAtMaturity = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)   { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)               { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                 { m.EndCash = &v }
func (m *Message) SetLegalConfirm(v bool)               { m.LegalConfirm = &v }
func (m *Message) SetTotNoAllocs(v int)                 { m.TotNoAllocs = &v }
func (m *Message) SetLastFragment(v bool)               { m.LastFragment = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)             { m.NoAllocs = v }

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
	return enum.BeginStringFIX44, "J", r
}
