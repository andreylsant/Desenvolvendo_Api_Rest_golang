package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/database"
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/model"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func ExibirPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade []model.Personalidade

	 err := database.DB.Find(&personalidade)
	 if err != nil{
		http.Error(w, "Request invalid", 500)
		return
	 }
	json.NewEncoder(w).Encode(personalidade)
}

func ExibirPersonalidadePorID(w http.ResponseWriter, r *http.Request) {
	var personalidade []model.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
