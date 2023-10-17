package main

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"Backend/parser"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

func main() {
	//routes
	http.HandleFunc("/ejecutar", ejecutar_peticion)

	//crea el servidor
	fmt.Println("Run server: http://localhost:5000")
	http.ListenAndServe("localhost:5000", nil)
}

func ejecutar_peticion(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) //permisos de cors

	//leo el json que recibo
	contenido, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	//fmt.Println(string(contenido))
	respuesta := ejecutar_analizador(string(contenido))

	//Escribo el json para enviar
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["nombre"] = "201901055 - Angel"
	resp["result"] = respuesta
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "*")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")
	(*w).Header().Set("Access-Control-Max-Age", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "*")
}

func ejecutar_analizador(code string) string {
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST
	//crear el entorno
	var globalenvioment environment.Environment = environment.NewEnvironment(nil, "GLOBAL")
	//crear el generador
	var generador generator.Generator
	generador = generator.NewGenerator()
	generador.MainCode = true
	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Ejecutar(&Ast, globalenvioment)
	}

	return string(Ast.GetPrint())
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}
