package main

import (
	"fmt"
	"net/http"
	"os/user"
)

var username string

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	username = u.Username
	fmt.Printf("Running server as %s...\n", username)
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, I'm %s!", username)
}
