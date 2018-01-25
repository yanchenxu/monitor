package server

type ServerInfo struct {
	CPU  int    `json:"cpu"`
	Disk int    `json:"disk"`
	IP   string `json:"ip"`
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
