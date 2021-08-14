package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Welcome to first web learning! </>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":7000", nil)
}