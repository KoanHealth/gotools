package fixgo

type TernaryArgument interface {
	int | int64 | string
}

func Ternary[V TernaryArgument](c bool, v1, v2 V) V {
	if c {
		return v1
	} else {
		return v2
	}
}
