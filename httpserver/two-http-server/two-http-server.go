package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello request host: %v", r.Host)
}

func createHTTPServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%v", port),
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleFunc)

	firstServer := createHTTPServer("8080", router)
	secondServer := createHTTPServer("8081", router)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		log.Printf("firstServer: %+v", firstServer)
		log.Fatal(firstServer.ListenAndServe())
		wg.Done()
	}()
	go func() {
		log.Printf("second server: %+v", secondServer)
		log.Fatal(secondServer.ListenAndServe())
		wg.Done()
	}()

	log.Printf("wait")
	wg.Wait()
	log.Printf("after wait")

}
