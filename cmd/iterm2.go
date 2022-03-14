/*
Copyright © 2022 tools tom <tom@lezotte.net>

*/
package cmd

import (
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

// iterm2Cmd represents the iterm2 command
var iterm2Cmd = &cobra.Command{
	Use:   "iterm2",
	Short: "Tasks to help with iTerm2",
	Long:  `Tasks to help with iTerm2`,
	Run: func(cmd *cobra.Command, args []string) {
		hostsFlag, _ := cmd.Flags().GetBool("hosts")
		if hostsFlag {
			getVUMC()
		}
	},
}

func init() {
	rootCmd.AddCommand(iterm2Cmd)

	// Here you will define your flags and configuration settings.
	iterm2Cmd.Flags().BoolP("hosts", "d", false, "Update dynamic hosts")
}

var (
	oriIterm2  = "https://ori.app.vumc.org/inventory/cli/iterm2-hosts.php"
)

func getVUMC() {
	fullPath := path.Join(homeDir, "/Library/Application Support/iTerm2/DynamicProfiles/iterm2-hosts.json")
	fmt.Printf("Generating new %s file\n", "iterm2-hosts.json")
	fmt.Printf("Generating new %s file\n", fullPath)
	//resp, err := http.Get(oriIterm2)
	//if err != nil {
	//	// handle err
	//}
	//defer resp.Body.Close()
}
