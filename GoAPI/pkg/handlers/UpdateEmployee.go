package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/mocks"
	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/models"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var UpdateEmployee models.Employee
	json.Unmarshal(body, &UpdateEmployee)
	// Iterate over all the mock Employees

	for index, employee := range mocks.Employees {
		if employee.Id == id {
			// update and send response when employee ID matches dynamic id
			employee.Name = UpdateEmployee.Name
			employee.Deprtment = UpdateEmployee.Deprtment

			mocks.Employees[index] = employee

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}

}
