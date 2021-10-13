package commands

import (
	"github.com/rawnly/youtrack/api"
	"github.com/rawnly/youtrack/util"
	"github.com/spf13/cobra"
)

type LogCmdOptions struct {
	Text string `json:"text"`
}

func GetLogTimeCommand(storage *util.Storage) *cobra.Command {
	opts := &LogCmdOptions{}
	
	cmd := &cobra.Command{
		Use: "log <time> <issue>",
		Short: "Log issue spent time",
		Example: "youtrack log 1d YT-30",
		Args: cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			wk, err := api.CreateWorkItem(args[1], opts.Text, args[0])(storage)

			if err != nil {
				cmd.PrintErr(err)
				return
			}

			cmd.Println("Successfully Reported!")
			cmd.Println(wk)
		},
	}

	cmd.Flags().StringVarP(&opts.Text, "text", "t", "", "Describe your workitem")

	return cmd
}
