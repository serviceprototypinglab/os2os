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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
	"os"
)


// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Download your project",
	Long: `Download the templates of all the objects or resources of your project.
		   The templates of the objects (deployments, pods, services, ...) will be save
		   in the path indicated (./templates by default)`,
	Run: func(cmd *cobra.Command, args []string) {
		export1(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")
	//exportCmd.PersistentFlags().StringVarP(&Path, "path","", "./templates", "path where export the templates")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func export(cmd *cobra.Command, args []string) {


	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(Path, os.FileMode(0777)) //All permision??
	changeProject(Project)
	//TODO Do it for all the objects.

	/* "cluster", "imagestreamimage", "petset", "componentstatus"
	objectsOc := []string{"buildconfig", "build", "configmap", "daemonset","daemonset","deployment", "deploymentconfig",
		"event","endpoints","horizontalpodautoscaler","imagestream","imagestreamtag","ingress","group","job",
		"limitrange","node","namespace","pod","persistentvolume","persistentvolumeclaim","policy","project","quota",
		"resourcequota","replicaset","replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}*/
	//iterate for each Openshift resource
	for _, typeObject := range ObjectsOc {
		typeString := getObjects(typeObject)
		if typeString != "" {
			//Create a folder for each resource
			os.Mkdir(Path+"/"+typeObject, os.FileMode(0777))
			//Take all the names of the resource
			namesDeployments := filterTableFirstColumn(typeString)
			for _, v := range namesDeployments {
				//Export template of the resource
				exportObject(typeObject, v)
			}
		}
	}
	fmt.Println("Templates created")
}

func export1(cmd *cobra.Command, args []string) {


	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(Path, os.FileMode(0777)) //All permision??
	changeProject(Project)

	for _, typeObject := range ObjectsOc {
		typeString := getObjects1(typeObject)
		items := typeString
		// TODO get items
		fmt.Println(items)
		byt := []byte(typeString)
		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		//fmt.Println(dat["items"].([]interface{})[0].(map[string]interface{})["kind"])
		/*for i, v := range dat["items"]{
			fmt.Println(v["kind"])
		}*/

		fmt.Println("-----")
		fmt.Println(dat)
		if items != "" {
			//Create a folder for each resource
			os.Mkdir(Path+"/"+typeObject, os.FileMode(0777))
			//Take all the names of the resource
			for i, _ := range items {
				//write json
				fmt.Println(i)
				break
			}
		}
	}
	fmt.Println("Templates created")
}

func getObjects1(typeObject string) string {
	CmdGetDeployments := exec.Command("oc", "get", typeObject, "-o", "json")
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("getObjects error in type " + typeObject)
		return ""
	}
	//checkErrorMessage(err, "Error running get " + typeObject)
	return string(CmdOut)
}

func loginCluster(cluster, username, password string) {
	username = "--username=" + username
	password = "--password=" + password
	CmdLogin := exec.Command("oc", "login", cluster, username, password)
	CmdOut, err := CmdLogin.Output()
	checkErrorMessage(err, "Error running login")
	fmt.Println(string(CmdOut))
}

//From the type of a resource return a table with the info
func getObjects(typeObject string) string {
	CmdGetDeployments := exec.Command("oc", "get", typeObject)
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("getObjects error in type " + typeObject)
		return ""
	}
	//checkErrorMessage(err, "Error running get " + typeObject)
	return string(CmdOut)
}

func changeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	checkErrorMessage(err, "Error running: change project")
	fmt.Println(string(CmdProjectOut))
}

//From the table with all the info, it filter
func filterTableFirstColumn(table string) []string {
	OutPutStrings := strings.Split(table,"\n")
	res := make([]string, 0)
	for _, v := range OutPutStrings {
		if v != "" {
			nameObject := strings.Fields(v)[0]
			if nameObject != "" && nameObject != "NAME" {
				res = append(res, nameObject)
			}
		}
	}
	return res
}

func exportObject(typeObject, nameObject string) {
	CmdGetDeployments := exec.Command("oc", "export", typeObject, nameObject, "-o", "json")
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("Error with the object " + typeObject + " called " + nameObject)
		return
	}
	//checkError(err)
	f, err := os.Create(Path+"/"+typeObject+"/"+ nameObject+".json")
	//checkError(err)
	if err != nil {
		fmt.Println("Error with the object " + typeObject + " called " + nameObject)
		return
	}
	f.WriteString(string(CmdOut))
	f.Sync()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkErrorMessage(err error, message string){
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}