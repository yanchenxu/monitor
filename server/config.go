package server

import (
	"time"
)

var cfg *Config

type Config struct {
	ID            string        `yaml:"id"`
	IP            string        `yaml:"ip"`
	MsgnetURL     []string      `yaml:"msgneturl"`
	ReportTimeDur time.Duration `yaml:"reporttimedur"`
}
