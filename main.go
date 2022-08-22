package main

import (
	"fmt"
	"log"
	"net/http"
)



func handleForm(w http.ResponseWriter ,r *http.Request){
	if err := r.ParseForm(); err !=nil{
		fmt.Fprintf(w,"parseform() err : %v", err)
	}

	fmt.Fprintf(w, "POST successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name : %s\n", name)
	fmt.Fprintf(w, "address : %s\n", address)

}




func handleHello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 error found", http.StatusNotFound)
	}

	if r.Method != "GET"{
		http.Error(w, "method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w,"Hello")


}

func main() {

	fileServer := http.FileServer(http.Dir("./statics"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	fmt.Printf("port is running at 8080\n")

	if err :=	http.ListenAndServe(":8080",nil) ; err != nil {
		log.Fatal(err)
	}
	
	
}