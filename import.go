package main

import(
  "fmt"
  "os"
  "bufio"
  "log"
)

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
    fmt.Println("sentence: ", sentence)
  }
  return nil
}
