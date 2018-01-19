package server

import (
	"time"
)

var cfg *Config

type Config struct {
	ID            string        `yaml:"id"`
	MsgnetURL     []string      `yaml:"msgneturl"`
	ReportTimeDur time.Duration `yaml:"reporttimedur"`
}
