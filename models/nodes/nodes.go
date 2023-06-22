package nodes

import (
	"cydeaos/models/net"
	"fmt"
)

var (
	ErrComputerOffline = fmt.Errorf("computer offline")
)

type Computer struct {
	Hostname   string       `json:"host"`
	IP         string       `json:"ip"`
	Owner      string       `json:"owner"`
	Online     bool         `json:"online"`
	FileSystem *FileSystem  `json:"-"`
	Daemons    []net.Daemon `json:"daemons"`
}

func (c *Computer) GetPort(port int) (*net.Port, error) {
	for _, d := range c.Daemons {
		for _, p := range d.Ports() {
			if p.Number == port {
				return p, nil
			}
		}
	}

	return nil, fmt.Errorf("port not found")
}

func (c *Computer) ListPorts() []*net.Port {
	var ports []*net.Port

	for _, d := range c.Daemons {
		for _, p := range d.Ports() {
			ports = append(ports, p)
		}
	}

	return ports
}

func (c *Computer) CheckStatus() error {
	if !c.Online {
		return ErrComputerOffline
	}

	return nil
}
