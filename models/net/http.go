package net

import (
	"cydeaos/models/nodes"
	"fmt"
)

type HttpDaemon struct {
	*daemon
}

func NewHttpDaemon(computer *nodes.Computer, ports ...*Port) *HttpDaemon {
	if len(ports) < 1 {
		ports = append(ports, NewPort(80))
	}

	return &HttpDaemon{
		daemon: &daemon{
			Name:        "http",
			Online:      true,
			computerRef: computer,
			ports:       ports,
		},
	}
}

func (h *HttpDaemon) HandleRequest(data interface{}) (interface{}, error) {
	path := nodes.FilePath(data.(string))

	file, err := h.computerRef.FileSystem.Get(path)
	if err != nil {
		return nil, err
	}

	if file.Type != nodes.File {
		return nil, fmt.Errorf("'%s' is not a file", path)
	}

	return file.Content, nil
}
