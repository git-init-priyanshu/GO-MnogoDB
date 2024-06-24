package routes

import (
	"encoding/json"
	"fmt"
	"go-MongoDB/controllers"
	"go-MongoDB/models"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmpService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmpService) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

	// Getting emploee id
	id := mux.Vars(r)["id"]

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	emp, err := controller.FindByEmployeeId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = emp
	w.WriteHeader(http.StatusOK)
	log.Printf("Employee: %s", emp)
}
func (svc *EmpService) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	emps, err := controller.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = emps
	w.WriteHeader(http.StatusOK)
	log.Printf("Employee: %s", emps)
}
func (svc *EmpService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

  // Getting body
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Changing the id
	emp.EmployeeId = uuid.NewString()

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	id, err := controller.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = emp.EmployeeId
	w.WriteHeader(http.StatusOK)
	log.Printf("Employee created with id: %s %s", id, emp)
}
func (svc *EmpService) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

	// Getting emploee id
	id := mux.Vars(r)["id"]
  if id==""{
    w.WriteHeader(http.StatusBadRequest)
    fmt.Println("Invalid empoyee id")
    res.Error = "Invalid empoyee id"
    return
  }

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

  emp.EmployeeId = id

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := controller.UpdateEmployeeById(id, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = count
	w.WriteHeader(http.StatusOK)
  log.Printf("Employees updated: %d", count)
}
func (svc *EmpService) DeleteEmpoyee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

	// Getting emploee id
	id := mux.Vars(r)["id"]
  if id==""{
    w.WriteHeader(http.StatusBadRequest)
    fmt.Println("Invalid empoyee id")
    res.Error = "Invalid empoyee id"
    return
  }

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := controller.DeleteEmployeeById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = count
	w.WriteHeader(http.StatusOK)
  log.Printf("Employees deleted: %d", count)
}
func (svc *EmpService) DeleteAllEmpoyees(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(&res)

	// Calling controller
	controller := controllers.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := controller.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error", err)
		res.Error = err.Error()
		return
	}

	// Response
	res.Data = count
	w.WriteHeader(http.StatusOK)
  log.Printf("Employees deleted: %d", count)
}
