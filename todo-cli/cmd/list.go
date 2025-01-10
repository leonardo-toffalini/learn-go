package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var All bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list out the items of the todo list",
	Long:  `list out the items of the todo list with some nice formatting`,
	Args:  cobra.ExactArgs(0),
	Run:   listFunc,
}

func listFunc(cmd *cobra.Command, args []string) {
	content, err := os.ReadFile("todo_list.csv")
	if err != nil {
		fmt.Println("Error opening todo_list.csv file ")
		return
	}

	r := csv.NewReader(strings.NewReader(string(content)))
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	rows, _ := r.ReadAll()

	for n, row := range rows {
		if n == 0 {
			if All {
				fmt.Fprintln(w, strings.Join(row, "\t"))
			} else {
				fmt.Fprint(w, "ID\t")        // ID
				fmt.Fprint(w, row[0]+"\t")   // Task
				fmt.Fprintln(w, row[2]+"\t") // Created
			}
			continue
		}

		t, e := time.Parse(time.RFC3339, row[2]) // absolute time
		if e != nil {
			log.Fatal(e)
		}
		dt := time.Now().Sub(t) // delta time
		prettyDt := timediff.TimeDiff(time.Now().Add(-dt))

		if All {
			fmt.Fprint(w, strconv.Itoa(n)+"\t") // ID
			fmt.Fprint(w, row[0]+"\t")          // Task
			fmt.Fprint(w, row[1]+"\t")          // Done
			fmt.Fprintln(w, prettyDt)           // Created
		} else {
			if row[1] == "false" {
				fmt.Fprint(w, strconv.Itoa(n)+"\t") // ID
				fmt.Fprint(w, row[0]+"\t")          // Task
				fmt.Fprintln(w, prettyDt)           // Created
			}
		}

	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVarP(&All, "all", "a", false, "this is the all flag")
}
