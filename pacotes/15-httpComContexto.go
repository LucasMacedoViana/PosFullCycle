package pacotes

import (
	"context"
	"io"
	"net/http"
	"time"
)

func HttpComContexto() {
	// O pacote http permite que você passe um contexto para as requisições HTTP
	// Isso é útil para cancelar requisições quando o cliente fecha a conexão ou quando o servidor demora muito para responder
	// O contexto é passado como um argumento para o método http.NewRequest
	// O método http.NewRequest retorna um objeto do tipo *http.Request

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
