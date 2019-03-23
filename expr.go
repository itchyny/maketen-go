package maketen

// Expr ...
type Expr interface{}

// BinOp ...
type BinOp struct {
	op       Operator
	lhs, rhs Expr
}

// Eval expression.
func Eval(e Expr) Num {
	switch e := e.(type) {
	case *BinOp:
		return e.op.Apply(Eval(e.lhs), Eval(e.rhs))
	case Num:
		return e
	default:
		panic(e)
	}
}
