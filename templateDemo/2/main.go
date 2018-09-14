package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("Hello {{.UserName}}!\n")
	p := Person{UserName: "HuangJian"}
	t.Execute(os.Stdout, p)
}
