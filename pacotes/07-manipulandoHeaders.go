package pacotes

import "net/http"

func Http3() {
	http.HandleFunc("/", buscaCepHeandler)
	http.ListenAndServe(":8080", nil)
}

func buscaCepHeandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro": "cep deve ser informado"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá Mundo!"))
}
