package commands

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/rawnly/youtrack/api"
	"github.com/rawnly/youtrack/util"
	"github.com/spf13/cobra"
	"io"
	"strconv"
)

func GetIssueCmd(storage *util.Storage) *cobra.Command  {
	return &cobra.Command{
		Use: "issue <issueId>",
		Short: "Get issue info",
		Args: cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			issue, err := api.FetchIssue(args[0], storage)

			if err != nil {
				cmd.PrintErr(err.Error())
				return
			}

			estimation, err := api.GetPeriodicIssueCustomField(*issue, "Estimation", *storage)

			if err != nil {
				cmd.PrintErr(err)
				return
			}

			spentTime, err := api.GetPeriodicIssueCustomField(*issue, "Spent time", *storage)

			if err != nil {
				cmd.PrintErr(err)
				return
			}

			if err := PrintIssue(cmd.OutOrStdout(), *issue, *estimation, *spentTime); err != nil {
				cmd.PrintErr(err.Error())
				return
			}
		},
	}
}


func PrintIssue(w io.Writer, issue api.Issue, estimation api.PeriodIssueCustomField, spentTime api.PeriodIssueCustomField) error {
	return util.Template {
		Data: map[string]interface{} {
			"issue": issue,
			"estimation": estimation.Value,
			"spentTime": spentTime.Value,
			"timeDiff": strconv.Itoa(int(spentTime.Value.Minutes - estimation.Value.Minutes)),
		},
		Template: heredoc.Doc(`

		{{ color "yellow+b" "["}}{{ color "yellow+b" .issue.IdReadable }}{{ color "yellow+b" "]"}} {{ color "yellow+b" .issue.Summary }}
		{{ dim .issue.Description }}
		-------------------
		Reported by: {{ .issue.Reporter.Login }}
		Updated by: {{ .issue.Updater.Login }}
		-------------------
		Status:{{if .issue.Resolved }} {{ color "green" "RESOLVED" }} {{ else }} {{ bgYellow " UNRESOLVED " }} {{ end }}
		{{ .issue.CommentsCount }} Comments
		-------------------
		Spent Time: {{if gt .spentTime.Minutes .estimation.Minutes }}{{ .spentTime.Presentation }} (+{{ bgYellow .timeDiff }}) {{ else }}{{ .spentTime.Presentation }} {{end}}
		Estimation: {{ .estimation.Presentation }}
	`),
	}.Execute(w)
}