package main

import (
	"github.com/Netflix/go-env"
	"github.com/rawnly/youtrack/commands"
	"github.com/rawnly/youtrack/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


func main() {
	storage := new(util.Storage)
	storage, err := storage.Init();
	
  var rootCmd = &cobra.Command{ Use: "youtrack" }
	rootCmd.AddCommand(commands.GetIssueCmd(storage))
	rootCmd.AddCommand(commands.GetAuthCommand(storage))
	rootCmd.AddCommand(commands.GetLogTimeCommand(storage))
	rootCmd.AddCommand(commands.GetWorkItemsCommand(storage))


  logrus.SetOutput(rootCmd.OutOrStdout())

  var e Environment

  es, err := env.UnmarshalFromEnviron(&e)

  if err != nil {
    logrus.Fatal(err.Error())
  }

  e.Extras = es

  if e.Debug == 1 {
    logrus.SetLevel(logrus.DebugLevel)
  } else {
    logrus.SetLevel(logrus.WarnLevel)
  }

	if err != nil { logrus.Fatal(err) }

	_ = rootCmd.Execute()
}
