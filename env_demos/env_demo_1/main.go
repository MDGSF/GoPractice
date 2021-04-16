package main

import (
	"fmt"
	"os"
)

// env ./env_demo_1
// env -i ./env_demo_1
// env --unset=name ./env_demo_1
func main() {
	fmt.Println(os.Getenv("name"))
}
