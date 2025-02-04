//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoMsgTypes is a repeating group in Logon
type NoMsgTypes struct {
	//RefMsgType is a non-required field for NoMsgTypes.
	RefMsgType *string `fix:"372"`
	//MsgDirection is a non-required field for NoMsgTypes.
	MsgDirection *string `fix:"385"`
}

func (m *NoMsgTypes) SetRefMsgType(v string)   { m.RefMsgType = &v }
func (m *NoMsgTypes) SetMsgDirection(v string) { m.MsgDirection = &v }

//Message is a Logon FIX Message
type Message struct {
	FIXMsgType string `fix:"A"`
	fix42.Header
	//EncryptMethod is a required field for Logon.
	EncryptMethod int `fix:"98"`
	//HeartBtInt is a required field for Logon.
	HeartBtInt int `fix:"108"`
	//RawDataLength is a non-required field for Logon.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Logon.
	RawData *string `fix:"96"`
	//ResetSeqNumFlag is a non-required field for Logon.
	ResetSeqNumFlag *bool `fix:"141"`
	//MaxMessageSize is a non-required field for Logon.
	MaxMessageSize *int `fix:"383"`
	//NoMsgTypes is a non-required field for Logon.
	NoMsgTypes []NoMsgTypes `fix:"384,omitempty"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEncryptMethod(v int)       { m.EncryptMethod = v }
func (m *Message) SetHeartBtInt(v int)          { m.HeartBtInt = v }
func (m *Message) SetRawDataLength(v int)       { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)          { m.RawData = &v }
func (m *Message) SetResetSeqNumFlag(v bool)    { m.ResetSeqNumFlag = &v }
func (m *Message) SetMaxMessageSize(v int)      { m.MaxMessageSize = &v }
func (m *Message) SetNoMsgTypes(v []NoMsgTypes) { m.NoMsgTypes = v }

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
	return enum.BeginStringFIX42, "A", r
}
