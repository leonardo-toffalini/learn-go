/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"time"
	"todo-cli/cmd"
)

func main() {
  t := time.Now().Format(time.RFC3339)
  fmt.Println(t)

	cmd.Execute()
}
