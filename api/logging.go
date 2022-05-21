package main

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

//methode recoit une interface et retourne une reponse a notre requette
func (a *App) createLoggingRouter(out io.Writer) http.Handler {
	return handlers.LoggingHandler(out, a.Router)
}
