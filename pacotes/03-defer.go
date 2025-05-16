package pacotes

func Defers() {
	defer println("defer 1")
	defer println("defer 2")
	println("execução")
}
