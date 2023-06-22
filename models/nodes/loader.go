package nodes

import (
	"encoding/json"
	"os"
)

type daemonTemplate struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type fileTemplate struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Contents json.RawMessage `json:"contents"`
}

type computerTemplate struct {
	Hostname string           `json:"host"`
	Daemons  []daemonTemplate `json:"daemons"`
	Files    []fileTemplate   `json:"files"`
}

func FromTemplate(f *os.File) (*Computer, error) {
	tmp := computerTemplate{}

	err := json.NewDecoder(f).Decode(&tmp)
	if err != nil {
		return nil, err
	}

	c := &Computer{
		Hostname: tmp.Hostname,
		Online:   true,
		FileSystem: &FileSystem{
			Root: NewDirectory("/", nil),
		},
	}

	return c, nil
}
