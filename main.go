package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/k0marov/newsbot/frontend"
	"log"
)

func main() {
	log.Println("Starting bot...")
	frontend.StartBot()
}
