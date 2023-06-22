package nodes

type FileType int

const (
	File FileType = iota
	Directory
)

type FileEntry struct {
	Name    string      `json:"name"`
	Type    FileType    `json:"type"`
	Content interface{} `json:"content"`
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
