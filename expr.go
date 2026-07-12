package maketen

import "strings"

// Expr is an arithmetic expression, either a Num or a BinOp.
type Expr interface {
	String() string
	isExpr()
}

// BinOp is a binary operation applying an Operator to two sub-expressions.
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
	bo, ok := e.(*BinOp)
	return ok && bo.op.isOneOf('+', '-')
}

func isBinOp(e Expr) bool {
	_, ok := e.(*BinOp)
	return ok
}

// Eval an expression.
func Eval(e Expr) *Num {
	switch e := e.(type) {
	case *BinOp:
		l, r := Eval(e.lhs), Eval(e.rhs)
		if l == nil || r == nil {
			return nil
		}
		return e.op.apply(l, r)
	case *Num:
		return e
	default:
		panic(e)
	}
}
