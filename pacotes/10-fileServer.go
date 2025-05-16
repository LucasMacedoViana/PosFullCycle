package pacotes

import (
	"log"
	"net/http"
)

func FileServer() {
	fileServer := http.FileServer(http.Dir("./public"))
	m := http.NewServeMux()
	m.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8000", m))
}
