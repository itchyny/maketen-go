package maketen

// Solve maketen.
func Solve(x, y, z, w *Num) []Expr {
	var exprs []Expr
	res := make(map[string]struct{})
	ten := NewInt(10)
	for _, a := range operators {
		for _, b := range operators {
			for _, c := range operators {
				for _, t := range trees(x, y, z, w, a, b, c) {
					if e := Eval(t); e != nil && e.Cmp(ten) == 0 {
						s := t.String()
						if _, ok := res[s]; !ok {
							res[s] = struct{}{}
							exprs = append(exprs, t)
						}
					}
				}
			}
		}
	}
	return exprs
}

func trees(x, y, z, w *Num, a, b, c Operator) []Expr {
	return []Expr{
		&BinOp{a, &BinOp{b, &BinOp{c, x, y}, z}, w},
		&BinOp{a, &BinOp{b, x, &BinOp{c, y, z}}, w},
		&BinOp{a, &BinOp{b, x, y}, &BinOp{c, z, w}},
		&BinOp{a, x, &BinOp{b, &BinOp{c, y, z}, w}},
		&BinOp{a, x, &BinOp{b, y, &BinOp{c, z, w}}},
	}
}
