package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

var DumpFlag bool

func Dump(ic *Intcode) {
	var b strings.Builder

	for i, m := range ic.mem {
		if i % 20 == 0 {
			b.WriteByte('|')
		} else {
			b.WriteByte(' ')
		}
		if i == ic.pc {
			b.WriteByte('*')
			b.WriteString(strconv.Itoa(m))
			b.WriteByte('*')
		} else {
			b.WriteString(strconv.Itoa(m))
		}
	}
	_, _ = fmt.Fprintln(ic.Dump, b.String())
}