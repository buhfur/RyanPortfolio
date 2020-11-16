// Portfolio website for my projects written in Golang


package main

import (
	"fmt"
	"ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "home page")
}

func main(){

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

