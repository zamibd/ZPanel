//go:build !with_gvisor

package core

import (
	"github.com/sagernet/sing-box/adapter/endpoint"
	"github.com/sagernet/sing-box/adapter/service"
	"github.com/sagernet/sing-box/dns"
)

func registerTailscaleEndpoint(_ *endpoint.Registry) {}

func registerTailscaleTransport(_ *dns.TransportRegistry) {}

func registerDERPService(_ *service.Registry) {}
