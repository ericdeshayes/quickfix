//Package derivativesecuritylistupdatereport msg type = BR.
package derivativesecuritylistupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativesecuritydefinition"
	"github.com/quickfixgo/quickfix/fix50sp1/relsymderivsecupdgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a DerivativeSecurityListUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BR"`
	fixt11.Header
	//SecurityReqID is a non-required field for DerivativeSecurityListUpdateReport.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for DerivativeSecurityListUpdateReport.
	SecurityResponseID *string `fix:"322"`
	//SecurityRequestResult is a non-required field for DerivativeSecurityListUpdateReport.
	SecurityRequestResult *int `fix:"560"`
	//SecurityUpdateAction is a non-required field for DerivativeSecurityListUpdateReport.
	SecurityUpdateAction *string `fix:"980"`
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//DerivativeSecurityDefinition Component
	derivativesecuritydefinition.DerivativeSecurityDefinition
	//TotNoRelatedSym is a non-required field for DerivativeSecurityListUpdateReport.
	TotNoRelatedSym *int `fix:"393"`
	//LastFragment is a non-required field for DerivativeSecurityListUpdateReport.
	LastFragment *bool `fix:"893"`
	//RelSymDerivSecUpdGrp Component
	relsymderivsecupdgrp.RelSymDerivSecUpdGrp
	//ApplicationSequenceControl Component
	applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)   { m.SecurityResponseID = &v }
func (m *Message) SetSecurityRequestResult(v int)   { m.SecurityRequestResult = &v }
func (m *Message) SetSecurityUpdateAction(v string) { m.SecurityUpdateAction = &v }
func (m *Message) SetTotNoRelatedSym(v int)         { m.TotNoRelatedSym = &v }
func (m *Message) SetLastFragment(v bool)           { m.LastFragment = &v }

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
	return enum.ApplVerID_FIX50SP1, "BR", r
}
