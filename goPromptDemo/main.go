package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func executor(in string) {
	fmt.Printf("in = %v\n", in)
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("go prompt demo"),
	)
	p.Run()
}
