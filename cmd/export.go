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
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
	"os"
)

var Path string
var Project string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Download your project",
	Long: `Download the templates of all the objects or resources of your project.
		   The templates of the objects (deployments, pods, services, ...) will be save
		   in the path indicated (./templates by default)`,
	Run: func(cmd *cobra.Command, args []string) {
		export(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")
	exportCmd.PersistentFlags().StringVarP(&Path, "path","", "./templates", "path where export the templates")
	exportCmd.PersistentFlags().StringVarP(&Project, "project", "p", "myproject", "name of the Openshift project")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func export(cmd *cobra.Command, args []string) {
	os.Mkdir(Path, os.FileMode(0777)) //All permision

	changeProject(Project)
	//TODO Do it for all the objects.
	objectsOc := []string{"deployment", "service"}
	for _, typeObject := range objectsOc {
		typetString := getObjects(typeObject)
		os.Mkdir(Path+"/"+typeObject, os.FileMode(0777))
		namesDeployments := filterTableFirstColumn(typetString)
		for _, v := range namesDeployments {
			exportObject(typeObject, v)
		}
	}

}

func getObjects(typeObject string) string {
	if typeObject == "deployment" {
		CmdGetDeployments := exec.Command("oc", "get", "deployments")
		CmdOut, err := CmdGetDeployments.Output()
		if err != nil {
			fmt.Println("Error running CmdGetDeployments")
			fmt.Println(err)
			panic(err)
		}
		return string(CmdOut)
	}

	if typeObject == "service" {
		CmdGetDeployments := exec.Command("oc", "get", "services")
		CmdOut, err := CmdGetDeployments.Output()
		if err != nil {
			fmt.Println("Error running CmdGetDeployments")
			fmt.Println(err)
			panic(err)
		}
		return string(CmdOut)
	}

	return ""
}

func changeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	if err != nil {
		fmt.Println("Error running change project")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(CmdProjectOut))
}

func filterTableFirstColumn(table string) []string {
	OutPutStrings := strings.Split(table,"\n")
	res := make([]string, 0)
	//fmt.Println(string(CmdGetDeploymentsOut))
	//fmt.Println(OutPutStrings)
	for _, v := range OutPutStrings {
		if v != "" {
			nameObject := strings.Fields(v)[0]
			if nameObject != "" && nameObject != "NAME"{
				res = append(res, nameObject)
			}
		}
	}
	return res
}

func exportObject(typeObject, nameObject string) {
	if typeObject == "deployment" {
		CmdGetDeployments := exec.Command("oc", "export", typeObject, nameObject, "-o", "json")
		CmdOut, err := CmdGetDeployments.Output()
		check(err)
		f, err := os.Create(Path+"/"+typeObject+"/"+ nameObject+".json")
		check(err)
		f.WriteString(string(CmdOut))

		f.Sync()
	}

	if typeObject == "service" {
		CmdGetDeployments := exec.Command("oc", "export", typeObject, nameObject, "-o", "json")
		CmdOut, err := CmdGetDeployments.Output()
		check(err)
		f, err := os.Create(Path+"/"+typeObject+"/"+ nameObject+".json")
		check(err)
		f.WriteString(string(CmdOut))
		f.Sync()
	}
}

func check(err error){
	if err != nil {
		panic(err)
	}
}

