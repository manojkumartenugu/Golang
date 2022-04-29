package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/mocks"
)

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Read the Dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//Iterate Over all the mock Books

	for index, employee := range mocks.Employees {
		if employee.Id == id {
			// Delete employee and send response if the empoyee Id matches dynamic Id
			//mocks.Employees = append(mocks.Employees[:index], mocks.Employees[index+1:]...)
			mocks.Employees = append(mocks.Employees[:index], mocks.Employees[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}

	}

}
