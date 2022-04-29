package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/mocks"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id perameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//
	for _, employee := range mocks.Employees {
		if employee.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(employee)
			break
		}
	}
}
