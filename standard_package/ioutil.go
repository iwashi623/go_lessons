package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// content, err := ioutil.ReadFile("ioutil.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
