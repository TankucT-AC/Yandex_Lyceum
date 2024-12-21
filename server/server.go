package server

import (
	"calculator/pkg/calculate"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type SuccessResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusInternalServerError, calculate.ErrInternalServer.Error())
		return
	}
	var req = Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	result, err := calculate.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, calculate.ErrInvalidExpression) {
			errorResponse(w, http.StatusUnprocessableEntity, calculate.ErrInvalidExpression.Error())
		} else {
			errorResponse(w, http.StatusInternalServerError, calculate.ErrInternalServer.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := SuccessResponse{Result: result}
	json.NewEncoder(w).Encode(response)
}

func errorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (a *Application) RunServer() error {
	fmt.Println("Сервер запущен! Вводите запрос через другой терминал)")
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
