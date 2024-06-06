package console

import (
	"fmt"
)

func Printf(format string, a ...any) {
	_, _ = fmt.Printf(format, a...)
}
