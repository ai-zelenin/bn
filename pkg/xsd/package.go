package xsd

type Package struct {
	Name    string
	Files   []*File
	Imports []string
}

func NewPackage(name string) *Package {
	return &Package{
		Name:    name,
		Files:   make([]*File, 0),
		Imports: make([]string, 0),
	}
}

func (p *Package) AddType(t *Type) {
	f := NewFileFromType(t)
	p.AddFile(f)
}

func (p *Package) AddFile(f *File) {
	f.PackageName = p.Name
	f.Imports = p.Imports
	p.Files = append(p.Files, f)
}
