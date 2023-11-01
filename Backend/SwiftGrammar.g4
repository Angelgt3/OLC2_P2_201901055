grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Backend/interfaces"
    import "Backend/environment"
    import "Backend/expressions"
    import "Backend/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;
//bloque de isntrucciones
block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;
/* 
//bloque de isntrucciones
block2 returns [[]interface{} blk2]
@init{
    $blk2 = []interface{}{}
    var listInt []IInstruction2Context
  }
: ins2+=instruction2+
    {
        listInt = localctx.(*Block2Context).GetIns()
        for _, e := range listInt {
            $blk2 = append($blk2, e.GetInst())
        }
    }
;
*/
//instrucciones
instruction returns [interfaces.Instruction inst]
: printstmt PCOMA?                                                  { $inst = $printstmt.prnt}
| variablestmt PCOMA?                                               { $inst = $variablestmt.vari}
| guardstmt                                                         { $inst = $guardstmt.guinst }
| ifstmt                                                            { $inst = $ifstmt.ifinst }
| switchstmt                                                        { $inst = $switchstmt.swinst }
| whilestmt                                                         { $inst = $whilestmt.whileinst }
| forstmt                                                           { $inst = $forstmt.forinst }
| funvecstmt                                                        { $inst = $funvecstmt.fvecisnt }
| BREAK PCOMA?                                                      { $inst = instructions.NewBreak($BREAK.line, $BREAK.pos) }
| CONTINUE PCOMA?                                                   { $inst = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos) }
| RETURN expr PCOMA?                                                { $inst = instructions.NewReturn($RETURN.line, $RETURN.pos, $expr.e) }
| funcallstmt                                                       { $inst = $funcallstmt.cfun }
//globales
| funcstmt                                                          { $inst = $funcstmt.fun }
| structCreation                                                    { $inst = $structCreation.dec }
;

//STRUCT
structCreation returns[interfaces.Instruction dec]
: STRUCT ID LLAVEIZQ listStructDec LLAVEDER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
;

listStructDec returns[[]interface{} l]
: list=listStructDec VAR ID DOSP typestmt {
        var arr []interface{}
        newParams := environment.NewStructType($ID.text, $typestmt.type)
        arr = append($list.l, newParams)
        $l = arr
    }
| VAR ID DOSP typestmt {
        var arr []interface{}
        newParams := environment.NewStructType($ID.text, $typestmt.type)
        arr = append(arr, newParams)
        $l = arr
    }
|  { $l = []interface{}{} }
;


// FUNCIONES- DECLARACION
funcstmt returns [interfaces.Instruction fun]
: FUNC ID  PARIZQ PARDER LLAVEIZQ block LLAVEDER                                    { $fun = instructions.NewDeclarationFunc($FUNC.line, $FUNC.pos, $ID.text, environment.NULL,  nil ,$block.blk) }
| FUNC ID PARIZQ listParamsFunc PARDER LLAVEIZQ block LLAVEDER                      { $fun = instructions.NewDeclarationFunc($FUNC.line, $FUNC.pos, $ID.text, environment.NULL, $listParamsFunc.lpf, $block.blk) }
| FUNC ID  PARIZQ PARDER FLECHA typestmt LLAVEIZQ block LLAVEDER                    { $fun = instructions.NewDeclarationFunc($FUNC.line, $FUNC.pos, $ID.text, $typestmt.type,  nil ,$block.blk) }
| FUNC ID PARIZQ listParamsFunc PARDER FLECHA typestmt LLAVEIZQ block LLAVEDER      { $fun = instructions.NewDeclarationFunc($FUNC.line, $FUNC.pos, $ID.text, $typestmt.type, $listParamsFunc.lpf, $block.blk) }
//| FUNC ID PARIZQ listParamsFunc PARDER FLECHA LLAVEIZQ block LLAVEDER            { $fun = instructions.NewDeclarationFunc($FUNC.line, $FUNC.pos, $ID.text, environment.STRUCT, $listParamsFunc.lpf, $block.blk) }
;

listParamsFunc returns[[]interface{} lpf]
: list=listParamsFunc COMA id1=(ID|GBAJO)? id2=ID DOSP INOUT? typestmt {
    var arr []interface{}
    newParam := instructions.NewParamsDeclaration($id2.line, $id2.pos, $id1.text,$id2.text,$INOUT.text, $typestmt.type)
    arr = append($list.lpf, newParam)
    $lpf = arr
    }
| id1=(ID|GBAJO)? id2=ID DOSP INOUT? typestmt {
    $lpf = []interface{}{}
    newParam := instructions.NewParamsDeclaration($id2.line, $id2.pos, $id1.text,$id2.text,$INOUT.text, $typestmt.type)
    $lpf = append($lpf, newParam)
    }
|  { $lpf = []interface{}{} }
;


//LLAMADA DE FUNCIONES (instrucion)
funcallstmt returns[interfaces.Instruction cfun]
: ID PARIZQ PARDER                                          { $cfun = instructions.NewFunCall($ID.line, $ID.pos, $ID.text, nil) }
| ID PARIZQ listParams PARDER                               { $cfun = instructions.NewFunCall($ID.line, $ID.pos, $ID.text, $listParams.l) }
;


//LLAMADA DE FUNCIONES (expresion)
funcallestmt returns[interfaces.Expression cefun]
: ID PARIZQ PARDER                                          { $cefun = expressions.NewFunCallE($ID.line, $ID.pos, $ID.text, nil) }
| ID PARIZQ listParams PARDER                               { $cefun = expressions.NewFunCallE($ID.line, $ID.pos, $ID.text, $listParams.l) }
;


//PRINT
printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER                                  { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
| PRINT PARIZQ listParams PARDER                            { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,expressions.NewArray($PARIZQ.line, $PARIZQ.pos, $listParams.l))}
;

//DECLARACIONES
variablestmt returns [interfaces.Instruction vari]
//variables
: VAR ID DOSP typestmt IG expr                                  { $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $typestmt.type, $expr.e, true) }
| VAR ID IG expr                                                { $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, environment.NULL, $expr.e, true) }
| VAR ID DOSP typestmt INTCE                                    { $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $typestmt.type, expressions.NewPrimitive($ID.line, $ID.pos, nil, environment.NULL), true) }
//inmutables
| LET ID IG primitives                                          { $vari = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.NULL, $primitives.p, false) }
| LET ID DOSP typestmt IG primitives                            { $vari = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $typestmt.type, $primitives.p, false) }
| LET ID IG expr                                                { $vari = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.NULL, $expr.e, false) }
| LET ID DOSP typestmt IG expr                                  { $vari = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $typestmt.type, $expr.e, false) }
| LET ID IG expr                                                { $vari = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.NULL, $expr.e, false) }
//asignacion
| ID IG expr                                                    { $vari = instructions.NewAssigment($ID.line, $ID.pos, $ID.text, $expr.e) }
| ID CORIZQ e1=expr CORDER IG e2=expr                           { $vari = instructions.NewAssigmentVec($ID.line, $ID.pos, $ID.text, $e1.e, $e2.e) }
| ID e1=expr IG e2=expr                                         { $vari = instructions.NewAssigmentMa($ID.line, $ID.pos, $ID.text, $e1.e, $e2.e) }
//Operador de asignacion
| ID ADD IG expr                                                { $vari = instructions.NewSumAssigment($ID.line, $ID.pos, $ID.text, $expr.e) }
| ID SUB IG expr                                                { $vari = instructions.NewResAssigment($ID.line, $ID.pos, $ID.text, $expr.e) }
//vectores
| VAR ID DOSP typestmt IG expr                                  { $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $typestmt.type, $expr.e, true) }
| VAR ID DOSP typestmt IG CORIZQ CORDER 
{
    $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $typestmt.type, 
    expressions.NewArray($CORIZQ.line, $CORIZQ.pos, nil),
    true)
}
//Struct
| VAR ID IG expr                                                { $vari = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, environment.STRUCT , $expr.e,true) }
;

//TIPOS
typestmt returns [environment.TipoExpresion type]
: INT                       { $type = environment.INTEGER }
| FLOAT                     { $type = environment.FLOAT }
| TSTRING                   { $type = environment.STRING }
| BOOL                      { $type = environment.BOOLEAN }
| CHAR                      { $type = environment.CHAR }
| CORIZQ type2stmt CORDER   { $type = $type2stmt.type2 }
;

type2stmt returns [environment.TipoExpresion type2]
: INT                       { $type2 = environment.A_INTEGER }
| FLOAT                     { $type2 = environment.A_FLOAT }
| TSTRING                   { $type2 = environment.A_STRING }
| BOOL                      { $type2 = environment.A_BOOLEAN }
| CHAR                      { $type2 = environment.A_CHAR }
| CORIZQ type2stmt CORDER   { $type2 = $type2stmt.type2 }
;

//Funciones vector
funvecstmt returns [interfaces.Instruction fvecisnt]
: ID PUNTO APPEND PARIZQ expr PARDER            { $fvecisnt = instructions.NewAppend($ID.line, $ID.pos, $ID.text , $expr.e) }
| ID PUNTO REMOVELAST PARIZQ PARDER             { $fvecisnt = instructions.NewRemoveLast($ID.line, $ID.pos, $ID.text) }
| ID PUNTO REMOVE PARIZQ AT DOSP expr PARDER    { $fvecisnt = instructions.NewRemove($ID.line, $ID.pos, $ID.text , $expr.e) }
;

//GUARD
guardstmt returns [interfaces.Instruction guinst]
:GUARD expr ELSE LLAVEIZQ block LLAVEDER        { $guinst = instructions.NewGuard($GUARD.line, $GUARD.pos, $expr.e, $block.blk) }
;

//IF
ifstmt returns [interfaces.Instruction ifinst]
: IF expr LLAVEIZQ block LLAVEDER                                                   { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil,nil) }
| IF expr LLAVEIZQ bif=block LLAVEDER ELSE LLAVEIZQ belse=block LLAVEDER            { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $bif.blk, $belse.blk,nil) }
| IF expr LLAVEIZQ bif=block LLAVEDER elifs                                         { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $bif.blk, nil, $elifs.elifinst) }
| IF expr LLAVEIZQ bif=block LLAVEDER elifs ELSE LLAVEIZQ belse=block LLAVEDER      { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $bif.blk, $belse.blk, $elifs.elifinst) }
;
//ELSE IF
elifs returns [[]interface{} elifinst]
:celif=elifs ELSE IF expr LLAVEIZQ block LLAVEDER 
{
    var arrif []interface{}
    arrif = append($celif.elifinst, instructions.NewElif($ELSE.line, $ELSE.pos, $expr.e, $block.blk))
    $elifinst = arrif
}   
| ELSE IF expr LLAVEIZQ block LLAVEDER
{
    $elifinst = []interface{}{}
    $elifinst = append($elifinst, instructions.NewElif($ELSE.line, $ELSE.pos, $expr.e, $block.blk))
    
}
;
//SWITCH
switchstmt returns [interfaces.Instruction swinst]
: SWITCH expr LLAVEIZQ cases DEFAULT DOSP block LLAVEDER        { $swinst = instructions.NewSwitch($SWITCH.line, $SWITCH.pos, $expr.e, $cases.casesinst, $block.blk) }
| SWITCH expr LLAVEIZQ cases LLAVEDER                           { $swinst = instructions.NewSwitch($SWITCH.line, $SWITCH.pos, $expr.e, $cases.casesinst, nil) }
;
//CASES
cases returns [[]interface{} casesinst]
:ccases=cases CASE expr DOSP block
{
    var arrcase []interface{}
    arrcase = append($ccases.casesinst, instructions.NewCase($CASE.line, $CASE.pos, $expr.e, $block.blk))
    $casesinst = arrcase
}   
|CASE expr DOSP block
{
    $casesinst = []interface{}{}
    $casesinst = append($casesinst, instructions.NewCase($CASE.line, $CASE.pos, $expr.e, $block.blk))
    
}
;

//WHILE
whilestmt returns [interfaces.Instruction whileinst]
: WHILE expr LLAVEIZQ block LLAVEDER                        { $whileinst = instructions.NewWhile($WHILE.line, $WHILE.pos, $expr.e, $block.blk) }
;

//FOR
forstmt returns [interfaces.Instruction forinst]
: FOR ID IN e1=expr TRESP e2=expr LLAVEIZQ block LLAVEDER                     { $forinst = instructions.NewForRange($FOR.line, $FOR.pos,$ID.text ,$e1.e,$e2.e ,$block.blk) }
| FOR ID IN expr LLAVEIZQ block LLAVEDER                                                  { $forinst = instructions.NewFor($FOR.line, $FOR.pos,$ID.text ,$expr.e ,$block.blk) }
;

//EXPRESIONES
expr returns [interfaces.Expression e]
//operaciones
: left=expr op=(MUL|DIV|MOD) right=expr                     { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr                         { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr                    { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr                    { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr                       { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr                               { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr                                { $e = expressions.NewOperationBinary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| op=NOT left=expr                                          { $e = expressions.NewOperationUnary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text) }
| op=SUB left=expr                                          { $e = expressions.NewOperationUnary($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text) }
| PARIZQ expr PARDER                                        { $e = $expr.e }
| ID                                                        { $e = expressions.NewCallVariable($ID.line, $ID.pos, $ID.text)}   
| primitives                                                { $e = $primitives.p}
//casteos
| INT PARIZQ expr PARDER                                  { $e = expressions.NewCastingInt($INT.line, $INT.pos, $expr.e ) }
| FLOAT PARIZQ expr PARDER                                { $e = expressions.NewCastingFloat($FLOAT.line, $FLOAT.pos, $expr.e ) }
| TSTRING PARIZQ expr PARDER                              { $e = expressions.NewCastingString($TSTRING.line, $TSTRING.pos, $expr.e ) }
//cosas de array
| listArray                                                 { $e = $listArray.p}
| CORIZQ listParams CORDER                                  { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| ID PUNTO COUNT                                            { $e = expressions.NewCount($ID.line, $ID.pos, $ID.text) }
| ID PUNTO ISEMPTY                                          { $e = expressions.NewIsEmpty($ID.line, $ID.pos, $ID.text) }
| listAceso                                                 { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listAceso.l) }
| ID PARIZQ listStructExp PARDER                            { $e = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
//Funcion
| funcallestmt PCOMA                                        { $e = $funcallestmt.cefun }
;

//EXPRESIONES PRIMITIVAS
primitives returns [interfaces.Expression p]
: NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| CHARACTER
    {
        str := $CHARACTER.text
        $p = expressions.NewPrimitive($CHARACTER.line, $CHARACTER.pos, str[1:len(str)-1],environment.CHAR)
    }
| STRING
    {
        str := $STRING.text
        $p = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                    
| TRU                                                       { $p = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL                                                       { $p = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| NIL                                                       { $p = expressions.NewPrimitive($NIL.line, $NIL.pos, nil, environment.NULL) }
;

//  LISTA DE PARAMETROS
listParams returns[[]interface{} l]
: list=listParams COMA AMP? expr 
{
    var arr []interface{}
    arr = append($list.l, $expr.e)
    $l = arr
}   
| AMP? expr 
{
    $l = []interface{}{}
    $l = append($l, $expr.e)
}
| list=listParams COMA expr DOSP AMP? id=expr
{
    var arr []interface{}
    arr = append($list.l, $id.e)
    $l = arr
}   
| expr PCOMA AMP? id=expr
{
    $l = []interface{}{}
    $l = append($l, $id.e)
}
;

//ACCEDER ARRAY
listArray returns[interfaces.Expression p]
: list = listArray CORIZQ expr CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expr.e) }
| ID { $p = expressions.NewCallVariable($ID.line, $ID.pos, $ID.text)}
//aceder struct
| list = listArray PUNTO ID { $p = expressions.NewStructAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $ID.text)  }
;

//Asignacion de vectores
listAceso returns[[]interface{} l]
: list=listAceso CORIZQ expr CORDER
{
    var arr []interface{}
    arr = append($list.l, $expr.e)
    $l = arr
}   
| CORIZQ expr CORDER
{
    $l = []interface{}{}
    $l = append($l, $expr.e)
}
;
//Lista de expresiones Struct
listStructExp returns[[]interface{} l]
: list=listStructExp COMA ID DOSP expr {
    var arr []interface{}
    StrExp := environment.NewStructContent($ID.text, $expr.e)
    arr = append($list.l, StrExp)
    $l = arr
}   
| ID DOSP expr{
    var arr []interface{}
    StrExp := environment.NewStructContent($ID.text, $expr.e)
    arr = append(arr, StrExp)
    $l = arr
}
|   {
    $l = []interface{}{}
}
;