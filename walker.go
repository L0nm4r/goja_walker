package gojawalker

import (
	"errors"
	"log"

	"github.com/dop251/goja/ast"
)

// visitor interface
type IVisitor interface {
	Enter(n ast.Node) IVisitor
	Exit(n ast.Node)
}

// error list
var ErrNodeIsNil = errors.New("node is nil")
var ErrVisitorIsNil = errors.New("visitor cann't be nil")

func Walk(v IVisitor, n ast.Node) error {
	if n == nil {
		return ErrNodeIsNil
	}

	if v == nil {
		return ErrVisitorIsNil
	}

	// visit
	if v = v.Enter(n); v == nil {
		return errors.New("visitor Enter cann't return nil")
	}
	defer v.Exit(n)

	switch n := n.(type) {
	// Nodes
	case *ast.Program:
		if n.Body != nil {
			for i := 0; i < len(n.Body); i++ {
				Walk(v, n.Body[i])
			}
		}
	// statement
	case *ast.BadStatement:
		return nil
	case *ast.BlockStatement:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
	case *ast.BranchStatement:
		return nil
	case *ast.CaseStatement:
		Walk(v, n.Test)
		if n.Consequent != nil {
			for i := 0; i < len(n.Consequent); i++ {
				Walk(v, n.Consequent[i])
			}
		}
	case *ast.CatchStatement:
		Walk(v, n.Parameter)
		if n.Body != nil {
			Walk(v, n.Body)
		}
	case *ast.DebuggerStatement:
		return nil
	case *ast.DoWhileStatement:
		Walk(v, n.Test)
		Walk(v, n.Body)
	case *ast.EmptyStatement:
		return nil
	case *ast.ExpressionStatement:
		Walk(v, n.Expression)
	case *ast.ForInStatement:
		Walk(v, n.Into)
		Walk(v, n.Source)
		Walk(v, n.Body)
	case *ast.ForOfStatement:
		Walk(v, n.Into)
		Walk(v, n.Source)
		Walk(v, n.Body)
	case *ast.ForStatement:
		Walk(v, n.Initializer)
		Walk(v, n.Test)
		Walk(v, n.Update)
		Walk(v, n.Body)
	case *ast.IfStatement:
		Walk(v, n.Test)
		Walk(v, n.Consequent)
		Walk(v, n.Alternate)
	case *ast.LabelledStatement:
		Walk(v, n.Label)
		Walk(v, n.Statement)
	case *ast.ReturnStatement:
		Walk(v, n.Argument)
	case *ast.SwitchStatement:
		Walk(v, n.Discriminant)
		if n.Body != nil {
			for i := 0; i < len(n.Body); i++ {
				Walk(v, n.Body[i])
			}
		}
	case *ast.ThrowStatement:
		Walk(v, n.Argument)
	case *ast.TryStatement:
		Walk(v, n.Body)
		Walk(v, n.Catch)
		Walk(v, n.Finally)
	case *ast.VariableStatement:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
	case *ast.LexicalDeclaration:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
	case *ast.WhileStatement:
		Walk(v, n.Test)
		Walk(v, n.Body)
	case *ast.WithStatement:
		Walk(v, n.Object)
		Walk(v, n.Body)
	case *ast.FunctionDeclaration:
		Walk(v, n.Function)
	case *ast.ClassDeclaration:
		Walk(v, n.Class)

	// expression
	case *ast.Binding:
		Walk(v, n.Target)
		Walk(v, n.Initializer)
	case *ast.YieldExpression:
		Walk(v, n.Argument)
	case *ast.AwaitExpression:
		Walk(v, n.Argument)
	case *ast.ArrayLiteral:
		if n.Value != nil {
			for i := 0; i < len(n.Value); i++ {
				Walk(v, n.Value[i])
			}
		}
	case *ast.ArrayPattern:
		if n.Elements != nil {
			for i := 0; i < len(n.Elements); i++ {
				Walk(v, n.Elements[i])
			}
		}
		Walk(v, n.Rest)
	case *ast.AssignExpression:
		Walk(v, n.Left)
		Walk(v, n.Right)
	case *ast.BadExpression:
		return nil
	case *ast.BinaryExpression:
		Walk(v, n.Left)
		Walk(v, n.Right)
	case *ast.BracketExpression:
		Walk(v, n.Left)
		Walk(v, n.Member)
	case *ast.CallExpression:
		Walk(v, n.Callee)
		if n.ArgumentList != nil {
			for i := 0; i < len(n.ArgumentList); i++ {
				Walk(v, n.ArgumentList[i])
			}
		}
	case *ast.ConditionalExpression:
		Walk(v, n.Test)
		Walk(v, n.Consequent)
		Walk(v, n.Alternate)
	case *ast.DotExpression:
		Walk(v, n.Left)
		// Walk(v, &n.Identifier)
	case *ast.PrivateDotExpression:
		Walk(v, n.Left)
	// case *ast.OptionalChain:
	// 	Walk(v, n.Expression)
	// case *ast.OptionalChain:
	// 	Walk(v, n.Expression)
	case *ast.FunctionLiteral:
		Walk(v, n.ParameterList)
		Walk(v, n.Body)
	case *ast.ClassLiteral:
		Walk(v, n.SuperClass)
		if n.Body != nil {
			for i := 0; i < len(n.Body); i++ {
				Walk(v, n.Body[i])
			}
		}
	case *ast.ExpressionBody:
		Walk(v, n.Expression)
	case *ast.ArrowFunctionLiteral:
		Walk(v, n.ParameterList)
		Walk(v, n.Body)
	case *ast.Identifier:
		return nil
	case *ast.PrivateIdentifier:
		return nil
	case *ast.NewExpression:
		Walk(v, n.Callee)
		if n.ArgumentList != nil {
			for i := 0; i < len(n.ArgumentList); i++ {
				Walk(v, n.ArgumentList[i])
			}
		}
	case *ast.NullLiteral:
		return nil
	case *ast.NumberLiteral:
		return nil
	case *ast.ObjectLiteral:
		if n.Value != nil {
			for i := 0; i < len(n.Value); i++ {
				Walk(v, n.Value[i])
			}
		}
	case *ast.ObjectPattern:
		if n.Properties != nil {
			for i := 0; i < len(n.Properties); i++ {
				Walk(v, n.Properties[i])
			}
		}
		Walk(v, n.Rest)
	case *ast.ParameterList:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
		Walk(v, n.Rest)
	case *ast.PropertyShort:
		Walk(v, n.Initializer)
	case *ast.PropertyKeyed:
		Walk(v, n.Key)
		Walk(v, n.Value)
	case *ast.SequenceExpression:
		if n.Sequence != nil {
			for i := 0; i < len(n.Sequence); i++ {
				Walk(v, n.Sequence[i])
			}
		}
	case *ast.StringLiteral:
		return nil
	case *ast.TemplateElement:
		return nil
	case *ast.TemplateLiteral:
		Walk(v, n.Tag)
		if n.Elements != nil {
			for i := 0; i < len(n.Elements); i++ {
				Walk(v, n.Elements[i])
			}
		}
		if n.Expressions != nil {
			for i := 0; i < len(n.Expressions); i++ {
				Walk(v, n.Expressions[i])
			}
		}
	case *ast.ThisExpression:
		return nil
	case *ast.SuperExpression:
		return nil
	case *ast.UnaryExpression:
		Walk(v, n.Operand)
	case *ast.MetaProperty:
		return nil

	// declaration
	case *ast.VariableDeclaration:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
	case *ast.FieldDefinition:
		Walk(v, n.Key)
		Walk(v, n.Initializer)
	case *ast.MethodDefinition:
		Walk(v, n.Key)
		Walk(v, n.Body)
	case *ast.ClassStaticBlock:
		Walk(v, n.Block)

	// for loop
	case *ast.ForLoopInitializerExpression:
		Walk(v, n.Expression)
	case *ast.ForLoopInitializerVarDeclList:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				Walk(v, n.List[i])
			}
		}
	case *ast.ForLoopInitializerLexicalDecl:
		Walk(v, &n.LexicalDeclaration)
	case *ast.ForIntoVar:
		Walk(v, n.Binding)
	case *ast.ForDeclaration:
		Walk(v, n.Target)
	case *ast.ForIntoExpression:
		Walk(v, n.Expression)

	// unexpeted
	default:
		log.Fatalf("unexpected: cann't determine node type: %s", n)
	}

	// 深度遍历
	return nil
}
