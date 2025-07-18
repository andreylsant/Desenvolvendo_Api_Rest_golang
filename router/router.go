package router

import (
	"log"
	"net/http"
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/controller"
	"github.com/gorilla/mux"
)

func HandleRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("Get")
	r.HandleFunc("/personalidades", controller.ExibirPersonalidade).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controller.ExibirPersonalidadePorID).Methods("Get")
	r.HandleFunc("/criandopersonalidades", controller.CriarPersonalidade).Methods("Post")
	r.HandleFunc("/deletapersonalidades", controller.DeletaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}