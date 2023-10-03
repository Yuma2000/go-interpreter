package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// Eval 関数は、ASTノードを評価して、その結果のオブジェクトを返す．
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// 文
	case *ast.Program:
		return evalStatements(node.Statements)
		
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	
	// 式
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}

// nativeBoolToBooleanObject 関数は、Goのbool値をMonkeyのBooleanオブジェクトに変換する．
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}
