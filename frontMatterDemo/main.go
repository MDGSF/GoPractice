package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	dir := "/home/huangjian/git/huangjian/mdgsf.github.io/_posts"
	//dir := "dir"
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range entries {
		fmt.Println("\nstart process file =", fileInfo.Name())
		if strings.HasSuffix(fileInfo.Name(), ".md") {
			path := filepath.Join(dir, fileInfo.Name())

			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
				continue
			}

			lines := strings.Split(string(data), "\n")

			var newData string

			processed := false
			for idx, line := range lines {

				fmt.Printf("line = %v, processed = %v, data = %v\n", idx, processed, line)

				if processed {

					if idx == len(lines)-1 {
						if len(line) == 0 {
							break
						}
						if len(strings.TrimSpace(line)) == 0 {
							break
						}
					}

					newData += line + "\n"
				} else {

					if strings.HasPrefix(line, "tags:") {

						linearr := strings.Split(line, ":")
						tagsContent := linearr[1]
						tagsContent = strings.TrimSpace(tagsContent)

						if strings.Contains(tagsContent, "[") && strings.Contains(tagsContent, "]") {
							tagsContent = strings.TrimLeft(tagsContent, "[")
							tagsContent = strings.TrimRight(tagsContent, "]")
						}

						newline := "tags: ["
						arr := strings.Split(tagsContent, ",")
						first := true
						for _, tag := range arr {
							newtag := strings.TrimSpace(tag)
							if len(newtag) == 0 {
								continue
							}

							newtag = strings.ToLower(newtag)
							b := []byte(newtag)
							b[0] = b[0] - ('a' - 'A')
							newtag = string(b)

							if first {
								newline = newline + newtag
								first = false
							} else {
								newline += ","
								newline = newline + newtag
							}
						}
						newline += "]"

						newData += newline + "\n"

						processed = true

					} else {
						newData += line + "\n"
					}

				}
			}

			ioutil.WriteFile(path, []byte(newData), fileInfo.Mode())

		}
	}

}
