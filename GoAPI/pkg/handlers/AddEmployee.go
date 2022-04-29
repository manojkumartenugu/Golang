package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/mocks"
	"github.com/manojkumartenugu/Golangpractice/GoAPI/pkg/models"
)

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	//Read to the requested body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var Employee models.Employee
	json.Unmarshal(body, &Employee)
	//Append to the employee mocks
	Employee.Id = rand.Intn(100)
	mocks.Employees = append(mocks.Employees, Employee)
	//.mocks.Employees = append(mocks.Employees, Employee)
	//send a 201 created responce
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
