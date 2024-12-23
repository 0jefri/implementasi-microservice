package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"tiket-online/models"
	"tiket-online/services"

	"github.com/google/uuid"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()
	log.Printf("[INFO] Creating order with ID: %s\n", order.ID)
	if err := h.service.CreateOrder(order); err != nil {
		log.Printf("[ERROR] Failed to create order: %v\n", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	createdOrder, err := h.service.GetOrderDetail(order.ID)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch created order: %v\n", err)
		http.Error(w, "Failed to fetch created order", http.StatusInternalServerError)
		return
	}

	log.Printf("[DEBUG] Final Response: %+v\n", createdOrder)

	response := models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Order created successfully",
		Data:    createdOrder,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	log.Println("[ERROR] ID is required")
	if id == "" {
		errorResponse := models.ErrorResponse{
			Status:  "error",
			Message: "ID is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	log.Printf("[INFO] Retrieving order with ID: %s\n", id)
	order, err := h.service.GetOrderDetail(id)
	if err != nil {
		log.Printf("[ERROR] Order not found: %v\n", err)
		errorResponse := models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Order retrieved successfully",
		Data:    order,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("[ERROR] ID is required")
		errorResponse := models.ErrorResponse{
			Status:  "error",
			Message: "ID is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		errorResponse := models.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	log.Printf("[INFO] Updating order with ID: %s\n", id)
	if err := h.service.UpdateOrder(id, updatedOrder); err != nil {
		log.Printf("[ERROR] Failed to update order: %v\n", err)
		errorResponse := models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Order updated successfully",
		Data:    updatedOrder,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
