package maketen

import "strings"

// Expr ...
type Expr interface {
	String() string
	isExpr()
}

// BinOp ...
type BinOp struct {
	op       Operator
	lhs, rhs Expr
}

func (*BinOp) isExpr() {}

// String implements Stringer.
func (bo *BinOp) String() string {
	lparen := bo.op.isOneOf('*', '/') && isAddOrSub(bo.lhs)
	rparen := bo.op.isOneOf('-', '*') && isAddOrSub(bo.rhs) || bo.op.isOneOf('/') && isBinOp(bo.rhs)
	var s strings.Builder
	if lparen {
		s.WriteByte('(')
	}
	s.WriteString(bo.lhs.String())
	if lparen {
		s.WriteByte(')')
	}
	s.WriteByte(' ')
	s.WriteByte(bo.op.symbol)
	s.WriteByte(' ')
	if rparen {
		s.WriteByte('(')
	}
	s.WriteString(bo.rhs.String())
	if rparen {
		s.WriteByte(')')
	}
	return s.String()
}

func isAddOrSub(e Expr) bool {
	switch e := e.(type) {
	case *BinOp:
		return e.op.isOneOf('+', '-')
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
func Eval(e Expr) *Num {
	switch e := e.(type) {
	case *BinOp:
		return e.op.apply(Eval(e.lhs), Eval(e.rhs))
	case *Num:
		return e
	default:
		panic(e)
	}
}
