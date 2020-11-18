// Portfolio website for my projects written in Golang
// Used free template for the html on "https://www.w3schools.com/w3css/w3css_templates.asp"

package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {

	stream, err := ioutil.ReadFile("index.html")

	if err != nil { log.Fatal(err) }

	indexFile := string(stream)

	fmt.Fprintf(w, indexFile)

	//should this be in the main func?

	
	
}

func main(){

	//Tells the webserver to create the endpoint "/" and display handler

	http.HandleFunc("/",indexHandler)
	// Wraps the ListenAndServe function to handle any errors

	log.Fatal(http.ListenAndServe(":8080", nil))

	//ListenAndServe() always returns an error
}

