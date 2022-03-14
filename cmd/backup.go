/*
Copyright © 2022 tools tom <tom@lezotte.net>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup some macos features",
	Long: `Backup some macos features`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backup called")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.
	backupCmd.Flags().BoolP("config", "c", false, "Backup macos configurations")
	backupCmd.Flags().BoolP("software", "s", false, "Backup software lists")
}
