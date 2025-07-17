package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/model"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Home")
}

func ExibirPersonalidade(w http.ResponseWriter, r *http.Request) {
	personalidade := []model.Personalidade{
		{"1", "name1","descricao1"},
		{"2","name2","descricao2"},
	}
	json.NewEncoder(w).Encode(personalidade)
}

func ExibirPersonalidadePorID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	for _, personalidade := range model.Personalidades{
		if personalidade.Id == id{
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
