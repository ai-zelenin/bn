package xsd

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type CodeGen struct {
	cfg     Config
	tmpl    *template.Template
	typeMap map[string]string
}

func NewCodeGen(cfg Config) (*CodeGen, error) {
	tmpl, err := template.New("").Funcs(map[string]any{
		"title": Title,
	}).ParseGlob(cfg.TemplateGlob)
	if err != nil {
		return nil, err
	}

	return &CodeGen{
		cfg:  cfg,
		tmpl: tmpl,
	}, nil
}

func (c *CodeGen) Generate(schemas []xsd.Schema) error {
	_ = os.Mkdir(c.cfg.OutDir, os.ModePerm)
	p := NewPackage(c.cfg.PackageName)
	p.Imports = c.cfg.Imports
	for _, schema := range schemas {

		for name, t := range schema.Types {
			if name.Space == c.cfg.TargetSpace {
				switch v := t.(type) {
				case *xsd.ComplexType:
					p.AddType(c.createComplexType(v))
				}
			}
		}
	}
	for _, file := range p.Files {
		err := c.generateFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CodeGen) generateFile(file *File) error {
	complexTypeTemplate := c.tmpl.Lookup("complex_type.tmpl")
	if complexTypeTemplate == nil {
		return fmt.Errorf("no temlate for complex_type")
	}
	f, err := os.Create(filepath.Join(c.cfg.OutDir, file.Filename))
	if err != nil {
		return err
	}
	defer f.Close()

	return complexTypeTemplate.Execute(f, file)
}

func (c *CodeGen) createComplexType(ct *xsd.ComplexType) *Type {
	name := makeTypeName(ct)
	t := NewType(name)
	if len(ct.Elements) > 2 {
		for _, element := range ct.Elements {
			fmt.Println(element.Name.Local)
		}
	}

	t.AddField(&Field{
		Name: "XMLName",
		Type: "xml.Name",
		Tag:  fmt.Sprintf("`xml:\"%s\"`", ct.Name.Local),
	})
	for _, attribute := range ct.Attributes {
		f := c.createField(attribute)

		t.AddField(f)
	}
	return t
}

func (c *CodeGen) createField(attribute xsd.Attribute) *Field {
	if attribute.Name.Local == "default" {
		fmt.Printf("%T\n %#v", attribute.Type, attribute)
	}
	switch v := attribute.Type.(type) {
	case *xsd.ComplexType:
		return &Field{
			Name: Title(v.Name.Local),
			Type: makeTypeName(v.Base),
		}
	case xsd.Builtin:
		return &Field{
			Name: Title(attribute.Name.Local),
			Type: makeTypeName(v),
		}
	}
	return nil
}
