package pacotes

import (
	"io"
	"net/http"
)

func Http() {
	httpGet()
}

func httpGet() {
	// manipulaçao de requisições HTTP
	request, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(result))
	
}
