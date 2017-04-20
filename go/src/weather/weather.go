package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//	"net/http"
	//	"time"
)

type Data struct {
	Main Main `json:"main"`
	//Weather Weather `json:"weather"`
	//Wind    Wind    `json:"wind"`
}

type Main struct {
	Temp     float64 `"json:"temp"`
	Humidity int     `json:"humidity"`
}

func main() {
	var data Data
	plan := []byte(`
	{
		"temp" : "280.32",
		"humidity: : "81"
	}
	`)
	err := json.Unmarshal(plan, &data)
	if err != nil {
	}
	fmt.Println(data)
}
