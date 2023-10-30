package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
)

type Break struct {
	Lin int
	Col int
}

func NewBreak(lin int, col int) Break {
	Instr := Break{lin, col}
	return Instr
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return "break"
}

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	Instr := Continue{lin, col}
	return Instr
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return "continue"
}

type Return struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewReturn(lin int, col int, expr interfaces.Expression) Return {
	Instr := Return{lin, col, expr}
	return Instr
}

func (r Return) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	if env.(environment.Environment).ReturnLbl == "" {
		fmt.Println("Return fuera de funcion")
		return result
	}
	result = r.Expresion.Ejecutar(ast, env, gen)
	if result.Type == environment.BOOLEAN {
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.PrintTrue()
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.PrintFalse()
		gen.AddLabel(newLabel)
		gen.AddPrintf("c", "10")
		gen.AddBr()
	} else {
		gen.AddSetStack("(int)P", result.Value)
	}
	gen.AddGoto(env.(environment.Environment).ReturnLbl)

	return result
}
