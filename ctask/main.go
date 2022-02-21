package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go_tour/ctask/cmd"
	"github.com/go_tour/ctask/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
