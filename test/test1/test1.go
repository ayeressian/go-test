package test1

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ayeressian/go-test2/test/test1/handlers"
	"github.com/gorilla/mux"
)

func Test1() {
	appLog := log.New(os.Stdout, "echo_server", log.LstdFlags)
	h := handlers.NewProducts(appLog)
	
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", h.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", h.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", h.AddProduct)

	server := &http.Server{
		Addr: ":1234",
		Handler: sm,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			appLog.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	appLog.Printf("Gracefully closing %v", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(tc)
}