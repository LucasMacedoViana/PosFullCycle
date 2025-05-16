package pacotes

import "net/http"

func Mux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", blog{title: "Olá Blog!"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
