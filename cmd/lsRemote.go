package cmd

import (
	"fmt"

	"github.com/Eldius/go-version-manager/versions"
	"github.com/spf13/cobra"
)

// lsRemoteCmd represents the lsRemote command
var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List available Go versions",
	Long:  `List available Go versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsRemote called")
		versions := versions.ListAvailableVersions()
		for _, v := range versions {
			fmt.Printf("- %s\n", v.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
