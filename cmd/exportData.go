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

//TODO Create a json with the information of the volumes in PathData" +

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"encoding/json"
	"strings"
	//"github.com/openshift/origin/test/extended/util/db"
)

// exportDataCmd represents the exportData command
var exportDataCmd = &cobra.Command{
	Use:   "exportData",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exportData called")
		exportData(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(exportDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func exportData(cmd *cobra.Command, args []string) {

	//TODO change to clusterFrom
	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(PathData, os.FileMode(0777)) //All permision??
	changeProject(ProjectFrom)

	var dat map[string]interface{}
	typeObject := "pods"
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
			os.Mkdir(PathData, os.FileMode(0777))

			var a [] map[string]interface{}

			//Take the name of the object
			for i := range items {
				var podName string
				nameObjectOsAux, ok :=
					items[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)
				if ok {
					podName = nameObjectOsAux
				} else {
					podName = typeObject + string(i)

				}
				//Create a folder for each deployment
				deploymentName, rsName := getDeploymentReplicaSet(podName)
				os.Mkdir(PathData+"/"+deploymentName, os.FileMode(0777))
				os.Mkdir(PathData+"/"+deploymentName+"/"+podName, os.FileMode(0777))
				//fmt.Println(podName)
				var volumeName string
				volumesAux, ok :=
					items[i].(map[string]interface{})["spec"].(map[string]interface{})["volumes"].([]interface{})
				if ok {
					for j := range volumesAux {
						volumeName = volumesAux[j].(map[string]interface{})["name"].(string)
						//fmt.Println(volumeName)
						descriptionVolume := volumesAux[j].(map[string]interface{})
						//fmt.Println(descriptionVolume)
						volumesMountAuxs, ok1 := items[i].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})
						for u := range volumesMountAuxs {
							if ok1 {
								volumesMountAux := volumesMountAuxs[u].(map[string]interface{})["volumeMounts"].([]interface{})
								for k := range volumesMountAux {
									nameVolumeMount := volumesMountAux[k].(map[string]interface{})["name"].(string)
									if nameVolumeMount == volumeName {
										descriptionVolumeMount := volumesMountAux[k].(map[string]interface{})
										mountPath := volumesMountAux[k].(map[string]interface{})["mountPath"].(string)
										pathVolume := PathData+"/"+deploymentName+"/"+podName + "/" + volumeName
										os.Mkdir(pathVolume, os.FileMode(0777))
										aux := createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
											descriptionVolume, descriptionVolumeMount)
										a = append(a, aux)
										os.Mkdir(pathVolume + "/data", os.FileMode(0777))
										exportDataFromVolume(podName, pathVolume, mountPath)
									}
								}
							}
						}
					}
				}
			}
			f, err3 := os.Create(PathData +"/data.json")
			if err3 != nil {
				fmt.Println("Error creating data.json")
				fmt.Println(err3)
			} else {
				objectOs, err2 := json.Marshal(a)
				if err2 != nil {
					fmt.Println("Error creating the json object")

					fmt.Println(err2)
				} else {
					f.WriteString(string(objectOs))
					f.Sync()
					fmt.Println("Created  data.json in" + PathData )
				}
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}

func getDeploymentReplicaSet(pod string) (string, string) {
	auxString := strings.Split(pod, "-")
	deploymentName := auxString[0]
	replicaSetName := deploymentName + "-" + auxString[1]
	return deploymentName, replicaSetName
}

func exportDataFromVolume(pod string, path string, mountPath string) {
	a := "oc rsync " + pod + ":" + mountPath + "/" +  " " + path + "/data"
	fmt.Println(a)
	cmdExportData := exec.Command("oc", "rsync", pod + ":" + mountPath + "/", path + "/data")
	cmdExportOut, err := cmdExportData.Output()
	if err != nil {
		fmt.Println("Error migrating " + a)
		fmt.Println(err)
	} else {
		fmt.Println(string(cmdExportOut))
	}
}

func createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName string,
	descriptionVolume, descriptionVolumeMount map[string]interface{}) map[string]interface{} {

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["pathVolume"] = pathVolume
	m["volumeName"] = volumeName
	m["podName"] = podName
	m["mountPath"] = mountPath
	m["rsName"] = rsName
	m["deploymentName"] = deploymentName
	m["descriptionVolume"] = descriptionVolume
	m["descriptionVolumeMount"] = descriptionVolumeMount

	f, err3 := os.Create(pathVolume + "/data.json")

	if err3 != nil {
		fmt.Println("Error creating data.json")
		fmt.Println(err3)
	} else {
		objectOs, err2 := json.Marshal(m)
		if err2 != nil {
			fmt.Println("Error creating the json object")
			fmt.Println(err2)
		} else {
			f.WriteString(string(objectOs))
			f.Sync()
			fmt.Println("Created  data.json in " + pathVolume)
		}
	}
	return m
}

