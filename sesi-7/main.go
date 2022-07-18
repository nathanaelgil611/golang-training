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

var p struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 20
	p.Water = (rand.Intn(max-min) + min)
	p.Wind = (rand.Intn(max-min) + min)
	p.WaterStatus = getWaterStatus(p.Water)
	p.WindStatus = getWindStatus(p.Wind)
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				p.Water = (rand.Intn(max-min) + min)
				p.Wind = (rand.Intn(max-min) + min)
				p.WaterStatus = getWaterStatus(p.Water)
				p.WindStatus = getWindStatus(p.Wind)
				fmt.Println(p.Water, p.Wind)
			case <-quit:
				ticker.Stop()
				return
			}
		}

	}()

	r := mux.NewRouter()
	r.HandleFunc("/users-url", getUserURL).Methods(http.MethodGet)
	r.HandleFunc("/weather-status", weatherStatus).Methods(http.MethodGet)

	http.Handle("/", r)

	http.ListenAndServe(PORT, nil)

}

func getWindStatus(val int) string {
	if val < 6 {
		return "Aman"
	} else if val < 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func getWaterStatus(val int) string {
	if val < 5 {
		return "Aman"
	} else if val < 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func weatherStatus(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./html/template.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, p)
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
