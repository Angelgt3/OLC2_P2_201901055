package generator

import (
	"fmt"
)

type Generator struct {
	Temporal        int           //cantidad de temporales
	Label           int           //cantidad de etiquetas
	Code            []interface{} //areglo de codigo
	FinalCode       []interface{} //C3D
	Natives         []interface{} //funciones nativas
	FuncCode        []interface{} //funciones
	TempList        []interface{} //temporales
	PrintStringFlag bool          //impresion
	BreakLabel      string        //break
	ContinueLabel   string        //continue
	MainCode        bool
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:        0,
		Label:           0,
		BreakLabel:      "",
		ContinueLabel:   "",
		PrintStringFlag: true,
		MainCode:        false,
	}
	return generator
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

func (g Generator) GetTemps() []interface{} {
	return g.TempList
}

func (g *Generator) SetMainFlag(newVal bool) {
	g.MainCode = newVal
}

/*
########################
# Manejo de Temporales
########################
*/

func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	g.TempList = append(g.TempList, temp)
	return temp
}

func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

func (g *Generator) AddLabel(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, Label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, Label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "heap["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "heap["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = heap["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = heap["+index+"];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "stack["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "stack["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = stack["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = stack["+index+"];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "void "+target+"() {\n")
	} else {
		g.FuncCode = append(g.FuncCode, "void "+target+"() {\n")
	}
}

func (g *Generator) AddEnd() {
	if g.MainCode {
		g.Code = append(g.Code, "\treturn;\n")
		g.Code = append(g.Code, "}\n\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\treturn;\n")
		g.FuncCode = append(g.FuncCode, "}\n\n")
	}
}

func (g *Generator) GenerateFinalCode() {
	// agregar headers
	g.FinalCode = append(g.FinalCode, "/*------HEADER------*/\n")
	g.FinalCode = append(g.FinalCode, "#include <stdio.h>\n")
	g.FinalCode = append(g.FinalCode, "#include <math.h>\n")
	g.FinalCode = append(g.FinalCode, "double heap[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double stack[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double P;\n")
	g.FinalCode = append(g.FinalCode, "double H;\n")
	// agregacion de temporales
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		g.FinalCode = append(g.FinalCode, "double ")
		tmpDec := fmt.Sprintf("%v", tempArr[0])
		tempArr = tempArr[1:]
		for _, s := range tempArr {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";\n\n"
		g.FinalCode = append(g.FinalCode, tmpDec)
	}
	// agregar funciones nativas
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------NATIVES------*/\n")
		for _, s := range g.Natives {
			g.FinalCode = append(g.FinalCode, s)
		}
	}
	// agregar funciones
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------FUNCTIONS------*/\n")
		for _, s := range g.FuncCode {
			g.FinalCode = append(g.FinalCode, s)
		}
	}
	// agregar main
	g.FinalCode = append(g.FinalCode, "/*------MAIN------*/\n")
	g.FinalCode = append(g.FinalCode, "int main() {\n")
	g.FinalCode = append(g.FinalCode, "\tP = 0; H = 0;\n\n")
	for _, s := range g.Code {
		g.FinalCode = append(g.FinalCode, "\t"+s.(string))
	}
	g.FinalCode = append(g.FinalCode, "\n\treturn 0;\n}\n")
}
