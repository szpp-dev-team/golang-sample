package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// go run を実行した場所がカレントディレクトリ扱いになるので注意

	b, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
