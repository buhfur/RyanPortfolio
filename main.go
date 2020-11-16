// Portfolio website for my projects written in Golang


package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//pageTitle := "Ryan Portfolio Home Page"

	pageBody, err := ioutil.ReadFile("index.html")

	if err != nil { log.Fatal(err) }

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, pageBody)

}

func main(){

	//Tells the webserver to create the endpoint "/" and display handler

	http.HandleFunc("/",indexHandler)

	// Wraps the ListenAndServe function to handle any errors

	log.Fatal(http.ListenAndServe(":8080", nil))

	//ListenAndServe() always returns an error
}

