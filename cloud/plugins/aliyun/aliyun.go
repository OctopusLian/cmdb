package aliyun

import (
	"cmdb/cloud"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type AliCloud struct {
	Addr    string
	Key     string
	Secrect string
	Region  string
}

func (a *AliCloud) Name() string {
	return "aliyun"
}

func (a *AliCloud) Init(addr string, key string, secrect string, region string) {
	a.Addr = addr
	a.Key = key
	a.Secrect = secrect
	a.Region = region
}

func (a *AliCloud) TestConnect() error {
	client, err := ecs.NewClientWithAccessKey(a.Region, a.Key, a.Secrect)
	if err == nil {
		request := ecs.CreateDescribeRegionsRequest()
		request.Scheme = "https"
		_, err = client.DescribeRegions(request)
	}
	return err
}

func (a *AliCloud) statusTransform(status string) string {
	statusMap := map[string]string{
		"Running":  cloud.Running,
		"Stopped":  cloud.Stopped,
		"Starting": cloud.Starting,
		"Stopping": cloud.Stopping,
	}

	if text, ok := statusMap[status]; ok {
		return text
	}
	return cloud.Unknow
}

func (a *AliCloud) GetInstances() []*cloud.Instance {
	var limit int = 100
	client, err := ecs.NewClientWithAccessKey(a.Region, a.Key, a.Secrect)
	if err != nil {
		return nil
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(limit)

	response, err := client.DescribeInstances(request)

	instances := make([]*cloud.Instance, response.TotalCount)
	for index, instance := range response.Instances.Instance {
		publicAddrs := make([]string, 0)
		if instance.EipAddress.IpAddress != "" {
			publicAddrs = append(publicAddrs, instance.EipAddress.IpAddress)
		}
		if len(instance.PublicIpAddress.IpAddress) > 0 {
			publicAddrs = append(publicAddrs, instance.PublicIpAddress.IpAddress...)
		}
		privateAddrs := make([]string, 0)
		if len(instance.InnerIpAddress.IpAddress) > 0 {
			privateAddrs = append(privateAddrs, instance.InnerIpAddress.IpAddress...)
		}
		if len(instance.VpcAttributes.PrivateIpAddress.IpAddress) > 0 {
			privateAddrs = append(privateAddrs, instance.VpcAttributes.PrivateIpAddress.IpAddress...)
		}
		instances[index] = &cloud.Instance{
			Key:          instance.InstanceId,
			UUID:         instance.SerialNumber,
			Name:         instance.InstanceName,
			OS:           instance.OSName,
			CPU:          instance.Cpu,
			Memory:       instance.Memory / 1024,
			PublicAddrs:  publicAddrs,
			PrivateAddrs: privateAddrs,
			Status:       a.statusTransform(instance.Status),
			CreatedTime:  instance.CreationTime,
			ExpiredTime:  instance.ExpiredTime,
		}
	}
	return instances
}

func (a *AliCloud) StartInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.Region, a.Key, a.Secrect)
	if err == nil {
		request := ecs.CreateStartInstanceRequest()
		request.Scheme = "https"
		request.InstanceId = uuid

		_, err = client.StartInstance(request)
	}

	return err
}

func (a *AliCloud) StopInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.Region, a.Key, a.Secrect)
	if err == nil {

		request := ecs.CreateStopInstanceRequest()
		request.Scheme = "https"
		request.InstanceId = uuid

		_, err = client.StopInstance(request)
	}

	return err
}

func (a *AliCloud) RestartInstance(uuid string) error {
	client, err := ecs.NewClientWithAccessKey(a.Region, a.Key, a.Secrect)
	if err == nil {
		request := ecs.CreateRebootInstanceRequest()
		request.Scheme = "https"
		request.InstanceId = uuid

		_, err = client.RebootInstance(request)
	}

	return err
}

func init() {
	cloud.DefaultManager.Register(new(AliCloud))
}
