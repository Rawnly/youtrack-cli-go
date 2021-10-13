package commands

import (
	"github.com/rawnly/youtrack/api"
	"github.com/rawnly/youtrack/util"
	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"
)

func GetWorkItemsCommand(storage *util.Storage) *cobra.Command {
	return &cobra.Command{
		Use: "workitems get <issueId>",
		Short: "Get Issue workitems",
		Args: cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "get":
				workItems, err := api.GetWorkItems(args[1])(storage)

				if err != nil {
					cmd.PrintErr(err)
					return
				}


				table := tablewriter.NewWriter(cmd.OutOrStdout())
				table.SetHeader([]string{"Id", "Time", "Text", "Type", "User"})

				for _, wk := range *workItems {
					data := []string{
						wk.Id,
						wk.Duration.Presentation,
						wk.Text,
						wk.Type.Name,
						wk.Author.FullName,
					}

					table.Append(data)
				}

				table.Render()

				break
			}
		},
	}
}
