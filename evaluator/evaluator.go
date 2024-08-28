package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, environment *object.Enivronment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, environment)

	case *ast.LetStatement:
		var value = Eval(node.Value, environment)
		if isError(value) {
			return value
		}
		environment.Set(node.Name.Value, value)

	case *ast.BlockStatement:
		return evalBlockStatement(node, environment)

	case *ast.IfExpression:
		return evalIfExpression(node, environment)

	case *ast.ReturnStatement:
		var value = Eval(node.ReturnValue, environment)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}

	case *ast.ExpressionStatement:
		return Eval(node.Expression, environment)

	case *ast.PrefixExpression:
		var right = Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		var left = Eval(node.Left, environment)
		if isError(left) {
			return left
		}
		var right = Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.Identifier:
		return evalIdentifier(node, environment)

	case *ast.FunctionLiteral:
		var parameters = node.Parameters
		var body = node.Body
		return &object.Function{
			Parameters:  parameters,
			Enivronment: environment,
			Body:        body,
		}

	case *ast.CallExpression:
		var function = Eval(node.Function, environment)
		if isError(function) {
			return function
		}
		var arguments = evalExpressions(node.Arguments, environment)
		if len(arguments) == 1 && isError(arguments[0]) {
			return arguments[0]
		}
		return applyFunction(function, arguments)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.ArrayLiteral:
		var elements = evalExpressions(node.Elements, environment)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.IndexExpression:
		var left = Eval(node.Left, environment)
		if isError(left) {
			return left
		}
		var index = Eval(node.Index, environment)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *ast.HashLiteral:
		return evalHashLiteral(node, environment)

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}
