/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
