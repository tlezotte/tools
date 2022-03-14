/*
Copyright © 2022 tools tom <tom@lezotte.net>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iterm2Cmd represents the iterm2 command
var iterm2Cmd = &cobra.Command{
	Use:   "iterm2",
	Short: "Tasks to help with iTerm2",
	Long: `Tasks to help with iTerm2`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iterm2 called")
	},
}

func init() {
	rootCmd.AddCommand(iterm2Cmd)

	// Here you will define your flags and configuration settings.
	iterm2Cmd.Flags().BoolP("hosts", "d", false, "Update dynamic hosts")
}
