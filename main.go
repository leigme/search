/*
Copyright © 2024 leig HERE <leigme@gmail.com>
*/
package main

import (
	"log"

	"github.com/leigme/search/cmd"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

func main() {
	cmd.Execute()
}
