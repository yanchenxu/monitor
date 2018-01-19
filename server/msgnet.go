package server

import (
	"github.com/bocheninc/base/log"
	"github.com/bocheninc/base/utils"
	msgnet "github.com/bocheninc/msg-net/peer"
)

const (
	nodeStatusMsg = 107
)

// Message represents the message transfer in msg-net
type Message struct {
	Cmd     uint8
	Payload []byte
}

// Serialize message to bytes
func (m *Message) Serialize() []byte {
	return utils.Serialize(*m)
}

// Deserialize bytes to message
func (m *Message) Deserialize(data []byte) {
	utils.Deserialize(data, m)
}

// MsgHandler handles the message of the msg-net
type MsgHandler func(src string, dst string, payload, sig []byte) error

// NewMsgnet start client msg-net service and returns a msg-net peer
func NewMsgnet(id string, routeAddress []string, fn MsgHandler, logOutPath string) *msgnet.Peer {
	// msg-net services
	if len(routeAddress) > 0 {
		msgnet.SetLogOut(logOutPath)
		msgnetPeer := msgnet.NewPeer(id, routeAddress, fn)
		msgnetPeer.Start()
		log.Debug("Msg-net Service Start ...")
		return msgnetPeer
	}
	return nil
}
