package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000;

func main(){
	listenAt := fmt.Sprintf(":%d", port);

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Println(req)
		fmt.Fprint(res, "Hello World")
	})

	log.Printf("Open the following URL in the browser: http://localhost:%d\n", port)

	log.Fatal(http.ListenAndServe(listenAt, nil))
}