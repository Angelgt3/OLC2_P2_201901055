package environment

import "fmt"

type Environment struct {
	Anterior       interface{}
	Tabla_variable map[string]Symbol
	Id             string
	Size           map[string]int
}

func NewEnvironment(ant interface{}, id string) Environment {
	env := Environment{
		Anterior:       ant,
		Tabla_variable: make(map[string]Symbol),
		Id:             id,
		Size:           make(map[string]int),
	}
	env.Size["size"] = 0
	return env
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

func (env Environment) SaveVariable(id string, mut bool, tipo TipoExpresion) Symbol {
	if variable, ok := env.Tabla_variable[id]; ok {
		fmt.Println("La variable "+id+" ya existe ", variable)
		return env.Tabla_variable[id]
	}
	env.Tabla_variable[id] = Symbol{Lin: 0, Col: 0, Tipo: tipo, Posicion: env.Size["size"], Mutable: mut}
	env.Size["size"] = env.Size["size"] + 1
	return env.Tabla_variable[id]
}

func (env Environment) GetVariable(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla_variable[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Posicion: 0}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla_variable[id]; ok {
			tmpEnv.Tabla_variable[id] = value
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Posicion: 0}
}

func (env Environment) SaveFunc(id string, value InstF) {

}

func (env Environment) GetFunc(id string) InstF {

	return InstF{0, 0, "", NULL, nil, nil}
}

func (env Environment) SaveStruct(id string, list []interface{}) {

}

func (env Environment) GetStruct(id string) Symbol {

	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Posicion: 0}
}
