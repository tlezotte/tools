/*
Copyright © 2022 tools tom <tom@lezotte.net>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
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
		localFlag, _ := cmd.Flags().GetBool("local")
		icloudFlag, _ := cmd.Flags().GetBool("icloud")
		remoteFlag, _ := cmd.Flags().GetBool("remote")
		if localFlag {
			apullLocal()
		} else if icloudFlag {
			apullIcloud()
		} else if remoteFlag {
			apullRemote()
		} else {
			apullLocal()
			apullIcloud()
			apullRemote()
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
	remoteServer   = "ori1020lp.hs.it.vumc.io"
	remoteProjects = []string{
		"ansible",
		"ansible_dev/lezottt",
		"ansible_dev/ori_versions",
	}
	homeDir, _ = os.UserHomeDir()
)

// get a list of directories in a directory
func getProjects(parent string) []string {
	var projects []string
	directories, err := ioutil.ReadDir(parent)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range directories {
		if fileInfo.IsDir() {
			projects = append(projects, fileInfo.Name())
		}
	}
	return projects
}

// Pull the lastest version for local repositories
func apullLocal() {
	pathName := "code/ansible/ansible"
	fullPath := path.Join(homeDir, pathName)

	os.Chdir(fullPath)
	fmt.Printf("Updating local - ansible ... ")
	script.Exec("git checkout master")
	script.Exec("git pull")
	fmt.Printf("✅\n")
}

// Pull the lastest version for remote repositories
func apullRemote() {
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

// Pull the lastest version for icloud repositories
func apullIcloud() {
	pathName := "Library/Mobile Documents/com~apple~CloudDocs/_Code/VICTR"
	fullPath := path.Join(homeDir, pathName)

	for _, project := range getProjects(fullPath) {
		fmt.Printf("Updating iCloud - %s ... ", project)
		os.Chdir(path.Join(fullPath, project))
		script.Exec("git pull")
		fmt.Printf("✅\n")
	}
}
