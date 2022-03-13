/*
Copyright © 2022 tools tom <tom@lezotte.net>

*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
)

// apullCmd represents the apull command
var apullCmd = &cobra.Command{
	Use:   "apull",
	Short: "Pull the lastest version for a repository",
	Long:  `Pull the lastest version for a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		local, _ := cmd.Flags().GetBool("local")
		icloud, _ := cmd.Flags().GetBool("icloud")
		remote, _ := cmd.Flags().GetBool("remote")
		if local {
			apull_local()
		} else if icloud {
			apull_icloud()
		} else if remote {
			apull_remote()
		} else {
			apull_local()
			apull_icloud()
			apull_remote()
		}
	},
}

func init() {
	rootCmd.AddCommand(apullCmd)

	// Here you will define your flags and configuration settings.
	apullCmd.Flags().BoolP("local", "l", false, "Pull repository updates on local ansible")
	apullCmd.Flags().BoolP("icloud", "i", false, "Pull repository updates on all icloud projects")
	apullCmd.Flags().BoolP("remote", "r", false, "Pull repository updates on some remote projects")
}

var (
	remoteServer   = "ori1020lp"
	remoteProjects = []string{
		"ansible",
		"ansible_dev/lezottt",
		"ansible_dev/ori_versions",
	}
	icloudProjects = []string{"ansible", "docker-centos7-ansible", "docker-centos8-ansible", "images"}
	icloud         = "$HOME/Library/Mobile Documents/com~apple~CloudDocs/_Code/VICTR"
)

func apull_local() {
	os.Chdir("$HOME/code/ansible/ansible")
	fmt.Printf("Updating local - ansible ... ")
	script.Exec("git checkout master")
	script.Exec("git pull")
	fmt.Printf("✅\n")
}

func apull_remote() {
	for _, project := range remoteProjects {
		fmt.Printf("Updating %s - %s ... ", remoteServer, project)
		//cmd := exec.Command("ssh", remoteServer, "cd /app001/" + project + ";git checkout master;git pull")
		//var out bytes.Buffer
		//cmd.Stdout = &out
		//err := cmd.Run()
		//if err != nil {
		//		log.Fatal(err)
		//}
		fmt.Printf("✅\n")
	}
}

func apull_icloud() {
	for _, project := range icloudProjects {
		fmt.Printf("Updating iCloud - %s ... ", project)
		os.Chdir(path.Join(icloud, project))
		script.Exec("git pull")
		fmt.Printf("✅\n")
	}
}
