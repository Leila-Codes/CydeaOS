package nodes

import "strings"

type FilePath string

func (p *FilePath) DirectoryPath() FilePath {
	i := strings.LastIndex(string(*p), "/")
	return FilePath(string(*p)[0:i])
}

func (p *FilePath) FileName() string {
	i := strings.LastIndex(string(*p), "/")
	return string(*p)[i+1:]
}

func (p *FilePath) Parts() []string {
	return strings.Split(string(*p), "/")
}
