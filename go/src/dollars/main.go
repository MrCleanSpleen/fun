package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	amnt := os.Args[1]
	curen1 := os.Args[2]
	curen2 := os.Args[3]
	CurrencyTransfer(curen2, curen1, amnt)
}
func CurrencyTransfer(cur1 string, cur2 string, amount string) {
	type Format struct {
		Base  string             `json:"base"`
		Date  string             `json:"date"`
		Rates map[string]float64 `json:"rates"`
	}
	format2 := Format{}
	get_json("http://api.fixer.io/latest", &format2)
	curen1 := format2.Rates[cur1]
	curen2 := format2.Rates[cur2]
	num, _ := strconv.ParseFloat(amount, 64)
	rate := curen1 / curen2
	cur3 := rate * num
	fmt.Println(num, cur2, "transfers to", cur3, cur1, "at a rate of", rate)
}

//http://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func get_json(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
