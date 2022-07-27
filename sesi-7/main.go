package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"sesi-7/model"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

const jsonPath = "./static/weather.json"
const htmlPath = "./html/template.html"

var dataWeather struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	go GenerateJson()
	r := mux.NewRouter()
	r.HandleFunc("/users-url", getUserURL).Methods(http.MethodGet)
	r.HandleFunc("/weather-status", weatherStatus).Methods(http.MethodGet)

	http.Handle("/", r)

	http.ListenAndServe(PORT, nil)

}

func GenerateJson() {

	for {
		dataWeather.Water = (rand.Intn(100))
		dataWeather.Wind = (rand.Intn(100))
		dataWeather.WaterStatus = GetWaterStatus(dataWeather.Water)
		dataWeather.WindStatus = GetWindStatus(dataWeather.Wind)
		jsonData, _ := json.Marshal(&dataWeather)
		ioutil.WriteFile(jsonPath, jsonData, os.ModePerm)

		time.Sleep(15 * time.Second)
	}
}

func GetWindStatus(val int) string {
	if val < 6 {
		return "Aman"
	} else if val < 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func GetWaterStatus(val int) string {
	if val < 5 {
		return "Aman"
	} else if val < 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func weatherStatus(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(file, &dataWeather)
	templates, err := template.ParseFiles(htmlPath)
	if err != nil {
		log.Fatal(err)
	}

	templates.Execute(w, dataWeather)
}

func getUserURL(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []model.UserURL
	json.Unmarshal(responseData, &responseObject)

	jsonData, _ := json.Marshal(&responseObject)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}
