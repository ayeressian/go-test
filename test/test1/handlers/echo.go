package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Echo struct {
	l *log.Logger
}

func NewEcho(l *log.Logger) *Echo {
	return &Echo{l: l}
}

func (e *Echo) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var body []byte
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Incorrect Body", http.StatusBadRequest)
		return
	}
	e.l.Printf("echo %v", string(body))
	response.Write(body)
}
