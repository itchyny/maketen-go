package maketen

// Operator ...
type Operator struct {
	str   rune
	Apply func(*Num, *Num) *Num
}

// String implements Stringer.
func (op Operator) String() string {
	return string(op.str)
}

func (op Operator) isOneOf(cs ...rune) bool {
	for _, r := range cs {
		if op.str == r {
			return true
		}
	}
	return false
}

var zero = NewZero()

var operators = []Operator{
	{'+', func(l, r *Num) *Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Add(l, r)
	}},
	{'-', func(l, r *Num) *Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Sub(l, r)
	}},
	{'*', func(l, r *Num) *Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Mul(l, r)
	}},
	{'/', func(l, r *Num) *Num {
		if l == nil || r == nil {
			return nil
		}
		if r.Cmp(zero) == 0 {
			return nil
		}
		return NewExpr().Quo(l, r)
	}},
}
