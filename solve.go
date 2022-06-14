package maketen

var ten = NewInt(10)

// Solve maketen.
func Solve(ns ...*Num) []Expr {
	switch len(ns) {
	case 0:
		return nil
	case 1:
		if ns[0].Cmp(ten) != 0 {
			return nil
		}
		return []Expr{ns[0]}
	}
	var exprs []Expr
	res := make(map[string]struct{})
	oiter := newOpsIter(len(ns) - 1)
	for ops, ok := oiter.next(); ok; ops, ok = oiter.next() {
		eiter := newExprIter(ns, ops)
		for t, ok := eiter.next(); ok; t, ok = eiter.next() {
			if e := Eval(t); e != nil && e.Cmp(ten) == 0 {
				s := t.String()
				if _, ok := res[s]; !ok {
					res[s] = struct{}{}
					exprs = append(exprs, t)
				}
			}
		}
	}
	return exprs
}

func newOpsIter(n int) opsIter {
	ops := make([]Operator, n)
	for i := 0; i < n-1; i++ {
		ops[i] = operators[0]
	}
	return opsIter(ops)
}

type opsIter []Operator

func (ops opsIter) next() ([]Operator, bool) {
	var ok bool
	for i := len(ops) - 1; i >= 0; i-- {
		if ops[i], ok = nextOp(ops[i]); ok {
			return ops, true
		}
	}
	return nil, false
}

func nextOp(op Operator) (Operator, bool) {
	switch op.symbol {
	case '\x00':
		return operators[0], true
	case '+':
		return operators[1], true
	case '-':
		return operators[2], true
	case '*':
		return operators[3], true
	default:
		return operators[0], false
	}
}

func newExprIter(ns []*Num, ops []Operator) exprIter {
	switch len(ns) {
	case 1:
		return &unitIter{expr: ns[0]}
	case 2:
		return &unitIter{expr: &BinOp{ops[0], ns[0], ns[1]}}
	default:
		return &treeIter{ns: ns, ops: ops}
	}
}

type exprIter interface {
	next() (Expr, bool)
}

type unitIter struct {
	expr Expr
	done bool
}

func (iter *unitIter) next() (Expr, bool) {
	if iter.done {
		return nil, false
	}
	iter.done = true
	return iter.expr, true
}

type treeIter struct {
	l, r exprIter
	le   Expr
	ns   []*Num
	ops  []Operator
	i    int
}

func (iter *treeIter) next() (Expr, bool) {
	if iter.l == nil {
		if iter.i++; iter.i >= len(iter.ns) {
			return nil, false
		}
		iter.l = newExprIter(iter.ns[:iter.i], iter.ops[:iter.i-1])
	}
	if iter.le == nil {
		var ok bool
		if iter.le, ok = iter.l.next(); !ok {
			iter.l = nil
			return iter.next()
		}
	}
	if iter.r == nil {
		iter.r = newExprIter(iter.ns[iter.i:], iter.ops[iter.i:])
	}
	if r, ok := iter.r.next(); ok {
		return &BinOp{iter.ops[iter.i-1], iter.le, r}, true
	}
	iter.r = nil
	iter.le = nil
	return iter.next()
}
