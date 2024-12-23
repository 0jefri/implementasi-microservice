package main

import (
	"fmt"
	"log"
	"net/http"
	"tiket-online/handlers"
	"tiket-online/repository"
	"tiket-online/services"
)

func main() {
	orderRepo := repository.NewOrderRepository()
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		case http.MethodGet:
			orderHandler.GetOrderDetail(w, r)
		case http.MethodPut:
			orderHandler.UpdateOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8083"
	fmt.Printf("Starting server on http://localhost%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
