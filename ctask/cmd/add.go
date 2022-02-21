package cmd

import (
	"fmt"
	"strings"

	"github.com/go_tour/ctask/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong while adding the task.")
			return
		}
		fmt.Printf("Added the task: %s", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
