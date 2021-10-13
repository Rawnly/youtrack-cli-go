package main

import (
	"github.com/rawnly/youtrack/commands"
	"github.com/rawnly/youtrack/util"
	"github.com/spf13/cobra"
)


func main() {
	storage := new(util.Storage)
	storage, err := storage.Init();

	if err != nil { panic(err) }

	var rootCmd = &cobra.Command{ Use: "youtrack" }
	rootCmd.AddCommand(commands.GetIssueCmd(storage))
	rootCmd.AddCommand(commands.GetAuthCommand(storage))
	rootCmd.AddCommand(commands.GetLogTimeCommand(storage))
	rootCmd.AddCommand(commands.GetWorkItemsCommand(storage))

	_ = rootCmd.Execute()
}