package main

import (
	"fmt"
	"github.com/andreylsant/Desenvolvendo_Api_Rest_golang/router"
)

func main() {
	fmt.Println("iniciando servidor")
	router.HandleRouter()
}