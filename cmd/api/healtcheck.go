package main

import (
	"fmt"
	"net/http"
)

func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	jsResponse := `{"status": "available", "environment": %q, "version": %q}`
	jsResponse = fmt.Sprintf(jsResponse, a.config.env, version)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsResponse))
}
