package cmd

import (
    "fmt"

    "github.com/Eldius/gvm/hooks"
    "github.com/Eldius/gvm/installer"
    "github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
    Use:   "use",
    Short: "Configure system to use an specific version",
    Long: `Configure system to use an specific version. For example:

gvm install go1.15.1

gvm install 1.15.1

`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if err := installer.Use(args[0]); err != nil {
            fmt.Println("Failed to setup version", args[0])
            fmt.Println(err.Error())
            return
        }
        for _, h := range hooks.ListHooks() {
            hooks.ExecuteHook(h)
        }
    },
}

func init() {
    rootCmd.AddCommand(useCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // useCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
