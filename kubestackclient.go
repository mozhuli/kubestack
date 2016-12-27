package main

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/hyperhq/kubestack/pkg/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Interface interface {
	// Pods returns a pod interface
	Pods() Pods
	// Networks returns a network interface
	Networks() Networks
	//Subnets returns a subnets interface
	Subnets() Subnets
	// ProviderName returns the network provider ID.
	ProviderName() string
	// CheckTenantID
	CheckTenantID(tenantID string) (bool, error)
}

type Pods interface {
	// Setup pod
	SetupPod(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime, subnetID string) error
	// Teardown pod
	TeardownPod(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime string) error
	// Status of pod
	PodStatus(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime string) (string, error)
}

// Networks is an abstract, pluggable interface for network segment
type Networks interface {
	// Get network by networkName
	GetNetwork(networkName string) (*types.Network, error)
	// Get network by networkID
	GetNetworkByID(networkID string) (*types.Network, error)
	// Create network
	CreateNetwork(network *types.Network) error
	// Update network
	UpdateNetwork(network *types.Network) error
	// Delete network by networkID
	DeleteNetwork(networkID string) error
	// List network by tenantID
	ListNetworks(tenantID string) ([]*types.Network, error)
}

type Subnets interface {
	ListSubnets(networkID string) ([]*types.Subnet, error)
	CreateSubnet(subnet *types.Subnet) error
	DeleteSubnet(subnetID string, networkID string) error
	UpdateSubnet(subnet *types.Subnet) error
	ConnectSubnets(subnet1, subnet2 *types.Subnet) error
	DisconnectSubnets(subnet1, subnet2 *types.Subnet) error
	GetRouterNS(subnetID string) (string, error)
}

type NeutronProvider struct {
	server string

	networkClient      types.NetworksClient
	podClient          types.PodsClient
	loadbalancerClient types.LoadBalancersClient
	subnetClient       types.SubnetsClient
}

func InitNeutronProviders(remoteAddr string) (*NeutronProvider, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		glog.Errorf("Connect network provider %s failed: %v", remoteAddr, err)
		return nil, err
	}

	networkClient := types.NewNetworksClient(conn)
	podClient := types.NewPodsClient(conn)
	lbClient := types.NewLoadBalancersClient(conn)
	subnetClient := types.NewSubnetsClient(conn)
	resp, err := networkClient.Active(
		context.Background(),
		&types.ActiveRequest{},
	)
	if err != nil || !resp.Result {
		glog.Errorf("Active network provider %s failed: %v", remoteAddr, err)
		return nil, err
	}
	return &NeutronProvider{
		server:             remoteAddr,
		podClient:          podClient,
		loadbalancerClient: lbClient,
		networkClient:      networkClient,
		subnetClient:       subnetClient,
	}, nil
}

// Network interface is self
func (r *NeutronProvider) Networks() Networks {
	return r
}

// Pods interface is self
func (r *NeutronProvider) Pods() Pods {
	return r
}

// Subnets interface is self
func (r *NeutronProvider) Subnets() Subnets {
	return r
}

// Get network by networkName
func (r *NeutronProvider) GetNetwork(networkName string) (*types.Network, error) {
	if networkName == "" {
		return nil, errors.New("networkName is null")
	}

	resp, err := r.networkClient.GetNetwork(
		context.Background(),
		&types.GetNetworkRequest{
			Name: networkName,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider get network %s failed: %v", networkName, err)
		return nil, err
	}

	return resp.Network, nil
}

// Get network by networkID
func (r *NeutronProvider) GetNetworkByID(networkID string) (*types.Network, error) {
	if networkID == "" {
		return nil, errors.New("networkID is null")
	}

	resp, err := r.networkClient.GetNetwork(
		context.Background(),
		&types.GetNetworkRequest{
			Id: networkID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider get network %s failed: %v", networkID, err)
		return nil, err
	}

	return resp.Network, nil
}

// Create network
func (r *NeutronProvider) CreateNetwork(network *types.Network) error {
	resp, err := r.networkClient.CreateNetwork(
		context.Background(),
		&types.CreateNetworkRequest{
			Network: network,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider create network %s failed: %v", network.Name, err)
		return err
	}

	return nil
}

// Update network
func (r *NeutronProvider) UpdateNetwork(network *types.Network) error {
	resp, err := r.networkClient.UpdateNetwork(
		context.Background(),
		&types.UpdateNetworkRequest{
			Network: network,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider update network %s failed: %v", network.Name, err)
		return err
	}

	return nil
}

// Delete network by networkName
func (r *NeutronProvider) DeleteNetwork(networkID string) error {
	if networkID == "" {
		return errors.New("networkName is null")
	}

	resp, err := r.networkClient.DeleteNetwork(
		context.Background(),
		&types.DeleteNetworkRequest{
			NetworkID: networkID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider delete network %s failed: %v", networkID, err)
		return err
	}

	return nil
}

//List all Networks by tenantID
func (r *NeutronProvider) ListNetworks(tenantID string) ([]*types.Network, error) {
	resp, err := r.networkClient.ListNetworks(
		context.Background(),
		&types.ListNetworkRequest{
			TenantID: tenantID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider list networks by tenantID %s failed: %v", tenantID, err)
		return nil, err
	}

	return resp.Networks, nil
}

//List all subnets in the  network
func (r *NeutronProvider) ListSubnets(networkID string) ([]*types.Subnet, error) {
	resp, err := r.subnetClient.ListSubnets(
		context.Background(),
		&types.ListSubnetsRequest{
			NetworkID: networkID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider list network %s 's subnets failed: %v", networkID, err)
		return nil, err
	}

	return resp.Subnets, nil
}

func (r *NeutronProvider) CreateSubnet(subnet *types.Subnet) error {
	resp, err := r.subnetClient.CreateSubnet(
		context.Background(),
		&types.CreateSubnetRequest{
			Subnet: subnet,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider create subnet %s  failed: %v", subnet.Name, err)
		return err
	}

	return nil
}

// Delete a subnet from a network
func (r *NeutronProvider) DeleteSubnet(subnetID string, networkID string) error {
	resp, err := r.subnetClient.DeleteSubnet(
		context.Background(),
		&types.DeleteSubnetRequest{
			SubnetID:  subnetID,
			NetworkID: networkID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider delete subnet %s  failed: %v", subnetID, err)
		return err
	}

	return nil
}

//Connect two subnets
func (r *NeutronProvider) ConnectSubnets(subnet1, subnet2 *types.Subnet) error {
	resp, err := r.subnetClient.ConnectSubnets(
		context.Background(),
		&types.ConnectSubnetsRequest{
			Subnet1: subnet1,
			Subnet2: subnet2,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider Connet two subnets %s ,%s failed: %v", subnet1.Name, subnet2.Name, err)
		return err
	}

	return nil
}

//Disonnect two subnets
func (r *NeutronProvider) DisconnectSubnets(subnet1, subnet2 *types.Subnet) error {
	resp, err := r.subnetClient.DisconnectSubnets(
		context.Background(),
		&types.DisconnectSubnetsRequest{
			Subnet1: subnet1,
			Subnet2: subnet2,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider Disconnet two subnets %s ,%s failed: %v", subnet1.Name, subnet2.Name, err)
		return err
	}

	return nil
}

// Get routerNS
func (r *NeutronProvider) GetRouterNS(subnetID string) (string, error) {
	if subnetID == "" {
		return "", errors.New("subnetID is null")
	}

	resp, err := r.subnetClient.GetRouterNS(
		context.Background(),
		&types.GetRouterNSRequest{
			SubnetID: subnetID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("SubnetProvider get routerNS failed: %v", err)
		return "", err
	}

	return resp.RouterNS, nil
}

//Update subnet
func (r *NeutronProvider) UpdateSubnet(subnet *types.Subnet) error {
	resp, err := r.subnetClient.UpdateSubnet(
		context.Background(),
		&types.UpdateSubnetRequest{
			Subnet: subnet,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider update subnet %s  failed: %v", subnet.Name, err)
		return err
	}

	return nil
}

// Setup pod
func (r *NeutronProvider) SetupPod(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime, subnetID string) error {
	resp, err := r.podClient.SetupPod(
		context.Background(),
		&types.SetupPodRequest{
			PodName:             podName,
			Namespace:           namespace,
			PodInfraContainerID: podInfraContainerID,
			ContainerRuntime:    containerRuntime,
			Network:             network,
			SubnetID:            subnetID,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider SetupPod %s failed: %v", podName, err)
		return err
	}

	return nil
}

// Teardown pod
func (r *NeutronProvider) TeardownPod(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime string) error {
	resp, err := r.podClient.TeardownPod(
		context.Background(),
		&types.TeardownPodRequest{
			PodName:             podName,
			Namespace:           namespace,
			PodInfraContainerID: podInfraContainerID,
			ContainerRuntime:    containerRuntime,
			Network:             network,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider TeardownPod %s failed: %v", podName, err)
		return err
	}

	return nil
}

// Status of pod
func (r *NeutronProvider) PodStatus(podName, namespace, podInfraContainerID string, network *types.Network, containerRuntime string) (string, error) {
	resp, err := r.podClient.PodStatus(
		context.Background(),
		&types.PodStatusRequest{
			PodName:             podName,
			Namespace:           namespace,
			PodInfraContainerID: podInfraContainerID,
			ContainerRuntime:    containerRuntime,
			Network:             network,
		},
	)
	if err != nil || resp.Error != "" {
		if err == nil {
			err = errors.New(resp.Error)
		}
		glog.Warningf("NetworkProvider TeardownPod %s failed: %v", podName, err)
		return "", err
	}

	return resp.Ip, nil
}

func main() {
	neutron, err := InitNeutronProviders("10.100.100.244:4237")
	if err != nil {
		fmt.Errorf("init client fail")
	}
	/*networkName := "br-ex"
	getNetworkResponse, err := neutron.Networks().GetNetwork(networkName)
	if err != nil {
		glog.Errorf("NetworkProvider get network %s failed: %v", networkName, err)
		return
	}
	fmt.Println("%v", getNetworkResponse)*/

	testNetName := "testnet1"
	var subnets []*types.Subnet
	subnet := &types.Subnet{
		Name:    "net1-sub1",
		Cidr:    "192.168.11.0/24",
		Gateway: "192.168.11.1",
	}
	subnets = append(subnets, subnet)
	subnet = &types.Subnet{
		Name:    "net1-sub2",
		Cidr:    "192.168.12.0/24",
		Gateway: "192.168.12.1",
	}
	subnets = append(subnets, subnet)
	network := &types.Network{
		Name:    testNetName,
		Subnets: subnets,
	}

	err = neutron.Networks().CreateNetwork(network)
	if err != nil {
		glog.Errorf("NetworkProvider create network failed: ", err)
		return
	}
	/*defer func() {
		err = neutron.Networks().DeleteNetwork(testNetName)
		if err != nil {
			glog.Errorf("NetworkProvider delete network filed: ", err)
		}
	}()*/

	/*getNetworkResponse, err = neutron.Networks().GetNetwork("testnet3")
	if err != nil {
		glog.Errorf("NetworkProvider get network failed: ", err)
		return
	}
	fmt.Println("%v", getNetworkResponse)*/

	/*routerNS, err := neutron.Subnets().GetRouterNS("ca95888f-5f4e-4337-9a0d-43bace66373e")
	if err != nil {
		glog.Errorf("Provider get routerNS failed: ", err)
		return
	}
	fmt.Println("%s", routerNS)

	/*network := &types.Network{
		Uid:  "c2f383a7-1c80-4f71-b987-39214b1597a2",
		Name: "testnet02",
	}
	err = neutron.Networks().UpdateNetwork(network)
	if err != nil {
		glog.Errorf("NetworkProvider update network failed: ", err)
		return
	}*/

	/*ListNetworkResponse, err := neutron.Networks().ListNetworks("414454afc37f4ff395d06e61167f8108")
	if err != nil {
		glog.Errorf("NetworkProvider list Networks failed: ", err)
		return
	}
	fmt.Println("%v", ListNetworkResponse)*/

	/*err = neutron.Networks().DeleteNetwork("754c5f95-c175-423e-98e0-7178559336dd")
	fmt.Println("%v", err)
	if err != nil {
		glog.Errorf("NetworkProvider delete network failed: ", err)
		return
	}

	/*listSubnetResponse, err := neutron.Subnets().ListSubnets("c084cb41-cf08-4d19-abb7-64b3d112baf2")
	if err != nil {
		glog.Errorf("NetworkProvider list subnets failed: ", err)
		return
	}
	fmt.Println("%v", listSubnetResponse)*/

	/*subnet := &types.Subnet{
		NetworkID: "e4b375ca-5834-4efa-ab9d-44ab9425091b",
		Name:      "net3-subnet2",
		Cidr:      "192.168.31.0/24",
		Gateway:   "192.168.31.1",
	}
	err = neutron.Subnets().CreateSubnet(subnet)
	if err != nil {
		glog.Errorf("NetworkProvider create subnets failed: ", err)
		return
	}

	// gateway can not update
	/*subnet := &types.Subnet{
		Uid:  "1fcbfbc3-fd75-49e6-a981-d5c0ba595d7b",
		Name: "subnet03",
	}
	err = neutron.Subnets().UpdateSubnet(subnet)
	if err != nil {
		glog.Errorf("NetworkProvider create subnets failed: ", err)
		return
	}*/
	/*err = neutron.Subnets().DeleteSubnet("1fcbfbc3-fd75-49e6-a981-d5c0ba595d7b", "c2f383a7-1c80-4f71-b987-39214b1597a2")
	if err != nil {
		glog.Errorf("NetworkProvider delete subnets failed: ", err)
		return
	}*/

	/*subnet1 := &types.Subnet{
		NetworkID: "78708fae-e1ef-4f4f-981f-2cb220114f5d",
		Tenantid:  "414454afc37f4ff395d06e61167f8108",
		Name:      "net1-sub1",
		Cidr:      "192.168.11.0/24",
	}

	subnet2 := &types.Subnet{
		NetworkID: "78708fae-e1ef-4f4f-981f-2cb220114f5d",
		Tenantid:  "414454afc37f4ff395d06e61167f8108",
		Name:      "net1-sub2",
		Cidr:      "192.168.12.0/24",
	}

	/*err = neutron.Subnets().ConnectSubnets(subnet1, subnet2)
	if err != nil {
		glog.Errorf("NetworkProvider Connet subnets failed: ", err)
		return
	}*/
	/*err = neutron.Subnets().DisconnectSubnets(subnet1, subnet2)
	if err != nil {
		glog.Errorf("NetworkProvider Connet subnets failed: ", err)
		return
	}

	/*err = neutron.Pods().SetupPod("testPodName1", "testNamespace", "c17601f05328", getNetworkResponse, "docker", "731eafab-1e4f-4ed7-8139-1dd6eb64b878")
	if err != nil {
		glog.Errorf("NetworkProvier create setup pod failed:%v", err)
		return
	}

	/*err = neutron.Pods().SetupPod("testPodName3", "testNamespace3", "828db65632e0", getNetworkResponse, "docker")
	if err != nil {
		glog.Errorf("NetworkProvier create setup pod failed:%v", err)
		return
	}*/

}
