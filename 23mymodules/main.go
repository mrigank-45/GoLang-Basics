package main

// go get -u github.com/gorilla/mux
// can do "go build ." to make mymodules.exe
// go "mod vendor" to make vendor folder (like npm install)
// start server with "go run main.go" from packages from web
// start with "go run -mod=vendor main.go" to use vendor folder packages

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") // route

	log.Fatal(http.ListenAndServe(":4000", r)) // server at port 4000
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
