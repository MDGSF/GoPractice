package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type Person struct {
	name string
}

func main() {
	t := template.New("test")
	t, _ = t.Parse("Hello {{.name}}!")
	p := Person{name: "Mary"}

	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	t.Execute(buf, p)

	fmt.Println(buf.String())
}
