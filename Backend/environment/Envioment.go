package environment

type Environment struct {
	Anterior         interface{}
	Tabla_variable   map[string]Symbol
	mutable_variable map[string]bool
	Structs          map[string]Symbol
	Tabla_funciones  map[string]InstF
	Id               string
}

func NewEnvironment(ant interface{}, id string) Environment {
	return Environment{
		Anterior:         ant,
		Tabla_variable:   make(map[string]Symbol),
		mutable_variable: make(map[string]bool),
		Structs:          make(map[string]Symbol),
		Tabla_funciones:  make(map[string]InstF),
		Id:               id,
	}
}

// obtengo entorno actual (string)
func (env Environment) GetEntorno() string {
	return env.Id
}

func GetEntornoGlobal(env interface{}) Environment {
	//obteniendo entorno global
	var tmpEnvGbl Environment

	return tmpEnvGbl
}

func (env Environment) SaveVariable(id string, mut bool, value Symbol) {

}

func (env Environment) GetVariable(id string) Symbol {

	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {

	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}

func (env Environment) SaveFunc(id string, value InstF) {

}

func (env Environment) GetFunc(id string) InstF {

	return InstF{0, 0, "", NULL, nil, nil}
}

func (env Environment) SaveStruct(id string, list []interface{}) {

}

func (env Environment) GetStruct(id string) Symbol {

	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}
