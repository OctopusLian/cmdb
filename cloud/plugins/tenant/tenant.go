package tenant

import (
	"cmdb/cloud"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TenantCloud struct {
	Addr       string
	Key        string
	Secrect    string
	Region     string
	credential *common.Credential
	profile    *profile.ClientProfile
}

func (t *TenantCloud) Name() string {
	return "tenantyun"
}

func (t *TenantCloud) Init(addr string, key string, secrect string, region string) {
	t.Addr = addr
	t.Key = key
	t.Secrect = secrect
	t.Region = region

	t.credential = common.NewCredential(key, secrect)

	t.profile = profile.NewClientProfile()
	t.profile.HttpProfile.Endpoint = addr

}

func (t *TenantCloud) TestConnect() error {
	client, err := cvm.NewClient(t.credential, t.Region, t.profile)
	if err == nil {
		request := cvm.NewDescribeRegionsRequest()
		_, err = client.DescribeRegions(request)
	}

	return err
}

func (t *TenantCloud) statusTransform(status string) string {
	statusMap := map[string]string{
		"PENDING":       cloud.StatusPending,
		"LAUNCH_FAILED": cloud.StatusLaunchFailed,
		"RUNNING":       cloud.StatusRunning,
		"STOPPED":       cloud.StatusStopped,
		"STARTING":      cloud.StatusStarting,
		"STOPPING":      cloud.StatusStopping,
		"REBOOTING":     cloud.StatusRebooting,
		"SHUTDOWN":      cloud.StatusShutdown,
		"TERMINATING":   cloud.StatusTerminating,
	}
	if text, ok := statusMap[status]; ok {
		return text
	}
	return cloud.StatusUnknow
}

func (t *TenantCloud) GetInstances() []*cloud.Instance {
	var limit int64 = 100
	client, err := cvm.NewClient(t.credential, t.Region, t.profile)
	if err != nil {
		return nil
	}

	request := cvm.NewDescribeInstancesRequest()
	request.Limit = &limit
	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil
	}

	instances := make([]*cloud.Instance, *response.Response.TotalCount)
	for index, instance := range response.Response.InstanceSet {
		publicAddrs := make([]string, len(instance.PublicIpAddresses))
		privateAddrs := make([]string, len(instance.PrivateIpAddresses))
		for i, addr := range instance.PublicIpAddresses {
			publicAddrs[i] = *addr
		}
		for i, addr := range instance.PrivateIpAddresses {
			privateAddrs[i] = *addr
		}
		instances[index] = &cloud.Instance{
			UUID:         *instance.Uuid,
			Name:         *instance.InstanceName,
			OS:           *instance.OsName,
			CPU:          int(*instance.CPU),
			Mem:          *instance.Memory * 1024,
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       t.statusTransform(*instance.InstanceState),
			CreatedTime:  *instance.CreatedTime,
			ExpiredTime:  *instance.ExpiredTime,
		}
	}
	return instances
}

func (t *TenantCloud) StartInstance(uuid string) error {
	client, err := cvm.NewClient(t.credential, t.Region, t.profile)
	if err == nil {
		request := cvm.NewStartInstancesRequest()
		request.InstanceIds = []*string{&uuid}
		_, err = client.StartInstances(request)
	}
	return err
}

func (t *TenantCloud) StopInstance(uuid string) error {
	client, err := cvm.NewClient(t.credential, t.Region, t.profile)

	if err == nil {
		request := cvm.NewStopInstancesRequest()
		request.InstanceIds = []*string{&uuid}
		_, err = client.StopInstances(request)
	}
	return err
}

func (t *TenantCloud) RestartInstance(uuid string) error {
	client, err := cvm.NewClient(t.credential, t.Region, t.profile)

	if err == nil {
		request := cvm.NewRebootInstancesRequest()
		request.InstanceIds = []*string{&uuid}
		_, err = client.RebootInstances(request)
	}
	return err
}

func init() {
	cloud.DefaultManager.Register(new(TenantCloud))
}
