%{

//TODO Put your favorite license here
		
// yacc source generated by ebnf2y[1]
// at 2014-07-25 14:37:06.973454437 +0200 CEST
//
//  $ ebnf2y -o nquads.y -pkg parser parser.ebnf
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
// 
//   [1]: http://modernc.org/ebnf2y

package parser //TODO real package name

//TODO required only be the demo _dump function
import (
	"bytes"
	"fmt"
	"strings"

	"github.com/skycoin/cx/cmd/goyacc/util"
)

%}

%union {
	item interface{} //TODO insert real field(s)
}

%token	DACCENT
%token	DOT
%token	EOL
%token	IRIREF
%token	LABEL
%token	LANGTAG
%token	STRING

%type	<item> 	/*TODO real type(s), if/where applicable */
	DACCENT
	DOT
	EOL
	IRIREF
	LABEL
	LANGTAG
	STRING

%type	<item> 	/*TODO real type(s), if/where applicable */
	GraphLabel
	Literal
	Literal1
	Literal11
	Object
	Predicate
	SourceFile
	SourceFile1
	SourceFile2
	SourceFile3
	Start
	Statement
	Statement1
	Subject

/*TODO %left, %right, ... declarations */

%start Start

%%

GraphLabel:
	IRIREF
	{
		$$ = $1 //TODO 1
	}
|	LABEL
	{
		$$ = $1 //TODO 2
	}

Literal:
	STRING Literal1
	{
		$$ = []Literal{$1, $2} //TODO 3
	}

Literal1:
	/* EMPTY */
	{
		$$ = nil //TODO 4
	}
|	Literal11
	{
		$$ = $1 //TODO 5
	}

Literal11:
	DACCENT IRIREF
	{
		$$ = []Literal11{$1, $2} //TODO 6
	}
|	LANGTAG
	{
		$$ = $1 //TODO 7
	}

Object:
	IRIREF
	{
		$$ = $1 //TODO 8
	}
|	LABEL
	{
		$$ = $1 //TODO 9
	}
|	Literal
	{
		$$ = $1 //TODO 10
	}

Predicate:
	IRIREF
	{
		$$ = $1 //TODO 11
	}

SourceFile:
	SourceFile1 SourceFile2 SourceFile3
	{
		$$ = []SourceFile{$1, $2, $3} //TODO 12
	}

SourceFile1:
	/* EMPTY */
	{
		$$ = nil //TODO 13
	}
|	Statement
	{
		$$ = $1 //TODO 14
	}

SourceFile2:
	/* EMPTY */
	{
		$$ = []SourceFile2(nil) //TODO 15
	}
|	SourceFile2 EOL Statement
	{
		$$ = append($1.([]SourceFile2), $2, $3) //TODO 16
	}

SourceFile3:
	/* EMPTY */
	{
		$$ = nil //TODO 17
	}
|	EOL
	{
		$$ = $1 //TODO 18
	}

Start:
	SourceFile
	{
		_parserResult = $1 //TODO 19
	}

Statement:
	Subject Predicate Object Statement1 DOT
	{
		$$ = []Statement{$1, $2, $3, $4, $5} //TODO 20
	}

Statement1:
	/* EMPTY */
	{
		$$ = nil //TODO 21
	}
|	GraphLabel
	{
		$$ = $1 //TODO 22
	}

Subject:
	IRIREF
	{
		$$ = $1 //TODO 23
	}
|	LABEL
	{
		$$ = $1 //TODO 24
	}

%%

//TODO remove demo stuff below

var _parserResult interface{}

type (
	GraphLabel interface{}
	Literal interface{}
	Literal1 interface{}
	Literal11 interface{}
	Object interface{}
	Predicate interface{}
	SourceFile interface{}
	SourceFile1 interface{}
	SourceFile2 interface{}
	SourceFile3 interface{}
	Start interface{}
	Statement interface{}
	Statement1 interface{}
	Subject interface{}
)
	
func _dump() {
	s := fmt.Sprintf("%#v", _parserResult)
	s = strings.Replace(s, "%", "%%", -1)
	s = strings.Replace(s, "{", "{%i\n", -1)
	s = strings.Replace(s, "}", "%u\n}", -1)
	s = strings.Replace(s, ", ", ",\n", -1)
	var buf bytes.Buffer
	util.IndentFormatter(&buf, ". ").Format(s)
	buf.WriteString("\n")
	a := strings.Split(buf.String(), "\n")
	for _, v := range a {
		if strings.HasSuffix(v, "(nil)") || strings.HasSuffix(v, "(nil),") {
			continue
		}
	
		fmt.Println(v)
	}
}

// End of demo stuff
