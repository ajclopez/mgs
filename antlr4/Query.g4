grammar Query;

input
	: query EOF
	;

query
	: left=query logicalOp=(AND | OR) right=query #opQuery
	| LPAREN query RPAREN #priorityQuery
	| criteria #atomQuery
	;

criteria
	: key op value
	| key
	;

key
	: IDENTIFIER
	| NEG_IDENTIFIER
	;

value
	: IDENTIFIER
	| STRING
	| ENCODED_STRING
	;

op
	: EQ
	| NE
	| GT
	| GTE
	| LT
	| LTE
	;

STRING
	: '\'' StringCharacter* '\''
	;

fragment StringCharacter
	: ~["\\\r\n]
	| '\\' EscapeSequence
	| LineContinuation
	;

fragment EscapeSequence
	: CharacterEscapeSequence
    | HexEscapeSequence
    | UnicodeEscapeSequence
    ;

fragment CharacterEscapeSequence
	: SingleEscapeCharacter
	| NonEscapeCharacter
	;

fragment HexEscapeSequence
	: 'x' HexDigit HexDigit
	;

fragment UnicodeEscapeSequence
	: 'u' HexDigit HexDigit HexDigit HexDigit
	;

fragment SingleEscapeCharacter
	: ['"\\bfnrtv]
	;

fragment NonEscapeCharacter
	: ~['"\\bfnrtv0-9xu\r\n]
	;

fragment EscapeCharacter
	: SingleEscapeCharacter
	| DecimalDigit
	| [xu]
	;

fragment LineContinuation
	: '\\' LineTerminatorSequence
	;

fragment LineTerminatorSequence
	: '\r\n'
	| LineTerminator
	;

fragment DecimalDigit
	: [0-9]
	;

fragment HexDigit
	: [0-9a-fA-F]
	;

fragment OctalDigit
	: [0-7]
	;

AND
	: 'AND'
	| 'and'
	;

OR
	: 'OR'
	| 'or'
	;

LPAREN
	: '('
	;

RPAREN
	: ')'
	;

EQ
	: '='
	;

NE
	: '!='
	;

GT
	: '>'
	;

GTE
	: '>='
	;

LT
	: '<'
	;

LTE
	: '<='
	;

IDENTIFIER
	: [A-Za-z0-9.:_-]+
	;

NEG_IDENTIFIER
	: '!'[A-Za-z0-9.:_-]+
	;

ENCODED_STRING
	: ~([ \\[\]<>!=()])+
	;

LineTerminator
	: [\r\n\u2028\u2029] -> channel(HIDDEN)
	;

WS
	: [ \t\r\n]+ -> skip
	;