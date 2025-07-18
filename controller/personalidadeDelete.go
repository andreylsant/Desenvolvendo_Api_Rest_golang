package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/database"
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/model"
	"github.com/gorilla/mux"
)

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade []model.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]

	err := database.DB.Delete(&personalidade, id)
	if err != nil{
		http.Error(w, "Request Invalid", 500)
		return
	}

	json.NewEncoder(w).Encode(personalidade)
}
