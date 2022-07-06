package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"sesi-4/database"
	"sesi-4/model"
	"sesi-4/services"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

// GET /orders (untuk get all orders)
// GET /orders/id (untuk get order by id)
// POST /orders (untuk create order)
// PUT /orders/id (untuk update order)
// DELETE /orders/id (untuk delete order by id)

func main() {

	database.DatabaseInit()
	defer database.CloseDatabase()

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/orders", getAllOrder).Methods(http.MethodGet)
	r.HandleFunc("/orders/{id}", getOrder).Methods(http.MethodGet)
	r.HandleFunc("/orders", createOrder).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}", updateOrder).Methods(http.MethodPut)
	r.HandleFunc("/orders/{id}", deleteOrder).Methods(http.MethodDelete)

	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)

}

func getAllOrder(w http.ResponseWriter, r *http.Request) {
	order := services.NewOrderService()
	newOrder := order.GetAll()

	jsonData, _ := json.Marshal(&newOrder)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	order := services.NewOrderService()

	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	newOrder := order.GetByID(orderID)

	jsonData, _ := json.Marshal(&newOrder)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func createOrder(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var order model.Order

	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	fmt.Println(order)
	orderSvc := services.NewOrderService()

	orderSvc.Create(order)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Register")
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	order := services.NewOrderService()

	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	fmt.Println(orderID)

	order.Delete(orderID)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "Success Delete")
}

func updateOrder(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var order model.Order
	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	orderSvc := services.NewOrderService()

	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])

	orderSvc.Update(&order, orderID)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Update")
}
