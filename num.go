package maketen

import "math/big"

// Num ...
type Num big.Rat

// Add ...
func (n *Num) Add(l, r *Num) *Num {
	return (*Num)(((*big.Rat)(n)).Add((*big.Rat)(l), (*big.Rat)(r)))
}

// Sub ...
func (n *Num) Sub(l, r *Num) *Num {
	return (*Num)(((*big.Rat)(n)).Sub((*big.Rat)(l), (*big.Rat)(r)))
}

// Mul ...
func (n *Num) Mul(l, r *Num) *Num {
	return (*Num)(((*big.Rat)(n)).Mul((*big.Rat)(l), (*big.Rat)(r)))
}

// Quo ...
func (n *Num) Quo(l, r *Num) *Num {
	return (*Num)(((*big.Rat)(n)).Quo((*big.Rat)(l), (*big.Rat)(r)))
}

// Cmp ...
func (n *Num) Cmp(m *Num) int {
	return ((*big.Rat)(n)).Cmp((*big.Rat)(m))
}

// String implements Stringer.
func (n *Num) String() string {
	return ((*big.Rat)(n)).String()
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
