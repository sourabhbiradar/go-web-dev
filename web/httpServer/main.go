package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)  // writing http status 
	w.Write([]byte("Hey! There")) // writng response in 'w'
}

func main(){
	// HTTP Server
	fmt.Println("\tHTTP Server")

	http.HandleFunc("/",handler) // handler func for root URL path '/'

	fmt.Println("server started at port 8080")
	http.ListenAndServe(":8080",nil)

	// remember to handle errors
}