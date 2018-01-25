package server

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

func (s *Server) getServerInfo() (*ServerInfo, error) {
	c, err := cpu.Percent(2*time.Second, false)
	if err != nil {
		return nil, err
	}
	d, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &ServerInfo{
		CPU:  int(c[0]),
		Disk: int(d.UsedPercent),
		IP:   cfg.IP,
	}, nil
}

func (s *Server) getServerStatus() error {
	c, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return err
	}

	ds, err := disk.IOCounters("/dev/sda1")
	if err != nil {
		return err
	}

	n, err := net.IOCounters(false)
	if err != nil {
		return err
	}

	f := func(arr []uint64, arg uint64) []uint64 {
		arr = append(arr, arg)
		if len(arr) > 10 {
			return arr[1:]
		}
		return arr
	}

	s.serverStatus.CPUUse = f(s.serverStatus.CPUUse, uint64(c[0]))
	s.serverStatus.DiskUse = uint64(d.UsedPercent)
	s.serverStatus.ReadBytes = f(s.serverStatus.ReadBytes, ds["sda1"].ReadBytes)
	s.serverStatus.ReadCount = f(s.serverStatus.ReadCount, ds["sda1"].ReadCount)
	s.serverStatus.RecvBytes = f(s.serverStatus.RecvBytes, n[0].BytesRecv)
	s.serverStatus.SentBytes = f(s.serverStatus.SentBytes, n[0].BytesSent)
	s.serverStatus.WriteBytes = f(s.serverStatus.WriteBytes, ds["sda1"].WriteBytes)
	s.serverStatus.WriteCount = f(s.serverStatus.WriteCount, ds["sda1"].WriteCount)
	return nil
}
