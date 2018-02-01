package server

type ServerInfo struct {
	CPU    int    `json:"cpu"`
	Disk   int    `json:"disk"`
	IP     string `json:"ip"`
	Status string `json:"status"`
}

type ServerStatus struct {
	CPUUse     []uint64 `json:"cpuUse"`
	DiskUse    uint64   `json:"diskUse"`
	ReadBytes  []uint64 `json:"readBytes"`  //磁盘读取流量
	ReadCount  []uint64 `json:"readCount"`  //磁盘读次数
	RecvBytes  []uint64 `json:"recvBytes"`  //流量接受（全网络出口）
	SentBytes  []uint64 `json:"sentBytes"`  //流量发送（全网络入口）
	WriteBytes []uint64 `json:"writeBytes"` //磁盘写入流量
	WriteCount []uint64 `json:"writeCount"` //磁盘写次数
}

func (s *ServerStatus) Default() {
	s = &ServerStatus{
		CPUUse:     []uint64{0},
		DiskUse:    0,
		ReadBytes:  []uint64{0},
		ReadCount:  []uint64{0},
		RecvBytes:  []uint64{0},
		SentBytes:  []uint64{0},
		WriteBytes: []uint64{0},
		WriteCount: []uint64{0},
	}
}

type tmpStatus struct {
	readBytes  uint64
	readCount  uint64
	recvBytes  uint64
	sentBytes  uint64
	writeBytes uint64
	writeCount uint64
}
