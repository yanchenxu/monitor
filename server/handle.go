package server

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

func GetServerInfo() (*ServerInfo, error) {
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
		IP:   GetLocalIP(),
	}, nil
}

func GetServerStatus() (*ServerStatus, error) {
	c, err := cpu.Percent(2*time.Second, false)
	if err != nil {
		return nil, err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	ds, err := disk.IOCounters("/dev/sda1")
	if err != nil {
		return nil, err
	}

	n, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	return &ServerStatus{
		CPUUse:     int(c[0]),
		DiskUse:    int(d.UsedPercent),
		ReadBytes:  ds["sda1"].ReadBytes,
		ReadCount:  ds["sda1"].ReadCount,
		RecvBytes:  n[0].BytesRecv,
		SentBytes:  n[0].BytesSent,
		WriteBytes: ds["sda1"].WriteBytes,
		WriteCount: ds["sda1"].WriteCount,
	}, nil
}
