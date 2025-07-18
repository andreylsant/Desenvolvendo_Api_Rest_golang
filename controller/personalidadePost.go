package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/database"
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/model"
)

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade []model.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	
	err := database.DB.Create(&personalidade)
	if err != nil{
		http.Error(w, err.Error.Error(), 500)
	}

	json.NewEncoder(w).Encode(personalidade)
}
