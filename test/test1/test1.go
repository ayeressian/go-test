package test1

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ayeressian/go-test2/test/test1/handlers"
)

func Test1() {
	appLog := log.New(os.Stdout, "echo_server", log.LstdFlags)
	h := handlers.NewProducts(appLog)
	sm := http.NewServeMux()
	sm.Handle("/", h)

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