package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type Array struct {
	Lin        int
	Col        int
	ValoresExp []interface{}
}

func NewArray(lin int, col int, list []interface{}) Array {
	exp := Array{lin, col, list}
	return exp
}

func (p Array) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("CREANDO ARRAY")
	var result environment.Value
	var arrVal = []interface{}{}

	for _, s := range p.ValoresExp {
		var indexExp = s.(interfaces.Expression).Ejecutar(ast, env, gen)
		arrVal = append(arrVal, indexExp)
	}
	result.Type = environment.ARRAY
	result.ArrValue = arrVal
	result.ArrSize = len(arrVal)
	return result
}

type ArrayAccess struct {
	Lin   int
	Col   int
	Array interfaces.Expression
	Index []interface{}
}

func NewArrayAccess(lin int, col int, array interfaces.Expression, index []interface{}) ArrayAccess {
	exp := ArrayAccess{lin, col, array, index}
	return exp
}

func (p ArrayAccess) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	result := p.Array.Ejecutar(ast, env, gen)
	In := 0
	Nn := result.ArrSize
	var FinalIndex int
	lerror := gen.NewLabel()
	laccess := gen.NewLabel()
	lsalida := gen.NewLabel()
	//validacion array
	if result.Type >= environment.ARRAY && result.Type <= environment.A_CHAR {
		//agregando primera dimension
		firstIndex := p.Index[0].(interfaces.Expression).Ejecutar(ast, env, gen)
		i, _ := strconv.Atoi(firstIndex.Value)
		FinalIndex = i - In
		copiedSlice := make([]interface{}, len(p.Index))
		for i, item := range p.Index {
			if i != 0 {
				copiedSlice[i] = item
			}
		}
		copiedSlice = copiedSlice[1:]
		//agregando N dimension
		for _, ind := range copiedSlice {
			firstIndex := ind.(interfaces.Expression).Ejecutar(ast, env, gen)
			j, _ := strconv.Atoi(firstIndex.Value)
			FinalIndex = FinalIndex*Nn + j - In
		}

		newTmp := gen.NewTemp()
		tmp := gen.NewTemp()
		gen.AddGetStack(tmp, "(int)"+result.Value)
		gen.AddExpression(newTmp, result.Value, strconv.Itoa(FinalIndex), "+")
		gen.AddExpression(newTmp, newTmp, "1", "+")
		newTmp2 := gen.NewTemp()

		terror := gen.NewTemp()
		gen.AddGetStack(terror, "(int)"+newTmp)
		gen.AddIf(terror, strconv.Itoa(0), "==", lerror)
		gen.AddGoto(laccess)

		gen.AddLabel(lerror)
		gen.PrintOutIndex()
		gen.PrintNil()
		gen.AddGoto(lsalida)
		//ast.SetError("Fuera de indice", p.Col, p.Lin, env.(environment.Environment).GetEntorno())

		gen.AddLabel(laccess)
		gen.AddGetStack(newTmp2, "(int)"+newTmp)
		gen.AddLabel(lsalida)
		result = environment.NewValue(newTmp2, true, environment.INTEGER)
	} else {
		ast.SetError("No es array para realizar este acesso", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
	}
	return result
}
