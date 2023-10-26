package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"finapp/bootstrap"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
