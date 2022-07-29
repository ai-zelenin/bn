package xsd

import (
	"aqwari.net/xml/xsd"
	"strings"
)

var BuiltinTypeMap = []string{
	xsd.AnyType:       "string",
	xsd.AnySimpleType: "string",
	xsd.ENTITIES:      "[]string",
	xsd.ENTITY:        "string",
	xsd.ID:            "string",
	xsd.IDREF:         "string",
	xsd.IDREFS:        "[]string",
	xsd.NCName:        "string",
	xsd.NMTOKEN:       "string",
	xsd.NMTOKENS:      "[]string",
	xsd.NOTATION:      "[]string",
	xsd.Name:          "string",
	xsd.QName:         "xml.Name",
	xsd.AnyURI:        "string",
	xsd.Base64Binary:  "[]byte",
	xsd.Boolean:       "bool",
	xsd.Byte:          "byte",
	xsd.Date:          "time.Time",
	xsd.DateTime:      "time.Time",
	xsd.Decimal:       "float64",
	xsd.Double:        "float64",
	// the "duration" built-in is especially broken, so we
	// don't parse it at all.
	xsd.Duration:           "string",
	xsd.Float:              "float32",
	xsd.GDay:               "time.Time",
	xsd.GMonth:             "time.Time",
	xsd.GMonthDay:          "time.Time",
	xsd.GYear:              "time.Time",
	xsd.GYearMonth:         "time.Time",
	xsd.HexBinary:          "[]byte",
	xsd.Int:                "int",
	xsd.Integer:            "int",
	xsd.Language:           "string",
	xsd.Long:               "int64",
	xsd.NegativeInteger:    "int",
	xsd.NonNegativeInteger: "int",
	xsd.NormalizedString:   "string",
	xsd.NonPositiveInteger: "int",
	xsd.PositiveInteger:    "int",
	xsd.Short:              "int",
	xsd.String:             "string",
	xsd.Time:               "time.Time",
	xsd.Token:              "string",
	xsd.UnsignedByte:       "byte",
	xsd.UnsignedInt:        "uint",
	xsd.UnsignedLong:       "uint64",
	xsd.UnsignedShort:      "uint",
}

func makeTypeName(t xsd.Type) string {
	switch v := t.(type) {
	case *xsd.ComplexType:
		return strings.Title(v.Name.Local)
	case xsd.Builtin:
		return BuiltinTypeMap[v]
	default:
		panic("unexpected type")
	}
}

func Title(str string) string {
	return strings.Title(str)
}
