package maketen

import "strings"

// Expr ...
type Expr interface {
	String() string
}

// BinOp ...
type BinOp struct {
	op       Operator
	lhs, rhs Expr
}

// String implements Stringer.
func (bo *BinOp) String() string {
	lparen := bo.op.isMulOrDiv() && isAddOrSub(bo.lhs)
	rparen := bo.op.isSubOrMul() && isAddOrSub(bo.rhs) || bo.op.str == '/' && isBinOp(bo.rhs)
	s := new(strings.Builder)
	if lparen {
		s.WriteRune('(')
	}
	s.WriteString(bo.lhs.String())
	if lparen {
		s.WriteRune(')')
	}
	s.WriteRune(' ')
	s.WriteString(bo.op.String())
	s.WriteRune(' ')
	if rparen {
		s.WriteRune('(')
	}
	s.WriteString(bo.rhs.String())
	if rparen {
		s.WriteRune(')')
	}
	return s.String()
}

func isAddOrSub(e Expr) bool {
	switch e := e.(type) {
	case *BinOp:
		return e.op.isAddOrSub()
	default:
		return false
	}
}

func isBinOp(e Expr) bool {
	switch e.(type) {
	case *BinOp:
		return true
	default:
		return false
	}
}

// Eval expression.
func Eval(e Expr) Num {
	switch e := e.(type) {
	case *BinOp:
		return e.op.Apply(Eval(e.lhs), Eval(e.rhs))
	case Num:
		return e
	default:
		panic(e)
	}
}
