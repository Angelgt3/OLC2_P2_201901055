package environment

// simbolos
type Symbol struct {
	Lin      int
	Col      int
	Tipo     TipoExpresion
	Posicion int
}

// funciones
type InstF struct {
	Lin        int
	Col        int
	Id         string
	Tipo       TipoExpresion
	Parametros []interface{}
	Bloque     []interface{}
}
