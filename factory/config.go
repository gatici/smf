// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/*
 * AMF Configuration Factory
 */

package factory

import (
	"time"

	"github.com/free5gc/logger_util"
	"github.com/free5gc/openapi/models"
)

const (
	SMF_EXPECTED_CONFIG_VERSION        = "1.0.0"
	UE_ROUTING_EXPECTED_CONFIG_VERSION = "1.0.0"
)

type Config struct {
	Info          *Info               `yaml:"info"`
	Configuration *Configuration      `yaml:"configuration"`
	Logger        *logger_util.Logger `yaml:"logger"`
}

type Info struct {
	Version     string `yaml:"version,omitempty"`
	Description string `yaml:"description,omitempty"`
}

const (
	SMF_DEFAULT_IPV4     = "127.0.0.2"
	SMF_DEFAULT_PORT     = "8000"
	SMF_DEFAULT_PORT_INT = 8000
)

type Configuration struct {
	SmfName              string               `yaml:"smfName,omitempty"`
	Sbi                  *Sbi                 `yaml:"sbi,omitempty"`
	PFCP                 *PFCP                `yaml:"pfcp,omitempty"`
	NrfUri               string               `yaml:"nrfUri,omitempty"`
	UserPlaneInformation UserPlaneInformation `yaml:"userplane_information"`
	ServiceNameList      []string             `yaml:"serviceNameList,omitempty"`
	SNssaiInfo           []SnssaiInfoItem     `yaml:"snssaiInfos,omitempty"`
	ULCL                 bool                 `yaml:"ulcl,omitempty"`
}

type SnssaiInfoItem struct {
	SNssai   *models.Snssai      `yaml:"sNssai"`
	DnnInfos []SnssaiDnnInfoItem `yaml:"dnnInfos"`
}

type SnssaiDnnInfoItem struct {
	Dnn      string `yaml:"dnn"`
	DNS      DNS    `yaml:"dns"`
	UESubnet string `yaml:"ueSubnet"`
}

type Sbi struct {
	Scheme       string `yaml:"scheme"`
	TLS          *TLS   `yaml:"tls"`
	RegisterIPv4 string `yaml:"registerIPv4,omitempty"` // IP that is registered at NRF.
	// IPv6Addr string `yaml:"ipv6Addr,omitempty"`
	BindingIPv4 string `yaml:"bindingIPv4,omitempty"` // IP used to run the server in the node.
	Port        int    `yaml:"port,omitempty"`
}

type TLS struct {
	PEM string `yaml:"pem,omitempty"`
	Key string `yaml:"key,omitempty"`
}

type PFCP struct {
	Addr string `yaml:"addr,omitempty"`
	Port uint16 `yaml:"port,omitempty"`
}

type DNS struct {
	IPv4Addr string `yaml:"ipv4,omitempty"`
	IPv6Addr string `yaml:"ipv6,omitempty"`
}

type Path struct {
	DestinationIP   string   `yaml:"DestinationIP,omitempty"`
	DestinationPort string   `yaml:"DestinationPort,omitempty"`
	UPF             []string `yaml:"UPF,omitempty"`
}

type UERoutingInfo struct {
	SUPI     string `yaml:"SUPI,omitempty"`
	AN       string `yaml:"AN,omitempty"`
	PathList []Path `yaml:"PathList,omitempty"`
}

// RouteProfID is string providing a Route Profile identifier.
type RouteProfID string

// RouteProfile maintains the mapping between RouteProfileID and ForwardingPolicyID of UPF
type RouteProfile struct {
	// Forwarding Policy ID of the route profile
	ForwardingPolicyID string `yaml:"forwardingPolicyID,omitempty"`
}

// PfdContent represents the flow of the application
type PfdContent struct {
	// Identifies a PFD of an application identifier.
	PfdID string `yaml:"pfdID,omitempty"`
	// Represents a 3-tuple with protocol, server ip and server port for
	// UL/DL application traffic.
	FlowDescriptions []string `yaml:"flowDescriptions,omitempty"`
	// Indicates a URL or a regular expression which is used to match the
	// significant parts of the URL.
	Urls []string `yaml:"urls,omitempty"`
	// Indicates an FQDN or a regular expression as a domain name matching
	// criteria.
	DomainNames []string `yaml:"domainNames,omitempty"`
}

// PfdDataForApp represents the PFDs for an application identifier
type PfdDataForApp struct {
	// Identifier of an application.
	AppID string `yaml:"applicationId"`
	// PFDs for the application identifier.
	Pfds []PfdContent `yaml:"pfds"`
	// Caching time for an application identifier.
	CachingTime *time.Time `yaml:"cachingTime,omitempty"`
}

type RoutingConfig struct {
	Info          *Info                        `yaml:"info"`
	UERoutingInfo []*UERoutingInfo             `yaml:"ueRoutingInfo"`
	RouteProf     map[RouteProfID]RouteProfile `yaml:"routeProfile,omitempty"`
	PfdDatas      []*PfdDataForApp             `yaml:"pfdDataForApp,omitempty"`
}

// UserPlaneInformation describe core network userplane information
type UserPlaneInformation struct {
	UPNodes map[string]UPNode `yaml:"up_nodes"`
	Links   []UPLink          `yaml:"links"`
}

// UPNode represent the user plane node
type UPNode struct {
	Type                 string                     `yaml:"type"`
	NodeID               string                     `yaml:"node_id"`
	ANIP                 string                     `yaml:"an_ip"`
	Dnn                  string                     `yaml:"dnn"`
	SNssaiInfos          []models.SnssaiUpfInfoItem `yaml:"sNssaiUpfInfos,omitempty"`
	InterfaceUpfInfoList []InterfaceUpfInfoItem     `yaml:"interfaces,omitempty"`
}

type InterfaceUpfInfoItem struct {
	InterfaceType   models.UpInterfaceType `yaml:"interfaceType"`
	Endpoints       []string               `yaml:"endpoints"`
	NetworkInstance string                 `yaml:"networkInstance"`
}

type UPLink struct {
	A string `yaml:"A"`
	B string `yaml:"B"`
}

func (c *Config) GetVersion() string {
	if c.Info != nil && c.Info.Version != "" {
		return c.Info.Version
	}
	return ""
}

func (r *RoutingConfig) GetVersion() string {
	if r.Info != nil && r.Info.Version != "" {
		return r.Info.Version
	}
	return ""
}
