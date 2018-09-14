package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse(`
	empty pipeline if demo :
	{{if ""}}
	no output.
	{{end}}
	\n`))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse(`
	not empty pipeline if demo :
	{{if "anythings"}}
	I have content.
	{{end}}
	\n`))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse(`
	not empty pipeline if demo :
	{{if "anythings"}}
		if part.
	{{else}}
		else part.
	{{end}}
	\n`))
	tIfElse.Execute(os.Stdout, nil)
}
