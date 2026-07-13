package maketen

import "iter"

var ten = NewNum(10)

// Solve maketen.
func Solve(ns ...*Num) iter.Seq[Expr] {
	return func(yield func(Expr) bool) {
		for e, v := range exprSeq(ns) {
			if v.Cmp(ten) == 0 && !yield(e) {
				return
			}
		}
	}
}

// exprSeq yields every valid expression over ns together with its value.
func exprSeq(ns []*Num) iter.Seq2[Expr, *Num] {
	return func(yield func(Expr, *Num) bool) {
		if len(ns) == 1 {
			yield(ns[0], ns[0])
			return
		}
		for i := 1; i < len(ns); i++ {
			for l, lv := range exprSeq(ns[:i]) {
				for r, rv := range exprSeq(ns[i:]) {
					for _, op := range operators {
						// Skip re-associable duplicates: for the associative
						// operators, a same-precedence right operand builds a
						// tree that renders identically to a left-leaning one.
						if op.isOneOf('+', '*') && prec(r) == op.prec() {
							continue
						}
						// Skip invalid subexpressions (division by zero); no
						// expression built on them can be a solution.
						v := op.apply(lv, rv)
						if v == nil {
							continue
						}
						if !yield(&BinOp{op, l, r}, v) {
							return
						}
					}
				}
			}
		}
	}
}
