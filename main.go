package main

import (
	"fmt"

	"github.com/luiz/go-rest-api/database"
	"github.com/luiz/go-rest-api/models"
	"github.com/luiz/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
