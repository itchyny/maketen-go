package maketen

import "math/big"

// Num ...
type Num big.Rat

func (*Num) isExpr() {}

func (n *Num) rat() *big.Rat {
	return (*big.Rat)(n)
}

// Add ...
func (n *Num) Add(l, r *Num) *Num {
	return (*Num)(n.rat().Add(l.rat(), r.rat()))
}

// Sub ...
func (n *Num) Sub(l, r *Num) *Num {
	return (*Num)(n.rat().Sub(l.rat(), r.rat()))
}

// Mul ...
func (n *Num) Mul(l, r *Num) *Num {
	return (*Num)(n.rat().Mul(l.rat(), r.rat()))
}

// Quo ...
func (n *Num) Quo(l, r *Num) *Num {
	return (*Num)(n.rat().Quo(l.rat(), r.rat()))
}

// Cmp ...
func (n *Num) Cmp(m *Num) int {
	return n.rat().Cmp(m.rat())
}

// String implements Stringer.
func (n *Num) String() string {
	r := n.rat()
	if r.Denom().Cmp(big.NewInt(1)) == 0 {
		return r.Num().String()
	}
	return r.String()
}

// NewExpr  ...
func NewExpr() *Num {
	return (*Num)(new(big.Rat))
}

// NewZero ...
func NewZero() *Num {
	return (*Num)(big.NewRat(0, 1))
}

// NewInt ...
func NewInt(i int) *Num {
	return (*Num)(big.NewRat(int64(i), 1))
}
