package cmd

import (
	"github.com/Eldius/go-version-manager/hooks"
	"github.com/spf13/cobra"
)

// hooksAddCmd represents the hooksAdd command
var hooksAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new setup hook",
	Long: `Adds a new setup hook. For example:

go-version-manager hooks add <hook>

where 'hook' is a script file or a command
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hooks.AddHook(args[0])
	},
}

func init() {
	hookCmd.AddCommand(hooksAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hooksAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hooksAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
