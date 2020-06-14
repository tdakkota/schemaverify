package object

import (
	schema "github.com/lestrrat-go/jsschema"
	"go/ast"
)

type Definition struct {
	SchemaName string
	*schema.Schema
}

type Pair struct {
	Object     *ast.TypeSpec
	Definition Definition
}

func (p Pair) StructName() string {
	return p.Object.Name.Name
}

func (p Pair) IsSchemaResolved() bool {
	return p.Definition.Schema != nil
}
