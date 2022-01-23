package cloud

type Cloud interface {
	Name() string
	Init(string, string, string, string)
	TestConnect() error
	GetInstances() []*Instance
	StartInstance(string) error
	StopInstance(string) error
	RestartInstance(string) error
}
