package nodes

import (
	"errors"
)

var (
	ErrNotFound                = errors.New("file/folder not found")
	ErrDestinationNotDirectory = errors.New("destination is not a directory")
)

func (f *FileEntry) IsDirectory() bool {
	return f.Type == Directory
}

type FileSystem struct {
	Root *FileEntry `json:"-"`
}

func (fs *FileSystem) findRecursive(dir *FileEntry, name string) *FileEntry {
	if dir.IsDirectory() {
		for _, f := range dir.Content.([]FileEntry) {
			if f.Name == name {
				return &f
			} else if f.IsDirectory() {
				return fs.findRecursive(&f, name)
			}
		}
	} else if dir.Name == name {
		return dir
	}

	return nil
}

func (fs *FileSystem) Get(path FilePath) (*FileEntry, error) {
	parts := path.Parts()

	current := fs.Root
	for _, part := range parts {
		res := fs.findRecursive(current, part)
		if res != nil {
			current = res
		} else {
			return nil, ErrNotFound
		}
	}

	return current, nil
}

func (fs *FileSystem) Create(path FilePath, entry *FileEntry) error {
	destPath := path.DirectoryPath()

	dest, err := fs.Get(destPath)
	if err != nil {
		return err
	}

	if !dest.IsDirectory() {
		return ErrDestinationNotDirectory
	}

	dest.Content = append(dest.Content.([]FileEntry), *entry)
	return nil
}

func (fs *FileSystem) Update(path FilePath, entry *FileEntry) error {
	dest, err := fs.Get(path.DirectoryPath())
	if err != nil {
		return err
	}

	if !dest.IsDirectory() {
		return ErrDestinationNotDirectory
	}

	for _, f := range dest.Content.([]FileEntry) {
		if f.Name == path.FileName() {
			if entry.Type != f.Type {
				return errors.New("cannot change file type")
			}

			f.Name = entry.Name
			f.Content = entry.Content

			return nil
		}
	}

	return ErrNotFound
}

func (fs *FileSystem) Delete(path FilePath) error {
	dest, err := fs.Get(path.DirectoryPath())
	if err != nil {
		return err
	}

	if !dest.IsDirectory() {
		return ErrDestinationNotDirectory
	}

	for i, f := range dest.Content.([]FileEntry) {
		if f.Name == path.FileName() {
			dest.Content = append(dest.Content.([]FileEntry)[:i], dest.Content.([]FileEntry)[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}
