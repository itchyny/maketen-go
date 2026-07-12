package maketen

// Operator is a binary arithmetic operator, identified by its symbol.
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

// prec reports the precedence of op: 2 for '*' and '/', 1 for '+' and '-'.
func (op Operator) prec() int {
	if op.isOneOf('*', '/') {
		return 2
	}
	return 1
}

var zero = NewInt(0)

var operators = []Operator{
	{'+', func(l, r *Num) *Num {
		return NewNum().Add(l, r)
	}},
	{'-', func(l, r *Num) *Num {
		return NewNum().Sub(l, r)
	}},
	{'*', func(l, r *Num) *Num {
		return NewNum().Mul(l, r)
	}},
	{'/', func(l, r *Num) *Num {
		if r.Cmp(zero) == 0 {
			return nil
		}
		return NewNum().Quo(l, r)
	}},
}
