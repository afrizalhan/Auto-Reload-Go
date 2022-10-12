package service

import (
	"auto-reload-go/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

func Reloader() {
	for {
		water := rand.Intn(100-1) + 1
		wind := rand.Intn(100-1) + 1

		status := models.Status{}

		status.Stats.Water = water
		status.Stats.Wind = wind

		jsonData, err := json.Marshal(status)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("status.json", jsonData, 0644)

		if err != nil {
			panic(err)
		}

		time.Sleep(15 * time.Second)
	}
}


func ReadJSON() map[string]interface{} {
	jsonData, err := ioutil.ReadFile("status.json")
	if err != nil {
		panic(err)
	}

	var status models.Status

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		panic(err)
	}

	water := status.Stats.Water
	wind := status.Stats.Wind

	var waterStatus, windStatus string

	waterStatus = waterCondition(water)
	windStatus = windCondition(wind)

	data := map[string]interface{}{
		"statusWater": waterStatus,
		"statusWind":  windStatus,
		"water":       water,
		"wind":        wind,
	}

	// template, err := template.ParseFiles("./template/template.html")
	// if err != nil {
	// 	panic(err)
	// }
	// template.Execute(w, data)

	return data
}

func waterCondition(height int) string {
	var status string
	if height <= 5 {
		status = "Aman"
	} else if height >= 6 && height <= 8 {
		status = "Siaga"
	} else {
		status = "Bahaya"
	}
	return status
}

func windCondition(condition int) string{
	var status string
	if condition <= 5 {
		status = "Aman"
	} else if condition >= 6 && condition <= 8 {
		status = "Siaga"
	} else {
		status = "Bahaya"
	}
	return status
}