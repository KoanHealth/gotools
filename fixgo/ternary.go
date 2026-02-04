package fixgo

type TernaryArgument interface{}

func Ternary[V TernaryArgument](c bool, v1, v2 V) V {
	if c {
		return v1
	} else {
		return v2
	}
}
