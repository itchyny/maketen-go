package maketen

import "math/big"

// Num is a number represented as an exact rational.
type Num big.Rat

// NewNum returns a new Num with the integer value i.
func NewNum(i int) *Num {
	return (*Num)(big.NewRat(int64(i), 1))
}

func (*Num) isExpr() {}

func (n *Num) rat() *big.Rat {
	return (*big.Rat)(n)
}

// Add sets n to the sum l+r and returns n.
func (n *Num) Add(l, r *Num) *Num {
	return (*Num)(n.rat().Add(l.rat(), r.rat()))
}

// Sub sets n to the difference l-r and returns n.
func (n *Num) Sub(l, r *Num) *Num {
	return (*Num)(n.rat().Sub(l.rat(), r.rat()))
}

// Mul sets n to the product l*r and returns n.
func (n *Num) Mul(l, r *Num) *Num {
	return (*Num)(n.rat().Mul(l.rat(), r.rat()))
}

// Quo sets n to the quotient l/r and returns n. It panics if r is zero.
func (n *Num) Quo(l, r *Num) *Num {
	return (*Num)(n.rat().Quo(l.rat(), r.rat()))
}

// Cmp compares n and m, returning -1 if n < m, 0 if n == m, and +1 if n > m.
func (n *Num) Cmp(m *Num) int {
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
