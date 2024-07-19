package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/object"
	"github.com/MBATheGamer/lang_core/parser"
	"github.com/MBATheGamer/repl/evaluator"
)

func testEval(input string) object.Object {
	var environment = object.NewEnivronment()

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()

	return evaluator.Eval(program, environment)
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != evaluator.NULL {
		t.Errorf(
			"object is not NULL. got=%T (%+v)",
			obj,
			obj,
		)
		return false
	}

	return true
}
