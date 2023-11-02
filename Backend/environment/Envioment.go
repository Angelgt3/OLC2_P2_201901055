package environment

import "fmt"

type Environment struct {
	Anterior        interface{}
	Tabla_variable  map[string]Symbol
	Tabla_funciones map[string]InstF
	Structs         map[string]Symbol
	Id              string
	Size            map[string]int
	BreakLbl        string
	ContinueLbl     string
	ReturnLbl       string
}

func NewEnvironment(ant interface{}, id string, tipo TipoExpresion) Environment {
	env := Environment{
		Anterior:        ant,
		Tabla_variable:  make(map[string]Symbol),
		Tabla_funciones: make(map[string]InstF),
		Structs:         make(map[string]Symbol),
		Id:              id,
		Size:            make(map[string]int),
		BreakLbl:        "",
		ContinueLbl:     "",
		ReturnLbl:       "",
	}
	env.Size["size"] = 0
	return env
}

func (env Environment) GetSizeEntorno() string {
	return string(env.Size["size"])
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
	if variable, ok := env.Tabla_variable[id]; ok {
		fmt.Println("La variable "+id+" ya existe", variable)
		return
	}
	env.Tabla_funciones[id] = value
}

func (env Environment) GetFunc(id string) InstF {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if funcion, ok := tmpEnv.Tabla_funciones[id]; ok {
			return funcion
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La funcion", id, " no existe")
	return InstF{0, 0, "", NULL}

}

func (env Environment) NewVariable() Symbol {
	tempSym := Symbol{Lin: 0, Col: 0, Tipo: NULL, Posicion: env.Size["size"]}
	env.Size["size"] = env.Size["size"] + 1
	return tempSym
}

func (env Environment) SaveArray(id string, tipo TipoExpresion, arrSize int) Symbol {
	if variable, ok := env.Tabla_variable[id]; ok {
		fmt.Println("La variable "+id+" ya existe ", variable)
		return env.Tabla_variable[id]
	}
	env.Tabla_variable[id] = Symbol{Lin: 0, Col: 0, Tipo: tipo, Posicion: env.Size["size"], ArrSize: arrSize}
	env.Size["size"] = env.Size["size"] + 1
	return env.Tabla_variable[id]
}

func (env Environment) SaveStruct(id string, list []interface{}) {
	if _, ok := env.Structs[id]; ok {
		fmt.Println("El struct " + id + " ya existe")
		return
	}
	//env.Structs[id] = list
}

func (env Environment) GetStruct(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env

	for {
		if tmpStruct, ok := tmpEnv.Structs[id]; ok {
			return tmpStruct
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}

	fmt.Println("El struct ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Posicion: 0}
}
