package pacotes

import (
	"io/ioutil"
	"net/http"
)

func RequestCustom() {
	// O pacote http permite que você customize requisições HTTP
	// Você pode adicionar headers, cookies, autenticação, etc
	// Para isso, você deve criar um objeto do tipo http.Client
	// E então criar um objeto do tipo http.Request
	// E então chamar o método Do do objeto http.Client
	// O método Do retorna um objeto do tipo http.Response
	// O objeto http.Response tem um campo chamado Body que é um objeto do tipo io.ReadCloser
	// Você pode ler o corpo da resposta chamando o método ioutil.ReadAll do pacote ioutil

	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
