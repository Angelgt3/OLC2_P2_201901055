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
	newtmp1 := gen.NewTemp()
	newtmp2 := gen.NewTemp()

	var result, val environment.Value
	size := len(p.ValoresExp)

	gen.AddAssign(newtmp1, "H")
	gen.AddExpression(newtmp2, newtmp1, "1", "+")
	gen.AddSetHeap("(int)H", strconv.Itoa(size))
	gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")

	for _, s := range p.ValoresExp {
		val = s.(interfaces.Expression).Ejecutar(ast, env, gen)
		gen.AddSetHeap("(int)"+newtmp2, val.Value)
		gen.AddExpression(newtmp2, newtmp2, "1", "+")
	}
	result = environment.NewValue(newtmp1, true, environment.ARRAY)
	result.Type = environment.ARRAY
	return result
}

type ArrayAccess struct {
	Lin   int
	Col   int
	Array interfaces.Expression
	Index interfaces.Expression
}

func NewArrayAccess(lin int, col int, array interfaces.Expression, index interfaces.Expression) ArrayAccess {
	exp := ArrayAccess{lin, col, array, index}
	return exp
}

func (p ArrayAccess) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var tempArray, tempIndex, result environment.Value
	tempArray = p.Array.Ejecutar(ast, env, gen)
	tempIndex = p.Index.Ejecutar(ast, env, gen)
	if tempArray.Type >= environment.ARRAY && tempArray.Type <= environment.A_CHAR {
		//llamada
		newTmp := gen.NewTemp()
		lerror := gen.NewLabel()
		laccess := gen.NewLabel()
		lsalida := gen.NewLabel()
		gen.AddIf(tempIndex.Value, "0", "<", lerror)
		tmp := gen.NewTemp()
		gen.AddGetHeap(tmp, "(int)"+tempArray.Value)
		gen.AddIf(tempIndex.Value, tmp, ">", lerror)
		gen.AddGoto(laccess)

		//error fuera de index
		gen.AddLabel(lerror)
		gen.PrintOutIndex()
		gen.PrintNil()
		gen.AddGoto(lsalida)
		ast.SetError("Fuera de indice", p.Col, p.Lin, env.(environment.Environment).GetEntorno())

		gen.AddLabel(laccess)
		gen.AddExpression(newTmp, tempArray.Value, tempIndex.Value, "+")
		gen.AddExpression(newTmp, newTmp, "1", "+")
		newTmp2 := gen.NewTemp()
		gen.AddGetHeap(newTmp2, "(int)"+newTmp)
		gen.AddLabel(lsalida)
		result = environment.NewValue(newTmp2, true, tempArray.Type-7)
		result.Type = tempArray.Type - 7

	} else {
		ast.SetError("No es array para realizar este acesso", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
	}
	return result
}
