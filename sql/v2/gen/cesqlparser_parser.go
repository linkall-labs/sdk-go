<<<<<<< HEAD
// Code generated from CESQLParser.g4 by ANTLR 4.10.1. DO NOT EDIT.
=======
/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

// Code generated from CESQLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.
>>>>>>> bcb75b3 (wip: feat: supports data in sql)

package gen // CESQLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
<<<<<<< HEAD
var _ = sync.Once{}
=======

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 119,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 45, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 61, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 67, 10, 3, 3, 3, 3, 3, 7,
	3, 71, 10, 3, 12, 3, 14, 3, 74, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 81, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 99, 10, 11, 12, 11,
	14, 11, 102, 11, 11, 5, 11, 104, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 7, 12, 112, 10, 12, 12, 12, 14, 12, 115, 11, 12, 3, 12, 3, 12,
	3, 12, 2, 3, 4, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 10, 3, 2,
	13, 15, 3, 2, 16, 17, 3, 2, 18, 24, 3, 2, 9, 11, 3, 2, 33, 34, 4, 2, 33,
	33, 36, 36, 3, 2, 28, 29, 3, 2, 30, 31, 2, 128, 2, 24, 3, 2, 2, 2, 4, 44,
	3, 2, 2, 2, 6, 80, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12,
	86, 3, 2, 2, 2, 14, 88, 3, 2, 2, 2, 16, 90, 3, 2, 2, 2, 18, 92, 3, 2, 2,
	2, 20, 94, 3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26,
	7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 8, 3, 1, 2, 28, 29, 5, 12, 7, 2,
	29, 30, 5, 20, 11, 2, 30, 45, 3, 2, 2, 2, 31, 32, 7, 12, 2, 2, 32, 45,
	5, 4, 3, 14, 33, 34, 7, 17, 2, 2, 34, 45, 5, 4, 3, 13, 35, 36, 7, 26, 2,
	2, 36, 45, 5, 8, 5, 2, 37, 38, 7, 26, 2, 2, 38, 45, 5, 10, 6, 2, 39, 40,
	7, 4, 2, 2, 40, 41, 5, 4, 3, 2, 41, 42, 7, 5, 2, 2, 42, 45, 3, 2, 2, 2,
	43, 45, 5, 6, 4, 2, 44, 27, 3, 2, 2, 2, 44, 31, 3, 2, 2, 2, 44, 33, 3,
	2, 2, 2, 44, 35, 3, 2, 2, 2, 44, 37, 3, 2, 2, 2, 44, 39, 3, 2, 2, 2, 44,
	43, 3, 2, 2, 2, 45, 72, 3, 2, 2, 2, 46, 47, 12, 8, 2, 2, 47, 48, 9, 2,
	2, 2, 48, 71, 5, 4, 3, 9, 49, 50, 12, 7, 2, 2, 50, 51, 9, 3, 2, 2, 51,
	71, 5, 4, 3, 8, 52, 53, 12, 6, 2, 2, 53, 54, 9, 4, 2, 2, 54, 71, 5, 4,
	3, 7, 55, 56, 12, 5, 2, 2, 56, 57, 9, 5, 2, 2, 57, 71, 5, 4, 3, 5, 58,
	60, 12, 12, 2, 2, 59, 61, 7, 12, 2, 2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2,
	2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 7, 25, 2, 2, 63, 71, 5, 16, 9, 2, 64,
	66, 12, 9, 2, 2, 65, 67, 7, 12, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2,
	2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 7, 27, 2, 2, 69, 71, 5, 22, 12, 2, 70,
	46, 3, 2, 2, 2, 70, 49, 3, 2, 2, 2, 70, 52, 3, 2, 2, 2, 70, 55, 3, 2, 2,
	2, 70, 58, 3, 2, 2, 2, 70, 64, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 5, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	75, 81, 5, 14, 8, 2, 76, 81, 5, 18, 10, 2, 77, 81, 5, 16, 9, 2, 78, 81,
	5, 10, 6, 2, 79, 81, 5, 8, 5, 2, 80, 75, 3, 2, 2, 2, 80, 76, 3, 2, 2, 2,
	80, 77, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 7, 3, 2,
	2, 2, 82, 83, 9, 6, 2, 2, 83, 9, 3, 2, 2, 2, 84, 85, 7, 35, 2, 2, 85, 11,
	3, 2, 2, 2, 86, 87, 9, 7, 2, 2, 87, 13, 3, 2, 2, 2, 88, 89, 9, 8, 2, 2,
	89, 15, 3, 2, 2, 2, 90, 91, 9, 9, 2, 2, 91, 17, 3, 2, 2, 2, 92, 93, 7,
	32, 2, 2, 93, 19, 3, 2, 2, 2, 94, 103, 7, 4, 2, 2, 95, 100, 5, 4, 3, 2,
	96, 97, 7, 6, 2, 2, 97, 99, 5, 4, 3, 2, 98, 96, 3, 2, 2, 2, 99, 102, 3,
	2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 104, 3, 2, 2,
	2, 102, 100, 3, 2, 2, 2, 103, 95, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	105, 3, 2, 2, 2, 105, 106, 7, 5, 2, 2, 106, 21, 3, 2, 2, 2, 107, 108, 7,
	4, 2, 2, 108, 113, 5, 4, 3, 2, 109, 110, 7, 6, 2, 2, 110, 112, 5, 4, 3,
	2, 111, 109, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117,
	7, 5, 2, 2, 117, 23, 3, 2, 2, 2, 11, 44, 60, 66, 70, 72, 80, 100, 103,
	113,
}
var literalNames = []string{
	"", "", "'('", "')'", "','", "'''", "'\"'", "'AND'", "'OR'", "'XOR'", "'NOT'",
	"'*'", "'/'", "'%'", "'+'", "'-'", "'='", "'!='", "'>'", "'>='", "'<'",
	"'<>'", "'<='", "'LIKE'", "'EXISTS'", "'IN'", "'TRUE'", "'FALSE'",
}
var symbolicNames = []string{
	"", "SPACE", "LR_BRACKET", "RR_BRACKET", "COMMA", "SINGLE_QUOTE_SYMB",
	"DOUBLE_QUOTE_SYMB", "AND", "OR", "XOR", "NOT", "STAR", "DIVIDE", "MODULE",
	"PLUS", "MINUS", "EQUAL", "NOT_EQUAL", "GREATER", "GREATER_OR_EQUAL", "LESS",
	"LESS_GREATER", "LESS_OR_EQUAL", "LIKE", "EXISTS", "IN", "TRUE", "FALSE",
	"DQUOTED_STRING_LITERAL", "SQUOTED_STRING_LITERAL", "INTEGER_LITERAL",
	"IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "DATA_IDENTIFIER", "FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
}

var ruleNames = []string{
	"cesql", "expression", "atom", "identifier", "dataIdentifier", "functionIdentifier",
	"booleanLiteral", "stringLiteral", "integerLiteral", "functionParameterList",
	"setExpression",
}
>>>>>>> bcb75b3 (wip: feat: supports data in sql)

type CESQLParserParser struct {
	*antlr.BaseParser
}

var cesqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cesqlparserParserInit() {
	staticData := &cesqlparserParserStaticData
	staticData.literalNames = []string{
		"", "", "'('", "')'", "','", "'''", "'\"'", "'AND'", "'OR'", "'XOR'",
		"'NOT'", "'*'", "'/'", "'%'", "'+'", "'-'", "'='", "'!='", "'>'", "'>='",
		"'<'", "'<>'", "'<='", "'LIKE'", "'EXISTS'", "'IN'", "'TRUE'", "'FALSE'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "LR_BRACKET", "RR_BRACKET", "COMMA", "SINGLE_QUOTE_SYMB",
		"DOUBLE_QUOTE_SYMB", "AND", "OR", "XOR", "NOT", "STAR", "DIVIDE", "MODULE",
		"PLUS", "MINUS", "EQUAL", "NOT_EQUAL", "GREATER", "GREATER_OR_EQUAL",
		"LESS", "LESS_GREATER", "LESS_OR_EQUAL", "LIKE", "EXISTS", "IN", "TRUE",
		"FALSE", "DQUOTED_STRING_LITERAL", "SQUOTED_STRING_LITERAL", "INTEGER_LITERAL",
		"IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
	}
	staticData.ruleNames = []string{
		"cesql", "expression", "atom", "identifier", "functionIdentifier", "booleanLiteral",
		"stringLiteral", "integerLiteral", "functionParameterList", "setExpression",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 110, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 55, 8, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 1, 1, 1, 5, 1, 65, 8, 1, 10, 1,
		12, 1, 68, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 74, 8, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 90, 8, 8, 10, 8, 12, 8, 93, 9, 8, 3, 8, 95, 8, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 103, 8, 9, 10, 9, 12, 9, 106, 9, 9, 1, 9, 1, 9,
		1, 9, 0, 1, 2, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 8, 1, 0, 11, 13,
		1, 0, 14, 15, 1, 0, 16, 22, 1, 0, 7, 9, 1, 0, 31, 32, 2, 0, 31, 31, 33,
		33, 1, 0, 26, 27, 1, 0, 28, 29, 118, 0, 20, 1, 0, 0, 0, 2, 38, 1, 0, 0,
		0, 4, 73, 1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 79, 1,
		0, 0, 0, 12, 81, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 85, 1, 0, 0, 0, 18,
		98, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0,
		0, 23, 24, 6, 1, -1, 0, 24, 25, 3, 8, 4, 0, 25, 26, 3, 16, 8, 0, 26, 39,
		1, 0, 0, 0, 27, 28, 5, 10, 0, 0, 28, 39, 3, 2, 1, 11, 29, 30, 5, 15, 0,
		0, 30, 39, 3, 2, 1, 10, 31, 32, 5, 24, 0, 0, 32, 39, 3, 6, 3, 0, 33, 34,
		5, 2, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 3, 0, 0, 36, 39, 1, 0, 0, 0,
		37, 39, 3, 4, 2, 0, 38, 23, 1, 0, 0, 0, 38, 27, 1, 0, 0, 0, 38, 29, 1,
		0, 0, 0, 38, 31, 1, 0, 0, 0, 38, 33, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39,
		66, 1, 0, 0, 0, 40, 41, 10, 6, 0, 0, 41, 42, 7, 0, 0, 0, 42, 65, 3, 2,
		1, 7, 43, 44, 10, 5, 0, 0, 44, 45, 7, 1, 0, 0, 45, 65, 3, 2, 1, 6, 46,
		47, 10, 4, 0, 0, 47, 48, 7, 2, 0, 0, 48, 65, 3, 2, 1, 5, 49, 50, 10, 3,
		0, 0, 50, 51, 7, 3, 0, 0, 51, 65, 3, 2, 1, 3, 52, 54, 10, 9, 0, 0, 53,
		55, 5, 10, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0,
		0, 0, 56, 57, 5, 23, 0, 0, 57, 65, 3, 12, 6, 0, 58, 60, 10, 7, 0, 0, 59,
		61, 5, 10, 0, 0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0,
		0, 0, 62, 63, 5, 25, 0, 0, 63, 65, 3, 18, 9, 0, 64, 40, 1, 0, 0, 0, 64,
		43, 1, 0, 0, 0, 64, 46, 1, 0, 0, 0, 64, 49, 1, 0, 0, 0, 64, 52, 1, 0, 0,
		0, 64, 58, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 3, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 74, 3, 10, 5, 0,
		70, 74, 3, 14, 7, 0, 71, 74, 3, 12, 6, 0, 72, 74, 3, 6, 3, 0, 73, 69, 1,
		0, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74,
		5, 1, 0, 0, 0, 75, 76, 7, 4, 0, 0, 76, 7, 1, 0, 0, 0, 77, 78, 7, 5, 0,
		0, 78, 9, 1, 0, 0, 0, 79, 80, 7, 6, 0, 0, 80, 11, 1, 0, 0, 0, 81, 82, 7,
		7, 0, 0, 82, 13, 1, 0, 0, 0, 83, 84, 5, 30, 0, 0, 84, 15, 1, 0, 0, 0, 85,
		94, 5, 2, 0, 0, 86, 91, 3, 2, 1, 0, 87, 88, 5, 4, 0, 0, 88, 90, 3, 2, 1,
		0, 89, 87, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 5, 3, 0, 0, 97, 17, 1,
		0, 0, 0, 98, 99, 5, 2, 0, 0, 99, 104, 3, 2, 1, 0, 100, 101, 5, 4, 0, 0,
		101, 103, 3, 2, 1, 0, 102, 100, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 107, 108, 5, 3, 0, 0, 108, 19, 1, 0, 0, 0, 9, 38, 54, 60, 64,
		66, 73, 91, 94, 104,
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

// CESQLParserParserInit initializes any static state used to implement CESQLParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCESQLParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CESQLParserParserInit() {
	staticData := &cesqlparserParserStaticData
	staticData.once.Do(cesqlparserParserInit)
}

// NewCESQLParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCESQLParserParser(input antlr.TokenStream) *CESQLParserParser {
	CESQLParserParserInit()
	this := new(CESQLParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cesqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CESQLParser.g4"

	return this
}

// CESQLParserParser tokens.
const (
	CESQLParserParserEOF                                 = antlr.TokenEOF
	CESQLParserParserSPACE                               = 1
	CESQLParserParserLR_BRACKET                          = 2
	CESQLParserParserRR_BRACKET                          = 3
	CESQLParserParserCOMMA                               = 4
	CESQLParserParserSINGLE_QUOTE_SYMB                   = 5
	CESQLParserParserDOUBLE_QUOTE_SYMB                   = 6
	CESQLParserParserAND                                 = 7
	CESQLParserParserOR                                  = 8
	CESQLParserParserXOR                                 = 9
	CESQLParserParserNOT                                 = 10
	CESQLParserParserSTAR                                = 11
	CESQLParserParserDIVIDE                              = 12
	CESQLParserParserMODULE                              = 13
	CESQLParserParserPLUS                                = 14
	CESQLParserParserMINUS                               = 15
	CESQLParserParserEQUAL                               = 16
	CESQLParserParserNOT_EQUAL                           = 17
	CESQLParserParserGREATER                             = 18
	CESQLParserParserGREATER_OR_EQUAL                    = 19
	CESQLParserParserLESS                                = 20
	CESQLParserParserLESS_GREATER                        = 21
	CESQLParserParserLESS_OR_EQUAL                       = 22
	CESQLParserParserLIKE                                = 23
	CESQLParserParserEXISTS                              = 24
	CESQLParserParserIN                                  = 25
	CESQLParserParserTRUE                                = 26
	CESQLParserParserFALSE                               = 27
	CESQLParserParserDQUOTED_STRING_LITERAL              = 28
	CESQLParserParserSQUOTED_STRING_LITERAL              = 29
	CESQLParserParserINTEGER_LITERAL                     = 30
	CESQLParserParserIDENTIFIER                          = 31
	CESQLParserParserIDENTIFIER_WITH_NUMBER              = 32
	CESQLParserParserDATA_IDENTIFIER                     = 33
	CESQLParserParserFUNCTION_IDENTIFIER_WITH_UNDERSCORE = 34
)

// CESQLParserParser rules.
const (
	CESQLParserParserRULE_cesql                 = 0
	CESQLParserParserRULE_expression            = 1
	CESQLParserParserRULE_atom                  = 2
	CESQLParserParserRULE_identifier            = 3
	CESQLParserParserRULE_dataIdentifier        = 4
	CESQLParserParserRULE_functionIdentifier    = 5
	CESQLParserParserRULE_booleanLiteral        = 6
	CESQLParserParserRULE_stringLiteral         = 7
	CESQLParserParserRULE_integerLiteral        = 8
	CESQLParserParserRULE_functionParameterList = 9
	CESQLParserParserRULE_setExpression         = 10
)

// ICesqlContext is an interface to support dynamic dispatch.
type ICesqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCesqlContext differentiates from other interfaces.
	IsCesqlContext()
}

type CesqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCesqlContext() *CesqlContext {
	var p = new(CesqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_cesql
	return p
}

func (*CesqlContext) IsCesqlContext() {}

func NewCesqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CesqlContext {
	var p = new(CesqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_cesql

	return p
}

func (s *CesqlContext) GetParser() antlr.Parser { return s.parser }

func (s *CesqlContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CesqlContext) EOF() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserEOF, 0)
}

func (s *CesqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CesqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CesqlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitCesql(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) Cesql() (localctx ICesqlContext) {
	this := p
	_ = this

	localctx = NewCesqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CESQLParserParserRULE_cesql)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.expression(0)
	}
	{
		p.SetState(23)
		p.Match(CESQLParserParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

<<<<<<< HEAD
type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserIN, 0)
}

func (s *InExpressionContext) SetExpression() ISetExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetExpressionContext)
}

func (s *InExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserNOT, 0)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

=======
>>>>>>> bcb75b3 (wip: feat: supports data in sql)
type BinaryComparisonExpressionContext struct {
	*ExpressionContext
}

func NewBinaryComparisonExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryComparisonExpressionContext {
	var p = new(BinaryComparisonExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryComparisonExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryComparisonExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserEQUAL, 0)
}

func (s *BinaryComparisonExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserNOT_EQUAL, 0)
}

func (s *BinaryComparisonExpressionContext) LESS_GREATER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLESS_GREATER, 0)
}

func (s *BinaryComparisonExpressionContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserGREATER_OR_EQUAL, 0)
}

func (s *BinaryComparisonExpressionContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLESS_OR_EQUAL, 0)
}

func (s *BinaryComparisonExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLESS, 0)
}

func (s *BinaryComparisonExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserGREATER, 0)
}

func (s *BinaryComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBinaryComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionInvocationExpressionContext struct {
	*ExpressionContext
}

func NewFunctionInvocationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionInvocationExpressionContext {
	var p = new(FunctionInvocationExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionInvocationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvocationExpressionContext) FunctionIdentifier() IFunctionIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionIdentifierContext)
}

func (s *FunctionInvocationExpressionContext) FunctionParameterList() IFunctionParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterListContext)
}

func (s *FunctionInvocationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitFunctionInvocationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryLogicExpressionContext struct {
	*ExpressionContext
}

func NewUnaryLogicExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryLogicExpressionContext {
	var p = new(UnaryLogicExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryLogicExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryLogicExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserNOT, 0)
}

func (s *UnaryLogicExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryLogicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitUnaryLogicExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserIN, 0)
}

func (s *InExpressionContext) SetExpression() ISetExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetExpressionContext)
}

func (s *InExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserNOT, 0)
}

func (s *InExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitInExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomExpressionContext struct {
	*ExpressionContext
}

func NewAtomExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitAtomExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistsExpressionContext struct {
	*ExpressionContext
}

func NewExistsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistsExpressionContext {
	var p = new(ExistsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExistsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsExpressionContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserEXISTS, 0)
}

func (s *ExistsExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExistsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitExistsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryLogicExpressionContext struct {
	*ExpressionContext
}

func NewBinaryLogicExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLogicExpressionContext {
	var p = new(BinaryLogicExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryLogicExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLogicExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryLogicExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryLogicExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserAND, 0)
}

func (s *BinaryLogicExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserOR, 0)
}

func (s *BinaryLogicExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserXOR, 0)
}

func (s *BinaryLogicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBinaryLogicExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikeExpressionContext struct {
	*ExpressionContext
}

func NewLikeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeExpressionContext {
	var p = new(LikeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LikeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LikeExpressionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLIKE, 0)
}

func (s *LikeExpressionContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LikeExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserNOT, 0)
}

func (s *LikeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitLikeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

<<<<<<< HEAD
type FunctionInvocationExpressionContext struct {
	*ExpressionContext
}

func NewFunctionInvocationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionInvocationExpressionContext {
	var p = new(FunctionInvocationExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionInvocationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvocationExpressionContext) FunctionIdentifier() IFunctionIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionIdentifierContext)
}

func (s *FunctionInvocationExpressionContext) FunctionParameterList() IFunctionParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterListContext)
}

func (s *FunctionInvocationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitFunctionInvocationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

=======
>>>>>>> bcb75b3 (wip: feat: supports data in sql)
type BinaryMultiplicativeExpressionContext struct {
	*ExpressionContext
}

func NewBinaryMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryMultiplicativeExpressionContext {
	var p = new(BinaryMultiplicativeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryMultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryMultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryMultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryMultiplicativeExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserSTAR, 0)
}

func (s *BinaryMultiplicativeExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserDIVIDE, 0)
}

func (s *BinaryMultiplicativeExpressionContext) MODULE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserMODULE, 0)
}

func (s *BinaryMultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBinaryMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type DataExistsExpressionContext struct {
	*ExpressionContext
}

func NewDataExistsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataExistsExpressionContext {
	var p = new(DataExistsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DataExistsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataExistsExpressionContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserEXISTS, 0)
}

<<<<<<< HEAD
func (s *UnaryLogicExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}
=======
func (s *DataExistsExpressionContext) DataIdentifier() IDataIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataIdentifierContext)(nil)).Elem(), 0)
>>>>>>> bcb75b3 (wip: feat: supports data in sql)

	if t == nil {
		return nil
	}

	return t.(IDataIdentifierContext)
}

func (s *DataExistsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitDataExistsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryNumericExpressionContext struct {
	*ExpressionContext
}

func NewUnaryNumericExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryNumericExpressionContext {
	var p = new(UnaryNumericExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryNumericExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryNumericExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserMINUS, 0)
}

func (s *UnaryNumericExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryNumericExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitUnaryNumericExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExpressionContext struct {
	*ExpressionContext
}

func NewSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExpressionContext {
	var p = new(SubExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLR_BRACKET, 0)
}

func (s *SubExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserRR_BRACKET, 0)
}

func (s *SubExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitSubExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryAdditiveExpressionContext struct {
	*ExpressionContext
}

func NewBinaryAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryAdditiveExpressionContext {
	var p = new(BinaryAdditiveExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryAdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryAdditiveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryAdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryAdditiveExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserPLUS, 0)
}

func (s *BinaryAdditiveExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserMINUS, 0)
}

func (s *BinaryAdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBinaryAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CESQLParserParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CESQLParserParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionInvocationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(26)
			p.FunctionIdentifier()
		}
		{
			p.SetState(27)
			p.FunctionParameterList()
		}

	case 2:
		localctx = NewUnaryLogicExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(CESQLParserParserNOT)
		}
		{
			p.SetState(30)
			p.expression(12)
		}

	case 3:
		localctx = NewUnaryNumericExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Match(CESQLParserParserMINUS)
		}
		{
			p.SetState(32)
			p.expression(11)
		}

	case 4:
		localctx = NewExistsExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(CESQLParserParserEXISTS)
		}
		{
			p.SetState(34)
			p.Identifier()
		}

	case 5:
		localctx = NewDataExistsExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(CESQLParserParserEXISTS)
		}
		{
			p.SetState(36)
			p.DataIdentifier()
		}

	case 6:
		localctx = NewSubExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(CESQLParserParserLR_BRACKET)
		}
		{
			p.SetState(38)
			p.expression(0)
		}
		{
			p.SetState(39)
			p.Match(CESQLParserParserRR_BRACKET)
		}

	case 7:
		localctx = NewAtomExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.Atom()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(45)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CESQLParserParserSTAR)|(1<<CESQLParserParserDIVIDE)|(1<<CESQLParserParserMODULE))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)
					p.expression(7)
				}

			case 2:
				localctx = NewBinaryAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(48)
					_la = p.GetTokenStream().LA(1)

					if !(_la == CESQLParserParserPLUS || _la == CESQLParserParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(49)
					p.expression(6)
				}

			case 3:
				localctx = NewBinaryComparisonExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(51)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CESQLParserParserEQUAL)|(1<<CESQLParserParserNOT_EQUAL)|(1<<CESQLParserParserGREATER)|(1<<CESQLParserParserGREATER_OR_EQUAL)|(1<<CESQLParserParserLESS)|(1<<CESQLParserParserLESS_GREATER)|(1<<CESQLParserParserLESS_OR_EQUAL))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(52)
					p.expression(5)
				}

			case 4:
				localctx = NewBinaryLogicExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(54)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CESQLParserParserAND)|(1<<CESQLParserParserOR)|(1<<CESQLParserParserXOR))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(55)
					p.expression(3)
				}

			case 5:
				localctx = NewLikeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(58)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == CESQLParserParserNOT {
					{
						p.SetState(57)
						p.Match(CESQLParserParserNOT)
					}

				}
				{
					p.SetState(60)
					p.Match(CESQLParserParserLIKE)
				}
				{
					p.SetState(61)
					p.StringLiteral()
				}

			case 6:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CESQLParserParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(64)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == CESQLParserParserNOT {
					{
						p.SetState(63)
						p.Match(CESQLParserParserNOT)
					}

				}
				{
					p.SetState(66)
					p.Match(CESQLParserParserIN)
				}
				{
					p.SetState(67)
					p.SetExpression()
				}

			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanAtomContext struct {
	*AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierAtomContext struct {
	*AtomContext
}

func NewIdentifierAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierAtomContext {
	var p = new(IdentifierAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IdentifierAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierAtomContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitIdentifierAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringAtomContext struct {
	*AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerAtomContext struct {
	*AtomContext
}

func NewIntegerAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerAtomContext {
	var p = new(IntegerAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IntegerAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerAtomContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *IntegerAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitIntegerAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type DataIdentifierAtomContext struct {
	*AtomContext
}

func NewDataIdentifierAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataIdentifierAtomContext {
	var p = new(DataIdentifierAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DataIdentifierAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataIdentifierAtomContext) DataIdentifier() IDataIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataIdentifierContext)
}

func (s *DataIdentifierAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitDataIdentifierAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CESQLParserParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CESQLParserParserTRUE, CESQLParserParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.BooleanLiteral()
		}

	case CESQLParserParserINTEGER_LITERAL:
		localctx = NewIntegerAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.IntegerLiteral()
		}

	case CESQLParserParserDQUOTED_STRING_LITERAL, CESQLParserParserSQUOTED_STRING_LITERAL:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.StringLiteral()
		}

	case CESQLParserParserDATA_IDENTIFIER:
		localctx = NewDataIdentifierAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.DataIdentifier()
		}

	case CESQLParserParserIDENTIFIER, CESQLParserParserIDENTIFIER_WITH_NUMBER:
		localctx = NewIdentifierAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserIDENTIFIER, 0)
}

func (s *IdentifierContext) IDENTIFIER_WITH_NUMBER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserIDENTIFIER_WITH_NUMBER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CESQLParserParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CESQLParserParserIDENTIFIER || _la == CESQLParserParserIDENTIFIER_WITH_NUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDataIdentifierContext is an interface to support dynamic dispatch.
type IDataIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataIdentifierContext differentiates from other interfaces.
	IsDataIdentifierContext()
}

type DataIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataIdentifierContext() *DataIdentifierContext {
	var p = new(DataIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_dataIdentifier
	return p
}

func (*DataIdentifierContext) IsDataIdentifierContext() {}

func NewDataIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataIdentifierContext {
	var p = new(DataIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_dataIdentifier

	return p
}

func (s *DataIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *DataIdentifierContext) DATA_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserDATA_IDENTIFIER, 0)
}

func (s *DataIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitDataIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) DataIdentifier() (localctx IDataIdentifierContext) {
	this := p
	_ = this

	localctx = NewDataIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CESQLParserParserRULE_dataIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(CESQLParserParserDATA_IDENTIFIER)
	}

	return localctx
}

// IFunctionIdentifierContext is an interface to support dynamic dispatch.
type IFunctionIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionIdentifierContext differentiates from other interfaces.
	IsFunctionIdentifierContext()
}

type FunctionIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionIdentifierContext() *FunctionIdentifierContext {
	var p = new(FunctionIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_functionIdentifier
	return p
}

func (*FunctionIdentifierContext) IsFunctionIdentifierContext() {}

func NewFunctionIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionIdentifierContext {
	var p = new(FunctionIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_functionIdentifier

	return p
}

func (s *FunctionIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserIDENTIFIER, 0)
}

func (s *FunctionIdentifierContext) FUNCTION_IDENTIFIER_WITH_UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserFUNCTION_IDENTIFIER_WITH_UNDERSCORE, 0)
}

func (s *FunctionIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitFunctionIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) FunctionIdentifier() (localctx IFunctionIdentifierContext) {
	this := p
	_ = this

	localctx = NewFunctionIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CESQLParserParserRULE_functionIdentifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CESQLParserParserIDENTIFIER || _la == CESQLParserParserFUNCTION_IDENTIFIER_WITH_UNDERSCORE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CESQLParserParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CESQLParserParserTRUE || _la == CESQLParserParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTED_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserDQUOTED_STRING_LITERAL, 0)
}

func (s *StringLiteralContext) SQUOTED_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserSQUOTED_STRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CESQLParserParserRULE_stringLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CESQLParserParserDQUOTED_STRING_LITERAL || _la == CESQLParserParserSQUOTED_STRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserINTEGER_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CESQLParserParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(CESQLParserParserINTEGER_LITERAL)
	}

	return localctx
}

// IFunctionParameterListContext is an interface to support dynamic dispatch.
type IFunctionParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterListContext differentiates from other interfaces.
	IsFunctionParameterListContext()
}

type FunctionParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterListContext() *FunctionParameterListContext {
	var p = new(FunctionParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_functionParameterList
	return p
}

func (*FunctionParameterListContext) IsFunctionParameterListContext() {}

func NewFunctionParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterListContext {
	var p = new(FunctionParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_functionParameterList

	return p
}

func (s *FunctionParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterListContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLR_BRACKET, 0)
}

func (s *FunctionParameterListContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserRR_BRACKET, 0)
}

func (s *FunctionParameterListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionParameterListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *FunctionParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CESQLParserParserCOMMA)
}

func (s *FunctionParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CESQLParserParserCOMMA, i)
}

func (s *FunctionParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitFunctionParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) FunctionParameterList() (localctx IFunctionParameterListContext) {
	this := p
	_ = this

	localctx = NewFunctionParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CESQLParserParserRULE_functionParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(CESQLParserParserLR_BRACKET)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CESQLParserParserLR_BRACKET)|(1<<CESQLParserParserNOT)|(1<<CESQLParserParserMINUS)|(1<<CESQLParserParserEXISTS)|(1<<CESQLParserParserTRUE)|(1<<CESQLParserParserFALSE)|(1<<CESQLParserParserDQUOTED_STRING_LITERAL)|(1<<CESQLParserParserSQUOTED_STRING_LITERAL)|(1<<CESQLParserParserINTEGER_LITERAL)|(1<<CESQLParserParserIDENTIFIER))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(CESQLParserParserIDENTIFIER_WITH_NUMBER-32))|(1<<(CESQLParserParserDATA_IDENTIFIER-32))|(1<<(CESQLParserParserFUNCTION_IDENTIFIER_WITH_UNDERSCORE-32)))) != 0) {
		{
			p.SetState(93)
			p.expression(0)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CESQLParserParserCOMMA {
			{
				p.SetState(94)
				p.Match(CESQLParserParserCOMMA)
			}
			{
				p.SetState(95)
				p.expression(0)
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(103)
		p.Match(CESQLParserParserRR_BRACKET)
	}

	return localctx
}

// ISetExpressionContext is an interface to support dynamic dispatch.
type ISetExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetExpressionContext differentiates from other interfaces.
	IsSetExpressionContext()
}

type SetExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetExpressionContext() *SetExpressionContext {
	var p = new(SetExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CESQLParserParserRULE_setExpression
	return p
}

func (*SetExpressionContext) IsSetExpressionContext() {}

func NewSetExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetExpressionContext {
	var p = new(SetExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CESQLParserParserRULE_setExpression

	return p
}

func (s *SetExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SetExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserLR_BRACKET, 0)
}

func (s *SetExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SetExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *SetExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(CESQLParserParserRR_BRACKET, 0)
}

func (s *SetExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CESQLParserParserCOMMA)
}

func (s *SetExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CESQLParserParserCOMMA, i)
}

func (s *SetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CESQLParserVisitor:
		return t.VisitSetExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CESQLParserParser) SetExpression() (localctx ISetExpressionContext) {
	this := p
	_ = this

	localctx = NewSetExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CESQLParserParserRULE_setExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(CESQLParserParserLR_BRACKET)
	}
	{
		p.SetState(106)
		p.expression(0)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CESQLParserParserCOMMA {
		{
			p.SetState(107)
			p.Match(CESQLParserParserCOMMA)
		}
		{
			p.SetState(108)
			p.expression(0)
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(CESQLParserParserRR_BRACKET)
	}

	return localctx
}

func (p *CESQLParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CESQLParserParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
