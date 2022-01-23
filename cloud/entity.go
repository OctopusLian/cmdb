package cloud

const (
	Pending      = "创建中"
	LaunchFailed = "创建失败"
	Running      = "运行中"
	Stopped      = "关机"
	Starting     = "开机中"
	Stopping     = "关机中"
	Rebooting    = "重启中"
	ShutDown     = "停止销毁"
	Terminating  = "销毁中"
	Unknow       = "未知"
)

type Instance struct {
	Key          string
	UUID         string
	Name         string
	OS           string
	CPU          int
	Memory       int
	PublicAddrs  []string
	PrivateAddrs []string
	Status       string
	CreatedTime  string
	ExpiredTime  string
}
