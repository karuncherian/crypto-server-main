package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"crypto-server/service"
)

func NewHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/status", healthCheck)
	r.Get("/currency/{id}", getAllCurrency)

	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Ok"))
}

func getAllCurrency(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	var jsonBody []byte

	if param == "all" {
		res := service.GetAllCurrency()
		jsonBody, _ = json.Marshal(res)
	} else {
		res := service.GetAllCryptoPriceBySymbol(param)
		jsonBody, _ = json.Marshal(res)
	}

	w.Write(jsonBody)
}
