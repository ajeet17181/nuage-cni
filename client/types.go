package client

import (
	"github.com/containernetworking/cni/pkg/types"
	"net"
)

// Args will hold the network metadata
// info to be passed to CNI plugins
type Args struct {
	Mesos Mesos `json:"org.apache.mesos,omitempty"`
}

// Mesos will hold the network information
// that will be passed to CNI plugins by Mesos runtime
type Mesos struct {
	NetworkInfo NetworkInfo `json:"network_info"`
}

// NetworkInfo defines CNI network name
// and network metadata labels that will be passed by CNI
type NetworkInfo struct {
	Name   string `json:"name"`
	Labels struct {
		Labels []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"labels,omitempty"`
	} `json:"labels,omitempty"`
}

// NetConf stores the common network config for Nuage CNI plugin
type NetConf struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Args     Args   `json:"args"`
}

// K8sArgs is the valid CNI_ARGS used for Kubernetes
type K8sArgs struct {
	types.CommonArgs
	IP                         net.IP
	K8S_POD_NAME               types.UnmarshallableString
	K8S_POD_NAMESPACE          types.UnmarshallableString
	K8S_POD_INFRA_CONTAINER_ID types.UnmarshallableString
}

// NuageMetadata will hold metadata needed to resolve
// a port using Nuage defined overlay network
type NuageMetadata struct {
	Enterprise        string
	Domain            string
	Zone              string
	Network           string
	User              string
	PolicyGroup       string
	StaticIP          string
	RedirectionTarget string
}
