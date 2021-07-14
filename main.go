package main

import (
	"fmt"
	"gomw/ansi"
	"gomw/handlers"
	"gomw/middlewares"
	"gomw/routers"
	"log"
	"net/http"
	"os"
)

const url string = "0.0.0.0:8080"

func serverInit(server http.Handler) {
	fmtUrl := ansi.NewFormatter().
		Foreground(ansi.Blue).
		Mode(ansi.Underline).
		Mode(ansi.Bold).
		Applyf("http://%s/", url)
	base := fmt.Sprintf("🚀 Starting server at %s", fmtUrl)
	fmt.Fprintln(os.Stderr, base)
	fmt.Fprint(os.Stderr, "\n")

	log.Fatal(http.ListenAndServe(url, server))

}

func main() {
	mux := http.NewServeMux()

	auth := middlewares.Auth("test")
	handler := &routers.Router{
		Post: handlers.RunCmd,
	}
	mux.Handle("/api/run", auth(handler))

	server := middlewares.AddMiddleware(mux, middlewares.JSON(), middlewares.CORS("*"), middlewares.Logger())

	serverInit(server)
}
