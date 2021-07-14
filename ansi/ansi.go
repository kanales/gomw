package ansi

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	None = -1

	Bold          = 1
	Dim           = 2
	Italic        = 3
	Underline     = 4
	Blinking      = 5
	Inverse       = 7
	Invisible     = 8
	Striketrhough = 9

	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
)

type Formatter struct {
	fg    int
	bg    int
	modes []int
}

func NewFormatter() *Formatter {
	return &Formatter{
		-1,
		-1,
		make([]int, 0),
	}
}

func (fmtr *Formatter) Apply(s string) string {
	args := make([]string, 0)

	tostring := func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}

	for _, mode := range fmtr.modes {
		if mode != None {
			args = append(args, tostring(mode))
		}

	}

	if fmtr.fg != None {
		args = append(args, tostring(fmtr.fg))
	}

	if fmtr.bg != None {
		args = append(args, tostring(fmtr.bg))
	}

	argv := strings.Join(args, ";")
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", argv, s)
}

func (fmtr *Formatter) Applyf(s string, args ...interface{}) string {
	fs := fmt.Sprintf(s, args...)

	return fmtr.Apply(fs)
}

func (fmtr *Formatter) Foreground(code int) *Formatter {
	if 31 <= code || code <= 37 {
		fmtr.fg = code
	}

	return fmtr
}

func (fmtr *Formatter) Background(code int) *Formatter {
	if 31 <= code || code <= 37 {
		code += 10
		fmtr.bg = code
	}

	return fmtr
}

func (fmtr *Formatter) Mode(code int) *Formatter {
	if 0 <= code || code <= 9 {
		fmtr.modes = append(fmtr.modes, code)
	}
	return fmtr
}
