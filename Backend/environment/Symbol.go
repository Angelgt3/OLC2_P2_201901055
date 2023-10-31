package environment

// simbolos
type Symbol struct {
	Lin      int
	Col      int
	Tipo     TipoExpresion
	Posicion int
	Mutable  bool
}

// funciones
type InstF struct {
	Lin  int
	Col  int
	Id   string
	Tipo TipoExpresion
}

// Valor
type Value struct {
	Value      string
	IsTemp     bool
	Type       TipoExpresion
	TrueLabel  []interface{}
	FalseLabel []interface{}
	OutLabel   []interface{}
	IntValue   int
}

func NewValue(Val string, temp bool, typee TipoExpresion) Value {
	result := Value{
		Value:    Val,
		IsTemp:   temp,
		Type:     typee,
		IntValue: 0,
	}
	return result
}
