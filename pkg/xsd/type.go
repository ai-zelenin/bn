package xsd

type Type struct {
	Name     string
	Fields   []*Field
	Embedded []*Type
}

func NewType(name string, fields ...*Field) *Type {
	return &Type{
		Name:   name,
		Fields: fields,
	}
}

func (t *Type) AddField(f *Field) {
	if f != nil {
		t.Fields = append(t.Fields, f)
	}
}
func (t *Type) Embed(f *Type) {
	t.Embedded = append(t.Embedded, f)
}
