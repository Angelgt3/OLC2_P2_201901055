// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Backend/interfaces"
import "Backend/environment"
import "Backend/expressions"
import "Backend/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'var'",
		"'let'", "'true'", "'false'", "'print'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'break'", "'continue'", "'return'", "'while'",
		"'for'", "'in'", "'guard'", "'append'", "'removeLast'", "'remove'",
		"'isEmpty'", "'count'", "'repeating'", "'struct'", "'self'", "'mutating'",
		"'func'", "'at'", "'inout'", "'nil'", "", "", "", "", "'!='", "'=='",
		"'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'%'", "'*'",
		"'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'",
		"'.'", "','", "'?'", "'->'", "'&'", "';'", "'_'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "TSTRING", "BOOL", "CHAR", "VAR", "LET", "TRU",
		"FAL", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK",
		"CONTINUE", "RETURN", "WHILE", "FOR", "IN", "GUARD", "APPEND", "REMOVELAST",
		"REMOVE", "ISEMPTY", "COUNT", "REPEATING", "STRUCT", "SELF", "MUTATING",
		"FUNC", "AT", "INOUT", "NIL", "NUMBER", "CHARACTER", "STRING", "ID",
		"DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MOD", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "CORIZQ", "CORDER", "DOSP", "PUNTO", "COMA", "INTCE", "FLECHA",
		"AMP", "PCOMA", "GBAJO", "TRESP", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "structCreation", "listStructDec", "funcstmt",
		"listParamsFunc", "funcallstmt", "funcallestmt", "printstmt", "variablestmt",
		"typestmt", "type2stmt", "funvecstmt", "guardstmt", "ifstmt", "elifs",
		"switchstmt", "cases", "whilestmt", "forstmt", "expr", "primitives",
		"listParams", "listArray", "listAceso", "listStructExp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 769, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 60, 8, 1, 11, 1, 12, 1, 61,
		1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 68, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 74,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		98, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 109, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 122, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 139, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 148, 8, 4, 10, 4, 12, 4, 151, 9,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 195, 8, 5, 1, 6, 1, 6, 3,
		6, 199, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 204, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 210, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 215, 8, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 220, 8, 6, 1, 6, 1, 6, 1, 6, 5, 6, 225, 8, 6, 10, 6, 12, 6, 228, 9,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 240,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		252, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 266, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 377,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 394, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 411, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 437,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 485, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 505, 8, 16, 10, 16, 12, 16, 508,
		9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 527, 8, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 5, 18, 543, 8, 18, 10, 18, 12, 18, 546, 9, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 575, 8, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 3, 21, 642, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 679, 8, 21, 10,
		21, 12, 21, 682, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 696, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 707, 8, 23, 10, 23,
		12, 23, 710, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 726, 8, 24, 10, 24,
		12, 24, 729, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 743, 8, 25, 10, 25, 12, 25, 746,
		9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 755, 8,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 764, 8, 26,
		10, 26, 12, 26, 767, 9, 26, 1, 26, 0, 9, 8, 12, 32, 36, 42, 46, 48, 50,
		52, 27, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 0, 6, 2, 0, 39, 39, 68, 68, 1,
		0, 50, 52, 1, 0, 53, 54, 2, 0, 46, 46, 48, 48, 2, 0, 47, 47, 49, 49, 1,
		0, 40, 41, 840, 0, 54, 1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4, 121, 1, 0, 0,
		0, 6, 123, 1, 0, 0, 0, 8, 138, 1, 0, 0, 0, 10, 194, 1, 0, 0, 0, 12, 209,
		1, 0, 0, 0, 14, 239, 1, 0, 0, 0, 16, 251, 1, 0, 0, 0, 18, 265, 1, 0, 0,
		0, 20, 376, 1, 0, 0, 0, 22, 393, 1, 0, 0, 0, 24, 410, 1, 0, 0, 0, 26, 436,
		1, 0, 0, 0, 28, 438, 1, 0, 0, 0, 30, 484, 1, 0, 0, 0, 32, 486, 1, 0, 0,
		0, 34, 526, 1, 0, 0, 0, 36, 528, 1, 0, 0, 0, 38, 547, 1, 0, 0, 0, 40, 574,
		1, 0, 0, 0, 42, 641, 1, 0, 0, 0, 44, 695, 1, 0, 0, 0, 46, 697, 1, 0, 0,
		0, 48, 711, 1, 0, 0, 0, 50, 730, 1, 0, 0, 0, 52, 754, 1, 0, 0, 0, 54, 55,
		3, 2, 1, 0, 55, 56, 5, 0, 0, 1, 56, 57, 6, 0, -1, 0, 57, 1, 1, 0, 0, 0,
		58, 60, 3, 4, 2, 0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1,
		0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 6, 1, -1, 0, 64,
		3, 1, 0, 0, 0, 65, 67, 3, 18, 9, 0, 66, 68, 5, 67, 0, 0, 67, 66, 1, 0,
		0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 6, 2, -1, 0, 70,
		122, 1, 0, 0, 0, 71, 73, 3, 20, 10, 0, 72, 74, 5, 67, 0, 0, 73, 72, 1,
		0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 6, 2, -1, 0, 76,
		122, 1, 0, 0, 0, 77, 78, 3, 28, 14, 0, 78, 79, 6, 2, -1, 0, 79, 122, 1,
		0, 0, 0, 80, 81, 3, 30, 15, 0, 81, 82, 6, 2, -1, 0, 82, 122, 1, 0, 0, 0,
		83, 84, 3, 34, 17, 0, 84, 85, 6, 2, -1, 0, 85, 122, 1, 0, 0, 0, 86, 87,
		3, 38, 19, 0, 87, 88, 6, 2, -1, 0, 88, 122, 1, 0, 0, 0, 89, 90, 3, 40,
		20, 0, 90, 91, 6, 2, -1, 0, 91, 122, 1, 0, 0, 0, 92, 93, 3, 26, 13, 0,
		93, 94, 6, 2, -1, 0, 94, 122, 1, 0, 0, 0, 95, 97, 5, 16, 0, 0, 96, 98,
		5, 67, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 122, 6, 2, -1, 0, 100, 102, 5, 17, 0, 0, 101, 103, 5, 67, 0, 0, 102,
		101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 122,
		6, 2, -1, 0, 105, 106, 5, 18, 0, 0, 106, 108, 3, 42, 21, 0, 107, 109, 5,
		67, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0,
		0, 110, 111, 6, 2, -1, 0, 111, 122, 1, 0, 0, 0, 112, 113, 3, 14, 7, 0,
		113, 114, 6, 2, -1, 0, 114, 122, 1, 0, 0, 0, 115, 116, 3, 10, 5, 0, 116,
		117, 6, 2, -1, 0, 117, 122, 1, 0, 0, 0, 118, 119, 3, 6, 3, 0, 119, 120,
		6, 2, -1, 0, 120, 122, 1, 0, 0, 0, 121, 65, 1, 0, 0, 0, 121, 71, 1, 0,
		0, 0, 121, 77, 1, 0, 0, 0, 121, 80, 1, 0, 0, 0, 121, 83, 1, 0, 0, 0, 121,
		86, 1, 0, 0, 0, 121, 89, 1, 0, 0, 0, 121, 92, 1, 0, 0, 0, 121, 95, 1, 0,
		0, 0, 121, 100, 1, 0, 0, 0, 121, 105, 1, 0, 0, 0, 121, 112, 1, 0, 0, 0,
		121, 115, 1, 0, 0, 0, 121, 118, 1, 0, 0, 0, 122, 5, 1, 0, 0, 0, 123, 124,
		5, 29, 0, 0, 124, 125, 5, 39, 0, 0, 125, 126, 5, 57, 0, 0, 126, 127, 3,
		8, 4, 0, 127, 128, 5, 58, 0, 0, 128, 129, 6, 3, -1, 0, 129, 7, 1, 0, 0,
		0, 130, 131, 6, 4, -1, 0, 131, 132, 5, 6, 0, 0, 132, 133, 5, 39, 0, 0,
		133, 134, 5, 61, 0, 0, 134, 135, 3, 22, 11, 0, 135, 136, 6, 4, -1, 0, 136,
		139, 1, 0, 0, 0, 137, 139, 6, 4, -1, 0, 138, 130, 1, 0, 0, 0, 138, 137,
		1, 0, 0, 0, 139, 149, 1, 0, 0, 0, 140, 141, 10, 3, 0, 0, 141, 142, 5, 6,
		0, 0, 142, 143, 5, 39, 0, 0, 143, 144, 5, 61, 0, 0, 144, 145, 3, 22, 11,
		0, 145, 146, 6, 4, -1, 0, 146, 148, 1, 0, 0, 0, 147, 140, 1, 0, 0, 0, 148,
		151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 9, 1,
		0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 32, 0, 0, 153, 154, 5, 39,
		0, 0, 154, 155, 5, 55, 0, 0, 155, 156, 5, 56, 0, 0, 156, 157, 5, 57, 0,
		0, 157, 158, 3, 2, 1, 0, 158, 159, 5, 58, 0, 0, 159, 160, 6, 5, -1, 0,
		160, 195, 1, 0, 0, 0, 161, 162, 5, 32, 0, 0, 162, 163, 5, 39, 0, 0, 163,
		164, 5, 55, 0, 0, 164, 165, 3, 12, 6, 0, 165, 166, 5, 56, 0, 0, 166, 167,
		5, 57, 0, 0, 167, 168, 3, 2, 1, 0, 168, 169, 5, 58, 0, 0, 169, 170, 6,
		5, -1, 0, 170, 195, 1, 0, 0, 0, 171, 172, 5, 32, 0, 0, 172, 173, 5, 39,
		0, 0, 173, 174, 5, 55, 0, 0, 174, 175, 5, 56, 0, 0, 175, 176, 5, 65, 0,
		0, 176, 177, 3, 22, 11, 0, 177, 178, 5, 57, 0, 0, 178, 179, 3, 2, 1, 0,
		179, 180, 5, 58, 0, 0, 180, 181, 6, 5, -1, 0, 181, 195, 1, 0, 0, 0, 182,
		183, 5, 32, 0, 0, 183, 184, 5, 39, 0, 0, 184, 185, 5, 55, 0, 0, 185, 186,
		3, 12, 6, 0, 186, 187, 5, 56, 0, 0, 187, 188, 5, 65, 0, 0, 188, 189, 3,
		22, 11, 0, 189, 190, 5, 57, 0, 0, 190, 191, 3, 2, 1, 0, 191, 192, 5, 58,
		0, 0, 192, 193, 6, 5, -1, 0, 193, 195, 1, 0, 0, 0, 194, 152, 1, 0, 0, 0,
		194, 161, 1, 0, 0, 0, 194, 171, 1, 0, 0, 0, 194, 182, 1, 0, 0, 0, 195,
		11, 1, 0, 0, 0, 196, 198, 6, 6, -1, 0, 197, 199, 7, 0, 0, 0, 198, 197,
		1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 5, 39,
		0, 0, 201, 203, 5, 61, 0, 0, 202, 204, 5, 34, 0, 0, 203, 202, 1, 0, 0,
		0, 203, 204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 3, 22, 11, 0,
		206, 207, 6, 6, -1, 0, 207, 210, 1, 0, 0, 0, 208, 210, 6, 6, -1, 0, 209,
		196, 1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 226, 1, 0, 0, 0, 211, 212,
		10, 3, 0, 0, 212, 214, 5, 63, 0, 0, 213, 215, 7, 0, 0, 0, 214, 213, 1,
		0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 5, 39, 0,
		0, 217, 219, 5, 61, 0, 0, 218, 220, 5, 34, 0, 0, 219, 218, 1, 0, 0, 0,
		219, 220, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 3, 22, 11, 0, 222,
		223, 6, 6, -1, 0, 223, 225, 1, 0, 0, 0, 224, 211, 1, 0, 0, 0, 225, 228,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 13, 1, 0,
		0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 5, 39, 0, 0, 230, 231, 5, 55, 0,
		0, 231, 232, 5, 56, 0, 0, 232, 240, 6, 7, -1, 0, 233, 234, 5, 39, 0, 0,
		234, 235, 5, 55, 0, 0, 235, 236, 3, 46, 23, 0, 236, 237, 5, 56, 0, 0, 237,
		238, 6, 7, -1, 0, 238, 240, 1, 0, 0, 0, 239, 229, 1, 0, 0, 0, 239, 233,
		1, 0, 0, 0, 240, 15, 1, 0, 0, 0, 241, 242, 5, 39, 0, 0, 242, 243, 5, 55,
		0, 0, 243, 244, 5, 56, 0, 0, 244, 252, 6, 8, -1, 0, 245, 246, 5, 39, 0,
		0, 246, 247, 5, 55, 0, 0, 247, 248, 3, 46, 23, 0, 248, 249, 5, 56, 0, 0,
		249, 250, 6, 8, -1, 0, 250, 252, 1, 0, 0, 0, 251, 241, 1, 0, 0, 0, 251,
		245, 1, 0, 0, 0, 252, 17, 1, 0, 0, 0, 253, 254, 5, 10, 0, 0, 254, 255,
		5, 55, 0, 0, 255, 256, 3, 42, 21, 0, 256, 257, 5, 56, 0, 0, 257, 258, 6,
		9, -1, 0, 258, 266, 1, 0, 0, 0, 259, 260, 5, 10, 0, 0, 260, 261, 5, 55,
		0, 0, 261, 262, 3, 46, 23, 0, 262, 263, 5, 56, 0, 0, 263, 264, 6, 9, -1,
		0, 264, 266, 1, 0, 0, 0, 265, 253, 1, 0, 0, 0, 265, 259, 1, 0, 0, 0, 266,
		19, 1, 0, 0, 0, 267, 268, 5, 6, 0, 0, 268, 269, 5, 39, 0, 0, 269, 270,
		5, 61, 0, 0, 270, 271, 3, 22, 11, 0, 271, 272, 5, 45, 0, 0, 272, 273, 3,
		42, 21, 0, 273, 274, 6, 10, -1, 0, 274, 377, 1, 0, 0, 0, 275, 276, 5, 6,
		0, 0, 276, 277, 5, 39, 0, 0, 277, 278, 5, 45, 0, 0, 278, 279, 3, 42, 21,
		0, 279, 280, 6, 10, -1, 0, 280, 377, 1, 0, 0, 0, 281, 282, 5, 6, 0, 0,
		282, 283, 5, 39, 0, 0, 283, 284, 5, 61, 0, 0, 284, 285, 3, 22, 11, 0, 285,
		286, 5, 64, 0, 0, 286, 287, 6, 10, -1, 0, 287, 377, 1, 0, 0, 0, 288, 289,
		5, 7, 0, 0, 289, 290, 5, 39, 0, 0, 290, 291, 5, 45, 0, 0, 291, 292, 3,
		44, 22, 0, 292, 293, 6, 10, -1, 0, 293, 377, 1, 0, 0, 0, 294, 295, 5, 7,
		0, 0, 295, 296, 5, 39, 0, 0, 296, 297, 5, 61, 0, 0, 297, 298, 3, 22, 11,
		0, 298, 299, 5, 45, 0, 0, 299, 300, 3, 44, 22, 0, 300, 301, 6, 10, -1,
		0, 301, 377, 1, 0, 0, 0, 302, 303, 5, 7, 0, 0, 303, 304, 5, 39, 0, 0, 304,
		305, 5, 45, 0, 0, 305, 306, 3, 42, 21, 0, 306, 307, 6, 10, -1, 0, 307,
		377, 1, 0, 0, 0, 308, 309, 5, 7, 0, 0, 309, 310, 5, 39, 0, 0, 310, 311,
		5, 61, 0, 0, 311, 312, 3, 22, 11, 0, 312, 313, 5, 45, 0, 0, 313, 314, 3,
		42, 21, 0, 314, 315, 6, 10, -1, 0, 315, 377, 1, 0, 0, 0, 316, 317, 5, 7,
		0, 0, 317, 318, 5, 39, 0, 0, 318, 319, 5, 45, 0, 0, 319, 320, 3, 42, 21,
		0, 320, 321, 6, 10, -1, 0, 321, 377, 1, 0, 0, 0, 322, 323, 5, 39, 0, 0,
		323, 324, 5, 45, 0, 0, 324, 325, 3, 42, 21, 0, 325, 326, 6, 10, -1, 0,
		326, 377, 1, 0, 0, 0, 327, 328, 5, 39, 0, 0, 328, 329, 5, 59, 0, 0, 329,
		330, 3, 42, 21, 0, 330, 331, 5, 60, 0, 0, 331, 332, 5, 45, 0, 0, 332, 333,
		3, 42, 21, 0, 333, 334, 6, 10, -1, 0, 334, 377, 1, 0, 0, 0, 335, 336, 5,
		39, 0, 0, 336, 337, 3, 42, 21, 0, 337, 338, 5, 45, 0, 0, 338, 339, 3, 42,
		21, 0, 339, 340, 6, 10, -1, 0, 340, 377, 1, 0, 0, 0, 341, 342, 5, 39, 0,
		0, 342, 343, 5, 53, 0, 0, 343, 344, 5, 45, 0, 0, 344, 345, 3, 42, 21, 0,
		345, 346, 6, 10, -1, 0, 346, 377, 1, 0, 0, 0, 347, 348, 5, 39, 0, 0, 348,
		349, 5, 54, 0, 0, 349, 350, 5, 45, 0, 0, 350, 351, 3, 42, 21, 0, 351, 352,
		6, 10, -1, 0, 352, 377, 1, 0, 0, 0, 353, 354, 5, 6, 0, 0, 354, 355, 5,
		39, 0, 0, 355, 356, 5, 61, 0, 0, 356, 357, 3, 22, 11, 0, 357, 358, 5, 45,
		0, 0, 358, 359, 3, 42, 21, 0, 359, 360, 6, 10, -1, 0, 360, 377, 1, 0, 0,
		0, 361, 362, 5, 6, 0, 0, 362, 363, 5, 39, 0, 0, 363, 364, 5, 61, 0, 0,
		364, 365, 3, 22, 11, 0, 365, 366, 5, 45, 0, 0, 366, 367, 5, 59, 0, 0, 367,
		368, 5, 60, 0, 0, 368, 369, 6, 10, -1, 0, 369, 377, 1, 0, 0, 0, 370, 371,
		5, 6, 0, 0, 371, 372, 5, 39, 0, 0, 372, 373, 5, 45, 0, 0, 373, 374, 3,
		42, 21, 0, 374, 375, 6, 10, -1, 0, 375, 377, 1, 0, 0, 0, 376, 267, 1, 0,
		0, 0, 376, 275, 1, 0, 0, 0, 376, 281, 1, 0, 0, 0, 376, 288, 1, 0, 0, 0,
		376, 294, 1, 0, 0, 0, 376, 302, 1, 0, 0, 0, 376, 308, 1, 0, 0, 0, 376,
		316, 1, 0, 0, 0, 376, 322, 1, 0, 0, 0, 376, 327, 1, 0, 0, 0, 376, 335,
		1, 0, 0, 0, 376, 341, 1, 0, 0, 0, 376, 347, 1, 0, 0, 0, 376, 353, 1, 0,
		0, 0, 376, 361, 1, 0, 0, 0, 376, 370, 1, 0, 0, 0, 377, 21, 1, 0, 0, 0,
		378, 379, 5, 1, 0, 0, 379, 394, 6, 11, -1, 0, 380, 381, 5, 2, 0, 0, 381,
		394, 6, 11, -1, 0, 382, 383, 5, 3, 0, 0, 383, 394, 6, 11, -1, 0, 384, 385,
		5, 4, 0, 0, 385, 394, 6, 11, -1, 0, 386, 387, 5, 5, 0, 0, 387, 394, 6,
		11, -1, 0, 388, 389, 5, 59, 0, 0, 389, 390, 3, 24, 12, 0, 390, 391, 5,
		60, 0, 0, 391, 392, 6, 11, -1, 0, 392, 394, 1, 0, 0, 0, 393, 378, 1, 0,
		0, 0, 393, 380, 1, 0, 0, 0, 393, 382, 1, 0, 0, 0, 393, 384, 1, 0, 0, 0,
		393, 386, 1, 0, 0, 0, 393, 388, 1, 0, 0, 0, 394, 23, 1, 0, 0, 0, 395, 396,
		5, 1, 0, 0, 396, 411, 6, 12, -1, 0, 397, 398, 5, 2, 0, 0, 398, 411, 6,
		12, -1, 0, 399, 400, 5, 3, 0, 0, 400, 411, 6, 12, -1, 0, 401, 402, 5, 4,
		0, 0, 402, 411, 6, 12, -1, 0, 403, 404, 5, 5, 0, 0, 404, 411, 6, 12, -1,
		0, 405, 406, 5, 59, 0, 0, 406, 407, 3, 24, 12, 0, 407, 408, 5, 60, 0, 0,
		408, 409, 6, 12, -1, 0, 409, 411, 1, 0, 0, 0, 410, 395, 1, 0, 0, 0, 410,
		397, 1, 0, 0, 0, 410, 399, 1, 0, 0, 0, 410, 401, 1, 0, 0, 0, 410, 403,
		1, 0, 0, 0, 410, 405, 1, 0, 0, 0, 411, 25, 1, 0, 0, 0, 412, 413, 5, 39,
		0, 0, 413, 414, 5, 62, 0, 0, 414, 415, 5, 23, 0, 0, 415, 416, 5, 55, 0,
		0, 416, 417, 3, 42, 21, 0, 417, 418, 5, 56, 0, 0, 418, 419, 6, 13, -1,
		0, 419, 437, 1, 0, 0, 0, 420, 421, 5, 39, 0, 0, 421, 422, 5, 62, 0, 0,
		422, 423, 5, 24, 0, 0, 423, 424, 5, 55, 0, 0, 424, 425, 5, 56, 0, 0, 425,
		437, 6, 13, -1, 0, 426, 427, 5, 39, 0, 0, 427, 428, 5, 62, 0, 0, 428, 429,
		5, 25, 0, 0, 429, 430, 5, 55, 0, 0, 430, 431, 5, 33, 0, 0, 431, 432, 5,
		61, 0, 0, 432, 433, 3, 42, 21, 0, 433, 434, 5, 56, 0, 0, 434, 435, 6, 13,
		-1, 0, 435, 437, 1, 0, 0, 0, 436, 412, 1, 0, 0, 0, 436, 420, 1, 0, 0, 0,
		436, 426, 1, 0, 0, 0, 437, 27, 1, 0, 0, 0, 438, 439, 5, 22, 0, 0, 439,
		440, 3, 42, 21, 0, 440, 441, 5, 12, 0, 0, 441, 442, 5, 57, 0, 0, 442, 443,
		3, 2, 1, 0, 443, 444, 5, 58, 0, 0, 444, 445, 6, 14, -1, 0, 445, 29, 1,
		0, 0, 0, 446, 447, 5, 11, 0, 0, 447, 448, 3, 42, 21, 0, 448, 449, 5, 57,
		0, 0, 449, 450, 3, 2, 1, 0, 450, 451, 5, 58, 0, 0, 451, 452, 6, 15, -1,
		0, 452, 485, 1, 0, 0, 0, 453, 454, 5, 11, 0, 0, 454, 455, 3, 42, 21, 0,
		455, 456, 5, 57, 0, 0, 456, 457, 3, 2, 1, 0, 457, 458, 5, 58, 0, 0, 458,
		459, 5, 12, 0, 0, 459, 460, 5, 57, 0, 0, 460, 461, 3, 2, 1, 0, 461, 462,
		5, 58, 0, 0, 462, 463, 6, 15, -1, 0, 463, 485, 1, 0, 0, 0, 464, 465, 5,
		11, 0, 0, 465, 466, 3, 42, 21, 0, 466, 467, 5, 57, 0, 0, 467, 468, 3, 2,
		1, 0, 468, 469, 5, 58, 0, 0, 469, 470, 3, 32, 16, 0, 470, 471, 6, 15, -1,
		0, 471, 485, 1, 0, 0, 0, 472, 473, 5, 11, 0, 0, 473, 474, 3, 42, 21, 0,
		474, 475, 5, 57, 0, 0, 475, 476, 3, 2, 1, 0, 476, 477, 5, 58, 0, 0, 477,
		478, 3, 32, 16, 0, 478, 479, 5, 12, 0, 0, 479, 480, 5, 57, 0, 0, 480, 481,
		3, 2, 1, 0, 481, 482, 5, 58, 0, 0, 482, 483, 6, 15, -1, 0, 483, 485, 1,
		0, 0, 0, 484, 446, 1, 0, 0, 0, 484, 453, 1, 0, 0, 0, 484, 464, 1, 0, 0,
		0, 484, 472, 1, 0, 0, 0, 485, 31, 1, 0, 0, 0, 486, 487, 6, 16, -1, 0, 487,
		488, 5, 12, 0, 0, 488, 489, 5, 11, 0, 0, 489, 490, 3, 42, 21, 0, 490, 491,
		5, 57, 0, 0, 491, 492, 3, 2, 1, 0, 492, 493, 5, 58, 0, 0, 493, 494, 6,
		16, -1, 0, 494, 506, 1, 0, 0, 0, 495, 496, 10, 2, 0, 0, 496, 497, 5, 12,
		0, 0, 497, 498, 5, 11, 0, 0, 498, 499, 3, 42, 21, 0, 499, 500, 5, 57, 0,
		0, 500, 501, 3, 2, 1, 0, 501, 502, 5, 58, 0, 0, 502, 503, 6, 16, -1, 0,
		503, 505, 1, 0, 0, 0, 504, 495, 1, 0, 0, 0, 505, 508, 1, 0, 0, 0, 506,
		504, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 33, 1, 0, 0, 0, 508, 506, 1,
		0, 0, 0, 509, 510, 5, 13, 0, 0, 510, 511, 3, 42, 21, 0, 511, 512, 5, 57,
		0, 0, 512, 513, 3, 36, 18, 0, 513, 514, 5, 15, 0, 0, 514, 515, 5, 61, 0,
		0, 515, 516, 3, 2, 1, 0, 516, 517, 5, 58, 0, 0, 517, 518, 6, 17, -1, 0,
		518, 527, 1, 0, 0, 0, 519, 520, 5, 13, 0, 0, 520, 521, 3, 42, 21, 0, 521,
		522, 5, 57, 0, 0, 522, 523, 3, 36, 18, 0, 523, 524, 5, 58, 0, 0, 524, 525,
		6, 17, -1, 0, 525, 527, 1, 0, 0, 0, 526, 509, 1, 0, 0, 0, 526, 519, 1,
		0, 0, 0, 527, 35, 1, 0, 0, 0, 528, 529, 6, 18, -1, 0, 529, 530, 5, 14,
		0, 0, 530, 531, 3, 42, 21, 0, 531, 532, 5, 61, 0, 0, 532, 533, 3, 2, 1,
		0, 533, 534, 6, 18, -1, 0, 534, 544, 1, 0, 0, 0, 535, 536, 10, 2, 0, 0,
		536, 537, 5, 14, 0, 0, 537, 538, 3, 42, 21, 0, 538, 539, 5, 61, 0, 0, 539,
		540, 3, 2, 1, 0, 540, 541, 6, 18, -1, 0, 541, 543, 1, 0, 0, 0, 542, 535,
		1, 0, 0, 0, 543, 546, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 544, 545, 1, 0,
		0, 0, 545, 37, 1, 0, 0, 0, 546, 544, 1, 0, 0, 0, 547, 548, 5, 19, 0, 0,
		548, 549, 3, 42, 21, 0, 549, 550, 5, 57, 0, 0, 550, 551, 3, 2, 1, 0, 551,
		552, 5, 58, 0, 0, 552, 553, 6, 19, -1, 0, 553, 39, 1, 0, 0, 0, 554, 555,
		5, 20, 0, 0, 555, 556, 5, 39, 0, 0, 556, 557, 5, 21, 0, 0, 557, 558, 3,
		42, 21, 0, 558, 559, 5, 69, 0, 0, 559, 560, 3, 42, 21, 0, 560, 561, 5,
		57, 0, 0, 561, 562, 3, 2, 1, 0, 562, 563, 5, 58, 0, 0, 563, 564, 6, 20,
		-1, 0, 564, 575, 1, 0, 0, 0, 565, 566, 5, 20, 0, 0, 566, 567, 5, 39, 0,
		0, 567, 568, 5, 21, 0, 0, 568, 569, 3, 42, 21, 0, 569, 570, 5, 57, 0, 0,
		570, 571, 3, 2, 1, 0, 571, 572, 5, 58, 0, 0, 572, 573, 6, 20, -1, 0, 573,
		575, 1, 0, 0, 0, 574, 554, 1, 0, 0, 0, 574, 565, 1, 0, 0, 0, 575, 41, 1,
		0, 0, 0, 576, 577, 6, 21, -1, 0, 577, 578, 5, 42, 0, 0, 578, 579, 3, 42,
		21, 15, 579, 580, 6, 21, -1, 0, 580, 642, 1, 0, 0, 0, 581, 582, 5, 54,
		0, 0, 582, 583, 3, 42, 21, 14, 583, 584, 6, 21, -1, 0, 584, 642, 1, 0,
		0, 0, 585, 586, 5, 55, 0, 0, 586, 587, 3, 42, 21, 0, 587, 588, 5, 56, 0,
		0, 588, 589, 6, 21, -1, 0, 589, 642, 1, 0, 0, 0, 590, 591, 5, 39, 0, 0,
		591, 642, 6, 21, -1, 0, 592, 593, 3, 44, 22, 0, 593, 594, 6, 21, -1, 0,
		594, 642, 1, 0, 0, 0, 595, 596, 5, 1, 0, 0, 596, 597, 5, 55, 0, 0, 597,
		598, 3, 42, 21, 0, 598, 599, 5, 56, 0, 0, 599, 600, 6, 21, -1, 0, 600,
		642, 1, 0, 0, 0, 601, 602, 5, 2, 0, 0, 602, 603, 5, 55, 0, 0, 603, 604,
		3, 42, 21, 0, 604, 605, 5, 56, 0, 0, 605, 606, 6, 21, -1, 0, 606, 642,
		1, 0, 0, 0, 607, 608, 5, 3, 0, 0, 608, 609, 5, 55, 0, 0, 609, 610, 3, 42,
		21, 0, 610, 611, 5, 56, 0, 0, 611, 612, 6, 21, -1, 0, 612, 642, 1, 0, 0,
		0, 613, 614, 3, 48, 24, 0, 614, 615, 6, 21, -1, 0, 615, 642, 1, 0, 0, 0,
		616, 617, 5, 59, 0, 0, 617, 618, 3, 46, 23, 0, 618, 619, 5, 60, 0, 0, 619,
		620, 6, 21, -1, 0, 620, 642, 1, 0, 0, 0, 621, 622, 5, 39, 0, 0, 622, 623,
		5, 62, 0, 0, 623, 624, 5, 27, 0, 0, 624, 642, 6, 21, -1, 0, 625, 626, 5,
		39, 0, 0, 626, 627, 5, 62, 0, 0, 627, 628, 5, 26, 0, 0, 628, 642, 6, 21,
		-1, 0, 629, 630, 3, 50, 25, 0, 630, 631, 6, 21, -1, 0, 631, 642, 1, 0,
		0, 0, 632, 633, 5, 39, 0, 0, 633, 634, 5, 55, 0, 0, 634, 635, 3, 52, 26,
		0, 635, 636, 5, 56, 0, 0, 636, 637, 6, 21, -1, 0, 637, 642, 1, 0, 0, 0,
		638, 639, 3, 16, 8, 0, 639, 640, 6, 21, -1, 0, 640, 642, 1, 0, 0, 0, 641,
		576, 1, 0, 0, 0, 641, 581, 1, 0, 0, 0, 641, 585, 1, 0, 0, 0, 641, 590,
		1, 0, 0, 0, 641, 592, 1, 0, 0, 0, 641, 595, 1, 0, 0, 0, 641, 601, 1, 0,
		0, 0, 641, 607, 1, 0, 0, 0, 641, 613, 1, 0, 0, 0, 641, 616, 1, 0, 0, 0,
		641, 621, 1, 0, 0, 0, 641, 625, 1, 0, 0, 0, 641, 629, 1, 0, 0, 0, 641,
		632, 1, 0, 0, 0, 641, 638, 1, 0, 0, 0, 642, 680, 1, 0, 0, 0, 643, 644,
		10, 22, 0, 0, 644, 645, 7, 1, 0, 0, 645, 646, 3, 42, 21, 23, 646, 647,
		6, 21, -1, 0, 647, 679, 1, 0, 0, 0, 648, 649, 10, 21, 0, 0, 649, 650, 7,
		2, 0, 0, 650, 651, 3, 42, 21, 22, 651, 652, 6, 21, -1, 0, 652, 679, 1,
		0, 0, 0, 653, 654, 10, 20, 0, 0, 654, 655, 7, 3, 0, 0, 655, 656, 3, 42,
		21, 21, 656, 657, 6, 21, -1, 0, 657, 679, 1, 0, 0, 0, 658, 659, 10, 19,
		0, 0, 659, 660, 7, 4, 0, 0, 660, 661, 3, 42, 21, 20, 661, 662, 6, 21, -1,
		0, 662, 679, 1, 0, 0, 0, 663, 664, 10, 18, 0, 0, 664, 665, 7, 5, 0, 0,
		665, 666, 3, 42, 21, 19, 666, 667, 6, 21, -1, 0, 667, 679, 1, 0, 0, 0,
		668, 669, 10, 17, 0, 0, 669, 670, 5, 44, 0, 0, 670, 671, 3, 42, 21, 18,
		671, 672, 6, 21, -1, 0, 672, 679, 1, 0, 0, 0, 673, 674, 10, 16, 0, 0, 674,
		675, 5, 43, 0, 0, 675, 676, 3, 42, 21, 17, 676, 677, 6, 21, -1, 0, 677,
		679, 1, 0, 0, 0, 678, 643, 1, 0, 0, 0, 678, 648, 1, 0, 0, 0, 678, 653,
		1, 0, 0, 0, 678, 658, 1, 0, 0, 0, 678, 663, 1, 0, 0, 0, 678, 668, 1, 0,
		0, 0, 678, 673, 1, 0, 0, 0, 679, 682, 1, 0, 0, 0, 680, 678, 1, 0, 0, 0,
		680, 681, 1, 0, 0, 0, 681, 43, 1, 0, 0, 0, 682, 680, 1, 0, 0, 0, 683, 684,
		5, 36, 0, 0, 684, 696, 6, 22, -1, 0, 685, 686, 5, 37, 0, 0, 686, 696, 6,
		22, -1, 0, 687, 688, 5, 38, 0, 0, 688, 696, 6, 22, -1, 0, 689, 690, 5,
		8, 0, 0, 690, 696, 6, 22, -1, 0, 691, 692, 5, 9, 0, 0, 692, 696, 6, 22,
		-1, 0, 693, 694, 5, 35, 0, 0, 694, 696, 6, 22, -1, 0, 695, 683, 1, 0, 0,
		0, 695, 685, 1, 0, 0, 0, 695, 687, 1, 0, 0, 0, 695, 689, 1, 0, 0, 0, 695,
		691, 1, 0, 0, 0, 695, 693, 1, 0, 0, 0, 696, 45, 1, 0, 0, 0, 697, 698, 6,
		23, -1, 0, 698, 699, 3, 42, 21, 0, 699, 700, 6, 23, -1, 0, 700, 708, 1,
		0, 0, 0, 701, 702, 10, 2, 0, 0, 702, 703, 5, 63, 0, 0, 703, 704, 3, 42,
		21, 0, 704, 705, 6, 23, -1, 0, 705, 707, 1, 0, 0, 0, 706, 701, 1, 0, 0,
		0, 707, 710, 1, 0, 0, 0, 708, 706, 1, 0, 0, 0, 708, 709, 1, 0, 0, 0, 709,
		47, 1, 0, 0, 0, 710, 708, 1, 0, 0, 0, 711, 712, 6, 24, -1, 0, 712, 713,
		5, 39, 0, 0, 713, 714, 6, 24, -1, 0, 714, 727, 1, 0, 0, 0, 715, 716, 10,
		3, 0, 0, 716, 717, 5, 59, 0, 0, 717, 718, 3, 42, 21, 0, 718, 719, 5, 60,
		0, 0, 719, 720, 6, 24, -1, 0, 720, 726, 1, 0, 0, 0, 721, 722, 10, 1, 0,
		0, 722, 723, 5, 62, 0, 0, 723, 724, 5, 39, 0, 0, 724, 726, 6, 24, -1, 0,
		725, 715, 1, 0, 0, 0, 725, 721, 1, 0, 0, 0, 726, 729, 1, 0, 0, 0, 727,
		725, 1, 0, 0, 0, 727, 728, 1, 0, 0, 0, 728, 49, 1, 0, 0, 0, 729, 727, 1,
		0, 0, 0, 730, 731, 6, 25, -1, 0, 731, 732, 5, 59, 0, 0, 732, 733, 3, 42,
		21, 0, 733, 734, 5, 60, 0, 0, 734, 735, 6, 25, -1, 0, 735, 744, 1, 0, 0,
		0, 736, 737, 10, 2, 0, 0, 737, 738, 5, 59, 0, 0, 738, 739, 3, 42, 21, 0,
		739, 740, 5, 60, 0, 0, 740, 741, 6, 25, -1, 0, 741, 743, 1, 0, 0, 0, 742,
		736, 1, 0, 0, 0, 743, 746, 1, 0, 0, 0, 744, 742, 1, 0, 0, 0, 744, 745,
		1, 0, 0, 0, 745, 51, 1, 0, 0, 0, 746, 744, 1, 0, 0, 0, 747, 748, 6, 26,
		-1, 0, 748, 749, 5, 39, 0, 0, 749, 750, 5, 61, 0, 0, 750, 751, 3, 42, 21,
		0, 751, 752, 6, 26, -1, 0, 752, 755, 1, 0, 0, 0, 753, 755, 6, 26, -1, 0,
		754, 747, 1, 0, 0, 0, 754, 753, 1, 0, 0, 0, 755, 765, 1, 0, 0, 0, 756,
		757, 10, 3, 0, 0, 757, 758, 5, 63, 0, 0, 758, 759, 5, 39, 0, 0, 759, 760,
		5, 61, 0, 0, 760, 761, 3, 42, 21, 0, 761, 762, 6, 26, -1, 0, 762, 764,
		1, 0, 0, 0, 763, 756, 1, 0, 0, 0, 764, 767, 1, 0, 0, 0, 765, 763, 1, 0,
		0, 0, 765, 766, 1, 0, 0, 0, 766, 53, 1, 0, 0, 0, 767, 765, 1, 0, 0, 0,
		38, 61, 67, 73, 97, 102, 108, 121, 138, 149, 194, 198, 203, 209, 214, 219,
		226, 239, 251, 265, 376, 393, 410, 436, 484, 506, 526, 544, 574, 641, 678,
		680, 695, 708, 725, 727, 744, 754, 765,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserTSTRING      = 3
	SwiftGrammarParserBOOL         = 4
	SwiftGrammarParserCHAR         = 5
	SwiftGrammarParserVAR          = 6
	SwiftGrammarParserLET          = 7
	SwiftGrammarParserTRU          = 8
	SwiftGrammarParserFAL          = 9
	SwiftGrammarParserPRINT        = 10
	SwiftGrammarParserIF           = 11
	SwiftGrammarParserELSE         = 12
	SwiftGrammarParserSWITCH       = 13
	SwiftGrammarParserCASE         = 14
	SwiftGrammarParserDEFAULT      = 15
	SwiftGrammarParserBREAK        = 16
	SwiftGrammarParserCONTINUE     = 17
	SwiftGrammarParserRETURN       = 18
	SwiftGrammarParserWHILE        = 19
	SwiftGrammarParserFOR          = 20
	SwiftGrammarParserIN           = 21
	SwiftGrammarParserGUARD        = 22
	SwiftGrammarParserAPPEND       = 23
	SwiftGrammarParserREMOVELAST   = 24
	SwiftGrammarParserREMOVE       = 25
	SwiftGrammarParserISEMPTY      = 26
	SwiftGrammarParserCOUNT        = 27
	SwiftGrammarParserREPEATING    = 28
	SwiftGrammarParserSTRUCT       = 29
	SwiftGrammarParserSELF         = 30
	SwiftGrammarParserMUTATING     = 31
	SwiftGrammarParserFUNC         = 32
	SwiftGrammarParserAT           = 33
	SwiftGrammarParserINOUT        = 34
	SwiftGrammarParserNIL          = 35
	SwiftGrammarParserNUMBER       = 36
	SwiftGrammarParserCHARACTER    = 37
	SwiftGrammarParserSTRING       = 38
	SwiftGrammarParserID           = 39
	SwiftGrammarParserDIF          = 40
	SwiftGrammarParserIG_IG        = 41
	SwiftGrammarParserNOT          = 42
	SwiftGrammarParserOR           = 43
	SwiftGrammarParserAND          = 44
	SwiftGrammarParserIG           = 45
	SwiftGrammarParserMAY_IG       = 46
	SwiftGrammarParserMEN_IG       = 47
	SwiftGrammarParserMAYOR        = 48
	SwiftGrammarParserMENOR        = 49
	SwiftGrammarParserMOD          = 50
	SwiftGrammarParserMUL          = 51
	SwiftGrammarParserDIV          = 52
	SwiftGrammarParserADD          = 53
	SwiftGrammarParserSUB          = 54
	SwiftGrammarParserPARIZQ       = 55
	SwiftGrammarParserPARDER       = 56
	SwiftGrammarParserLLAVEIZQ     = 57
	SwiftGrammarParserLLAVEDER     = 58
	SwiftGrammarParserCORIZQ       = 59
	SwiftGrammarParserCORDER       = 60
	SwiftGrammarParserDOSP         = 61
	SwiftGrammarParserPUNTO        = 62
	SwiftGrammarParserCOMA         = 63
	SwiftGrammarParserINTCE        = 64
	SwiftGrammarParserFLECHA       = 65
	SwiftGrammarParserAMP          = 66
	SwiftGrammarParserPCOMA        = 67
	SwiftGrammarParserGBAJO        = 68
	SwiftGrammarParserTRESP        = 69
	SwiftGrammarParserWHITESPACE   = 70
	SwiftGrammarParserCOMMENT      = 71
	SwiftGrammarParserLINE_COMMENT = 72
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s              = 0
	SwiftGrammarParserRULE_block          = 1
	SwiftGrammarParserRULE_instruction    = 2
	SwiftGrammarParserRULE_structCreation = 3
	SwiftGrammarParserRULE_listStructDec  = 4
	SwiftGrammarParserRULE_funcstmt       = 5
	SwiftGrammarParserRULE_listParamsFunc = 6
	SwiftGrammarParserRULE_funcallstmt    = 7
	SwiftGrammarParserRULE_funcallestmt   = 8
	SwiftGrammarParserRULE_printstmt      = 9
	SwiftGrammarParserRULE_variablestmt   = 10
	SwiftGrammarParserRULE_typestmt       = 11
	SwiftGrammarParserRULE_type2stmt      = 12
	SwiftGrammarParserRULE_funvecstmt     = 13
	SwiftGrammarParserRULE_guardstmt      = 14
	SwiftGrammarParserRULE_ifstmt         = 15
	SwiftGrammarParserRULE_elifs          = 16
	SwiftGrammarParserRULE_switchstmt     = 17
	SwiftGrammarParserRULE_cases          = 18
	SwiftGrammarParserRULE_whilestmt      = 19
	SwiftGrammarParserRULE_forstmt        = 20
	SwiftGrammarParserRULE_expr           = 21
	SwiftGrammarParserRULE_primitives     = 22
	SwiftGrammarParserRULE_listParams     = 23
	SwiftGrammarParserRULE_listArray      = 24
	SwiftGrammarParserRULE_listAceso      = 25
	SwiftGrammarParserRULE_listStructExp  = 26
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(55)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(58)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_variablestmt returns the _variablestmt rule contexts.
	Get_variablestmt() IVariablestmtContext

	// Get_guardstmt returns the _guardstmt rule contexts.
	Get_guardstmt() IGuardstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_switchstmt returns the _switchstmt rule contexts.
	Get_switchstmt() ISwitchstmtContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_funvecstmt returns the _funvecstmt rule contexts.
	Get_funvecstmt() IFunvecstmtContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_funcallstmt returns the _funcallstmt rule contexts.
	Get_funcallstmt() IFuncallstmtContext

	// Get_funcstmt returns the _funcstmt rule contexts.
	Get_funcstmt() IFuncstmtContext

	// Get_structCreation returns the _structCreation rule contexts.
	Get_structCreation() IStructCreationContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_variablestmt sets the _variablestmt rule contexts.
	Set_variablestmt(IVariablestmtContext)

	// Set_guardstmt sets the _guardstmt rule contexts.
	Set_guardstmt(IGuardstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_switchstmt sets the _switchstmt rule contexts.
	Set_switchstmt(ISwitchstmtContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_funvecstmt sets the _funvecstmt rule contexts.
	Set_funvecstmt(IFunvecstmtContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_funcallstmt sets the _funcallstmt rule contexts.
	Set_funcallstmt(IFuncallstmtContext)

	// Set_funcstmt sets the _funcstmt rule contexts.
	Set_funcstmt(IFuncstmtContext)

	// Set_structCreation sets the _structCreation rule contexts.
	Set_structCreation(IStructCreationContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PCOMA() antlr.TerminalNode
	Variablestmt() IVariablestmtContext
	Guardstmt() IGuardstmtContext
	Ifstmt() IIfstmtContext
	Switchstmt() ISwitchstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Funvecstmt() IFunvecstmtContext
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	Funcallstmt() IFuncallstmtContext
	Funcstmt() IFuncstmtContext
	StructCreation() IStructCreationContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	inst            interfaces.Instruction
	_printstmt      IPrintstmtContext
	_variablestmt   IVariablestmtContext
	_guardstmt      IGuardstmtContext
	_ifstmt         IIfstmtContext
	_switchstmt     ISwitchstmtContext
	_whilestmt      IWhilestmtContext
	_forstmt        IForstmtContext
	_funvecstmt     IFunvecstmtContext
	_BREAK          antlr.Token
	_CONTINUE       antlr.Token
	_RETURN         antlr.Token
	_expr           IExprContext
	_funcallstmt    IFuncallstmtContext
	_funcstmt       IFuncstmtContext
	_structCreation IStructCreationContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *InstructionContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *InstructionContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *InstructionContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *InstructionContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *InstructionContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_variablestmt() IVariablestmtContext { return s._variablestmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_funvecstmt() IFunvecstmtContext { return s._funvecstmt }

func (s *InstructionContext) Get_expr() IExprContext { return s._expr }

func (s *InstructionContext) Get_funcallstmt() IFuncallstmtContext { return s._funcallstmt }

func (s *InstructionContext) Get_funcstmt() IFuncstmtContext { return s._funcstmt }

func (s *InstructionContext) Get_structCreation() IStructCreationContext { return s._structCreation }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_variablestmt(v IVariablestmtContext) { s._variablestmt = v }

func (s *InstructionContext) Set_guardstmt(v IGuardstmtContext) { s._guardstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_switchstmt(v ISwitchstmtContext) { s._switchstmt = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_funvecstmt(v IFunvecstmtContext) { s._funvecstmt = v }

func (s *InstructionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *InstructionContext) Set_funcallstmt(v IFuncallstmtContext) { s._funcallstmt = v }

func (s *InstructionContext) Set_funcstmt(v IFuncstmtContext) { s._funcstmt = v }

func (s *InstructionContext) Set_structCreation(v IStructCreationContext) { s._structCreation = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) PCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPCOMA, 0)
}

func (s *InstructionContext) Variablestmt() IVariablestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariablestmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Funvecstmt() IFunvecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunvecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunvecstmtContext)
}

func (s *InstructionContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *InstructionContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *InstructionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *InstructionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InstructionContext) Funcallstmt() IFuncallstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncallstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncallstmtContext)
}

func (s *InstructionContext) Funcstmt() IFuncstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncstmtContext)
}

func (s *InstructionContext) StructCreation() IStructCreationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCreationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCreationContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(66)
				p.Match(SwiftGrammarParserPCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)

			var _x = p.Variablestmt()

			localctx.(*InstructionContext)._variablestmt = _x
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(72)
				p.Match(SwiftGrammarParserPCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_variablestmt().GetVari()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGuinst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwinst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(86)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhileinst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetForinst()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)

			var _x = p.Funvecstmt()

			localctx.(*InstructionContext)._funvecstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_funvecstmt().GetFvecisnt()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(95)

			var _m = p.Match(SwiftGrammarParserBREAK)

			localctx.(*InstructionContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(96)
				p.Match(SwiftGrammarParserPCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		localctx.(*InstructionContext).inst = instructions.NewBreak((func() int {
			if localctx.(*InstructionContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_BREAK().GetColumn()
			}
		}()))

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(100)

			var _m = p.Match(SwiftGrammarParserCONTINUE)

			localctx.(*InstructionContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(101)
				p.Match(SwiftGrammarParserPCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		localctx.(*InstructionContext).inst = instructions.NewContinue((func() int {
			if localctx.(*InstructionContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_CONTINUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_CONTINUE().GetColumn()
			}
		}()))

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(105)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*InstructionContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(107)
				p.Match(SwiftGrammarParserPCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		localctx.(*InstructionContext).inst = instructions.NewReturn((func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetColumn()
			}
		}()), localctx.(*InstructionContext).Get_expr().GetE())

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(112)

			var _x = p.Funcallstmt()

			localctx.(*InstructionContext)._funcallstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_funcallstmt().GetCfun()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(115)

			var _x = p.Funcstmt()

			localctx.(*InstructionContext)._funcstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_funcstmt().GetFun()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(118)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructCreationContext is an interface to support dynamic dispatch.
type IStructCreationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRUCT returns the _STRUCT token.
	Get_STRUCT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_STRUCT sets the _STRUCT token.
	Set_STRUCT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listStructDec returns the _listStructDec rule contexts.
	Get_listStructDec() IListStructDecContext

	// Set_listStructDec sets the _listStructDec rule contexts.
	Set_listStructDec(IListStructDecContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	ListStructDec() IListStructDecContext
	LLAVEDER() antlr.TerminalNode

	// IsStructCreationContext differentiates from other interfaces.
	IsStructCreationContext()
}

type StructCreationContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	dec            interfaces.Instruction
	_STRUCT        antlr.Token
	_ID            antlr.Token
	_listStructDec IListStructDecContext
}

func NewEmptyStructCreationContext() *StructCreationContext {
	var p = new(StructCreationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
	return p
}

func InitEmptyStructCreationContext(p *StructCreationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
}

func (*StructCreationContext) IsStructCreationContext() {}

func NewStructCreationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCreationContext {
	var p = new(StructCreationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structCreation

	return p
}

func (s *StructCreationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCreationContext) Get_STRUCT() antlr.Token { return s._STRUCT }

func (s *StructCreationContext) Get_ID() antlr.Token { return s._ID }

func (s *StructCreationContext) Set_STRUCT(v antlr.Token) { s._STRUCT = v }

func (s *StructCreationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructCreationContext) Get_listStructDec() IListStructDecContext { return s._listStructDec }

func (s *StructCreationContext) Set_listStructDec(v IListStructDecContext) { s._listStructDec = v }

func (s *StructCreationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *StructCreationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *StructCreationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *StructCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructCreationContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *StructCreationContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *StructCreationContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (p *SwiftGrammarParser) StructCreation() (localctx IStructCreationContext) {
	localctx = NewStructCreationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_structCreation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _m = p.Match(SwiftGrammarParserSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructCreationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(127)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructCreationContext).dec = instructions.NewStruct((func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructCreationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructCreationContext).Get_ID().GetText()
		}
	}()), localctx.(*StructCreationContext).Get_listStructDec().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructDecContext is an interface to support dynamic dispatch.
type IListStructDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructDecContext

	// Get_typestmt returns the _typestmt rule contexts.
	Get_typestmt() ITypestmtContext

	// SetList sets the list rule contexts.
	SetList(IListStructDecContext)

	// Set_typestmt sets the _typestmt rule contexts.
	Set_typestmt(ITypestmtContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSP() antlr.TerminalNode
	Typestmt() ITypestmtContext
	ListStructDec() IListStructDecContext

	// IsListStructDecContext differentiates from other interfaces.
	IsListStructDecContext()
}

type ListStructDecContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         []interface{}
	list      IListStructDecContext
	_ID       antlr.Token
	_typestmt ITypestmtContext
}

func NewEmptyListStructDecContext() *ListStructDecContext {
	var p = new(ListStructDecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
	return p
}

func InitEmptyListStructDecContext(p *ListStructDecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
}

func (*ListStructDecContext) IsListStructDecContext() {}

func NewListStructDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructDecContext {
	var p = new(ListStructDecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec

	return p
}

func (s *ListStructDecContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructDecContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructDecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructDecContext) GetList() IListStructDecContext { return s.list }

func (s *ListStructDecContext) Get_typestmt() ITypestmtContext { return s._typestmt }

func (s *ListStructDecContext) SetList(v IListStructDecContext) { s.list = v }

func (s *ListStructDecContext) Set_typestmt(v ITypestmtContext) { s._typestmt = v }

func (s *ListStructDecContext) GetL() []interface{} { return s.l }

func (s *ListStructDecContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *ListStructDecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListStructDecContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *ListStructDecContext) Typestmt() ITypestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypestmtContext)
}

func (s *ListStructDecContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *ListStructDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructDec(s)
	}
}

func (s *ListStructDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructDec(s)
	}
}

func (p *SwiftGrammarParser) ListStructDec() (localctx IListStructDecContext) {
	return p.listStructDec(0)
}

func (p *SwiftGrammarParser) listStructDec(_p int) (localctx IListStructDecContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructDecContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructDecContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, SwiftGrammarParserRULE_listStructDec, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(131)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.Typestmt()

			localctx.(*ListStructDecContext)._typestmt = _x
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructDecContext).Get_typestmt().GetType_())
		arr = append(arr, newParams)
		localctx.(*ListStructDecContext).l = arr

	case 2:
		localctx.(*ListStructDecContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructDecContext(p, _parentctx, _parentState)
			localctx.(*ListStructDecContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
			p.SetState(140)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(141)
				p.Match(SwiftGrammarParserVAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(142)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructDecContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Match(SwiftGrammarParserDOSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)

				var _x = p.Typestmt()

				localctx.(*ListStructDecContext)._typestmt = _x
			}

			var arr []interface{}
			newParams := environment.NewStructType((func() string {
				if localctx.(*ListStructDecContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructDecContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructDecContext).Get_typestmt().GetType_())
			arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
			localctx.(*ListStructDecContext).l = arr

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncstmtContext is an interface to support dynamic dispatch.
type IFuncstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_listParamsFunc returns the _listParamsFunc rule contexts.
	Get_listParamsFunc() IListParamsFuncContext

	// Get_typestmt returns the _typestmt rule contexts.
	Get_typestmt() ITypestmtContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_listParamsFunc sets the _listParamsFunc rule contexts.
	Set_listParamsFunc(IListParamsFuncContext)

	// Set_typestmt sets the _typestmt rule contexts.
	Set_typestmt(ITypestmtContext)

	// GetFun returns the fun attribute.
	GetFun() interfaces.Instruction

	// SetFun sets the fun attribute.
	SetFun(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext
	FLECHA() antlr.TerminalNode
	Typestmt() ITypestmtContext

	// IsFuncstmtContext differentiates from other interfaces.
	IsFuncstmtContext()
}

type FuncstmtContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	fun             interfaces.Instruction
	_FUNC           antlr.Token
	_ID             antlr.Token
	_block          IBlockContext
	_listParamsFunc IListParamsFuncContext
	_typestmt       ITypestmtContext
}

func NewEmptyFuncstmtContext() *FuncstmtContext {
	var p = new(FuncstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt
	return p
}

func InitEmptyFuncstmtContext(p *FuncstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt
}

func (*FuncstmtContext) IsFuncstmtContext() {}

func NewFuncstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncstmtContext {
	var p = new(FuncstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt

	return p
}

func (s *FuncstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncstmtContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FuncstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncstmtContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FuncstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncstmtContext) Get_block() IBlockContext { return s._block }

func (s *FuncstmtContext) Get_listParamsFunc() IListParamsFuncContext { return s._listParamsFunc }

func (s *FuncstmtContext) Get_typestmt() ITypestmtContext { return s._typestmt }

func (s *FuncstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *FuncstmtContext) Set_listParamsFunc(v IListParamsFuncContext) { s._listParamsFunc = v }

func (s *FuncstmtContext) Set_typestmt(v ITypestmtContext) { s._typestmt = v }

func (s *FuncstmtContext) GetFun() interfaces.Instruction { return s.fun }

func (s *FuncstmtContext) SetFun(v interfaces.Instruction) { s.fun = v }

func (s *FuncstmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FuncstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FuncstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FuncstmtContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *FuncstmtContext) FLECHA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLECHA, 0)
}

func (s *FuncstmtContext) Typestmt() ITypestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypestmtContext)
}

func (s *FuncstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncstmt(s)
	}
}

func (s *FuncstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncstmt(s)
	}
}

func (p *SwiftGrammarParser) Funcstmt() (localctx IFuncstmtContext) {
	localctx = NewFuncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_funcstmt)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(158)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncstmtContext).fun = instructions.NewDeclarationFunc((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, nil, localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.listParamsFunc(0)

			localctx.(*FuncstmtContext)._listParamsFunc = _x
		}
		{
			p.SetState(165)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(168)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncstmtContext).fun = instructions.NewDeclarationFunc((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*FuncstmtContext).Get_listParamsFunc().GetLpf(), localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)

			var _x = p.Typestmt()

			localctx.(*FuncstmtContext)._typestmt = _x
		}
		{
			p.SetState(177)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(179)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncstmtContext).fun = instructions.NewDeclarationFunc((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncstmtContext).Get_typestmt().GetType_(), nil, localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _x = p.listParamsFunc(0)

			localctx.(*FuncstmtContext)._listParamsFunc = _x
		}
		{
			p.SetState(186)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Typestmt()

			localctx.(*FuncstmtContext)._typestmt = _x
		}
		{
			p.SetState(189)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncstmtContext).fun = instructions.NewDeclarationFunc((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncstmtContext).Get_typestmt().GetType_(), localctx.(*FuncstmtContext).Get_listParamsFunc().GetLpf(), localctx.(*FuncstmtContext).Get_block().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsFuncContext is an interface to support dynamic dispatch.
type IListParamsFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId1 returns the id1 token.
	GetId1() antlr.Token

	// GetId2 returns the id2 token.
	GetId2() antlr.Token

	// Get_INOUT returns the _INOUT token.
	Get_INOUT() antlr.Token

	// SetId1 sets the id1 token.
	SetId1(antlr.Token)

	// SetId2 sets the id2 token.
	SetId2(antlr.Token)

	// Set_INOUT sets the _INOUT token.
	Set_INOUT(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListParamsFuncContext

	// Get_typestmt returns the _typestmt rule contexts.
	Get_typestmt() ITypestmtContext

	// SetList sets the list rule contexts.
	SetList(IListParamsFuncContext)

	// Set_typestmt sets the _typestmt rule contexts.
	Set_typestmt(ITypestmtContext)

	// GetLpf returns the lpf attribute.
	GetLpf() []interface{}

	// SetLpf sets the lpf attribute.
	SetLpf([]interface{})

	// Getter signatures
	DOSP() antlr.TerminalNode
	Typestmt() ITypestmtContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	INOUT() antlr.TerminalNode
	GBAJO() antlr.TerminalNode
	COMA() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext

	// IsListParamsFuncContext differentiates from other interfaces.
	IsListParamsFuncContext()
}

type ListParamsFuncContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	lpf       []interface{}
	list      IListParamsFuncContext
	id1       antlr.Token
	id2       antlr.Token
	_INOUT    antlr.Token
	_typestmt ITypestmtContext
}

func NewEmptyListParamsFuncContext() *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
	return p
}

func InitEmptyListParamsFuncContext(p *ListParamsFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
}

func (*ListParamsFuncContext) IsListParamsFuncContext() {}

func NewListParamsFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc

	return p
}

func (s *ListParamsFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsFuncContext) GetId1() antlr.Token { return s.id1 }

func (s *ListParamsFuncContext) GetId2() antlr.Token { return s.id2 }

func (s *ListParamsFuncContext) Get_INOUT() antlr.Token { return s._INOUT }

func (s *ListParamsFuncContext) SetId1(v antlr.Token) { s.id1 = v }

func (s *ListParamsFuncContext) SetId2(v antlr.Token) { s.id2 = v }

func (s *ListParamsFuncContext) Set_INOUT(v antlr.Token) { s._INOUT = v }

func (s *ListParamsFuncContext) GetList() IListParamsFuncContext { return s.list }

func (s *ListParamsFuncContext) Get_typestmt() ITypestmtContext { return s._typestmt }

func (s *ListParamsFuncContext) SetList(v IListParamsFuncContext) { s.list = v }

func (s *ListParamsFuncContext) Set_typestmt(v ITypestmtContext) { s._typestmt = v }

func (s *ListParamsFuncContext) GetLpf() []interface{} { return s.lpf }

func (s *ListParamsFuncContext) SetLpf(v []interface{}) { s.lpf = v }

func (s *ListParamsFuncContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *ListParamsFuncContext) Typestmt() ITypestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypestmtContext)
}

func (s *ListParamsFuncContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ListParamsFuncContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ListParamsFuncContext) INOUT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINOUT, 0)
}

func (s *ListParamsFuncContext) GBAJO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGBAJO, 0)
}

func (s *ListParamsFuncContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsFuncContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *ListParamsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsFunc(s)
	}
}

func (s *ListParamsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsFunc(s)
	}
}

func (p *SwiftGrammarParser) ListParamsFunc() (localctx IListParamsFuncContext) {
	return p.listParamsFunc(0)
}

func (p *SwiftGrammarParser) listParamsFunc(_p int) (localctx IListParamsFuncContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsFuncContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsFuncContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, SwiftGrammarParserRULE_listParamsFunc, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.SetState(198)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(197)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ListParamsFuncContext).id1 = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SwiftGrammarParserID || _la == SwiftGrammarParserGBAJO) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ListParamsFuncContext).id1 = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(200)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext).id2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserINOUT {
			{
				p.SetState(202)

				var _m = p.Match(SwiftGrammarParserINOUT)

				localctx.(*ListParamsFuncContext)._INOUT = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(205)

			var _x = p.Typestmt()

			localctx.(*ListParamsFuncContext)._typestmt = _x
		}

		localctx.(*ListParamsFuncContext).lpf = []interface{}{}

		newParam := instructions.NewParamsDeclaration((func() int {
			if localctx.(*ListParamsFuncContext).GetId2() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).GetId2().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).GetId2() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).GetId2().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).GetId1() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).GetId1().GetText()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).GetId2() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).GetId2().GetText()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_INOUT() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_INOUT().GetText()
			}
		}()), localctx.(*ListParamsFuncContext).Get_typestmt().GetType_())
		localctx.(*ListParamsFuncContext).lpf = append(localctx.(*ListParamsFuncContext).lpf, newParam)

	case 2:
		localctx.(*ListParamsFuncContext).lpf = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
			localctx.(*ListParamsFuncContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
			p.SetState(211)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(212)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(213)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ListParamsFuncContext).id1 = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserID || _la == SwiftGrammarParserGBAJO) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ListParamsFuncContext).id1 = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}
			{
				p.SetState(216)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListParamsFuncContext).id2 = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(217)
				p.Match(SwiftGrammarParserDOSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SwiftGrammarParserINOUT {
				{
					p.SetState(218)

					var _m = p.Match(SwiftGrammarParserINOUT)

					localctx.(*ListParamsFuncContext)._INOUT = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(221)

				var _x = p.Typestmt()

				localctx.(*ListParamsFuncContext)._typestmt = _x
			}

			var arr []interface{}
			newParam := instructions.NewParamsDeclaration((func() int {
				if localctx.(*ListParamsFuncContext).GetId2() == nil {
					return 0
				} else {
					return localctx.(*ListParamsFuncContext).GetId2().GetLine()
				}
			}()), (func() int {
				if localctx.(*ListParamsFuncContext).GetId2() == nil {
					return 0
				} else {
					return localctx.(*ListParamsFuncContext).GetId2().GetColumn()
				}
			}()), (func() string {
				if localctx.(*ListParamsFuncContext).GetId1() == nil {
					return ""
				} else {
					return localctx.(*ListParamsFuncContext).GetId1().GetText()
				}
			}()), (func() string {
				if localctx.(*ListParamsFuncContext).GetId2() == nil {
					return ""
				} else {
					return localctx.(*ListParamsFuncContext).GetId2().GetText()
				}
			}()), (func() string {
				if localctx.(*ListParamsFuncContext).Get_INOUT() == nil {
					return ""
				} else {
					return localctx.(*ListParamsFuncContext).Get_INOUT().GetText()
				}
			}()), localctx.(*ListParamsFuncContext).Get_typestmt().GetType_())
			arr = append(localctx.(*ListParamsFuncContext).GetList().GetLpf(), newParam)
			localctx.(*ListParamsFuncContext).lpf = arr

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncallstmtContext is an interface to support dynamic dispatch.
type IFuncallstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetCfun returns the cfun attribute.
	GetCfun() interfaces.Instruction

	// SetCfun sets the cfun attribute.
	SetCfun(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsFuncallstmtContext differentiates from other interfaces.
	IsFuncallstmtContext()
}

type FuncallstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	cfun        interfaces.Instruction
	_ID         antlr.Token
	_listParams IListParamsContext
}

func NewEmptyFuncallstmtContext() *FuncallstmtContext {
	var p = new(FuncallstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcallstmt
	return p
}

func InitEmptyFuncallstmtContext(p *FuncallstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcallstmt
}

func (*FuncallstmtContext) IsFuncallstmtContext() {}

func NewFuncallstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncallstmtContext {
	var p = new(FuncallstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcallstmt

	return p
}

func (s *FuncallstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncallstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncallstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncallstmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *FuncallstmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *FuncallstmtContext) GetCfun() interfaces.Instruction { return s.cfun }

func (s *FuncallstmtContext) SetCfun(v interfaces.Instruction) { s.cfun = v }

func (s *FuncallstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncallstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncallstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncallstmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *FuncallstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncallstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncallstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncallstmt(s)
	}
}

func (s *FuncallstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncallstmt(s)
	}
}

func (p *SwiftGrammarParser) Funcallstmt() (localctx IFuncallstmtContext) {
	localctx = NewFuncallstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_funcallstmt)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncallstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncallstmtContext).cfun = instructions.NewFunCall((func() int {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetText()
			}
		}()), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncallstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _x = p.listParams(0)

			localctx.(*FuncallstmtContext)._listParams = _x
		}
		{
			p.SetState(236)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncallstmtContext).cfun = instructions.NewFunCall((func() int {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncallstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncallstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncallstmtContext).Get_listParams().GetL())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncallestmtContext is an interface to support dynamic dispatch.
type IFuncallestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetCefun returns the cefun attribute.
	GetCefun() interfaces.Expression

	// SetCefun sets the cefun attribute.
	SetCefun(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsFuncallestmtContext differentiates from other interfaces.
	IsFuncallestmtContext()
}

type FuncallestmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	cefun       interfaces.Expression
	_ID         antlr.Token
	_listParams IListParamsContext
}

func NewEmptyFuncallestmtContext() *FuncallestmtContext {
	var p = new(FuncallestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcallestmt
	return p
}

func InitEmptyFuncallestmtContext(p *FuncallestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcallestmt
}

func (*FuncallestmtContext) IsFuncallestmtContext() {}

func NewFuncallestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncallestmtContext {
	var p = new(FuncallestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcallestmt

	return p
}

func (s *FuncallestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncallestmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncallestmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncallestmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *FuncallestmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *FuncallestmtContext) GetCefun() interfaces.Expression { return s.cefun }

func (s *FuncallestmtContext) SetCefun(v interfaces.Expression) { s.cefun = v }

func (s *FuncallestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncallestmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncallestmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncallestmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *FuncallestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncallestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncallestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncallestmt(s)
	}
}

func (s *FuncallestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncallestmt(s)
	}
}

func (p *SwiftGrammarParser) Funcallestmt() (localctx IFuncallestmtContext) {
	localctx = NewFuncallestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_funcallestmt)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncallestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncallestmtContext).cefun = expressions.NewFunCallE((func() int {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetText()
			}
		}()), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncallestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)

			var _x = p.listParams(0)

			localctx.(*FuncallestmtContext)._listParams = _x
		}
		{
			p.SetState(248)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncallestmtContext).cefun = expressions.NewFunCallE((func() int {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncallestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncallestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncallestmtContext).Get_listParams().GetL())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Get_PARIZQ returns the _PARIZQ token.
	Get_PARIZQ() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Set_PARIZQ sets the _PARIZQ token.
	Set_PARIZQ(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	prnt        interfaces.Instruction
	_PRINT      antlr.Token
	_expr       IExprContext
	_PARIZQ     antlr.Token
	_listParams IListParamsContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Get_PARIZQ() antlr.Token { return s._PARIZQ }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Set_PARIZQ(v antlr.Token) { s._PARIZQ = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_printstmt)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)

			var _x = p.expr(0)

			localctx.(*PrintstmtContext)._expr = _x
		}
		{
			p.SetState(256)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _m = p.Match(SwiftGrammarParserPARIZQ)

			localctx.(*PrintstmtContext)._PARIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)

			var _x = p.listParams(0)

			localctx.(*PrintstmtContext)._listParams = _x
		}
		{
			p.SetState(262)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), expressions.NewArray((func() int {
			if localctx.(*PrintstmtContext).Get_PARIZQ() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PARIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PARIZQ() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PARIZQ().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).Get_listParams().GetL()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariablestmtContext is an interface to support dynamic dispatch.
type IVariablestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Get_typestmt returns the _typestmt rule contexts.
	Get_typestmt() ITypestmtContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_primitives returns the _primitives rule contexts.
	Get_primitives() IPrimitivesContext

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// Set_typestmt sets the _typestmt rule contexts.
	Set_typestmt(ITypestmtContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_primitives sets the _primitives rule contexts.
	Set_primitives(IPrimitivesContext)

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// GetVari returns the vari attribute.
	GetVari() interfaces.Instruction

	// SetVari sets the vari attribute.
	SetVari(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSP() antlr.TerminalNode
	Typestmt() ITypestmtContext
	IG() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	INTCE() antlr.TerminalNode
	LET() antlr.TerminalNode
	Primitives() IPrimitivesContext
	CORIZQ() antlr.TerminalNode
	CORDER() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsVariablestmtContext differentiates from other interfaces.
	IsVariablestmtContext()
}

type VariablestmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	vari        interfaces.Instruction
	_VAR        antlr.Token
	_ID         antlr.Token
	_typestmt   ITypestmtContext
	_expr       IExprContext
	_LET        antlr.Token
	_primitives IPrimitivesContext
	e1          IExprContext
	e2          IExprContext
	_CORIZQ     antlr.Token
}

func NewEmptyVariablestmtContext() *VariablestmtContext {
	var p = new(VariablestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_variablestmt
	return p
}

func InitEmptyVariablestmtContext(p *VariablestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_variablestmt
}

func (*VariablestmtContext) IsVariablestmtContext() {}

func NewVariablestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablestmtContext {
	var p = new(VariablestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_variablestmt

	return p
}

func (s *VariablestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablestmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *VariablestmtContext) Get_ID() antlr.Token { return s._ID }

func (s *VariablestmtContext) Get_LET() antlr.Token { return s._LET }

func (s *VariablestmtContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *VariablestmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *VariablestmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VariablestmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *VariablestmtContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *VariablestmtContext) Get_typestmt() ITypestmtContext { return s._typestmt }

func (s *VariablestmtContext) Get_expr() IExprContext { return s._expr }

func (s *VariablestmtContext) Get_primitives() IPrimitivesContext { return s._primitives }

func (s *VariablestmtContext) GetE1() IExprContext { return s.e1 }

func (s *VariablestmtContext) GetE2() IExprContext { return s.e2 }

func (s *VariablestmtContext) Set_typestmt(v ITypestmtContext) { s._typestmt = v }

func (s *VariablestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VariablestmtContext) Set_primitives(v IPrimitivesContext) { s._primitives = v }

func (s *VariablestmtContext) SetE1(v IExprContext) { s.e1 = v }

func (s *VariablestmtContext) SetE2(v IExprContext) { s.e2 = v }

func (s *VariablestmtContext) GetVari() interfaces.Instruction { return s.vari }

func (s *VariablestmtContext) SetVari(v interfaces.Instruction) { s.vari = v }

func (s *VariablestmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *VariablestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VariablestmtContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *VariablestmtContext) Typestmt() ITypestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypestmtContext)
}

func (s *VariablestmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VariablestmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VariablestmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariablestmtContext) INTCE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINTCE, 0)
}

func (s *VariablestmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *VariablestmtContext) Primitives() IPrimitivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivesContext)
}

func (s *VariablestmtContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *VariablestmtContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *VariablestmtContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *VariablestmtContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *VariablestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVariablestmt(s)
	}
}

func (s *VariablestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVariablestmt(s)
	}
}

func (p *SwiftGrammarParser) Variablestmt() (localctx IVariablestmtContext) {
	localctx = NewVariablestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_variablestmt)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(271)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(), localctx.(*VariablestmtContext).Get_expr().GetE(), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*VariablestmtContext).Get_expr().GetE(), true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(285)
			p.Match(SwiftGrammarParserINTCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(), expressions.NewPrimitive((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), nil, environment.NULL), true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(288)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*VariablestmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)

			var _x = p.Primitives()

			localctx.(*VariablestmtContext)._primitives = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*VariablestmtContext).Get_primitives().GetP(), false)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(294)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*VariablestmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(298)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

			var _x = p.Primitives()

			localctx.(*VariablestmtContext)._primitives = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(), localctx.(*VariablestmtContext).Get_primitives().GetP(), false)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(302)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*VariablestmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*VariablestmtContext).Get_expr().GetE(), false)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(308)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*VariablestmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(312)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(), localctx.(*VariablestmtContext).Get_expr().GetE(), false)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(316)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*VariablestmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*VariablestmtContext).Get_expr().GetE(), false)

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(322)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewAssigment((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_expr().GetE())

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(327)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext).e1 = _x
		}
		{
			p.SetState(330)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext).e2 = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewAssigmentVec((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).GetE1().GetE(), localctx.(*VariablestmtContext).GetE2().GetE())

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(335)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext).e1 = _x
		}
		{
			p.SetState(337)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext).e2 = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewAssigmentMa((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).GetE1().GetE(), localctx.(*VariablestmtContext).GetE2().GetE())

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(341)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewSumAssigment((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_expr().GetE())

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(347)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewResAssigment((func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_expr().GetE())

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(353)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(357)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(), localctx.(*VariablestmtContext).Get_expr().GetE(), true)

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(361)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _x = p.Typestmt()

			localctx.(*VariablestmtContext)._typestmt = _x
		}
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*VariablestmtContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VariablestmtContext).Get_typestmt().GetType_(),
			expressions.NewArray((func() int {
				if localctx.(*VariablestmtContext).Get_CORIZQ() == nil {
					return 0
				} else {
					return localctx.(*VariablestmtContext).Get_CORIZQ().GetLine()
				}
			}()), (func() int {
				if localctx.(*VariablestmtContext).Get_CORIZQ() == nil {
					return 0
				} else {
					return localctx.(*VariablestmtContext).Get_CORIZQ().GetColumn()
				}
			}()), nil),
			true)

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(370)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VariablestmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VariablestmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _x = p.expr(0)

			localctx.(*VariablestmtContext)._expr = _x
		}
		localctx.(*VariablestmtContext).vari = instructions.NewDeclaration((func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VariablestmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VariablestmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VariablestmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VariablestmtContext).Get_ID().GetText()
			}
		}()), environment.STRUCT, localctx.(*VariablestmtContext).Get_expr().GetE(), true)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypestmtContext is an interface to support dynamic dispatch.
type ITypestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type2stmt returns the _type2stmt rule contexts.
	Get_type2stmt() IType2stmtContext

	// Set_type2stmt sets the _type2stmt rule contexts.
	Set_type2stmt(IType2stmtContext)

	// GetType_ returns the type_ attribute.
	GetType_() environment.TipoExpresion

	// SetType_ sets the type_ attribute.
	SetType_(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	TSTRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Type2stmt() IType2stmtContext
	CORDER() antlr.TerminalNode

	// IsTypestmtContext differentiates from other interfaces.
	IsTypestmtContext()
}

type TypestmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	type_      environment.TipoExpresion
	_type2stmt IType2stmtContext
}

func NewEmptyTypestmtContext() *TypestmtContext {
	var p = new(TypestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typestmt
	return p
}

func InitEmptyTypestmtContext(p *TypestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typestmt
}

func (*TypestmtContext) IsTypestmtContext() {}

func NewTypestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypestmtContext {
	var p = new(TypestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_typestmt

	return p
}

func (s *TypestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TypestmtContext) Get_type2stmt() IType2stmtContext { return s._type2stmt }

func (s *TypestmtContext) Set_type2stmt(v IType2stmtContext) { s._type2stmt = v }

func (s *TypestmtContext) GetType_() environment.TipoExpresion { return s.type_ }

func (s *TypestmtContext) SetType_(v environment.TipoExpresion) { s.type_ = v }

func (s *TypestmtContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TypestmtContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TypestmtContext) TSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTSTRING, 0)
}

func (s *TypestmtContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TypestmtContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHAR, 0)
}

func (s *TypestmtContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *TypestmtContext) Type2stmt() IType2stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType2stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType2stmtContext)
}

func (s *TypestmtContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *TypestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypestmt(s)
	}
}

func (s *TypestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypestmt(s)
	}
}

func (p *SwiftGrammarParser) Typestmt() (localctx ITypestmtContext) {
	localctx = NewTypestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_typestmt)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = environment.FLOAT

	case SwiftGrammarParserTSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)
			p.Match(SwiftGrammarParserTSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(384)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = environment.BOOLEAN

	case SwiftGrammarParserCHAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(386)
			p.Match(SwiftGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = environment.CHAR

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(388)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)

			var _x = p.Type2stmt()

			localctx.(*TypestmtContext)._type2stmt = _x
		}
		{
			p.SetState(390)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypestmtContext).type_ = localctx.(*TypestmtContext).Get_type2stmt().GetType2()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType2stmtContext is an interface to support dynamic dispatch.
type IType2stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_type2stmt returns the _type2stmt rule contexts.
	Get_type2stmt() IType2stmtContext

	// Set_type2stmt sets the _type2stmt rule contexts.
	Set_type2stmt(IType2stmtContext)

	// GetType2 returns the type2 attribute.
	GetType2() environment.TipoExpresion

	// SetType2 sets the type2 attribute.
	SetType2(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	TSTRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Type2stmt() IType2stmtContext
	CORDER() antlr.TerminalNode

	// IsType2stmtContext differentiates from other interfaces.
	IsType2stmtContext()
}

type Type2stmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	type2      environment.TipoExpresion
	_type2stmt IType2stmtContext
}

func NewEmptyType2stmtContext() *Type2stmtContext {
	var p = new(Type2stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_type2stmt
	return p
}

func InitEmptyType2stmtContext(p *Type2stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_type2stmt
}

func (*Type2stmtContext) IsType2stmtContext() {}

func NewType2stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type2stmtContext {
	var p = new(Type2stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_type2stmt

	return p
}

func (s *Type2stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Type2stmtContext) Get_type2stmt() IType2stmtContext { return s._type2stmt }

func (s *Type2stmtContext) Set_type2stmt(v IType2stmtContext) { s._type2stmt = v }

func (s *Type2stmtContext) GetType2() environment.TipoExpresion { return s.type2 }

func (s *Type2stmtContext) SetType2(v environment.TipoExpresion) { s.type2 = v }

func (s *Type2stmtContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *Type2stmtContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *Type2stmtContext) TSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTSTRING, 0)
}

func (s *Type2stmtContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *Type2stmtContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHAR, 0)
}

func (s *Type2stmtContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *Type2stmtContext) Type2stmt() IType2stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType2stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType2stmtContext)
}

func (s *Type2stmtContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *Type2stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type2stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type2stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterType2stmt(s)
	}
}

func (s *Type2stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitType2stmt(s)
	}
}

func (p *SwiftGrammarParser) Type2stmt() (localctx IType2stmtContext) {
	localctx = NewType2stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_type2stmt)
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = environment.A_INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = environment.A_FLOAT

	case SwiftGrammarParserTSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(399)
			p.Match(SwiftGrammarParserTSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = environment.A_STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(401)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = environment.A_BOOLEAN

	case SwiftGrammarParserCHAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(403)
			p.Match(SwiftGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = environment.A_CHAR

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(405)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)

			var _x = p.Type2stmt()

			localctx.(*Type2stmtContext)._type2stmt = _x
		}
		{
			p.SetState(407)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Type2stmtContext).type2 = localctx.(*Type2stmtContext).Get_type2stmt().GetType2()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunvecstmtContext is an interface to support dynamic dispatch.
type IFunvecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetFvecisnt returns the fvecisnt attribute.
	GetFvecisnt() interfaces.Instruction

	// SetFvecisnt sets the fvecisnt attribute.
	SetFvecisnt(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	AT() antlr.TerminalNode
	DOSP() antlr.TerminalNode

	// IsFunvecstmtContext differentiates from other interfaces.
	IsFunvecstmtContext()
}

type FunvecstmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	fvecisnt interfaces.Instruction
	_ID      antlr.Token
	_expr    IExprContext
}

func NewEmptyFunvecstmtContext() *FunvecstmtContext {
	var p = new(FunvecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funvecstmt
	return p
}

func InitEmptyFunvecstmtContext(p *FunvecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funvecstmt
}

func (*FunvecstmtContext) IsFunvecstmtContext() {}

func NewFunvecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunvecstmtContext {
	var p = new(FunvecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funvecstmt

	return p
}

func (s *FunvecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FunvecstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FunvecstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunvecstmtContext) Get_expr() IExprContext { return s._expr }

func (s *FunvecstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *FunvecstmtContext) GetFvecisnt() interfaces.Instruction { return s.fvecisnt }

func (s *FunvecstmtContext) SetFvecisnt(v interfaces.Instruction) { s.fvecisnt = v }

func (s *FunvecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FunvecstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *FunvecstmtContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *FunvecstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FunvecstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunvecstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FunvecstmtContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVELAST, 0)
}

func (s *FunvecstmtContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *FunvecstmtContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT, 0)
}

func (s *FunvecstmtContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *FunvecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunvecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunvecstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFunvecstmt(s)
	}
}

func (s *FunvecstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFunvecstmt(s)
	}
}

func (p *SwiftGrammarParser) Funvecstmt() (localctx IFunvecstmtContext) {
	localctx = NewFunvecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_funvecstmt)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunvecstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(SwiftGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)

			var _x = p.expr(0)

			localctx.(*FunvecstmtContext)._expr = _x
		}
		{
			p.SetState(417)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunvecstmtContext).fvecisnt = instructions.NewAppend((func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FunvecstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunvecstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(SwiftGrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunvecstmtContext).fvecisnt = instructions.NewRemoveLast((func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(426)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunvecstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(SwiftGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(SwiftGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)

			var _x = p.expr(0)

			localctx.(*FunvecstmtContext)._expr = _x
		}
		{
			p.SetState(433)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunvecstmtContext).fvecisnt = instructions.NewRemove((func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunvecstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunvecstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FunvecstmtContext).Get_expr().GetE())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGuinst returns the guinst attribute.
	GetGuinst() interfaces.Instruction

	// SetGuinst sets the guinst attribute.
	SetGuinst(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	guinst interfaces.Instruction
	_GUARD antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardstmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardstmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardstmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardstmtContext) GetGuinst() interfaces.Instruction { return s.guinst }

func (s *GuardstmtContext) SetGuinst(v interfaces.Instruction) { s.guinst = v }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(440)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(443)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardstmtContext).guinst = instructions.NewGuard((func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardstmtContext).Get_expr().GetE(), localctx.(*GuardstmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetBif returns the bif rule contexts.
	GetBif() IBlockContext

	// GetBelse returns the belse rule contexts.
	GetBelse() IBlockContext

	// Get_elifs returns the _elifs rule contexts.
	Get_elifs() IElifsContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetBif sets the bif rule contexts.
	SetBif(IBlockContext)

	// SetBelse sets the belse rule contexts.
	SetBelse(IBlockContext)

	// Set_elifs sets the _elifs rule contexts.
	Set_elifs(IElifsContext)

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Elifs() IElifsContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ifinst interfaces.Instruction
	_IF    antlr.Token
	_expr  IExprContext
	_block IBlockContext
	bif    IBlockContext
	belse  IBlockContext
	_elifs IElifsContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) Get_block() IBlockContext { return s._block }

func (s *IfstmtContext) GetBif() IBlockContext { return s.bif }

func (s *IfstmtContext) GetBelse() IBlockContext { return s.belse }

func (s *IfstmtContext) Get_elifs() IElifsContext { return s._elifs }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) SetBif(v IBlockContext) { s.bif = v }

func (s *IfstmtContext) SetBelse(v IBlockContext) { s.belse = v }

func (s *IfstmtContext) Set_elifs(v IElifsContext) { s._elifs = v }

func (s *IfstmtContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *IfstmtContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) Elifs() IElifsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifsContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_ifstmt)
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(448)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(450)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(455)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)

			var _x = p.Block()

			localctx.(*IfstmtContext).bif = _x
		}
		{
			p.SetState(457)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)

			var _x = p.Block()

			localctx.(*IfstmtContext).belse = _x
		}
		{
			p.SetState(461)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetBif().GetBlk(), localctx.(*IfstmtContext).GetBelse().GetBlk(), nil)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(464)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(466)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)

			var _x = p.Block()

			localctx.(*IfstmtContext).bif = _x
		}
		{
			p.SetState(468)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)

			var _x = p.elifs(0)

			localctx.(*IfstmtContext)._elifs = _x
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetBif().GetBlk(), nil, localctx.(*IfstmtContext).Get_elifs().GetElifinst())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(472)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(474)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)

			var _x = p.Block()

			localctx.(*IfstmtContext).bif = _x
		}
		{
			p.SetState(476)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)

			var _x = p.elifs(0)

			localctx.(*IfstmtContext)._elifs = _x
		}
		{
			p.SetState(478)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)

			var _x = p.Block()

			localctx.(*IfstmtContext).belse = _x
		}
		{
			p.SetState(481)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetBif().GetBlk(), localctx.(*IfstmtContext).GetBelse().GetBlk(), localctx.(*IfstmtContext).Get_elifs().GetElifinst())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElifsContext is an interface to support dynamic dispatch.
type IElifsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// GetCelif returns the celif rule contexts.
	GetCelif() IElifsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// SetCelif sets the celif rule contexts.
	SetCelif(IElifsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetElifinst returns the elifinst attribute.
	GetElifinst() []interface{}

	// SetElifinst sets the elifinst attribute.
	SetElifinst([]interface{})

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	Elifs() IElifsContext

	// IsElifsContext differentiates from other interfaces.
	IsElifsContext()
}

type ElifsContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	elifinst []interface{}
	celif    IElifsContext
	_ELSE    antlr.Token
	_expr    IExprContext
	_block   IBlockContext
}

func NewEmptyElifsContext() *ElifsContext {
	var p = new(ElifsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elifs
	return p
}

func InitEmptyElifsContext(p *ElifsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elifs
}

func (*ElifsContext) IsElifsContext() {}

func NewElifsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifsContext {
	var p = new(ElifsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elifs

	return p
}

func (s *ElifsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifsContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *ElifsContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *ElifsContext) GetCelif() IElifsContext { return s.celif }

func (s *ElifsContext) Get_expr() IExprContext { return s._expr }

func (s *ElifsContext) Get_block() IBlockContext { return s._block }

func (s *ElifsContext) SetCelif(v IElifsContext) { s.celif = v }

func (s *ElifsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ElifsContext) Set_block(v IBlockContext) { s._block = v }

func (s *ElifsContext) GetElifinst() []interface{} { return s.elifinst }

func (s *ElifsContext) SetElifinst(v []interface{}) { s.elifinst = v }

func (s *ElifsContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElifsContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *ElifsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElifsContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElifsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElifsContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElifsContext) Elifs() IElifsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifsContext)
}

func (s *ElifsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElifs(s)
	}
}

func (s *ElifsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElifs(s)
	}
}

func (p *SwiftGrammarParser) Elifs() (localctx IElifsContext) {
	return p.elifs(0)
}

func (p *SwiftGrammarParser) elifs(_p int) (localctx IElifsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewElifsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElifsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, SwiftGrammarParserRULE_elifs, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)

		var _m = p.Match(SwiftGrammarParserELSE)

		localctx.(*ElifsContext)._ELSE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)

		var _x = p.expr(0)

		localctx.(*ElifsContext)._expr = _x
	}
	{
		p.SetState(490)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)

		var _x = p.Block()

		localctx.(*ElifsContext)._block = _x
	}
	{
		p.SetState(492)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ElifsContext).elifinst = []interface{}{}
	localctx.(*ElifsContext).elifinst = append(localctx.(*ElifsContext).elifinst, instructions.NewElif((func() int {
		if localctx.(*ElifsContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*ElifsContext).Get_ELSE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ElifsContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*ElifsContext).Get_ELSE().GetColumn()
		}
	}()), localctx.(*ElifsContext).Get_expr().GetE(), localctx.(*ElifsContext).Get_block().GetBlk()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElifsContext(p, _parentctx, _parentState)
			localctx.(*ElifsContext).celif = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_elifs)
			p.SetState(495)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(496)

				var _m = p.Match(SwiftGrammarParserELSE)

				localctx.(*ElifsContext)._ELSE = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(497)
				p.Match(SwiftGrammarParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(498)

				var _x = p.expr(0)

				localctx.(*ElifsContext)._expr = _x
			}
			{
				p.SetState(499)
				p.Match(SwiftGrammarParserLLAVEIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(500)

				var _x = p.Block()

				localctx.(*ElifsContext)._block = _x
			}
			{
				p.SetState(501)
				p.Match(SwiftGrammarParserLLAVEDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			var arrif []interface{}
			arrif = append(localctx.(*ElifsContext).GetCelif().GetElifinst(), instructions.NewElif((func() int {
				if localctx.(*ElifsContext).Get_ELSE() == nil {
					return 0
				} else {
					return localctx.(*ElifsContext).Get_ELSE().GetLine()
				}
			}()), (func() int {
				if localctx.(*ElifsContext).Get_ELSE() == nil {
					return 0
				} else {
					return localctx.(*ElifsContext).Get_ELSE().GetColumn()
				}
			}()), localctx.(*ElifsContext).Get_expr().GetE(), localctx.(*ElifsContext).Get_block().GetBlk()))
			localctx.(*ElifsContext).elifinst = arrif

		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_cases returns the _cases rule contexts.
	Get_cases() ICasesContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_cases sets the _cases rule contexts.
	Set_cases(ICasesContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetSwinst returns the swinst attribute.
	GetSwinst() interfaces.Instruction

	// SetSwinst sets the swinst attribute.
	SetSwinst(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Cases() ICasesContext
	DEFAULT() antlr.TerminalNode
	DOSP() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	swinst  interfaces.Instruction
	_SWITCH antlr.Token
	_expr   IExprContext
	_cases  ICasesContext
	_block  IBlockContext
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchstmtContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchstmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchstmtContext) Get_cases() ICasesContext { return s._cases }

func (s *SwitchstmtContext) Get_block() IBlockContext { return s._block }

func (s *SwitchstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchstmtContext) Set_cases(v ICasesContext) { s._cases = v }

func (s *SwitchstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *SwitchstmtContext) GetSwinst() interfaces.Instruction { return s.swinst }

func (s *SwitchstmtContext) SetSwinst(v interfaces.Instruction) { s.swinst = v }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *SwitchstmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *SwitchstmtContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *SwitchstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *SwiftGrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_switchstmt)
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)

			var _m = p.Match(SwiftGrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(511)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)

			var _x = p.cases(0)

			localctx.(*SwitchstmtContext)._cases = _x
		}
		{
			p.SetState(513)
			p.Match(SwiftGrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)

			var _x = p.Block()

			localctx.(*SwitchstmtContext)._block = _x
		}
		{
			p.SetState(516)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).swinst = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), localctx.(*SwitchstmtContext).Get_cases().GetCasesinst(), localctx.(*SwitchstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(519)

			var _m = p.Match(SwiftGrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(521)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)

			var _x = p.cases(0)

			localctx.(*SwitchstmtContext)._cases = _x
		}
		{
			p.SetState(523)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).swinst = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), localctx.(*SwitchstmtContext).Get_cases().GetCasesinst(), nil)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// GetCcases returns the ccases rule contexts.
	GetCcases() ICasesContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// SetCcases sets the ccases rule contexts.
	SetCcases(ICasesContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCasesinst returns the casesinst attribute.
	GetCasesinst() []interface{}

	// SetCasesinst sets the casesinst attribute.
	SetCasesinst([]interface{})

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	DOSP() antlr.TerminalNode
	Block() IBlockContext
	Cases() ICasesContext

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	casesinst []interface{}
	ccases    ICasesContext
	_CASE     antlr.Token
	_expr     IExprContext
	_block    IBlockContext
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasesContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasesContext) GetCcases() ICasesContext { return s.ccases }

func (s *CasesContext) Get_expr() IExprContext { return s._expr }

func (s *CasesContext) Get_block() IBlockContext { return s._block }

func (s *CasesContext) SetCcases(v ICasesContext) { s.ccases = v }

func (s *CasesContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CasesContext) Set_block(v IBlockContext) { s._block = v }

func (s *CasesContext) GetCasesinst() []interface{} { return s.casesinst }

func (s *CasesContext) SetCasesinst(v []interface{}) { s.casesinst = v }

func (s *CasesContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CasesContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CasesContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *CasesContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CasesContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCases(s)
	}
}

func (p *SwiftGrammarParser) Cases() (localctx ICasesContext) {
	return p.cases(0)
}

func (p *SwiftGrammarParser) cases(_p int) (localctx ICasesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCasesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICasesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SwiftGrammarParserRULE_cases, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CasesContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(530)

		var _x = p.expr(0)

		localctx.(*CasesContext)._expr = _x
	}
	{
		p.SetState(531)
		p.Match(SwiftGrammarParserDOSP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(532)

		var _x = p.Block()

		localctx.(*CasesContext)._block = _x
	}

	localctx.(*CasesContext).casesinst = []interface{}{}
	localctx.(*CasesContext).casesinst = append(localctx.(*CasesContext).casesinst, instructions.NewCase((func() int {
		if localctx.(*CasesContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasesContext).Get_CASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CasesContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasesContext).Get_CASE().GetColumn()
		}
	}()), localctx.(*CasesContext).Get_expr().GetE(), localctx.(*CasesContext).Get_block().GetBlk()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCasesContext(p, _parentctx, _parentState)
			localctx.(*CasesContext).ccases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_cases)
			p.SetState(535)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(536)

				var _m = p.Match(SwiftGrammarParserCASE)

				localctx.(*CasesContext)._CASE = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(537)

				var _x = p.expr(0)

				localctx.(*CasesContext)._expr = _x
			}
			{
				p.SetState(538)
				p.Match(SwiftGrammarParserDOSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(539)

				var _x = p.Block()

				localctx.(*CasesContext)._block = _x
			}

			var arrcase []interface{}
			arrcase = append(localctx.(*CasesContext).GetCcases().GetCasesinst(), instructions.NewCase((func() int {
				if localctx.(*CasesContext).Get_CASE() == nil {
					return 0
				} else {
					return localctx.(*CasesContext).Get_CASE().GetLine()
				}
			}()), (func() int {
				if localctx.(*CasesContext).Get_CASE() == nil {
					return 0
				} else {
					return localctx.(*CasesContext).Get_CASE().GetColumn()
				}
			}()), localctx.(*CasesContext).Get_expr().GetE(), localctx.(*CasesContext).Get_block().GetBlk()))
			localctx.(*CasesContext).casesinst = arrcase

		}
		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhileinst returns the whileinst attribute.
	GetWhileinst() interfaces.Instruction

	// SetWhileinst sets the whileinst attribute.
	SetWhileinst(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	whileinst interfaces.Instruction
	_WHILE    antlr.Token
	_expr     IExprContext
	_block    IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhileinst() interfaces.Instruction { return s.whileinst }

func (s *WhilestmtContext) SetWhileinst(v interfaces.Instruction) { s.whileinst = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(547)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(548)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(549)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(551)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).whileinst = instructions.NewWhile((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetForinst returns the forinst attribute.
	GetForinst() interfaces.Instruction

	// SetForinst sets the forinst attribute.
	SetForinst(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	TRESP() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	forinst interfaces.Instruction
	_FOR    antlr.Token
	_ID     antlr.Token
	e1      IExprContext
	e2      IExprContext
	_block  IBlockContext
	_expr   IExprContext
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) GetE1() IExprContext { return s.e1 }

func (s *ForstmtContext) GetE2() IExprContext { return s.e2 }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ForstmtContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ForstmtContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ForstmtContext) GetForinst() interfaces.Instruction { return s.forinst }

func (s *ForstmtContext) SetForinst(v interfaces.Instruction) { s.forinst = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) TRESP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRESP, 0)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_forstmt)
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(554)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e1 = _x
		}
		{
			p.SetState(558)
			p.Match(SwiftGrammarParserTRESP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).e2 = _x
		}
		{
			p.SetState(560)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(562)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).forinst = instructions.NewForRange((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ForstmtContext).GetE1().GetE(), localctx.(*ForstmtContext).GetE2().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(565)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)

			var _x = p.expr(0)

			localctx.(*ForstmtContext)._expr = _x
		}
		{
			p.SetState(569)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(571)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForstmtContext).forinst = instructions.NewFor((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ForstmtContext).Get_expr().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_TSTRING returns the _TSTRING token.
	Get_TSTRING() antlr.Token

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_TSTRING sets the _TSTRING token.
	Set_TSTRING(antlr.Token)

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_primitives returns the _primitives rule contexts.
	Get_primitives() IPrimitivesContext

	// Get_listArray returns the _listArray rule contexts.
	Get_listArray() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Get_listAceso returns the _listAceso rule contexts.
	Get_listAceso() IListAcesoContext

	// Get_listStructExp returns the _listStructExp rule contexts.
	Get_listStructExp() IListStructExpContext

	// Get_funcallestmt returns the _funcallestmt rule contexts.
	Get_funcallestmt() IFuncallestmtContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_primitives sets the _primitives rule contexts.
	Set_primitives(IPrimitivesContext)

	// Set_listArray sets the _listArray rule contexts.
	Set_listArray(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// Set_listAceso sets the _listAceso rule contexts.
	Set_listAceso(IListAcesoContext)

	// Set_listStructExp sets the _listStructExp rule contexts.
	Set_listStructExp(IListStructExpContext)

	// Set_funcallestmt sets the _funcallestmt rule contexts.
	Set_funcallestmt(IFuncallestmtContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	NOT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	SUB() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ID() antlr.TerminalNode
	Primitives() IPrimitivesContext
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	TSTRING() antlr.TerminalNode
	ListArray() IListArrayContext
	CORIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	CORDER() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode
	ListAceso() IListAcesoContext
	ListStructExp() IListStructExpContext
	Funcallestmt() IFuncallestmtContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	e              interfaces.Expression
	left           IExprContext
	op             antlr.Token
	_expr          IExprContext
	_ID            antlr.Token
	_primitives    IPrimitivesContext
	_INT           antlr.Token
	_FLOAT         antlr.Token
	_TSTRING       antlr.Token
	_listArray     IListArrayContext
	_CORIZQ        antlr.Token
	_listParams    IListParamsContext
	_listAceso     IListAcesoContext
	_listStructExp IListStructExpContext
	_funcallestmt  IFuncallestmtContext
	right          IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_INT() antlr.Token { return s._INT }

func (s *ExprContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *ExprContext) Get_TSTRING() antlr.Token { return s._TSTRING }

func (s *ExprContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *ExprContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *ExprContext) Set_TSTRING(v antlr.Token) { s._TSTRING = v }

func (s *ExprContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_primitives() IPrimitivesContext { return s._primitives }

func (s *ExprContext) Get_listArray() IListArrayContext { return s._listArray }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) Get_listAceso() IListAcesoContext { return s._listAceso }

func (s *ExprContext) Get_listStructExp() IListStructExpContext { return s._listStructExp }

func (s *ExprContext) Get_funcallestmt() IFuncallestmtContext { return s._funcallestmt }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_primitives(v IPrimitivesContext) { s._primitives = v }

func (s *ExprContext) Set_listArray(v IListArrayContext) { s._listArray = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) Set_listAceso(v IListAcesoContext) { s._listAceso = v }

func (s *ExprContext) Set_listStructExp(v IListStructExpContext) { s._listStructExp = v }

func (s *ExprContext) Set_funcallestmt(v IFuncallestmtContext) { s._funcallestmt = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) Primitives() IPrimitivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivesContext)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *ExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *ExprContext) TSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTSTRING, 0)
}

func (s *ExprContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ExprContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ExprContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ExprContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *ExprContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserISEMPTY, 0)
}

func (s *ExprContext) ListAceso() IListAcesoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAcesoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAcesoContext)
}

func (s *ExprContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ExprContext) Funcallestmt() IFuncallestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncallestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncallestmtContext)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(641)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(577)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(578)

			var _x = p.expr(15)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperationUnary((func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()))

	case 2:
		{
			p.SetState(581)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)

			var _x = p.expr(14)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperationUnary((func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()))

	case 3:
		{
			p.SetState(585)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(587)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 4:
		{
			p.SetState(590)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCallVariable((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 5:
		{
			p.SetState(592)

			var _x = p.Primitives()

			localctx.(*ExprContext)._primitives = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_primitives().GetP()

	case 6:
		{
			p.SetState(595)

			var _m = p.Match(SwiftGrammarParserINT)

			localctx.(*ExprContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(596)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(598)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCastingInt((func() int {
			if localctx.(*ExprContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_INT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_INT().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_expr().GetE())

	case 7:
		{
			p.SetState(601)

			var _m = p.Match(SwiftGrammarParserFLOAT)

			localctx.(*ExprContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(604)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCastingFloat((func() int {
			if localctx.(*ExprContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FLOAT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FLOAT().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_expr().GetE())

	case 8:
		{
			p.SetState(607)

			var _m = p.Match(SwiftGrammarParserTSTRING)

			localctx.(*ExprContext)._TSTRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(610)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCastingString((func() int {
			if localctx.(*ExprContext).Get_TSTRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TSTRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TSTRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TSTRING().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_expr().GetE())

	case 9:
		{
			p.SetState(613)

			var _x = p.listArray(0)

			localctx.(*ExprContext)._listArray = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_listArray().GetP()

	case 10:
		{
			p.SetState(616)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(617)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(618)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listParams().GetL())

	case 11:
		{
			p.SetState(621)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(623)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCount((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 12:
		{
			p.SetState(625)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(627)
			p.Match(SwiftGrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewIsEmpty((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 13:
		{
			p.SetState(629)

			var _x = p.listAceso(0)

			localctx.(*ExprContext)._listAceso = _x
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listAceso().GetL())

	case 14:
		{
			p.SetState(632)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(633)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)

			var _x = p.listStructExp(0)

			localctx.(*ExprContext)._listStructExp = _x
		}
		{
			p.SetState(635)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewStructExp((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()), localctx.(*ExprContext).Get_listStructExp().GetL())

	case 15:
		{
			p.SetState(638)

			var _x = p.Funcallestmt()

			localctx.(*ExprContext)._funcallestmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_funcallestmt().GetCefun()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(680)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(678)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(643)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(644)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(645)

					var _x = p.expr(23)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(648)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(649)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(650)

					var _x = p.expr(22)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(653)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(654)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(655)

					var _x = p.expr(21)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(658)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(659)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(660)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(663)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(664)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(665)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(668)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(669)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(670)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(673)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(674)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(675)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperationBinary((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(682)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitivesContext is an interface to support dynamic dispatch.
type IPrimitivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_CHARACTER returns the _CHARACTER token.
	Get_CHARACTER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_CHARACTER sets the _CHARACTER token.
	Set_CHARACTER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	NUMBER() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	NIL() antlr.TerminalNode

	// IsPrimitivesContext differentiates from other interfaces.
	IsPrimitivesContext()
}

type PrimitivesContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	_NUMBER    antlr.Token
	_CHARACTER antlr.Token
	_STRING    antlr.Token
	_TRU       antlr.Token
	_FAL       antlr.Token
	_NIL       antlr.Token
}

func NewEmptyPrimitivesContext() *PrimitivesContext {
	var p = new(PrimitivesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_primitives
	return p
}

func InitEmptyPrimitivesContext(p *PrimitivesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_primitives
}

func (*PrimitivesContext) IsPrimitivesContext() {}

func NewPrimitivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivesContext {
	var p = new(PrimitivesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_primitives

	return p
}

func (s *PrimitivesContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivesContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivesContext) Get_CHARACTER() antlr.Token { return s._CHARACTER }

func (s *PrimitivesContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivesContext) Get_TRU() antlr.Token { return s._TRU }

func (s *PrimitivesContext) Get_FAL() antlr.Token { return s._FAL }

func (s *PrimitivesContext) Get_NIL() antlr.Token { return s._NIL }

func (s *PrimitivesContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivesContext) Set_CHARACTER(v antlr.Token) { s._CHARACTER = v }

func (s *PrimitivesContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivesContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *PrimitivesContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *PrimitivesContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *PrimitivesContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitivesContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitivesContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *PrimitivesContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHARACTER, 0)
}

func (s *PrimitivesContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *PrimitivesContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *PrimitivesContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *PrimitivesContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
}

func (s *PrimitivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrimitives(s)
	}
}

func (s *PrimitivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrimitives(s)
	}
}

func (p *SwiftGrammarParser) Primitives() (localctx IPrimitivesContext) {
	localctx = NewPrimitivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_primitives)
	p.SetState(695)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(683)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*PrimitivesContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivesContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*PrimitivesContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitivesContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case SwiftGrammarParserCHARACTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(685)

			var _m = p.Match(SwiftGrammarParserCHARACTER)

			localctx.(*PrimitivesContext)._CHARACTER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*PrimitivesContext).Get_CHARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivesContext).Get_CHARACTER().GetText()
			}
		}())
		localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitivesContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_CHARACTER().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivesContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_CHARACTER().GetColumn()
			}
		}()), str[1:len(str)-1], environment.CHAR)

	case SwiftGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(687)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*PrimitivesContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*PrimitivesContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivesContext).Get_STRING().GetText()
			}
		}())
		localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitivesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case SwiftGrammarParserTRU:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(689)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*PrimitivesContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitivesContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivesContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case SwiftGrammarParserFAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(691)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*PrimitivesContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitivesContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivesContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case SwiftGrammarParserNIL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(693)

			var _m = p.Match(SwiftGrammarParserNIL)

			localctx.(*PrimitivesContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrimitivesContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitivesContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivesContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivesContext).Get_NIL().GetColumn()
			}
		}()), nil, environment.NULL)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *SwiftGrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *SwiftGrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(698)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(708)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(701)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(702)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(703)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(710)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext
	PUNTO() antlr.TerminalNode

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expression
	list   IListArrayContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
	return p
}

func InitEmptyListArrayContext(p *ListArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *SwiftGrammarParser) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *SwiftGrammarParser) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(712)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ListArrayContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ListArrayContext).p = expressions.NewCallVariable((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(727)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(725)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(715)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(716)
					p.Match(SwiftGrammarParserCORIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(717)

					var _x = p.expr(0)

					localctx.(*ListArrayContext)._expr = _x
				}
				{
					p.SetState(718)
					p.Match(SwiftGrammarParserCORDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expr().GetE())

			case 2:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(721)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(722)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(723)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListArrayContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewStructAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), (func() string {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetText()
					}
				}()))

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(729)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListAcesoContext is an interface to support dynamic dispatch.
type IListAcesoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListAcesoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListAcesoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListAceso() IListAcesoContext

	// IsListAcesoContext differentiates from other interfaces.
	IsListAcesoContext()
}

type ListAcesoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListAcesoContext
	_expr  IExprContext
}

func NewEmptyListAcesoContext() *ListAcesoContext {
	var p = new(ListAcesoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAceso
	return p
}

func InitEmptyListAcesoContext(p *ListAcesoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAceso
}

func (*ListAcesoContext) IsListAcesoContext() {}

func NewListAcesoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAcesoContext {
	var p = new(ListAcesoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listAceso

	return p
}

func (s *ListAcesoContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAcesoContext) GetList() IListAcesoContext { return s.list }

func (s *ListAcesoContext) Get_expr() IExprContext { return s._expr }

func (s *ListAcesoContext) SetList(v IListAcesoContext) { s.list = v }

func (s *ListAcesoContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListAcesoContext) GetL() []interface{} { return s.l }

func (s *ListAcesoContext) SetL(v []interface{}) { s.l = v }

func (s *ListAcesoContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListAcesoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListAcesoContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListAcesoContext) ListAceso() IListAcesoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAcesoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAcesoContext)
}

func (s *ListAcesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAcesoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAcesoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListAceso(s)
	}
}

func (s *ListAcesoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListAceso(s)
	}
}

func (p *SwiftGrammarParser) ListAceso() (localctx IListAcesoContext) {
	return p.listAceso(0)
}

func (p *SwiftGrammarParser) listAceso(_p int) (localctx IListAcesoContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListAcesoContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListAcesoContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, SwiftGrammarParserRULE_listAceso, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(731)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(732)

		var _x = p.expr(0)

		localctx.(*ListAcesoContext)._expr = _x
	}
	{
		p.SetState(733)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ListAcesoContext).l = []interface{}{}
	localctx.(*ListAcesoContext).l = append(localctx.(*ListAcesoContext).l, localctx.(*ListAcesoContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(744)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAcesoContext(p, _parentctx, _parentState)
			localctx.(*ListAcesoContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listAceso)
			p.SetState(736)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(737)
				p.Match(SwiftGrammarParserCORIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(738)

				var _x = p.expr(0)

				localctx.(*ListAcesoContext)._expr = _x
			}
			{
				p.SetState(739)
				p.Match(SwiftGrammarParserCORDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			var arr []interface{}
			arr = append(localctx.(*ListAcesoContext).GetList().GetL(), localctx.(*ListAcesoContext).Get_expr().GetE())
			localctx.(*ListAcesoContext).l = arr

		}
		p.SetState(746)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructExpContext is an interface to support dynamic dispatch.
type IListStructExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructExpContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListStructExpContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	DOSP() antlr.TerminalNode
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListStructExp() IListStructExpContext

	// IsListStructExpContext differentiates from other interfaces.
	IsListStructExpContext()
}

type ListStructExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructExpContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListStructExpContext() *ListStructExpContext {
	var p = new(ListStructExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
	return p
}

func InitEmptyListStructExpContext(p *ListStructExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
}

func (*ListStructExpContext) IsListStructExpContext() {}

func NewListStructExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructExpContext {
	var p = new(ListStructExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp

	return p
}

func (s *ListStructExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructExpContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructExpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructExpContext) GetList() IListStructExpContext { return s.list }

func (s *ListStructExpContext) Get_expr() IExprContext { return s._expr }

func (s *ListStructExpContext) SetList(v IListStructExpContext) { s.list = v }

func (s *ListStructExpContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListStructExpContext) GetL() []interface{} { return s.l }

func (s *ListStructExpContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructExpContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListStructExpContext) DOSP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSP, 0)
}

func (s *ListStructExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListStructExpContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructExpContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ListStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructExp(s)
	}
}

func (s *ListStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructExp(s)
	}
}

func (p *SwiftGrammarParser) ListStructExp() (localctx IListStructExpContext) {
	return p.listStructExp(0)
}

func (p *SwiftGrammarParser) listStructExp(_p int) (localctx IListStructExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, SwiftGrammarParserRULE_listStructExp, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(754)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(748)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructExpContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(749)
			p.Match(SwiftGrammarParserDOSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(750)

			var _x = p.expr(0)

			localctx.(*ListStructExpContext)._expr = _x
		}

		var arr []interface{}
		StrExp := environment.NewStructContent((func() string {
			if localctx.(*ListStructExpContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructExpContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructExpContext).Get_expr().GetE())
		arr = append(arr, StrExp)
		localctx.(*ListStructExpContext).l = arr

	case 2:

		localctx.(*ListStructExpContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(765)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructExpContext(p, _parentctx, _parentState)
			localctx.(*ListStructExpContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructExp)
			p.SetState(756)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(757)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(758)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructExpContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(759)
				p.Match(SwiftGrammarParserDOSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(760)

				var _x = p.expr(0)

				localctx.(*ListStructExpContext)._expr = _x
			}

			var arr []interface{}
			StrExp := environment.NewStructContent((func() string {
				if localctx.(*ListStructExpContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructExpContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructExpContext).Get_expr().GetE())
			arr = append(localctx.(*ListStructExpContext).GetList().GetL(), StrExp)
			localctx.(*ListStructExpContext).l = arr

		}
		p.SetState(767)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 6:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 16:
		var t *ElifsContext = nil
		if localctx != nil {
			t = localctx.(*ElifsContext)
		}
		return p.Elifs_Sempred(t, predIndex)

	case 18:
		var t *CasesContext = nil
		if localctx != nil {
			t = localctx.(*CasesContext)
		}
		return p.Cases_Sempred(t, predIndex)

	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 23:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 24:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 25:
		var t *ListAcesoContext = nil
		if localctx != nil {
			t = localctx.(*ListAcesoContext)
		}
		return p.ListAceso_Sempred(t, predIndex)

	case 26:
		var t *ListStructExpContext = nil
		if localctx != nil {
			t = localctx.(*ListStructExpContext)
		}
		return p.ListStructExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsFunc_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Elifs_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Cases_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListAceso_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListStructExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
