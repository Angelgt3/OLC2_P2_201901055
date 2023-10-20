package environment

type Environment struct {
	Anterior       interface{}
	Tabla_variable map[string]Symbol
	Id             string
	Size           map[string]int
}

func NewEnvironment(ant interface{}, id string) Environment {
	return Environment{}
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
