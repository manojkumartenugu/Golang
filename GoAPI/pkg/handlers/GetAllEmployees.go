package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/mocks"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Employees)
}
