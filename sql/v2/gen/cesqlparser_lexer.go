// Code generated from CESQLParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CESQLParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cesqlparserlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cesqlparserlexerLexerInit() {
	staticData := &cesqlparserlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "DATA_IDENTIFIER", "FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
	}
	staticData.ruleNames = []string{
		"SPACE", "ID_LITERAL", "DQUOTA_STRING", "SQUOTA_STRING", "INT_DIGIT",
		"FN_LITERAL", "LR_BRACKET", "RR_BRACKET", "COMMA", "SINGLE_QUOTE_SYMB",
		"DOUBLE_QUOTE_SYMB", "QUOTE_SYMB", "AND", "OR", "XOR", "NOT", "STAR",
		"DIVIDE", "MODULE", "PLUS", "MINUS", "EQUAL", "NOT_EQUAL", "GREATER",
		"GREATER_OR_EQUAL", "LESS", "LESS_GREATER", "LESS_OR_EQUAL", "LIKE",
		"EXISTS", "IN", "TRUE", "FALSE", "DQUOTED_STRING_LITERAL", "SQUOTED_STRING_LITERAL",
		"INTEGER_LITERAL", "IDENTIFIER", "IDENTIFIER_WITH_NUMBER", "DATA_IDENTIFIER",
		"FUNCTION_IDENTIFIER_WITH_UNDERSCORE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 252, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 4, 0, 83, 8, 0,
		11, 0, 12, 0, 84, 1, 0, 1, 0, 1, 1, 4, 1, 90, 8, 1, 11, 1, 12, 1, 91, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 100, 8, 2, 10, 2, 12, 2, 103, 9,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 113, 8, 3, 10,
		3, 12, 3, 116, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 124, 8,
		5, 10, 5, 12, 5, 127, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 141, 8, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 4, 35, 217, 8,
		35, 11, 35, 12, 35, 218, 1, 36, 4, 36, 222, 8, 36, 11, 36, 12, 36, 223,
		1, 37, 4, 37, 227, 8, 37, 11, 37, 12, 37, 228, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 4, 38, 238, 8, 38, 11, 38, 12, 38, 239, 4, 38,
		242, 8, 38, 11, 38, 12, 38, 243, 1, 39, 1, 39, 5, 39, 248, 8, 39, 10, 39,
		12, 39, 251, 9, 39, 0, 0, 40, 1, 1, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13,
		2, 15, 3, 17, 4, 19, 5, 21, 6, 23, 0, 25, 7, 27, 8, 29, 9, 31, 10, 33,
		11, 35, 12, 37, 13, 39, 14, 41, 15, 43, 16, 45, 17, 47, 18, 49, 19, 51,
		20, 53, 21, 55, 22, 57, 23, 59, 24, 61, 25, 63, 26, 65, 27, 67, 28, 69,
		29, 71, 30, 73, 31, 75, 32, 77, 33, 79, 34, 1, 0, 9, 3, 0, 9, 10, 13, 13,
		32, 32, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 34, 34, 92, 92, 2, 0, 39,
		39, 92, 92, 1, 0, 48, 57, 1, 0, 65, 90, 2, 0, 65, 90, 95, 95, 2, 0, 65,
		90, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 261, 0, 1,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 82, 1, 0, 0,
		0, 3, 89, 1, 0, 0, 0, 5, 93, 1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 119, 1,
		0, 0, 0, 11, 121, 1, 0, 0, 0, 13, 128, 1, 0, 0, 0, 15, 130, 1, 0, 0, 0,
		17, 132, 1, 0, 0, 0, 19, 134, 1, 0, 0, 0, 21, 136, 1, 0, 0, 0, 23, 140,
		1, 0, 0, 0, 25, 142, 1, 0, 0, 0, 27, 146, 1, 0, 0, 0, 29, 149, 1, 0, 0,
		0, 31, 153, 1, 0, 0, 0, 33, 157, 1, 0, 0, 0, 35, 159, 1, 0, 0, 0, 37, 161,
		1, 0, 0, 0, 39, 163, 1, 0, 0, 0, 41, 165, 1, 0, 0, 0, 43, 167, 1, 0, 0,
		0, 45, 169, 1, 0, 0, 0, 47, 172, 1, 0, 0, 0, 49, 174, 1, 0, 0, 0, 51, 177,
		1, 0, 0, 0, 53, 179, 1, 0, 0, 0, 55, 182, 1, 0, 0, 0, 57, 185, 1, 0, 0,
		0, 59, 190, 1, 0, 0, 0, 61, 197, 1, 0, 0, 0, 63, 200, 1, 0, 0, 0, 65, 205,
		1, 0, 0, 0, 67, 211, 1, 0, 0, 0, 69, 213, 1, 0, 0, 0, 71, 216, 1, 0, 0,
		0, 73, 221, 1, 0, 0, 0, 75, 226, 1, 0, 0, 0, 77, 230, 1, 0, 0, 0, 79, 245,
		1, 0, 0, 0, 81, 83, 7, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 6,
		0, 0, 0, 87, 2, 1, 0, 0, 0, 88, 90, 7, 1, 0, 0, 89, 88, 1, 0, 0, 0, 90,
		91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 4, 1, 0, 0,
		0, 93, 101, 5, 34, 0, 0, 94, 95, 5, 92, 0, 0, 95, 100, 9, 0, 0, 0, 96,
		97, 5, 34, 0, 0, 97, 100, 5, 34, 0, 0, 98, 100, 8, 2, 0, 0, 99, 94, 1,
		0, 0, 0, 99, 96, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0,
		101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101,
		1, 0, 0, 0, 104, 105, 5, 34, 0, 0, 105, 6, 1, 0, 0, 0, 106, 114, 5, 39,
		0, 0, 107, 108, 5, 92, 0, 0, 108, 113, 9, 0, 0, 0, 109, 110, 5, 39, 0,
		0, 110, 113, 5, 39, 0, 0, 111, 113, 8, 3, 0, 0, 112, 107, 1, 0, 0, 0, 112,
		109, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 117, 118, 5, 39, 0, 0, 118, 8, 1, 0, 0, 0, 119, 120, 7, 4, 0, 0,
		120, 10, 1, 0, 0, 0, 121, 125, 7, 5, 0, 0, 122, 124, 7, 6, 0, 0, 123, 122,
		1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0,
		0, 0, 126, 12, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 129, 5, 40, 0, 0,
		129, 14, 1, 0, 0, 0, 130, 131, 5, 41, 0, 0, 131, 16, 1, 0, 0, 0, 132, 133,
		5, 44, 0, 0, 133, 18, 1, 0, 0, 0, 134, 135, 5, 39, 0, 0, 135, 20, 1, 0,
		0, 0, 136, 137, 5, 34, 0, 0, 137, 22, 1, 0, 0, 0, 138, 141, 3, 19, 9, 0,
		139, 141, 3, 21, 10, 0, 140, 138, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141,
		24, 1, 0, 0, 0, 142, 143, 5, 65, 0, 0, 143, 144, 5, 78, 0, 0, 144, 145,
		5, 68, 0, 0, 145, 26, 1, 0, 0, 0, 146, 147, 5, 79, 0, 0, 147, 148, 5, 82,
		0, 0, 148, 28, 1, 0, 0, 0, 149, 150, 5, 88, 0, 0, 150, 151, 5, 79, 0, 0,
		151, 152, 5, 82, 0, 0, 152, 30, 1, 0, 0, 0, 153, 154, 5, 78, 0, 0, 154,
		155, 5, 79, 0, 0, 155, 156, 5, 84, 0, 0, 156, 32, 1, 0, 0, 0, 157, 158,
		5, 42, 0, 0, 158, 34, 1, 0, 0, 0, 159, 160, 5, 47, 0, 0, 160, 36, 1, 0,
		0, 0, 161, 162, 5, 37, 0, 0, 162, 38, 1, 0, 0, 0, 163, 164, 5, 43, 0, 0,
		164, 40, 1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166, 42, 1, 0, 0, 0, 167, 168,
		5, 61, 0, 0, 168, 44, 1, 0, 0, 0, 169, 170, 5, 33, 0, 0, 170, 171, 5, 61,
		0, 0, 171, 46, 1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 48, 1, 0, 0, 0,
		174, 175, 5, 62, 0, 0, 175, 176, 5, 61, 0, 0, 176, 50, 1, 0, 0, 0, 177,
		178, 5, 60, 0, 0, 178, 52, 1, 0, 0, 0, 179, 180, 5, 60, 0, 0, 180, 181,
		5, 62, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183, 5, 60, 0, 0, 183, 184, 5, 61,
		0, 0, 184, 56, 1, 0, 0, 0, 185, 186, 5, 76, 0, 0, 186, 187, 5, 73, 0, 0,
		187, 188, 5, 75, 0, 0, 188, 189, 5, 69, 0, 0, 189, 58, 1, 0, 0, 0, 190,
		191, 5, 69, 0, 0, 191, 192, 5, 88, 0, 0, 192, 193, 5, 73, 0, 0, 193, 194,
		5, 83, 0, 0, 194, 195, 5, 84, 0, 0, 195, 196, 5, 83, 0, 0, 196, 60, 1,
		0, 0, 0, 197, 198, 5, 73, 0, 0, 198, 199, 5, 78, 0, 0, 199, 62, 1, 0, 0,
		0, 200, 201, 5, 84, 0, 0, 201, 202, 5, 82, 0, 0, 202, 203, 5, 85, 0, 0,
		203, 204, 5, 69, 0, 0, 204, 64, 1, 0, 0, 0, 205, 206, 5, 70, 0, 0, 206,
		207, 5, 65, 0, 0, 207, 208, 5, 76, 0, 0, 208, 209, 5, 83, 0, 0, 209, 210,
		5, 69, 0, 0, 210, 66, 1, 0, 0, 0, 211, 212, 3, 5, 2, 0, 212, 68, 1, 0,
		0, 0, 213, 214, 3, 7, 3, 0, 214, 70, 1, 0, 0, 0, 215, 217, 3, 9, 4, 0,
		216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218,
		219, 1, 0, 0, 0, 219, 72, 1, 0, 0, 0, 220, 222, 7, 7, 0, 0, 221, 220, 1,
		0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0,
		0, 224, 74, 1, 0, 0, 0, 225, 227, 7, 1, 0, 0, 226, 225, 1, 0, 0, 0, 227,
		228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 76, 1,
		0, 0, 0, 230, 231, 5, 68, 0, 0, 231, 232, 5, 65, 0, 0, 232, 233, 5, 84,
		0, 0, 233, 234, 5, 65, 0, 0, 234, 241, 1, 0, 0, 0, 235, 237, 5, 46, 0,
		0, 236, 238, 7, 8, 0, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 235,
		1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 78, 1, 0, 0, 0, 245, 249, 7, 5, 0, 0, 246, 248, 7, 6, 0, 0,
		247, 246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249,
		250, 1, 0, 0, 0, 250, 80, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 15, 0, 84,
		91, 99, 101, 112, 114, 125, 140, 218, 223, 228, 239, 243, 249, 1, 6, 0,
		0,
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

// CESQLParserLexerInit initializes any static state used to implement CESQLParserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCESQLParserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CESQLParserLexerInit() {
	staticData := &cesqlparserlexerLexerStaticData
	staticData.once.Do(cesqlparserlexerLexerInit)
}

// NewCESQLParserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCESQLParserLexer(input antlr.CharStream) *CESQLParserLexer {
	CESQLParserLexerInit()
	l := new(CESQLParserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cesqlparserlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CESQLParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CESQLParserLexer tokens.
const (
	CESQLParserLexerSPACE                               = 1
	CESQLParserLexerLR_BRACKET                          = 2
	CESQLParserLexerRR_BRACKET                          = 3
	CESQLParserLexerCOMMA                               = 4
	CESQLParserLexerSINGLE_QUOTE_SYMB                   = 5
	CESQLParserLexerDOUBLE_QUOTE_SYMB                   = 6
	CESQLParserLexerAND                                 = 7
	CESQLParserLexerOR                                  = 8
	CESQLParserLexerXOR                                 = 9
	CESQLParserLexerNOT                                 = 10
	CESQLParserLexerSTAR                                = 11
	CESQLParserLexerDIVIDE                              = 12
	CESQLParserLexerMODULE                              = 13
	CESQLParserLexerPLUS                                = 14
	CESQLParserLexerMINUS                               = 15
	CESQLParserLexerEQUAL                               = 16
	CESQLParserLexerNOT_EQUAL                           = 17
	CESQLParserLexerGREATER                             = 18
	CESQLParserLexerGREATER_OR_EQUAL                    = 19
	CESQLParserLexerLESS                                = 20
	CESQLParserLexerLESS_GREATER                        = 21
	CESQLParserLexerLESS_OR_EQUAL                       = 22
	CESQLParserLexerLIKE                                = 23
	CESQLParserLexerEXISTS                              = 24
	CESQLParserLexerIN                                  = 25
	CESQLParserLexerTRUE                                = 26
	CESQLParserLexerFALSE                               = 27
	CESQLParserLexerDQUOTED_STRING_LITERAL              = 28
	CESQLParserLexerSQUOTED_STRING_LITERAL              = 29
	CESQLParserLexerINTEGER_LITERAL                     = 30
	CESQLParserLexerIDENTIFIER                          = 31
	CESQLParserLexerIDENTIFIER_WITH_NUMBER              = 32
	CESQLParserLexerDATA_IDENTIFIER                     = 33
	CESQLParserLexerFUNCTION_IDENTIFIER_WITH_UNDERSCORE = 34
)
