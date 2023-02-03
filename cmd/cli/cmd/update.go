package cmd

import (
    "github.com/Eldius/gvm/updates"
    "github.com/spf13/cobra"
)

// checkupdatesCmd represents the checkupdates command
var checkupdatesCmd = &cobra.Command{
    Use:   "update",
    Short: "Fetch latest release",
    Long: `Fetch latest release.
For example:

gvm update
`,
    Run: func(_ *cobra.Command, _ []string) {
        updates.CheckForUpdates()
    },
}

func init() {
    rootCmd.AddCommand(checkupdatesCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // checkupdatesCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // checkupdatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
