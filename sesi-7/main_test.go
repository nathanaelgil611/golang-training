package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//testing GenerateJson()

func TestGenerateJson(t *testing.T) {
	var dataWeather struct {
		Water       int    `json:"water"`
		Wind        int    `json:"wind"`
		WaterStatus string `json:"water_status"`
		WindStatus  string `json:"wind_status"`
	}

	t.Run("success", func(t *testing.T) {
		go GenerateJson()
		time.Sleep(500 * time.Millisecond)
		// read from json file and write to webData
		path := "./static/weather.json"
		require.FileExists(t, path)

		file, err := ioutil.ReadFile(path)
		require.NoError(t, err)
		require.NotNil(t, file)

		json.Unmarshal(file, &dataWeather)
		require.NotEmpty(t, dataWeather.WindStatus)
		t.Log(dataWeather)
	})
}
