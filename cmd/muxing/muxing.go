package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", NameHandler).Methods("GET")
	router.HandleFunc("/bad", BadHandler).Methods("GET")
	router.HandleFunc("/data", DataHandler).Methods("POST")
	router.HandleFunc("/headers", HeadersHandler).Methods("POST")
	router.HandleFunc("/", NotDefinedHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars((r))
	fmt.Fprintf(w, "Hello, %s!", name["PARAM"])
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, `I got message:\n%s!`, string(p))
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		log.Fatal(err)
	}
	c := a + b
	w.Header().Set("a+b", strconv.Itoa(c))
}

func NotDefinedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
