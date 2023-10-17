package environment

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       []Error
	RTS          []RS
}

type Error struct {
	Col     int
	Lin     int
	Error   string
	Entorno string
}

type RS struct {
	Id           string
	Tipo_simbolo string
	Tipo_dato    TipoExpresion
	Entorno      string
	Lin          int
	Col          int
}

func NewAST(inst []interface{}, print string, err []Error) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() []Error {
	return a.Errors
}

func (a *AST) SetError(ToErr string, toEcol int, toElin int, entorno string) {
	a.Errors = append(a.Errors, Error{toEcol, toElin, ToErr, entorno})
}

func (a *AST) SetRs(id string, tipo_sim string, tipo_dato TipoExpresion, entorno string, lin int, col int) {
	a.RTS = append(a.RTS, RS{id, tipo_sim, tipo_dato, entorno, lin, col})
}

func (a *AST) GetRTS() []RS {
	return a.RTS
}
