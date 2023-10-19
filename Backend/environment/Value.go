package environment

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
