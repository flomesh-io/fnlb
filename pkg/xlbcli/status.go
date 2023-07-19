package xlbcli

type Status struct {
	CommonAPI
}

type ProcessGet struct {
	ProcessAttr []Process `json:"processAttr"`
}

type FilesystemGet struct {
	FilesystemAttr []Filesystem `json:"filesystemAttr"`
}

type DeviceGet struct {
	HostName     string
	MachineID    string
	BootID       string
	OS           string
	Kernel       string
	Architecture string
	Uptime       string
}

type Process struct {
	Pid          string
	User         string
	Priority     string
	Nice         string
	VirtMemory   string
	ResidentSize string
	SharedMemory string
	Status       string
	CPUUsage     string
	MemoryUsage  string
	ProcessTime  string `json:"time"`
	Command      string
}

type Filesystem struct {
	FileSystem string
	Fstype     string `json:"type"`
	Size       string
	Used       string
	Avail      string
	UsePercent string
	MountedOn  string
}
