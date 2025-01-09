/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task-description>",
	Short: "Add an entry to the todo list",
	Long:  `Nothing more to explain here, just add an entry to the todo list and that's it.`,
	Args:  cobra.ExactArgs(1),
	Run:   addFunc,
}

func addFunc(cmd *cobra.Command, args []string) {
  f, err := os.OpenFile("todo_list.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
  if err != nil {
    fmt.Println("Error opening todo_list.csv file")
    return
  }
  defer f.Close()

  _, err = f.WriteString(args[0] + ",false," + time.Now().Format(time.RFC3339) + "\n")
  if err != nil {
    fmt.Println("Error writing `" + args[0] + "` to todo_list.csv", err)
    return 
  }

  fmt.Println("A new todo list item with the following description created succesfully:\n\t'" + args[0] + "'")
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
