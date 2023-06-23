package nodes

import (
	"encoding/json"
)

type FileType int

const (
	File FileType = iota
	Directory
)

type fileHeader struct {
	Name    string          `json:"name"`
	Type    FileType        `json:"type"`
	Content json.RawMessage `json:"content"`
}

type FileEntry struct {
	Name    string      `json:"name"`
	Type    FileType    `json:"type"`
	Content interface{} `json:"content"`
}

func (f *FileEntry) UnmarshalJSON(data []byte) error {
	var header fileHeader
	err := json.Unmarshal(data, &header)
	if err != nil {
		return err
	}

	f.Name = header.Name
	f.Type = header.Type

	var (
		content interface{}
	)

	switch header.Type {
	case File:
		content = ""
	case Directory:
		content = make([]FileEntry, 0)
	}

	err = json.Unmarshal(header.Content, &content)
	if err != nil {
		return err
	}
	f.Content = content
	return nil
}

func NewFile(name string, content string) *FileEntry {
	return &FileEntry{
		Name:    name,
		Type:    File,
		Content: content,
	}
}

func NewDirectory(name string, content []FileEntry) *FileEntry {
	return &FileEntry{
		Name:    name,
		Type:    Directory,
		Content: content,
	}
}
