// Prototype (☺) 2017 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Copyright (©) 2017 Zürcher Hochschule für Angewandte Wissenschaften
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
		export(cmd, args)
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

	// List of type of objects to export
	/*	if len(ObjectsOc) == 0 {
		ObjectsOc = []string{"service", "deployment", "secrets", "configmap", "job", "namespace"}
	} else if ObjectsOc[0] == "default" {
		ObjectsOc = []string{"service", "deployment", "secrets", "configmap", "job", "namespace"}
	} else if ObjectsOc[0] == "all" {
		ObjectsOc = []string{"service", "buildconfig", "build", "configmap", "daemonset","daemonset","deployment",
			"deploymentconfig",
			"event","endpoints","horizontalpodautoscaler","imagestream","imagestreamtag","ingress","group","job",
			"limitrange","node","namespace","pod","persistentvolume","persistentvolumeclaim","policy","project","quota",
			"resourcequota","replicaset","replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}
	} else {
		ObjectsOc = strings.Split(ObjectsOc[0], ",")
	}*/

	ObjectsOc = getTypeObjects(ObjectsOc)

	// Login in the cluster and change the project
	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(Path, os.FileMode(0777)) //All permision??
	changeProject(ProjectFrom)

	for _, typeObject := range ObjectsOc {
		fmt.Println("Starting exporting the objects with kind: " + typeObject)
		var dat map[string]interface{}

		//Get all the objects for the type: typeObject and parse to json
		typeString := getObjects(typeObject)
		byt := []byte(typeString)
		if err1 := json.Unmarshal(byt, &dat); err1 != nil {
			fmt.Println("Error with the objects with type " + typeObject)
			fmt.Println("-------")
			if typeString != "" {
				fmt.Println(typeString)
			}
		} else {
			items := dat["items"].([]interface{})
			if len(items) > 0 {
				//Create a folder for each type
				os.Mkdir(Path+"/"+typeObject, os.FileMode(0777))
				//Take the name of the object
				for i := range items {
					var nameObjectOs string
					nameObjectOsAux, ok :=
						items[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)
					if ok {
						nameObjectOs = nameObjectOsAux
					} else {
						nameObjectOs = typeObject + string(i)

					}

					//Copy the json to a file
					objectOs, err2 := json.Marshal(items[i])

					if err2 != nil {
						fmt.Println("Error parsing json for the "  + typeObject + " called " + nameObjectOs)
					} else {
						f, err3 := os.Create(Path + "/" + typeObject + "/" + nameObjectOs + ".json")
						//checkError(err)
						if err3 != nil {
							fmt.Println("Error with the " + typeObject + " called " + nameObjectOs)
						} else {
							f.WriteString(string(objectOs))
							f.Sync()
							fmt.Println("Exported the " + typeObject + " called " + nameObjectOs)
						}
					}

				}
			} else {
				fmt.Println("No objects for the type " + typeObject)
			}
			fmt.Println("-----------")
		}
	}
	fmt.Println("Templates created")
}

func getObjects(typeObject string) string {
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

func changeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	checkErrorMessage(err, "Error running: change project")
	fmt.Println(string(CmdProjectOut))
}

func getTypeObjects(ObjectsTypes []string) []string {
	// List of type of objects to export
	if len(ObjectsTypes) == 0 {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job", "namespace"}
	} else if ObjectsTypes[0] == "default" {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job", "namespace"}
	} else if ObjectsOc[0] == "all" {
		ObjectsTypes = []string{"service", "buildconfig", "build", "configmap", "daemonset","daemonset",
			"deployment", "deploymentconfig", "event","endpoints","horizontalpodautoscaler","imagestream",
			"imagestreamtag","ingress","group","job", "limitrange","node","namespace","pod","persistentvolume",
			"persistentvolumeclaim","policy","project","quota", "resourcequota","replicaset",
			"replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}
	} else {
		ObjectsTypes = strings.Split(ObjectsTypes[0], ",")
	}
	return ObjectsTypes

}

//CHECK ERRORS

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

// ALL METHODS

func export1(cmd *cobra.Command, args []string) {


	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(Path, os.FileMode(0777)) //All permision??
	changeProject(ProjectFrom)
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
//From the type of a resource return a table with the info
func getObjects1(typeObject string) string {
	CmdGetDeployments := exec.Command("oc", "get", typeObject)
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("getObjects error in type " + typeObject)
		return ""
	}
	//checkErrorMessage(err, "Error running get " + typeObject)
	return string(CmdOut)
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
