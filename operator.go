package maketen

type operator struct {
	symbol byte
	apply  func(*Num, *Num) *Num
}

func (op operator) isOneOf(bs ...byte) bool {
	for _, b := range bs {
		if op.symbol == b {
			return true
		}
	}
	return false
}

func (op operator) prec() int {
	if op.isOneOf('*', '/') {
		return 2
	}
	return 1
}

var zero = NewNum(0)

var operators = []operator{
	{'+', func(l, r *Num) *Num {
		return NewNum(0).add(l, r)
	}},
	{'-', func(l, r *Num) *Num {
		return NewNum(0).sub(l, r)
	}},
	{'*', func(l, r *Num) *Num {
		return NewNum(0).mul(l, r)
	}},
	{'/', func(l, r *Num) *Num {
		if r.cmp(zero) == 0 {
			return nil
		}
		return NewNum(0).quo(l, r)
	}},
}
