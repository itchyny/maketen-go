package maketen

// Operator ...
type Operator struct {
	symbol byte
	apply  func(*Num, *Num) *Num
}

// String implements Stringer.
func (op Operator) String() string {
	return string(op.symbol)
}

func (op Operator) isOneOf(bs ...byte) bool {
	for _, b := range bs {
		if op.symbol == b {
			return true
		}
	}
	return false
}

var zero = NewInt(0)

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
