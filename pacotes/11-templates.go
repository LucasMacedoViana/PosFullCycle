package pacotes

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func Templates() {
	curso := Curso{"Go", 40}
	tmp := template.New("Curso template")
	tmp, _ = tmp.Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

func Templates2() {
	curso2 := Curso{"Go", 40}
	t := template.Must(template.New("Curso template").Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas."))

	err := t.Execute(os.Stdout, curso2)
	if err != nil {
		panic(err)
	}
}

type c []Curso

func Templates3() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, c{{"Go", 40}, {"Python", 35}, {"Java", 60}})
	if err != nil {
		panic(err)
	}

}
