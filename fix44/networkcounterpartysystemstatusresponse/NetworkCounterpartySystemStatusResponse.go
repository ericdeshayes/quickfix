//Package networkcounterpartysystemstatusresponse msg type = BD.
package networkcounterpartysystemstatusresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//NoCompIDs is a repeating group in NetworkCounterpartySystemStatusResponse
type NoCompIDs struct {
	//RefCompID is a non-required field for NoCompIDs.
	RefCompID *string `fix:"930"`
	//RefSubID is a non-required field for NoCompIDs.
	RefSubID *string `fix:"931"`
	//LocationID is a non-required field for NoCompIDs.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoCompIDs.
	DeskID *string `fix:"284"`
	//StatusValue is a non-required field for NoCompIDs.
	StatusValue *int `fix:"928"`
	//StatusText is a non-required field for NoCompIDs.
	StatusText *string `fix:"929"`
}

func (m *NoCompIDs) SetRefCompID(v string)  { m.RefCompID = &v }
func (m *NoCompIDs) SetRefSubID(v string)   { m.RefSubID = &v }
func (m *NoCompIDs) SetLocationID(v string) { m.LocationID = &v }
func (m *NoCompIDs) SetDeskID(v string)     { m.DeskID = &v }
func (m *NoCompIDs) SetStatusValue(v int)   { m.StatusValue = &v }
func (m *NoCompIDs) SetStatusText(v string) { m.StatusText = &v }

//Message is a NetworkCounterpartySystemStatusResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"BD"`
	fix44.Header
	//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
	NetworkStatusResponseType int `fix:"937"`
	//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
	NetworkRequestID *string `fix:"933"`
	//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
	NetworkResponseID string `fix:"932"`
	//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
	LastNetworkResponseID *string `fix:"934"`
	//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
	NoCompIDs []NoCompIDs `fix:"936"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetNetworkStatusResponseType(v int) { m.NetworkStatusResponseType = v }
func (m *Message) SetNetworkRequestID(v string)       { m.NetworkRequestID = &v }
func (m *Message) SetNetworkResponseID(v string)      { m.NetworkResponseID = v }
func (m *Message) SetLastNetworkResponseID(v string)  { m.LastNetworkResponseID = &v }
func (m *Message) SetNoCompIDs(v []NoCompIDs)         { m.NoCompIDs = v }

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
	return enum.BeginStringFIX44, "BD", r
}
