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

func (b Break) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	if env.(environment.Environment).BreakLbl == "" || env.(environment.Environment).BreakLbl == "global" {
		fmt.Println("Break fuera de un ciclo")
		return result
	}
	gen.AddGoto(env.(environment.Environment).BreakLbl)
	return result
}

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	Instr := Continue{lin, col}
	return Instr
}

func (c Continue) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	if env.(environment.Environment).ContinueLbl == "" || env.(environment.Environment).ContinueLbl == "global" {
		fmt.Println("Continue fuera de un ciclo")
		return result
	}
	gen.AddGoto(env.(environment.Environment).ContinueLbl)
	return result
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
	if env.(environment.Environment).ReturnLbl == "" || env.(environment.Environment).ReturnLbl == "global" {
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
