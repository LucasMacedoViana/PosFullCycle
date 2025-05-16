package ctx

import (
	"log"
	"net/http"
	"time"
)

func Server() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	log.Panicln("request iniciada")
	defer log.Println("request finalizada")
	select {
	case <-time.After(1 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	case <-c.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancela pelo cliente", http.StatusRequestTimeout)
	}
}
