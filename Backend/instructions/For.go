package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type ForRange struct {
	Lin        int
	Col        int
	Id         string
	Expresion  interfaces.Expression
	Expresion2 interfaces.Expression
	Bloque     []interface{}
}

func NewForRange(lin int, col int, id string, condition interfaces.Expression, condition2 interfaces.Expression, bloque []interface{}) ForRange {
	forInstr := ForRange{lin, col, id, condition, condition2, bloque}
	return forInstr
}

func (p ForRange) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO FOR")
	//entorno FOR
	var envfor environment.Environment
	envfor = environment.NewEnvironment(env.(environment.Environment), "FOR", environment.NULL)
	envfor.Size["size"] = envfor.Size["size"] + 1

	var condicion1, condicion2, result environment.Value
	condicion1 = p.Expresion.Ejecutar(ast, envfor, gen)
	condicion2 = p.Expresion2.Ejecutar(ast, envfor, gen)

	//declaro un contador temporal
	tempcontador := envfor.SaveVariable(p.Id, true, environment.INTEGER)
	gen.AddSetStack(strconv.Itoa(tempcontador.Posicion), condicion1.Value)
	newTemp := gen.NewTemp()
	gen.AddAssign(newTemp, "0")
	gen.AddBr()

	var OutLvls []interface{}
	inicio := gen.NewLabel()
	gen.AddLabel(inicio)
	newLabel := gen.NewLabel()
	trueLabel := gen.NewLabel()
	falseLabel := gen.NewLabel()
	envfor.ContinueLbl = inicio
	envfor.BreakLbl = falseLabel

	//actualizacion de variables
	variable := envfor.GetVariable(p.Id)                                               //variable
	newvalue := environment.NewValue(fmt.Sprintf("%v", 1), false, environment.INTEGER) // 1
	gen.AddExpression(newTemp, newTemp, newvalue.Value, "+")                           //suma
	newvalue.IntValue = condicion1.IntValue + newvalue.IntValue
	gen.AddSetStack(strconv.Itoa(variable.Posicion), newTemp) //guardo el valor en condicion 1

	//condicion
	gen.AddIf(newTemp, condicion2.Value, "<=", trueLabel)
	gen.AddGoto(falseLabel)
	gen.AddLabel(trueLabel)

	//instrucciones
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			respuesta := s.(interfaces.Instruction).Ejecutar(ast, envfor, gen)
			if respuesta != nil {
				//agregando etiquetas de salida
				for _, lvl := range respuesta.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		}
	}
	gen.AddGoto(inicio)
	//add false labels
	gen.AddLabel(falseLabel)

	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	gen.AddComment("FIN FOR")
	return result
}

type For struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewFor(lin int, col int, id string, condition interfaces.Expression, bloque []interface{}) For {
	forInstr := For{lin, col, id, condition, bloque}
	return forInstr
}

func (p For) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
