package  main

import(
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"Hello")
}

func postHandler(w http.ResponseWriter,req *http.Request){
	ID:=req.PathValue("id") // "id" --> {id} wildcard
   w.Write([]byte(fmt.Sprintf("hello %s",ID)))
}

func getAll(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"Hello get all")
}

func main(){
	// serveMux 
	fmt.Println("\tserveMux")
   
	router:=http.NewServeMux()

	router.HandleFunc("GET /hello",getHandler)

	router.HandleFunc("POST /hello/{id}",postHandler)

	router.HandleFunc("GET /hello/get/all",getAll)

	if err:=http.ListenAndServe(":8080",router);err !=nil{
		fmt.Println("Error staring server",err)
	}



}