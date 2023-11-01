package generator

import (
	"fmt"
)

type Generator struct {
	Temporal               int           //cantidad de temporales
	Label                  int           //cantidad de etiquetas
	Code                   []interface{} //areglo de codigo
	FinalCode              []interface{} //C3D
	Natives                []interface{} //funciones nativas
	FuncCode               []interface{} //funciones
	TempList               []interface{} //temporales
	PrintStringFlag        bool          //impresion
	NativetoStringFlag     bool          //to string
	NativeConcatStringFlag bool          //concat string
	BreakLabel             string        //break
	ContinueLabel          string        //continue
	MainCode               bool
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:               0,
		Label:                  0,
		BreakLabel:             "",
		ContinueLabel:          "",
		PrintStringFlag:        true,
		NativetoStringFlag:     true,
		NativeConcatStringFlag: true,
		MainCode:               true,
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

///////////////////////
// MANEJO DE TEMPORALES
///////////////////////

func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	g.TempList = append(g.TempList, temp)
	return temp
}

////////////////////
// MANEJO DE LABELS
////////////////////

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

////////
// GOTO
////////

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

//////
// IF
//////

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}

///////////////
// EXPRESIONES
///////////////

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

//////////
// PRINTS
//////////

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) PrintMathError() {
	g.AddPrintf("c", "77")  //M
	g.AddPrintf("c", "97")  //a
	g.AddPrintf("c", "116") //t
	g.AddPrintf("c", "104") //h
	g.AddPrintf("c", "69")  //E
	g.AddPrintf("c", "114") //r
	g.AddPrintf("c", "114") //r
	g.AddPrintf("c", "111") //o
	g.AddPrintf("c", "114") //r
	g.AddPrintf("c", "10")
}

func (g *Generator) PrintNil() {
	g.AddPrintf("c", "110") //N
	g.AddPrintf("c", "105") //i
	g.AddPrintf("c", "108") //l
	g.AddPrintf("c", "10")
}

/*
	func (g *Generator) PrintOutIndex() {
		g.AddPrintf("c", "79")
		g.AddPrintf("c", "117")
		g.AddPrintf("c", "116")
		g.AddPrintf("c", "79")
		g.AddPrintf("c", "102")
		g.AddPrintf("c", "73")
		g.AddPrintf("c", "110")
		g.AddPrintf("c", "100")
		g.AddPrintf("c", "101")
		g.AddPrintf("c", "120")
		g.AddPrintf("c", "10")
	}
*/
func (g *Generator) PrintOutIndex() {
	g.AddPrintf("%c", "66")  //B
	g.AddPrintf("%c", "111") //o
	g.AddPrintf("%c", "117") //u
	g.AddPrintf("%c", "110") //n
	g.AddPrintf("%c", "100") //d
	g.AddPrintf("%c", "115") //s
	g.AddPrintf("%c", "69")  //E
	g.AddPrintf("%c", "114") //r
	g.AddPrintf("%c", "114") //r
	g.AddPrintf("%c", "111") //o
	g.AddPrintf("%c", "114") //r
	g.AddPrintf("c", "10")
}

func (g *Generator) PrintTrue() {
	g.AddPrintf("c", "84")  //T
	g.AddPrintf("c", "114") //r
	g.AddPrintf("c", "117") //u
	g.AddPrintf("c", "101") //e
}

func (g *Generator) PrintFalse() {
	g.AddPrintf("c", "70")  //F
	g.AddPrintf("c", "97")  //a
	g.AddPrintf("c", "108") //l
	g.AddPrintf("c", "115") //s
	g.AddPrintf("c", "101") //e
}

////////
// HEAP
////////

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

/////////
// STACK
/////////

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

///////////
// LLAMADA
///////////

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
	}
}

//////////////
// UTILIDADES
//////////////

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
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

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "void "+target+"() {\n")
	} else {
		g.FuncCode = append(g.FuncCode, "void "+target+"() {\n")
	}
}

//////////////
// COMENTARIO
//////////////

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

func (g *Generator) GenerateFinalCode() {
	// agregar Header
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

//////////
// NATIVE
//////////

func (g *Generator) NativePrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel() //Label para salir de la funcion
		newLvl2 := g.NewLabel() //Label para la comparacion para buscar fin de cadena
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void printString() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];\n")
		g.Natives = append(g.Natives, "\t"+newLvl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = heap[(int)"+newTemp2+"];\n")
		g.Natives = append(g.Natives, "\tif("+newTemp3+" == -1) goto "+newLvl1+";\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char)"+newTemp3+");\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = "+newTemp2+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl1+":\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.PrintStringFlag = false
	}
}

func (g *Generator) NativeFloatToString() {
	if g.NativetoStringFlag {
		//generando temporales y etiquetas
		t0 := g.NewTemp()
		t1 := g.NewTemp()
		t2 := g.NewTemp()
		t3 := g.NewTemp()
		t4 := g.NewTemp()

		L0 := g.NewLabel()
		L1 := g.NewLabel()
		L2 := g.NewLabel()

		g.Natives = append(g.Natives, "void FloatString() {\n")
		g.Natives = append(g.Natives, "heap[(int)H] = -1;\n")
		g.Natives = append(g.Natives, "H = H + 1;\n")
		g.Natives = append(g.Natives, t0+" = P + 1;\n")
		g.Natives = append(g.Natives, t1+" = stack[(int)"+t0+"];\n")
		g.Natives = append(g.Natives, "if ( "+t1+" >= 0 ) goto "+L0+";\n")
		g.Natives = append(g.Natives, t1+" = 0 - "+t1+";\n")
		g.Natives = append(g.Natives, L0+": \n")
		g.Natives = append(g.Natives, t2+" = (int)"+t1+"% 10;\n")
		g.Natives = append(g.Natives, t2+" = "+t2+" + 48;\n")
		g.Natives = append(g.Natives, t1+" = "+t1+" / 10;\n")
		g.Natives = append(g.Natives, t4+" = (int)"+t1+" % 1;\n")
		g.Natives = append(g.Natives, "heap[(int)H]="+t2+";\n")
		g.Natives = append(g.Natives, "H = H + 1;\n")
		g.Natives = append(g.Natives, "if ( "+t1+" != 0 ) goto "+L0+";\n")
		g.Natives = append(g.Natives, t3+" = H;\n")
		g.Natives = append(g.Natives, t0+" = P + 1;\n")
		g.Natives = append(g.Natives, t1+" = stack[(int)"+t0+"];\n")
		g.Natives = append(g.Natives, "if ( "+t1+" >= 0 ) goto "+L1+";\n")
		g.Natives = append(g.Natives, "H=heap[(int)45];\n")
		g.Natives = append(g.Natives, "H = H + 1;\n")
		g.Natives = append(g.Natives, L1+": \n")
		g.Natives = append(g.Natives, t0+" = "+t3+";\n")
		g.Natives = append(g.Natives, L2+": \n")
		g.Natives = append(g.Natives, t0+" = "+t0+" + 1;\n")
		g.Natives = append(g.Natives, t2+" =heap[(int)"+t0+"];\n")
		g.Natives = append(g.Natives, "heap[(int)H]="+t2+";\n")
		g.Natives = append(g.Natives, "H = H + 1;\n")
		g.Natives = append(g.Natives, "if ( "+t2+" != -1 ) goto "+L2+";\n")
		g.Natives = append(g.Natives, "stack[(int)P]="+t3+";\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")

		//se genera la funcion FloatString

		g.NativetoStringFlag = false
	}
}

func (g *Generator) NativeConcatString() {
	if g.NativeConcatStringFlag {
		tmp1 := g.NewTemp()
		tmp2 := g.NewTemp()
		tmp3 := g.NewTemp()
		tmp4 := g.NewTemp()
		tmp5 := g.NewTemp()
		lvl1 := g.NewLabel()
		lvl2 := g.NewLabel()
		lvl3 := g.NewLabel()
		lvl4 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void dbrust_concatString() {\n")
		g.Natives = append(g.Natives, "\t"+tmp1+" = H;"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp2+" = P + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp4+" = stack[(int)"+tmp2+"];"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp3+" = P + 2;"+"\n")
		g.Natives = append(g.Natives, "\t"+lvl2+":"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp5+" = heap[(int)"+tmp4+"];"+"\n")
		g.Natives = append(g.Natives, "\t"+"if("+tmp5+" == -1) goto "+lvl3+";"+"\n")
		g.Natives = append(g.Natives, "\t"+"heap[(int)H] = "+tmp5+";"+"\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp4+" = "+tmp4+" + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+"goto "+lvl2+";"+"\n")
		g.Natives = append(g.Natives, "\t"+lvl3+":"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp4+" = stack[(int)"+tmp3+"];"+"\n")
		g.Natives = append(g.Natives, "\t"+lvl4+":"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp5+" = heap[(int)"+tmp4+"];"+"\n")
		g.Natives = append(g.Natives, "\t"+"if("+tmp5+" == -1) goto "+lvl1+";"+"\n")
		g.Natives = append(g.Natives, "\t"+"heap[(int)H] = "+tmp5+";"+"\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+tmp4+" = "+tmp4+" + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+"goto "+lvl4+";"+"\n")
		g.Natives = append(g.Natives, "\t"+lvl1+":"+"\n")
		g.Natives = append(g.Natives, "\t"+"heap[(int)H] = -1;"+"\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;"+"\n")
		g.Natives = append(g.Natives, "\t"+"stack[(int)P] = "+tmp1+";"+"\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.NativeConcatStringFlag = false
	}
}
