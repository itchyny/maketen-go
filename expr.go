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
	lparen := prec(bo.lhs) < prec(bo)
	rparen := prec(bo.rhs) < prec(bo) ||
		prec(bo.rhs) == prec(bo) && bo.op.isOneOf('-', '/')
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

// prec reports the precedence of e's root operator, or 3 for an atom.
func prec(e Expr) int {
	if bo, ok := e.(*BinOp); ok {
		return bo.op.prec()
	}
	return 3
}
