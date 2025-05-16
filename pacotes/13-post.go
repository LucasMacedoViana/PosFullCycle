package pacotes

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func Post() {
	// O método http.Post é um método de conveniência que faz uma requisição HTTP POST
	// Ele recebe três argumentos: a URL, o tipo de conteúdo e o corpo da requisição
	// O corpo da requisição deve ser um objeto do tipo io.Reader
	// O método http.Post retorna um objeto do tipo *http.Response
	// O objeto *http.Response tem um campo chamado Body que é um objeto do tipo io.ReadCloser
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"nome":"João"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
