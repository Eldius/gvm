package cmd

import (
	"fmt"
	"log"

	"github.com/Eldius/go-version-manager/versions"
	"github.com/spf13/cobra"
)

// lsRemoteCmd represents the lsRemote command
var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List available Go versions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsRemote called")
		versions := versions.ListAvailableVersions()
		for _, v := range versions {
			log.Printf("- %s => %s (%s)\n", v.Name, v.LinuxAmd64, v.Source)
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
