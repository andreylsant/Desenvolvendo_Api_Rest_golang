package router

import (
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/controller"
	"github.com/gorilla/mux"
)

func MyRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("Get")
	r.HandleFunc("/personalidades", controller.ExibirPersonalidade).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controller.ExibirPersonalidadePorID).Methods("Get")
}