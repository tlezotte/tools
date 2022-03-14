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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the lastest version for a repository",
	Long:  `Pull the lastest version for a repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		localFlag, _ := cmd.Flags().GetBool("local")
		icloudFlag, _ := cmd.Flags().GetBool("icloud")
		remoteFlag, _ := cmd.Flags().GetBool("remote")
		if localFlag {
			pullLocal()
		} else if icloudFlag {
			pullIcloud()
		} else if remoteFlag {
			pullRemote()
		} else {
			pullLocal()
			pullIcloud()
			pullRemote()
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.
	pullCmd.Flags().BoolP("local", "l", false, "Pull repository updates on local ansible")
	pullCmd.Flags().BoolP("icloud", "i", false, "Pull repository updates on all icloud projects")
	pullCmd.Flags().BoolP("remote", "r", false, "Pull repository updates on some remote projects")
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
func pullLocal() {
	pathName := "code/ansible/ansible"
	fullPath := path.Join(homeDir, pathName)

	os.Chdir(fullPath)
	fmt.Printf("Updating local - ansible ... ")
	script.Exec("git checkout master")
	script.Exec("git pull")
	fmt.Printf("✅\n")
}

// Pull the lastest version for remote repositories
func pullRemote() {
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
func pullIcloud() {
	pathName := "Library/Mobile Documents/com~apple~CloudDocs/_Code/VICTR"
	fullPath := path.Join(homeDir, pathName)

	for _, project := range getProjects(fullPath) {
		fmt.Printf("Updating iCloud - %s ... ", project)
		os.Chdir(path.Join(fullPath, project))
		script.Exec("git pull")
		fmt.Printf("✅\n")
	}
}
