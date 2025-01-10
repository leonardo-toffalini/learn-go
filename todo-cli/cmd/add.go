package cmd

import (
	"encoding/csv"
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
	f, err := os.OpenFile("todo_list.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening todo_list.csv file")
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting information of todo_list.csv")
	}

	// if file is newly created, add the column names
	if fileInfo.Size() == 0 {
		colNames := []string{"Task", "Done", "Created"}

		err = w.Write(colNames)
		if err != nil {
			fmt.Println("Error writing column names")
			return
		}
	}

	record := []string{args[0], "false", time.Now().Format(time.RFC3339)}

	err = w.Write(record)
	if err != nil {
		fmt.Println("Error writing into todo_list.csv", args[0])
		return
	}

	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("Error flushing data to CSV file", err)
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
