package environment

// Valores del struct
type StructType struct {
	Id   string
	Tipo TipoExpresion
}

func NewStructType(id string, tipo TipoExpresion) StructType {
	exp := StructType{id, tipo}
	return exp
}

type StructContent struct {
	Id        string
	Expresion interface{}
}

func NewStructContent(ide string, ex interface{}) StructContent {
	exp := StructContent{Id: ide, Expresion: ex}
	return exp
}
