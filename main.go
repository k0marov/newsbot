package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/k0marov/newsbot/backend"
	"github.com/k0marov/newsbot/frontend"
	"log"
)

func main() {
	svc := &backend.FakeService{}
	log.Println("Starting bot...")
	frontend.StartBot(svc)
}
