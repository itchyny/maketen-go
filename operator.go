package maketen

// Operator ...
type Operator struct {
	str   rune
	Apply func(Num, Num) Num
}

// String implements Stringer.
func (op Operator) String() string {
	return string(op.str)
}

func (op Operator) isAddOrSub() bool {
	return op.str == '+' || op.str == '-'
}

func (op Operator) isMulOrDiv() bool {
	return op.str == '*' || op.str == '/'
}

func (op Operator) isSubOrMul() bool {
	return op.str == '-' || op.str == '*'
}

var zero = NewZero()

var operators = []Operator{
	Operator{'+', func(l, r Num) Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Add(l, r)
	}},
	Operator{'-', func(l, r Num) Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Sub(l, r)
	}},
	Operator{'*', func(l, r Num) Num {
		if l == nil || r == nil {
			return nil
		}
		return NewExpr().Mul(l, r)
	}},
	Operator{'/', func(l, r Num) Num {
		if l == nil || r == nil {
			return nil
		}
		if r.Cmp(zero) == 0 {
			return nil
		}
		return NewExpr().Quo(l, r)
	}},
}
