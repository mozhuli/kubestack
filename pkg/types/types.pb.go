// Code generated by protoc-gen-gogo.
// source: types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	ActiveRequest
	ActivateResponse
	CommonResponse
	Subnet
	Route
	Network
	GetNetworkRequest
	GetNetworkResponse
	CreateNetworkRequest
	UpdateNetworkRequest
	DeleteNetworkRequest
	ListSubnetsRequest
	ListSubnetsResponse
	CreateSubnetRequest
	DeleteSubnetRequest
	UpdateSubnetRequest
	GetLoadBalancerRequest
	GetLoadBalancerResponse
	HostPort
	LoadBalancer
	CreateLoadBalancerRequest
	CreateLoadBalancerResponse
	UpdateLoadBalancerRequest
	UpdateLoadBalancerResponse
	DeleteLoadBalancerRequest
	CheckTenantIDRequest
	CheckTenantIDResponse
	SetupPodRequest
	TeardownPodRequest
	PodStatusRequest
	PodStatusResponse
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ActiveRequest struct {
}

func (m *ActiveRequest) Reset()         { *m = ActiveRequest{} }
func (m *ActiveRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveRequest) ProtoMessage()    {}

type ActivateResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *ActivateResponse) Reset()         { *m = ActivateResponse{} }
func (m *ActivateResponse) String() string { return proto.CompactTextString(m) }
func (*ActivateResponse) ProtoMessage()    {}

type CommonResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}

// Subnet is a representaion of a subnet
type Subnet struct {
	NetworkID  string   `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uid        string   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cidr       string   `protobuf:"bytes,4,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Gateway    string   `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Tenantid   string   `protobuf:"bytes,6,opt,name=tenantid,proto3" json:"tenantid,omitempty"`
	Dnsservers []string `protobuf:"bytes,7,rep,name=dnsservers" json:"dnsservers,omitempty"`
	Routes     []*Route `protobuf:"bytes,8,rep,name=routes" json:"routes,omitempty"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}

func (m *Subnet) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// Route is a representation of an advanced routing rule.
type Route struct {
	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Nexthop         string `protobuf:"bytes,2,opt,name=nexthop,proto3" json:"nexthop,omitempty"`
	DestinationCIDR string `protobuf:"bytes,3,opt,name=destinationCIDR,proto3" json:"destinationCIDR,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}

type Network struct {
	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid       string    `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	TenantID  string    `protobuf:"bytes,3,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	SegmentID int32     `protobuf:"varint,4,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	Subnets   []*Subnet `protobuf:"bytes,5,rep,name=subnets" json:"subnets,omitempty"`
	// Status of network
	// Valid value: Initializing, Active, Pending, Failed, Terminating
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}

func (m *Network) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type GetNetworkRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetNetworkRequest) Reset()         { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()    {}

type GetNetworkResponse struct {
	Network *Network `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Error   string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *GetNetworkResponse) Reset()         { *m = GetNetworkResponse{} }
func (m *GetNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*GetNetworkResponse) ProtoMessage()    {}

func (m *GetNetworkResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type CreateNetworkRequest struct {
	Network *Network `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
}

func (m *CreateNetworkRequest) Reset()         { *m = CreateNetworkRequest{} }
func (m *CreateNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkRequest) ProtoMessage()    {}

func (m *CreateNetworkRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type UpdateNetworkRequest struct {
	Network *Network `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
}

func (m *UpdateNetworkRequest) Reset()         { *m = UpdateNetworkRequest{} }
func (m *UpdateNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNetworkRequest) ProtoMessage()    {}

func (m *UpdateNetworkRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type DeleteNetworkRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
}

func (m *DeleteNetworkRequest) Reset()         { *m = DeleteNetworkRequest{} }
func (m *DeleteNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkRequest) ProtoMessage()    {}

type ListSubnetsRequest struct {
	NetworkID string `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
}

func (m *ListSubnetsRequest) Reset()         { *m = ListSubnetsRequest{} }
func (m *ListSubnetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSubnetsRequest) ProtoMessage()    {}

type ListSubnetsResponse struct {
	Subnets []*Subnet `protobuf:"bytes,1,rep,name=subnets" json:"subnets,omitempty"`
}

func (m *ListSubnetsResponse) Reset()         { *m = ListSubnetsResponse{} }
func (m *ListSubnetsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSubnetsResponse) ProtoMessage()    {}

func (m *ListSubnetsResponse) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type CreateSubnetRequest struct {
	Subnet *Subnet `protobuf:"bytes,1,opt,name=subnet" json:"subnet,omitempty"`
}

func (m *CreateSubnetRequest) Reset()         { *m = CreateSubnetRequest{} }
func (m *CreateSubnetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSubnetRequest) ProtoMessage()    {}

func (m *CreateSubnetRequest) GetSubnet() *Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type DeleteSubnetRequest struct {
	SubnetID  string `protobuf:"bytes,1,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	NetworkID string `protobuf:"bytes,2,opt,name=networkID,proto3" json:"networkID,omitempty"`
}

func (m *DeleteSubnetRequest) Reset()         { *m = DeleteSubnetRequest{} }
func (m *DeleteSubnetRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSubnetRequest) ProtoMessage()    {}

type UpdateSubnetRequest struct {
	Subnet *Subnet `protobuf:"bytes,1,opt,name=subnet" json:"subnet,omitempty"`
}

func (m *UpdateSubnetRequest) Reset()         { *m = UpdateSubnetRequest{} }
func (m *UpdateSubnetRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSubnetRequest) ProtoMessage()    {}

func (m *UpdateSubnetRequest) GetSubnet() *Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type GetLoadBalancerRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetLoadBalancerRequest) Reset()         { *m = GetLoadBalancerRequest{} }
func (m *GetLoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerRequest) ProtoMessage()    {}

type GetLoadBalancerResponse struct {
	LoadBalancer *LoadBalancer `protobuf:"bytes,1,opt,name=loadBalancer" json:"loadBalancer,omitempty"`
	Error        string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *GetLoadBalancerResponse) Reset()         { *m = GetLoadBalancerResponse{} }
func (m *GetLoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerResponse) ProtoMessage()    {}

func (m *GetLoadBalancerResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type HostPort struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ipaddress   string `protobuf:"bytes,2,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	ServicePort int32  `protobuf:"varint,3,opt,name=servicePort,proto3" json:"servicePort,omitempty"`
	TargetPort  int32  `protobuf:"varint,4,opt,name=targetPort,proto3" json:"targetPort,omitempty"`
}

func (m *HostPort) Reset()         { *m = HostPort{} }
func (m *HostPort) String() string { return proto.CompactTextString(m) }
func (*HostPort) ProtoMessage()    {}

// LoadBalancer is a replace of kube-proxy, so load-balancing can be handled
// by network providers so as to overcome iptables overhead
type LoadBalancer struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid  string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Vip  string `protobuf:"bytes,3,opt,name=vip,proto3" json:"vip,omitempty"`
	// Valid values: "TCP", "UDP", "HTTP", "HTTPS"
	LoadBalanceType string      `protobuf:"bytes,4,opt,name=loadBalanceType,proto3" json:"loadBalanceType,omitempty"`
	Status          string      `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	TenantID        string      `protobuf:"bytes,6,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	ExternalIPs     []string    `protobuf:"bytes,7,rep,name=externalIPs" json:"externalIPs,omitempty"`
	Subnets         []*Subnet   `protobuf:"bytes,8,rep,name=subnets" json:"subnets,omitempty"`
	Hosts           []*HostPort `protobuf:"bytes,9,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *LoadBalancer) Reset()         { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()    {}

func (m *LoadBalancer) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *LoadBalancer) GetHosts() []*HostPort {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type CreateLoadBalancerRequest struct {
	LoadBalancer *LoadBalancer `protobuf:"bytes,1,opt,name=loadBalancer" json:"loadBalancer,omitempty"`
	// Valid value: ClientIP, None
	Affinity string `protobuf:"bytes,2,opt,name=affinity,proto3" json:"affinity,omitempty"`
}

func (m *CreateLoadBalancerRequest) Reset()         { *m = CreateLoadBalancerRequest{} }
func (m *CreateLoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLoadBalancerRequest) ProtoMessage()    {}

func (m *CreateLoadBalancerRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type CreateLoadBalancerResponse struct {
	Vip   string `protobuf:"bytes,1,opt,name=vip,proto3" json:"vip,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *CreateLoadBalancerResponse) Reset()         { *m = CreateLoadBalancerResponse{} }
func (m *CreateLoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLoadBalancerResponse) ProtoMessage()    {}

type UpdateLoadBalancerRequest struct {
	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExternalIPs []string    `protobuf:"bytes,2,rep,name=externalIPs" json:"externalIPs,omitempty"`
	Hosts       []*HostPort `protobuf:"bytes,3,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *UpdateLoadBalancerRequest) Reset()         { *m = UpdateLoadBalancerRequest{} }
func (m *UpdateLoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateLoadBalancerRequest) ProtoMessage()    {}

func (m *UpdateLoadBalancerRequest) GetHosts() []*HostPort {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type UpdateLoadBalancerResponse struct {
	Vip   string `protobuf:"bytes,1,opt,name=vip,proto3" json:"vip,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *UpdateLoadBalancerResponse) Reset()         { *m = UpdateLoadBalancerResponse{} }
func (m *UpdateLoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateLoadBalancerResponse) ProtoMessage()    {}

type DeleteLoadBalancerRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteLoadBalancerRequest) Reset()         { *m = DeleteLoadBalancerRequest{} }
func (m *DeleteLoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLoadBalancerRequest) ProtoMessage()    {}

type CheckTenantIDRequest struct {
	TenantID string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
}

func (m *CheckTenantIDRequest) Reset()         { *m = CheckTenantIDRequest{} }
func (m *CheckTenantIDRequest) String() string { return proto.CompactTextString(m) }
func (*CheckTenantIDRequest) ProtoMessage()    {}

type CheckTenantIDResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *CheckTenantIDResponse) Reset()         { *m = CheckTenantIDResponse{} }
func (m *CheckTenantIDResponse) String() string { return proto.CompactTextString(m) }
func (*CheckTenantIDResponse) ProtoMessage()    {}

type SetupPodRequest struct {
	PodName             string   `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	Namespace           string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerRuntime    string   `protobuf:"bytes,3,opt,name=containerRuntime,proto3" json:"containerRuntime,omitempty"`
	PodInfraContainerID string   `protobuf:"bytes,4,opt,name=podInfraContainerID,proto3" json:"podInfraContainerID,omitempty"`
	Network             *Network `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *SetupPodRequest) Reset()         { *m = SetupPodRequest{} }
func (m *SetupPodRequest) String() string { return proto.CompactTextString(m) }
func (*SetupPodRequest) ProtoMessage()    {}

func (m *SetupPodRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type TeardownPodRequest struct {
	PodName             string   `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	Namespace           string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerRuntime    string   `protobuf:"bytes,3,opt,name=containerRuntime,proto3" json:"containerRuntime,omitempty"`
	PodInfraContainerID string   `protobuf:"bytes,4,opt,name=podInfraContainerID,proto3" json:"podInfraContainerID,omitempty"`
	Network             *Network `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *TeardownPodRequest) Reset()         { *m = TeardownPodRequest{} }
func (m *TeardownPodRequest) String() string { return proto.CompactTextString(m) }
func (*TeardownPodRequest) ProtoMessage()    {}

func (m *TeardownPodRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type PodStatusRequest struct {
	PodName             string   `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	Namespace           string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerRuntime    string   `protobuf:"bytes,3,opt,name=containerRuntime,proto3" json:"containerRuntime,omitempty"`
	PodInfraContainerID string   `protobuf:"bytes,4,opt,name=podInfraContainerID,proto3" json:"podInfraContainerID,omitempty"`
	Network             *Network `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *PodStatusRequest) Reset()         { *m = PodStatusRequest{} }
func (m *PodStatusRequest) String() string { return proto.CompactTextString(m) }
func (*PodStatusRequest) ProtoMessage()    {}

func (m *PodStatusRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type PodStatusResponse struct {
	Ip    string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *PodStatusResponse) Reset()         { *m = PodStatusResponse{} }
func (m *PodStatusResponse) String() string { return proto.CompactTextString(m) }
func (*PodStatusResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*ActiveRequest)(nil), "types.ActiveRequest")
	proto.RegisterType((*ActivateResponse)(nil), "types.ActivateResponse")
	proto.RegisterType((*CommonResponse)(nil), "types.CommonResponse")
	proto.RegisterType((*Subnet)(nil), "types.Subnet")
	proto.RegisterType((*Route)(nil), "types.Route")
	proto.RegisterType((*Network)(nil), "types.Network")
	proto.RegisterType((*GetNetworkRequest)(nil), "types.GetNetworkRequest")
	proto.RegisterType((*GetNetworkResponse)(nil), "types.GetNetworkResponse")
	proto.RegisterType((*CreateNetworkRequest)(nil), "types.CreateNetworkRequest")
	proto.RegisterType((*UpdateNetworkRequest)(nil), "types.UpdateNetworkRequest")
	proto.RegisterType((*DeleteNetworkRequest)(nil), "types.DeleteNetworkRequest")
	proto.RegisterType((*ListSubnetsRequest)(nil), "types.ListSubnetsRequest")
	proto.RegisterType((*ListSubnetsResponse)(nil), "types.ListSubnetsResponse")
	proto.RegisterType((*CreateSubnetRequest)(nil), "types.CreateSubnetRequest")
	proto.RegisterType((*DeleteSubnetRequest)(nil), "types.DeleteSubnetRequest")
	proto.RegisterType((*UpdateSubnetRequest)(nil), "types.UpdateSubnetRequest")
	proto.RegisterType((*GetLoadBalancerRequest)(nil), "types.GetLoadBalancerRequest")
	proto.RegisterType((*GetLoadBalancerResponse)(nil), "types.GetLoadBalancerResponse")
	proto.RegisterType((*HostPort)(nil), "types.HostPort")
	proto.RegisterType((*LoadBalancer)(nil), "types.LoadBalancer")
	proto.RegisterType((*CreateLoadBalancerRequest)(nil), "types.CreateLoadBalancerRequest")
	proto.RegisterType((*CreateLoadBalancerResponse)(nil), "types.CreateLoadBalancerResponse")
	proto.RegisterType((*UpdateLoadBalancerRequest)(nil), "types.UpdateLoadBalancerRequest")
	proto.RegisterType((*UpdateLoadBalancerResponse)(nil), "types.UpdateLoadBalancerResponse")
	proto.RegisterType((*DeleteLoadBalancerRequest)(nil), "types.DeleteLoadBalancerRequest")
	proto.RegisterType((*CheckTenantIDRequest)(nil), "types.CheckTenantIDRequest")
	proto.RegisterType((*CheckTenantIDResponse)(nil), "types.CheckTenantIDResponse")
	proto.RegisterType((*SetupPodRequest)(nil), "types.SetupPodRequest")
	proto.RegisterType((*TeardownPodRequest)(nil), "types.TeardownPodRequest")
	proto.RegisterType((*PodStatusRequest)(nil), "types.PodStatusRequest")
	proto.RegisterType((*PodStatusResponse)(nil), "types.PodStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Networks service

type NetworksClient interface {
	Active(ctx context.Context, in *ActiveRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	CheckTenantID(ctx context.Context, in *CheckTenantIDRequest, opts ...grpc.CallOption) (*CheckTenantIDResponse, error)
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error)
	CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateNetwork(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type networksClient struct {
	cc *grpc.ClientConn
}

func NewNetworksClient(cc *grpc.ClientConn) NetworksClient {
	return &networksClient{cc}
}

func (c *networksClient) Active(ctx context.Context, in *ActiveRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := grpc.Invoke(ctx, "/types.Networks/Active", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networksClient) CheckTenantID(ctx context.Context, in *CheckTenantIDRequest, opts ...grpc.CallOption) (*CheckTenantIDResponse, error) {
	out := new(CheckTenantIDResponse)
	err := grpc.Invoke(ctx, "/types.Networks/CheckTenantID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networksClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error) {
	out := new(GetNetworkResponse)
	err := grpc.Invoke(ctx, "/types.Networks/GetNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networksClient) CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Networks/CreateNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networksClient) UpdateNetwork(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Networks/UpdateNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networksClient) DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Networks/DeleteNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Networks service

type NetworksServer interface {
	Active(context.Context, *ActiveRequest) (*ActivateResponse, error)
	CheckTenantID(context.Context, *CheckTenantIDRequest) (*CheckTenantIDResponse, error)
	GetNetwork(context.Context, *GetNetworkRequest) (*GetNetworkResponse, error)
	CreateNetwork(context.Context, *CreateNetworkRequest) (*CommonResponse, error)
	UpdateNetwork(context.Context, *UpdateNetworkRequest) (*CommonResponse, error)
	DeleteNetwork(context.Context, *DeleteNetworkRequest) (*CommonResponse, error)
}

func RegisterNetworksServer(s *grpc.Server, srv NetworksServer) {
	s.RegisterService(&_Networks_serviceDesc, srv)
}

func _Networks_Active_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).Active(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Networks_CheckTenantID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckTenantIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).CheckTenantID(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Networks_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).GetNetwork(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Networks_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).CreateNetwork(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Networks_UpdateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).UpdateNetwork(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Networks_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NetworksServer).DeleteNetwork(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Networks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Networks",
	HandlerType: (*NetworksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Active",
			Handler:    _Networks_Active_Handler,
		},
		{
			MethodName: "CheckTenantID",
			Handler:    _Networks_CheckTenantID_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _Networks_GetNetwork_Handler,
		},
		{
			MethodName: "CreateNetwork",
			Handler:    _Networks_CreateNetwork_Handler,
		},
		{
			MethodName: "UpdateNetwork",
			Handler:    _Networks_UpdateNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _Networks_DeleteNetwork_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Subnets service

type SubnetsClient interface {
	ListSubnets(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error)
	CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateSubnet(ctx context.Context, in *UpdateSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteSubnet(ctx context.Context, in *DeleteSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type subnetsClient struct {
	cc *grpc.ClientConn
}

func NewSubnetsClient(cc *grpc.ClientConn) SubnetsClient {
	return &subnetsClient{cc}
}

func (c *subnetsClient) ListSubnets(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error) {
	out := new(ListSubnetsResponse)
	err := grpc.Invoke(ctx, "/types.Subnets/ListSubnets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetsClient) CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Subnets/CreateSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetsClient) UpdateSubnet(ctx context.Context, in *UpdateSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Subnets/UpdateSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetsClient) DeleteSubnet(ctx context.Context, in *DeleteSubnetRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Subnets/DeleteSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Subnets service

type SubnetsServer interface {
	ListSubnets(context.Context, *ListSubnetsRequest) (*ListSubnetsResponse, error)
	CreateSubnet(context.Context, *CreateSubnetRequest) (*CommonResponse, error)
	UpdateSubnet(context.Context, *UpdateSubnetRequest) (*CommonResponse, error)
	DeleteSubnet(context.Context, *DeleteSubnetRequest) (*CommonResponse, error)
}

func RegisterSubnetsServer(s *grpc.Server, srv SubnetsServer) {
	s.RegisterService(&_Subnets_serviceDesc, srv)
}

func _Subnets_ListSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListSubnetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SubnetsServer).ListSubnets(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Subnets_CreateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SubnetsServer).CreateSubnet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Subnets_UpdateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SubnetsServer).UpdateSubnet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Subnets_DeleteSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SubnetsServer).DeleteSubnet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Subnets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Subnets",
	HandlerType: (*SubnetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSubnets",
			Handler:    _Subnets_ListSubnets_Handler,
		},
		{
			MethodName: "CreateSubnet",
			Handler:    _Subnets_CreateSubnet_Handler,
		},
		{
			MethodName: "UpdateSubnet",
			Handler:    _Subnets_UpdateSubnet_Handler,
		},
		{
			MethodName: "DeleteSubnet",
			Handler:    _Subnets_DeleteSubnet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	GetLoadBalancer(ctx context.Context, in *GetLoadBalancerRequest, opts ...grpc.CallOption) (*GetLoadBalancerResponse, error)
	CreateLoadBalancer(ctx context.Context, in *CreateLoadBalancerRequest, opts ...grpc.CallOption) (*CreateLoadBalancerResponse, error)
	UpdateLoadBalancer(ctx context.Context, in *UpdateLoadBalancerRequest, opts ...grpc.CallOption) (*UpdateLoadBalancerResponse, error)
	DeleteLoadBalancer(ctx context.Context, in *DeleteLoadBalancerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) GetLoadBalancer(ctx context.Context, in *GetLoadBalancerRequest, opts ...grpc.CallOption) (*GetLoadBalancerResponse, error) {
	out := new(GetLoadBalancerResponse)
	err := grpc.Invoke(ctx, "/types.LoadBalancers/GetLoadBalancer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) CreateLoadBalancer(ctx context.Context, in *CreateLoadBalancerRequest, opts ...grpc.CallOption) (*CreateLoadBalancerResponse, error) {
	out := new(CreateLoadBalancerResponse)
	err := grpc.Invoke(ctx, "/types.LoadBalancers/CreateLoadBalancer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) UpdateLoadBalancer(ctx context.Context, in *UpdateLoadBalancerRequest, opts ...grpc.CallOption) (*UpdateLoadBalancerResponse, error) {
	out := new(UpdateLoadBalancerResponse)
	err := grpc.Invoke(ctx, "/types.LoadBalancers/UpdateLoadBalancer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) DeleteLoadBalancer(ctx context.Context, in *DeleteLoadBalancerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.LoadBalancers/DeleteLoadBalancer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	GetLoadBalancer(context.Context, *GetLoadBalancerRequest) (*GetLoadBalancerResponse, error)
	CreateLoadBalancer(context.Context, *CreateLoadBalancerRequest) (*CreateLoadBalancerResponse, error)
	UpdateLoadBalancer(context.Context, *UpdateLoadBalancerRequest) (*UpdateLoadBalancerResponse, error)
	DeleteLoadBalancer(context.Context, *DeleteLoadBalancerRequest) (*CommonResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_GetLoadBalancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LoadBalancersServer).GetLoadBalancer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _LoadBalancers_CreateLoadBalancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LoadBalancersServer).CreateLoadBalancer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _LoadBalancers_UpdateLoadBalancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LoadBalancersServer).UpdateLoadBalancer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _LoadBalancers_DeleteLoadBalancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(LoadBalancersServer).DeleteLoadBalancer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLoadBalancer",
			Handler:    _LoadBalancers_GetLoadBalancer_Handler,
		},
		{
			MethodName: "CreateLoadBalancer",
			Handler:    _LoadBalancers_CreateLoadBalancer_Handler,
		},
		{
			MethodName: "UpdateLoadBalancer",
			Handler:    _LoadBalancers_UpdateLoadBalancer_Handler,
		},
		{
			MethodName: "DeleteLoadBalancer",
			Handler:    _LoadBalancers_DeleteLoadBalancer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Pods service

type PodsClient interface {
	SetupPod(ctx context.Context, in *SetupPodRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	TeardownPod(ctx context.Context, in *TeardownPodRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	PodStatus(ctx context.Context, in *PodStatusRequest, opts ...grpc.CallOption) (*PodStatusResponse, error)
}

type podsClient struct {
	cc *grpc.ClientConn
}

func NewPodsClient(cc *grpc.ClientConn) PodsClient {
	return &podsClient{cc}
}

func (c *podsClient) SetupPod(ctx context.Context, in *SetupPodRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Pods/SetupPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) TeardownPod(ctx context.Context, in *TeardownPodRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := grpc.Invoke(ctx, "/types.Pods/TeardownPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) PodStatus(ctx context.Context, in *PodStatusRequest, opts ...grpc.CallOption) (*PodStatusResponse, error) {
	out := new(PodStatusResponse)
	err := grpc.Invoke(ctx, "/types.Pods/PodStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pods service

type PodsServer interface {
	SetupPod(context.Context, *SetupPodRequest) (*CommonResponse, error)
	TeardownPod(context.Context, *TeardownPodRequest) (*CommonResponse, error)
	PodStatus(context.Context, *PodStatusRequest) (*PodStatusResponse, error)
}

func RegisterPodsServer(s *grpc.Server, srv PodsServer) {
	s.RegisterService(&_Pods_serviceDesc, srv)
}

func _Pods_SetupPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetupPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PodsServer).SetupPod(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Pods_TeardownPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TeardownPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PodsServer).TeardownPod(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Pods_PodStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PodStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PodsServer).PodStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Pods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Pods",
	HandlerType: (*PodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupPod",
			Handler:    _Pods_SetupPod_Handler,
		},
		{
			MethodName: "TeardownPod",
			Handler:    _Pods_TeardownPod_Handler,
		},
		{
			MethodName: "PodStatus",
			Handler:    _Pods_PodStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
