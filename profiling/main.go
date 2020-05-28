package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
	}
	http.HandleFunc("/", handler)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var err error
	//add profiling here
	file, err := os.Create("profiling/cpu.prof")
	if err != nil {
		log.Fatal("err creating file: ", err)
	}

	err = pprof.StartCPUProfile(file)
	if err != nil {
		log.Fatal("err startCPUprofile: ", err)
	}
	defer pprof.StopCPUProfile()

	go func() {
		server.ListenAndServe()
	}()
	<-done
	log.Println("do shutdown")
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Println("error shutting down: ", err)
	}
}
