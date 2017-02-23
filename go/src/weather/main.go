package main
import (
	"fmt"
	"net/url"
	s "strings"
	"log"
	"net/http"
	"encoding/json"
)
func weather(loc []string) {
	units := "imperial"
	m := "main"
	s := "sys"
	loct :=	s.ToLower(loc)
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + loct + "&appid=6a4cbda1e239084151dea640a95d2a0c" + "&units=" + units
	type Data struct {
		Place	string
		Temp	string
		Country	string
		Temp_min	string
	}	
	req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return	
