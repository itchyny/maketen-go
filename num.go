package maketen

import "math/big"

// Num ...
type Num = *big.Rat // can this be an interface?

// NewExpr  ...
func NewExpr() Num {
	return new(big.Rat)
}

// NewZero ...
func NewZero() Num {
	return big.NewRat(0, 1)
}

// NewInt ...
func NewInt(i int) Num {
	return big.NewRat(int64(i), 1)
}
