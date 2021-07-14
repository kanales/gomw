package middlewares

import (
	"gomw/ansi"
	"log"
	"net/http"
	"os"
	"strconv"
)

type logger struct {
	http.ResponseWriter
	code int
	l    *log.Logger
}

func (l *logger) WriteHeader(code int) {
	l.code = code
	l.ResponseWriter.WriteHeader(code)
}

func (l *logger) Log(rw http.ResponseWriter, r *http.Request) {
	fg := ansi.None
	switch l.code / 100 {
	case 1:
		fg = ansi.Blue
	case 2:
		fg = ansi.Green
	case 3:
		fg = ansi.Yellow
	case 4:
		fg = ansi.Magenta
	case 5:
		fg = ansi.Red
	}

	codefmt := ansi.NewFormatter().
		Mode(ansi.Bold).
		Foreground(fg).
		Apply(strconv.FormatInt(int64(l.code), 10))

	if l.code/100 == 2 {
		codefmt += " ✅"
	}

	if l.code == 401 {
		codefmt += " 🔒"
	}

	l.l.Printf("%s  \t%s\t%s", r.Method, r.URL.Path, codefmt)
}

func Logger() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			l := logger{rw, 200, log.New(os.Stdout, "", log.Ldate|log.Ltime)}
			h.ServeHTTP(&l, r)
			l.Log(rw, r)
		})
	}
}
