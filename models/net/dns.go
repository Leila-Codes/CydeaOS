package net

import (
	"cydeaos/models/nodes"
	"errors"
)

var (
	ErrNodeNotFound          = errors.New("node not found")
	ErrNodeAlreadyRegistered = errors.New("node already registered")
)

type DnsDaemon struct {
	*daemon
	dns map[string]string
}

func NewDnsDaemon(ports ...*Port) *DnsDaemon {
	if len(ports) < 1 {
		ports = append(ports, NewPort(53))
	}

	return &DnsDaemon{
		daemon: &daemon{
			Name:  "dns",
			ports: ports,
		},
		dns: make(map[string]string),
	}
}

func (d *DnsDaemon) Register(node *nodes.Computer) error {
	if _, exists := d.dns[node.Hostname]; exists {
		return ErrNodeAlreadyRegistered
	}

	d.dns[node.Hostname] = node.IP
	return nil
}

func (d *DnsDaemon) Resolve(hostname string) (string, error) {
	if ip, exists := d.dns[hostname]; exists {
		return ip, nil
	}

	return "", ErrNodeNotFound
}

func (d *DnsDaemon) HandleRequest(data interface{}) (interface{}, error) {
	return d.Resolve(data.(string))
}
