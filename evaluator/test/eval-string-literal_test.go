package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

func TestEvalStringLiteral(t *testing.T) {
	var input = `"Hello, World!"`

	var evaluated = testEval(input)
	var str, ok = evaluated.(*object.String)

	if !ok {
		t.Fatalf(
			"object is not String. got=%T (%+v)",
			evaluated,
			evaluated,
		)
	}

	if str.Value != "Hello, World!" {
		t.Errorf(
			"String has wrong value. got=%q",
			str.Value,
		)
	}
}
