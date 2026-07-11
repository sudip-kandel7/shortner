package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){

	r := mux.NewRouter()
	
	

	http.Handle("/", r)
	fmt.Println("Server started at port: 8000")
	log.Fatal(http.ListenAndServe(":8000",r))

}