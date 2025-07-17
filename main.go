package main

import (
	"net/http"

	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.MyRouter()
	if err := http.ListenAndServe(":8000", r); err != nil{
		err.Error()
		return
	}
}