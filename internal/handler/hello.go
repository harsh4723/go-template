package handler

import (
	"encoding/json"
	"net/http"

	"go.template/internal/models"
	"go.template/internal/service"
)

type Decoder func(*http.Request) (interface{}, error)
type Encoder func(http.ResponseWriter, interface{}) error

type HelloHandler struct {
	svc service.HelloService
}

func NewHelloHandler(svc service.HelloService) *HelloHandler {
	return &HelloHandler{svc: svc}
}

func HelloDecoder(r *http.Request) (interface{}, error) {
	// do request validations here
	//var req models.HelloRequest
	name := r.URL.Query().Get("name")
	req := models.HelloRequest{Name: name}
	//err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	return models.HelloRequest{}, errors.Wrap(errBadRequest, "cannot decode request")
	// }

	return req, nil
}

func HelloEncoder(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

func (h *HelloHandler) SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := HelloDecoder(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		res, err := h.svc.SayHello(r.Context(), req.(models.HelloRequest))
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		err = HelloEncoder(w, res)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
	}
}
