/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/go_tour/ctask/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete.")
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
		// fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
