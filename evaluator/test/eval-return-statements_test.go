package evaluator_test

import "testing"

func TestEvalReturnStatements(t *testing.T) {
	var tests = []IntegerExpressionType{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}
