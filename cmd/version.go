package cmd

import (
    "fmt"
    "github.com/Eldius/gvm/config"
    "github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Shows app version info",
    Long:  `Shows app version info.`,
    Run: func(cmd *cobra.Command, args []string) {
        showVersionInfo()
    },
}

func showVersionInfo() {
    fmt.Printf(`
version:    %s
build date: %s
`,
        config.Version,
        config.BuildDate,
    )
}

func init() {
    rootCmd.AddCommand(versionCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // versionCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
