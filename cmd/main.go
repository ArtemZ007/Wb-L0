package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Получение параметров подключения
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}
	dbDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("myuser"), os.Getenv("mypassword"),
		os.Getenv("localhost"), os.Getenv("5432"),
		os.Getenv("mydb"))

	// Настройка подключения к Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Запуск веб-сервера для обработки API
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		order := &Order{}
		err := json.NewDecoder(r.Body).Decode(order)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// TODO: Process order
		log.Printf("Received new order: %v", order)
		err = json.NewEncoder(w).Encode(order)
		if err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
