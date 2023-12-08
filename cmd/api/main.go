package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goodvandro/go-intensivo/internal/entity"
)

func main() {
	// chi -> go router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", OrderHandler)
	// http.HandleFunc("/order", OrderHandler)
	http.ListenAndServe(":8888", r)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	order, _ := entity.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(order)
}
