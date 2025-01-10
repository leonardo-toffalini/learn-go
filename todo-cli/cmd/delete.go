/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <ID>",
	Short: "Delete a task by id",
	Long: `Delete a task by id`,
  Args: cobra.ExactArgs(1),
	Run: deleteFunc,
}

func deleteFunc(cmd *cobra.Command, args []string) {
	f, err := os.OpenFile("todo_list.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) // NOTE: mode shoule be os.ModeAppend
	if err != nil {
		fmt.Println("Error opening todo_list.csv file")
		return
	}
	defer f.Close()

	tempFile, err := os.OpenFile("temp.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) // NOTE: mode shoule be os.ModeAppend
	if err != nil {
		fmt.Println("Error opening todo_list.csv file")
		return
	}
	defer tempFile.Close()

	w := csv.NewWriter(tempFile)
	r := csv.NewReader(f)

	rows, _ := r.ReadAll()

	for rowID, row := range rows {
		if strconv.Itoa(rowID) == args[0] {
			continue
		}
		if err := w.Write(row); err != nil {
			fmt.Println(err)
		}
	}
  w.Flush()

  if err := os.Rename("temp.csv", "todo_list.csv"); err != nil {
		fmt.Printf("Error replacing original file: %v\n", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
