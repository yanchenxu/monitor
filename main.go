package main

import (
	"io/ioutil"
	"os"

	"github.com/bocheninc/base/log"
	"github.com/bocheninc/monitor/server"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 2 {
		log.Infoln("must ./monitor monitor.yaml")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Errorln("read agent config file error:", err)
		return
	}

	config := &server.Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Errorln("unmarshal error, error:", err)
		return
	}

	s := server.NewServer(config)
	s.Start()
}
