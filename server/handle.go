package server

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

var once sync.Once

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
		CPU:    int(c[0]),
		Disk:   int(d.UsedPercent),
		IP:     cfg.IP,
		Status: "运行中",
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
	// dd, err := disk.Partitions(true)
	// if err != nil {
	// 	return err
	// }

	// for _, v := range dd {
	// 	fmt.Println("----->", v, v.Mountpoint)
	// }

	ds, err := disk.IOCounters(cfg.Disk)
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

	diskName := strings.TrimPrefix(cfg.Disk, "/dev/")

	fmt.Println("--->", diskName, ds)
	//initialization tmpStatus
	if s.tmpStatus == nil {
		s.tmpStatus = &tmpStatus{
			readBytes:  ds[diskName].ReadBytes,
			readCount:  ds[diskName].ReadCount,
			recvBytes:  n[0].BytesRecv,
			sentBytes:  n[0].BytesSent,
			writeBytes: ds[diskName].WriteBytes,
			writeCount: ds[diskName].WriteCount,
		}
		s.serverStatus.Default()
		return nil
	}

	s.serverStatus.CPUUse = f(s.serverStatus.CPUUse, uint64(c[0]))
	s.serverStatus.DiskUse = uint64(d.UsedPercent)
	s.serverStatus.ReadBytes = f(s.serverStatus.ReadBytes, ds[diskName].ReadBytes-s.tmpStatus.readBytes)
	s.serverStatus.ReadCount = f(s.serverStatus.ReadCount, ds[diskName].ReadCount-s.tmpStatus.readCount)
	s.serverStatus.RecvBytes = f(s.serverStatus.RecvBytes, n[0].BytesRecv-s.tmpStatus.recvBytes)
	s.serverStatus.SentBytes = f(s.serverStatus.SentBytes, n[0].BytesSent-s.tmpStatus.sentBytes)
	s.serverStatus.WriteBytes = f(s.serverStatus.WriteBytes, ds[diskName].WriteBytes-s.tmpStatus.writeBytes)
	s.serverStatus.WriteCount = f(s.serverStatus.WriteCount, ds[diskName].WriteCount-s.tmpStatus.writeCount)

	s.tmpStatus = &tmpStatus{
		readBytes:  ds[diskName].ReadBytes,
		readCount:  ds[diskName].ReadCount,
		recvBytes:  n[0].BytesRecv,
		sentBytes:  n[0].BytesSent,
		writeBytes: ds[diskName].WriteBytes,
		writeCount: ds[diskName].WriteCount,
	}
	return nil
}
