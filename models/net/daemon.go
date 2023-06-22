package net

import "cydeaos/models/nodes"

type daemon struct {
	Name        string          `json:"name"`
	Online      bool            `json:"online"`
	computerRef *nodes.Computer `json:"-"`
	ports       []*Port         `json:"ports"`
}

type Daemon interface {
	HandleRequest(data interface{}) (interface{}, error)
	Ports() []*Port
}
