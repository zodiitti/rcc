package cmd

import (
	"encoding/json"

	"github.com/robocorp/rcc/cloud"
	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/operations"

	"github.com/spf13/cobra"
)

var userinfoCmd = &cobra.Command{
	Use:     "userinfo",
	Aliases: []string{"user"},
	Short:   "Query user information from Robocorp Cloud.",
	Long:    "Query user information from Robocorp Cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Debug {
			defer common.Stopwatch("Userinfo query lasted").Report()
		}
		account := operations.AccountByName(AccountName())
		if account == nil {
			common.Exit(1, "Error: Could not find account by name: %v", AccountName())
		}
		client, err := cloud.NewClient(account.Endpoint)
		if err != nil {
			common.Exit(2, "Error: Could not create client for endpoint: %v, reason: %v", account.Endpoint, err)
		}
		data, err := operations.UserinfoCommand(client, account)
		if err != nil {
			common.Exit(3, "Error: Could not get user information: %v", err)
		}
		nice, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			common.Exit(4, "Error: Could not format reply: %v", err)
		}
		common.Log("%s", nice)
	},
}

func init() {
	cloudCmd.AddCommand(userinfoCmd)
}