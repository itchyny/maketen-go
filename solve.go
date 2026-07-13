package maketen

import "iter"

var ten = NewNum(10)

// Solve maketen.
func Solve(ns ...*Num) iter.Seq[Expr] {
	return func(yield func(Expr) bool) {
		for t := range exprSeq(ns) {
			if e := Eval(t); e != nil && e.Cmp(ten) == 0 {
				if !yield(t) {
					return
				}
			}
		}
	}
}

func exprSeq(ns []*Num) iter.Seq[Expr] {
	return func(yield func(Expr) bool) {
		if len(ns) == 1 {
			yield(ns[0])
			return
		}
		for i := 1; i < len(ns); i++ {
			for l := range exprSeq(ns[:i]) {
				for r := range exprSeq(ns[i:]) {
					for _, op := range operators {
						// Skip re-associable duplicates: for the associative
						// operators, a same-precedence right operand builds a
						// tree that renders identically to a left-leaning one.
						if op.isOneOf('+', '*') && prec(r) == op.prec() {
							continue
						}
						if !yield(&BinOp{op, l, r}) {
							return
						}
					}
				}
			}
		}
	}
}
