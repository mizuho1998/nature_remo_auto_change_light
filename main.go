package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/mizuho1998/nature_remo_light/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getIl() (il int, err error) {
	TOKEN := os.Getenv("TOKEN")
	url := "https://api.nature.global/1/devices"

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	req.Header.Add("Authorization", "Bearer "+TOKEN)
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var device []model.Device
	err = json.Unmarshal(body, &device)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	il = device[0].NewestEvents.Il.Val

	return il, nil
}

func post() (status string, err error) {
	TOKEN := os.Getenv("TOKEN")
	SIGNAL_ID := os.Getenv("SIGNAL_ID")
	url := "https://api.nature.global/1/signals/" + SIGNAL_ID + "/send"

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+TOKEN)
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return res.Status, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	il, _ := getIl()
	if il > 50 {
		post()
	}
}
