package maketen

import "iter"

var ten = NewInt(10)

// Solve maketen.
func Solve(ns ...*Num) iter.Seq[Expr] {
	return func(yield func(Expr) bool) {
		switch len(ns) {
		case 0:
			return
		case 1:
			if ns[0].Cmp(ten) == 0 {
				yield(ns[0])
			}
			return
		}
		seen := make(map[string]struct{})
		for t := range exprSeq(ns) {
			if e := Eval(t); e != nil && e.Cmp(ten) == 0 {
				s := t.String()
				if _, ok := seen[s]; !ok {
					seen[s] = struct{}{}
					if !yield(t) {
						return
					}
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
						if !yield(&BinOp{op, l, r}) {
							return
						}
					}
				}
			}
		}
	}
}
