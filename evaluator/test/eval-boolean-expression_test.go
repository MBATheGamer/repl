package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

type BooleanExpressionType struct {
	input    string
	expected bool
}

func TestEvalBooleanExpression(t *testing.T) {
	var tests = []BooleanExpressionType{
		{"true", true},
		{"false", false},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	var result, ok = obj.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}
