package xsd

import "strings"

type File struct {
	PackageName string
	Filename    string
	Imports     []string
	Types       []*Type
}

func NewFile(filename string) *File {
	return &File{
		Filename: filename,
		Types:    make([]*Type, 0),
	}
}

func NewFileFromType(t *Type) *File {
	return &File{
		Filename: strings.ToLower(t.Name) + ".go",
		Types:    []*Type{t},
	}
}

func (f *File) AddType(t *Type) {
	f.Types = append(f.Types, t)
}
