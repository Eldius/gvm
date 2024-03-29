package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
    Use:   "hooks",
    Short: "Lists the configured setup hooks",
    Long: `Lists the configured setup hooks. For example:

gvm hooks
`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hook called")
    },
}

func init() {
    rootCmd.AddCommand(hookCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // hookCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // hookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
