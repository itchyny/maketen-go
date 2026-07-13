package maketen

import (
	"math/big"
	"strings"
)

// Expr is an arithmetic expression, either a Num or a BinOp.
type Expr interface {
	String() string
	prec() int
}

// Num is a number represented as an exact rational.
type Num big.Rat

// NewNum returns a new Num with the integer value i.
func NewNum(i int) *Num {
	return (*Num)(big.NewRat(int64(i), 1))
}

func (*Num) prec() int {
	return 3
}

func (n *Num) rat() *big.Rat {
	return (*big.Rat)(n)
}

func (n *Num) add(l, r *Num) *Num {
	return (*Num)(n.rat().Add(l.rat(), r.rat()))
}

func (n *Num) sub(l, r *Num) *Num {
	return (*Num)(n.rat().Sub(l.rat(), r.rat()))
}

func (n *Num) mul(l, r *Num) *Num {
	return (*Num)(n.rat().Mul(l.rat(), r.rat()))
}

func (n *Num) quo(l, r *Num) *Num {
	return (*Num)(n.rat().Quo(l.rat(), r.rat()))
}

func (n *Num) cmp(m *Num) int {
	return n.rat().Cmp(m.rat())
}

var one = big.NewInt(1)

// String implements Stringer.
func (n *Num) String() string {
	r := n.rat()
	if r.Denom().Cmp(one) == 0 {
		return r.Num().String()
	}
	return r.String()
}

// BinOp is a binary operation applying an operator to two sub-expressions.
type BinOp struct {
	op       operator
	lhs, rhs Expr
}

func (bo *BinOp) prec() int {
	return bo.op.prec()
}

// String implements Stringer.
func (bo *BinOp) String() string {
	lparen := bo.lhs.prec() < bo.prec()
	rparen := bo.rhs.prec() < bo.prec() ||
		bo.rhs.prec() == bo.prec() && bo.op.isOneOf('-', '/')
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
