/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
)

// apullCmd represents the apull command
var apullCmd = &cobra.Command{
	Use:   "apull",
	Short: "Pulls the lastest repository version",
	Long: `Pulss the lastest version of a group of Ansible repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		run_apull()
	},
}

func init() {
	rootCmd.AddCommand(apullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func run_apull() {
	remoteServer := "ori1020lp"
	remoteProjects := []string{"ansible", "ansible_dev/lezottt", "ansible_dev/ori_versions"}
	icloudProjects := []string{"ansible", "docker-centos7-ansible", "docker-centos8-ansible", "images"}
	//icloud := "$HOME/Library/Mobile Documents/com~apple~CloudDocs/_Code/VICTR"

	os.Chdir("$HOME/code/ansible/ansible")

	fmt.Printf("Updating MBP - ansible ... ")
	script.Exec("git checkout master")
	script.Exec("git pull")
	fmt.Printf("Done\n")

	for _, project := range remoteProjects {
		fmt.Printf("Updating %s - %s ... ", remoteServer, project)
		//cmd := exec.Command("ssh", remoteServer, "cd /app001/ansible;git checkout master;git pull")
		//var out bytes.Buffer
		//cmd.Stdout = &out
		//err := cmd.Run()
		//if err != nil {
		//		log.Fatal(err)
		//}
		fmt.Printf("Done\n")
	}

	for _, p := range icloudProjects {
		fmt.Printf("Updating iCloud - %s ... ", p)
		//os.Chdir(icloud.Join(x))
		//script.Exec("git pull")
		fmt.Printf("Done\n")
	}
}
