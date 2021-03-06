package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readSentence(sentence string) {
	index := strings.Index(sentence, "http")
	var exp, ep, p, temp string
	if index != -1 {
		exp = sentence[:(index - 1)]
		temp = sentence[index:]
		subIndex := strings.Index(temp, " ")
		if subIndex != -1 {
			ep = temp[:subIndex]
			p = temp[(subIndex + 1):]
		} else {
			ep = temp
		}
		entry := NewEntry(exp, ep, p)
		err := Store.SaveEntry(entry)
		if err != nil {
			log.Println("Error: ", err)
		}
	} else {
		fmt.Println("Endpoint URL not given")
	}
}

func ImportFile(filename string) error {
	fmt.Println("Importing cron file ...")
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sentence := scanner.Text()
		sentence = strings.TrimSpace(sentence)
		readSentence(sentence)
	}
	return nil
}
