package server

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/bocheninc/base/log"
	"github.com/bocheninc/msg-net/peer"
)

const (
	monitorPrefix   = "monitor"
	msgnetRPCPrefix = "__virtual"
)

type Server struct {
	peeID        string
	peer         *peer.Peer
	ticker       *time.Ticker
	serverStatus *ServerStatus
	tmpStatus    *tmpStatus
}

func NewServer(config *Config) *Server {
	log.Infoln("config: ", *config)
	cfg = config

	s := &Server{}
	s.peeID = fmt.Sprintf("%s:%s", monitorPrefix, cfg.ID)
	s.peer = NewMsgnet(s.peeID, cfg.MsgnetURL, s.msghandle, "")
	s.ticker = time.NewTicker(config.ReportTimeDur)
	s.serverStatus = &ServerStatus{}
	return s
}

func (s *Server) Start() {
	log.Infoln("server start ...")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if err := s.getServerStatus(); err != nil {
				log.Errorln("GetServerStatus func err: ", err)
			}
		}
	}()

	for {
		select {
		case <-s.ticker.C:
			m := make(map[string]interface{})
			serverInfo, err := s.getServerInfo()
			if err != nil {
				log.Errorln("GetServerInfo func err: ", err)
			}

			m["localServer"] = serverInfo
			m["serverStatus"] = s.serverStatus

			payload, err := json.Marshal(m)
			if err != nil {
				log.Errorln(err)
			}

			log.Debugln("send to msgnet msg", string(payload))
			msg := &Message{
				Cmd:     nodeStatusMsg,
				Payload: payload,
			}

			s.peer.Send(msgnetRPCPrefix, msg.Serialize(), nil)
		}
	}
}

func (s *Server) msghandle(src string, dst string, payload, sig []byte) error {
	return nil
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
