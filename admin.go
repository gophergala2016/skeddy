package main

import(
  "fmt"
	"log"
	"net/http"
)

func StartAdminInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Starting admin interface at http://%s\n", addr)
  http.ListenAndServe(addr, nil)
}
