package pacotes

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func Json() {
	//structParaJson()
	jsonParaStruct()

}

func structParaJson() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	//Json sempre retorna um slice de bytes
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

}

func jsonParaStruct() {
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var conta Conta
	err := json.Unmarshal(jsonPuro, &conta)
	if err != nil {
		println(err)
	}
	println(conta.Saldo)
}
