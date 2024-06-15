package evaluator_test

import "testing"

func TestEvalIntegerExpression(t *testing.T) {
	var tests = []TestIntegerExpression{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}
