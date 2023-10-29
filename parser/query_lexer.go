// Code generated from Query.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var QueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylexerLexerInit() {
	staticData := &QueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'('", "')'", "'='", "'!='", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "STRING", "AND", "OR", "LPAREN", "RPAREN", "EQ", "NE", "GT", "GTE",
		"LT", "LTE", "IDENTIFIER", "NEG_IDENTIFIER", "ENCODED_STRING", "LineTerminator",
		"WS",
	}
	staticData.RuleNames = []string{
		"STRING", "StringCharacter", "EscapeSequence", "CharacterEscapeSequence",
		"HexEscapeSequence", "UnicodeEscapeSequence", "SingleEscapeCharacter",
		"NonEscapeCharacter", "EscapeCharacter", "LineContinuation", "LineTerminatorSequence",
		"DecimalDigit", "HexDigit", "OctalDigit", "AND", "OR", "LPAREN", "RPAREN",
		"EQ", "NE", "GT", "GTE", "LT", "LTE", "IDENTIFIER", "NEG_IDENTIFIER",
		"ENCODED_STRING", "LineTerminator", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 176, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 5, 0, 62, 8, 0, 10,
		0, 12, 0, 65, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 73, 8, 1,
		1, 2, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 3, 1, 3, 3, 3, 82, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 3, 8, 101, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 3, 10, 109, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 123, 8, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 129, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 24, 4, 24, 151, 8, 24, 11, 24, 12, 24, 152, 1, 25,
		1, 25, 4, 25, 157, 8, 25, 11, 25, 12, 25, 158, 1, 26, 4, 26, 162, 8, 26,
		11, 26, 12, 26, 163, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28, 171, 8,
		28, 11, 28, 12, 28, 172, 1, 28, 1, 28, 0, 0, 29, 1, 1, 3, 0, 5, 0, 7, 0,
		9, 0, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29,
		2, 31, 3, 33, 4, 35, 5, 37, 6, 39, 7, 41, 8, 43, 9, 45, 10, 47, 11, 49,
		12, 51, 13, 53, 14, 55, 15, 57, 16, 1, 0, 11, 4, 0, 10, 10, 13, 13, 34,
		34, 92, 92, 9, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 118, 118, 12, 0, 10, 10, 13, 13, 34, 34, 39, 39, 48, 57,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 118, 120, 120, 2, 0,
		117, 117, 120, 120, 1, 0, 48, 57, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0,
		48, 55, 5, 0, 45, 46, 48, 58, 65, 90, 95, 95, 97, 122, 4, 0, 32, 33, 40,
		41, 60, 62, 91, 93, 3, 0, 10, 10, 13, 13, 8232, 8233, 3, 0, 9, 10, 13,
		13, 32, 32, 177, 0, 1, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 72, 1, 0, 0, 0, 5, 77,
		1, 0, 0, 0, 7, 81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 87, 1, 0, 0, 0, 13,
		93, 1, 0, 0, 0, 15, 95, 1, 0, 0, 0, 17, 100, 1, 0, 0, 0, 19, 102, 1, 0,
		0, 0, 21, 108, 1, 0, 0, 0, 23, 110, 1, 0, 0, 0, 25, 112, 1, 0, 0, 0, 27,
		114, 1, 0, 0, 0, 29, 122, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 130, 1,
		0, 0, 0, 35, 132, 1, 0, 0, 0, 37, 134, 1, 0, 0, 0, 39, 136, 1, 0, 0, 0,
		41, 139, 1, 0, 0, 0, 43, 141, 1, 0, 0, 0, 45, 144, 1, 0, 0, 0, 47, 146,
		1, 0, 0, 0, 49, 150, 1, 0, 0, 0, 51, 154, 1, 0, 0, 0, 53, 161, 1, 0, 0,
		0, 55, 165, 1, 0, 0, 0, 57, 170, 1, 0, 0, 0, 59, 63, 5, 39, 0, 0, 60, 62,
		3, 3, 1, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 67, 5,
		39, 0, 0, 67, 2, 1, 0, 0, 0, 68, 73, 8, 0, 0, 0, 69, 70, 5, 92, 0, 0, 70,
		73, 3, 5, 2, 0, 71, 73, 3, 19, 9, 0, 72, 68, 1, 0, 0, 0, 72, 69, 1, 0,
		0, 0, 72, 71, 1, 0, 0, 0, 73, 4, 1, 0, 0, 0, 74, 78, 3, 7, 3, 0, 75, 78,
		3, 9, 4, 0, 76, 78, 3, 11, 5, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		77, 76, 1, 0, 0, 0, 78, 6, 1, 0, 0, 0, 79, 82, 3, 13, 6, 0, 80, 82, 3,
		15, 7, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 8, 1, 0, 0, 0, 83,
		84, 5, 120, 0, 0, 84, 85, 3, 25, 12, 0, 85, 86, 3, 25, 12, 0, 86, 10, 1,
		0, 0, 0, 87, 88, 5, 117, 0, 0, 88, 89, 3, 25, 12, 0, 89, 90, 3, 25, 12,
		0, 90, 91, 3, 25, 12, 0, 91, 92, 3, 25, 12, 0, 92, 12, 1, 0, 0, 0, 93,
		94, 7, 1, 0, 0, 94, 14, 1, 0, 0, 0, 95, 96, 8, 2, 0, 0, 96, 16, 1, 0, 0,
		0, 97, 101, 3, 13, 6, 0, 98, 101, 3, 23, 11, 0, 99, 101, 7, 3, 0, 0, 100,
		97, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 18, 1, 0,
		0, 0, 102, 103, 5, 92, 0, 0, 103, 104, 3, 21, 10, 0, 104, 20, 1, 0, 0,
		0, 105, 106, 5, 13, 0, 0, 106, 109, 5, 10, 0, 0, 107, 109, 3, 55, 27, 0,
		108, 105, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 22, 1, 0, 0, 0, 110, 111,
		7, 4, 0, 0, 111, 24, 1, 0, 0, 0, 112, 113, 7, 5, 0, 0, 113, 26, 1, 0, 0,
		0, 114, 115, 7, 6, 0, 0, 115, 28, 1, 0, 0, 0, 116, 117, 5, 65, 0, 0, 117,
		118, 5, 78, 0, 0, 118, 123, 5, 68, 0, 0, 119, 120, 5, 97, 0, 0, 120, 121,
		5, 110, 0, 0, 121, 123, 5, 100, 0, 0, 122, 116, 1, 0, 0, 0, 122, 119, 1,
		0, 0, 0, 123, 30, 1, 0, 0, 0, 124, 125, 5, 79, 0, 0, 125, 129, 5, 82, 0,
		0, 126, 127, 5, 111, 0, 0, 127, 129, 5, 114, 0, 0, 128, 124, 1, 0, 0, 0,
		128, 126, 1, 0, 0, 0, 129, 32, 1, 0, 0, 0, 130, 131, 5, 40, 0, 0, 131,
		34, 1, 0, 0, 0, 132, 133, 5, 41, 0, 0, 133, 36, 1, 0, 0, 0, 134, 135, 5,
		61, 0, 0, 135, 38, 1, 0, 0, 0, 136, 137, 5, 33, 0, 0, 137, 138, 5, 61,
		0, 0, 138, 40, 1, 0, 0, 0, 139, 140, 5, 62, 0, 0, 140, 42, 1, 0, 0, 0,
		141, 142, 5, 62, 0, 0, 142, 143, 5, 61, 0, 0, 143, 44, 1, 0, 0, 0, 144,
		145, 5, 60, 0, 0, 145, 46, 1, 0, 0, 0, 146, 147, 5, 60, 0, 0, 147, 148,
		5, 61, 0, 0, 148, 48, 1, 0, 0, 0, 149, 151, 7, 7, 0, 0, 150, 149, 1, 0,
		0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 50, 1, 0, 0, 0, 154, 156, 5, 33, 0, 0, 155, 157, 7, 7, 0, 0, 156,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 52, 1, 0, 0, 0, 160, 162, 8, 8, 0, 0, 161, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		164, 54, 1, 0, 0, 0, 165, 166, 7, 9, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		6, 27, 0, 0, 168, 56, 1, 0, 0, 0, 169, 171, 7, 10, 0, 0, 170, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 175, 6, 28, 1, 0, 175, 58, 1, 0, 0, 0, 13, 0,
		63, 72, 77, 81, 100, 108, 122, 128, 152, 158, 163, 172, 2, 0, 1, 0, 6,
		0, 0,
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

// QueryLexerInit initializes any static state used to implement QueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLexerInit() {
	staticData := &QueryLexerLexerStaticData
	staticData.once.Do(querylexerLexerInit)
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	QueryLexerInit()
	l := new(QueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &QueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerSTRING         = 1
	QueryLexerAND            = 2
	QueryLexerOR             = 3
	QueryLexerLPAREN         = 4
	QueryLexerRPAREN         = 5
	QueryLexerEQ             = 6
	QueryLexerNE             = 7
	QueryLexerGT             = 8
	QueryLexerGTE            = 9
	QueryLexerLT             = 10
	QueryLexerLTE            = 11
	QueryLexerIDENTIFIER     = 12
	QueryLexerNEG_IDENTIFIER = 13
	QueryLexerENCODED_STRING = 14
	QueryLexerLineTerminator = 15
	QueryLexerWS             = 16
)
