package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"hello")
	})
	http.ListenAndServe(":8888",nil)
}
