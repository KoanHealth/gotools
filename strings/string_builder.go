package strings

import (
	"bytes"
	"fmt"
)

type StringBuilder struct {
	bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

func (sb *StringBuilder) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(sb, a...)
}

func (sb *StringBuilder) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(sb, format, a...)
}

func (sb *StringBuilder) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(sb, a...)
}
