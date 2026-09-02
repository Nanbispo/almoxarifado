package main

import (
	"almoxarifado/config"
	"log"
	"os"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	router, err := config.InitializeApp(db)
	if err != nil {
		log.Fatal("Erro ao inicializar a aplicacao:", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
