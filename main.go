package main

import (
	"github.com/RianAlvesAce/gestao_online/internal/server"
	"github.com/RianAlvesAce/gestao_online/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDB()
	server.InitApi()
}