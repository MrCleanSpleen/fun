package main
import (
	"os"
	"fmt"
	"encoding/json"
	"http"
	"time"
)
cur1 := os.Args[1]
amount := os.Args[2]
cur2 := os.Args[3]
url := fmt.Sprintf("http://api.fixer.io/latest")
type Format struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates map[string]float64 `json:"rates"`
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
data := get_json("http://api.fixer.io/latest", &Format{})


