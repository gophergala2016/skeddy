package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkPayload(payload string) string {
	if string(payload[0]) == "@" {
		payload_file := payload[1:]
		file, err := os.OpenFile(payload_file, os.O_RDONLY, 0444)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		reader, _ := ioutil.ReadAll(file)

		return string(reader)
	}
	return payload
}

func sendRequest(endpoint string, payload string) {
	raw := bytes.NewBuffer([]byte(payload))
	req, err := http.NewRequest("POST", endpoint, raw)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	log.Println("HTTP request sent", response)
}

func Dispatch(endpoint string, payload string) {
	log.Println("Dispatching", endpoint, payload)
	endpoint = strings.TrimSpace(endpoint)
	payload = strings.TrimSpace(payload)
	payload = checkPayload(payload)
	sendRequest(endpoint, payload)
}
