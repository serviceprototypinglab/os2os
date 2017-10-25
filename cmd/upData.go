// Copyright © 2017 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

// upDataCmd represents the upData command
var upDataCmd = &cobra.Command{
	Use:   "upData",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upData called")
		upData(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(upDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func upData(cmd *cobra.Command, args []string) {
	loginCluster(ClusterTo, UsernameFrom, PasswordFrom)

	//Get pods from deployment
	deploymentName := "A"
	//TODO GET ALL THE VOLUME TO MIGRATE
	podName := getPodName(deploymentName)
	//Copy the data there
	upDataToVolume(podName, podName, podName)
}

func getPodName(deploymentName string) string {
	return deploymentName
}

func upDataToVolume(podName, path, mountPath string) {
	a := "oc rsync " +  path + "/data"  +  " " + podName + ":" + mountPath
	fmt.Println(a)
	cmdUpData := exec.Command("oc", "rsync", path + "/data", podName + ":" + mountPath)
	cmdUpOut, err := cmdUpData.Output()
	if err != nil {
		fmt.Println("Error migrating " + a)
		fmt.Println(err)
	} else {
		fmt.Println(string(cmdUpOut))
	}
}