package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	"github.com/KenethSandoval/uigh/ui"
)

var (
	Username string
	Token    string
)

var rootCmd = &cobra.Command{
	Use:     "uigh",
	Short:   "A terminal UI for my github",
	Long:    "uigh allows you to browse and interact with github from your terminal",
	Example: "uigh --token <token> --username <username>",
	Run: func(cmd *cobra.Command, _ []string) {
		Username = getVariable(cmd, "Github username", "username", "GITHUB_USERNAME")
		Token = getVariable(cmd, "Github access token", "token", "GITHUB_TOKEN")
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: Token},
		)
		tc := oauth2.NewClient(ctx, ts)
		gh := github.NewClient(tc)
		if err := ui.NewProgram(Username, gh).Start(); err != nil {
			fmt.Println("Could not start uigh", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "", "Github username")
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Github personal access token")
}

func Execute() error {
	return rootCmd.Execute()
}

// getVariable return an empty string if the flags (Username, Token) are incorrect
func getVariable(cmd *cobra.Command, name string, param string, env string) string {
	value, err := cmd.PersistentFlags().GetString(param)
	if err != nil {
		fmt.Println("Could not get "+param, err)
		os.Exit(1)
	}

	if value != "" {
		return value
	}

	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	fmt.Println("You must pass your " + name + " using the --" + param + " flag or set the " + env + " environment variable")
	os.Exit(1)
	return ""
}
