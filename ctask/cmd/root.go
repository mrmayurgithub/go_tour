package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ctask",
	Short: "It's a CLI Task Manager",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
