package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	cur1 := os.Args[1]
	mount := os.Args[2]
	cur2 := os.Args[3]
	url := fmt.Sprintf("http://api.fixer.io/latest")
	type Format struct {
		Base  string             `json:"base"`
		Date  string             `json:"date"`
		Rates map[string]float64 `json:"rates"`
	}
	format2 := Format{}
	get_json("http://api.fixer.io/latest", &format2)
	fmt.Println(format2.Rates["USD"])
	curen1 := format2.Rates[cur1]
	curen2 := format2.Rates[cur2]
	curen1 float
	curen2 float
	rate := cur1 / cur2
	amount := float(mount)
	cur3 := rate * amount
	fmt.Println(amount, "of", cur1, "transfers to", cur3, cur2, "at a rate of", rate)
}

//http://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func get_json(url string, target interface{}) error {
	//if it works, it works
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
