package main

import (
	"net/http"
)

// set up a really simple handlers

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")

}
