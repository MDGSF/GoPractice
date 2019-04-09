package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/awalterschulze/gographviz"
)

func main() {
	fmt.Println("vim-go")
	gographviz.ParseString()
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)

	g.AddSubGraph("G", "G0", nil)

	g.AddSubGraph("G0", "Sub0", nil)

	g.AddNode("G", "Hello", nil)
	g.AddNode("G", "World", nil)
	g.AddEdge("Hello", "World", true, nil)
	s := g.String()
	fmt.Println(s)
	ioutil.WriteFile("test.dot", []byte(s), 0666)
	cmd := exec.Command("dot", "-Tpng", "test.dot", "-o", "test.png")
	cmd.Run()
}
