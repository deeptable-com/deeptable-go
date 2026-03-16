// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/deeptable-com/deeptable-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type File string            // Always "file"
type List string            // Always "list"
type StructuredSheet string // Always "structured_sheet"
type Table string           // Always "table"

func (c File) Default() File                       { return "file" }
func (c List) Default() List                       { return "list" }
func (c StructuredSheet) Default() StructuredSheet { return "structured_sheet" }
func (c Table) Default() Table                     { return "table" }

func (c File) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c StructuredSheet) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Table) MarshalJSON() ([]byte, error)           { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
