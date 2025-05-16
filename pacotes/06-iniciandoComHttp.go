package pacotes

import "net/http"

func Http2() {
	http.HandleFunc("/", buscaCep)
	http.ListenAndServe(":8080", nil)
}

func buscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}
