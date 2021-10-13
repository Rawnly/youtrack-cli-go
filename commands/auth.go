package commands

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/rawnly/youtrack/util"
	"github.com/spf13/cobra"
	"net/url"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func GetAuthCommand(storage *util.Storage) *cobra.Command {
	return &cobra.Command{
		Use: "auth",
		Short: "Authenticate your API",
		Args: cobra.NoArgs,
		Aliases: []string{"login"},
		Run: func(cmd *cobra.Command, args []string) {
			var qs = []*survey.Question{
				{
					Name: "token",
					Prompt: &survey.Input{
						Message:  "Paste your api token",
						Default: storage.Token,
					},
					Validate: survey.Required,
				},
				{
					Name: "url",
					Prompt: &survey.Input{
						Message: "Paste your Youtrack Url",
						Default: storage.Url,
					},
					Validate: func(ans interface{}) error {
						if err := survey.Required(ans); err != nil {
							return err
						}

						if str, ok := ans.(string); !ok || !isValidUrl(str) {
							return errors.New("Invalid Url")
						}

						return nil
					},
				},
			}

			answers := util.Storage{}

			if err := survey.Ask(qs, &answers); err != nil {
				cmd.PrintErrln(err)
				return
			}

			if err := answers.Save(); err != nil {
				cmd.PrintErrln(err)
				return
			}

			cmd.Println("Successfully updated!")
		},
	}
}