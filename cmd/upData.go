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
	"log"
	"path/filepath"
	"os"
	"encoding/json"
	"strings"
	"io/ioutil"
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
		//listVolumesPods()
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

	getAllValue()
	loginCluster(ClusterTo, UsernameTo, PasswordTo)
	os.Mkdir(PathData, os.FileMode(0777)) //All permission??
	changeProject(ProjectTo)


	data := readJsonData("./volumes")

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
			//Take the name of the object
			for i := range items {
				var podName string
				nameObjectOsAux, ok :=
					items[i].(map[string]interface{})["metadata"].
					(map[string]interface{})["name"].(string)
				if ok {
					podName = nameObjectOsAux
				} else {
					podName = typeObject + string(i)

				}
				//Create a folder for each deployment
				deploymentName, _ := getDeploymentReplicaSet(podName)

				//FIND DEPLOYMENT AND PROJECT NAME
				for _, a := range data {
					if a["deploymentName"] == deploymentName {
						path := "./volumes/" + deploymentName + "/" + a["podName"].(string) + "/" +
							a["volumeName"].(string)
						mountPath := a["mountPath"].(string)
						upDataToVolume(podName, path, mountPath)
					}
				}
			/*
				for _, v := range listDeployments() {
					deployment := getDeploymentName(v)
					if deployment == deploymentName {
						for _, podName := range listPods(deploymentName) {

							volumes := listVolumes(deploymentName+"/"+getPodNameFromPath(podName))
							for _, volumePath := range volumes {
								fmt.Println(volumePath)
							}
						}
					}
				}
				//fmt.Println(podName)
				/*var volumeName string
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
										createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
											descriptionVolume, descriptionVolumeMount)
										os.Mkdir(pathVolume + "/data", os.FileMode(0777))
										exportDataFromVolume(podName, pathVolume, mountPath)
									}
								}
							}
						}
					}
				}*/
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}


func getPodName(deploymentName string) string {
	return deploymentName
}

func upDataToVolume(podName, path, mountPath string) {
	a := "oc rsync " +  path + "/data/"  +  " " + podName + ":" + mountPath + "/"
	fmt.Println(a)
	cmdUpData := exec.Command("oc", "rsync", path + "/data/", podName + ":" + mountPath + "/")
	cmdUpOut, err := cmdUpData.Output()
	if err != nil {
		fmt.Println("Error migrating " + a)
		fmt.Println(err)
	} else {
		fmt.Println(string(cmdUpOut))
	}
}

func listPods(deploymentName string) []string{
	files, err := filepath.Glob("./volumes/"+deploymentName+"/*")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func listVolumes(podpath string) []string{
	files, err := filepath.Glob("./volumes/"+podpath+"/*")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func listDeployments() []string{
	files, err := filepath.Glob("./volumes/*")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func getDeploymentName(path string) string {
	auxString := strings.Split(path, "/")
	deploymentName := auxString[1]
	return deploymentName
}

func getPodNameFromPath(path string) string {
	auxString := strings.Split(path, "/")
	deploymentName := auxString[2]
	return deploymentName
}

func readJsonData(path string) []map[string]interface{} {
	plan, _ := ioutil.ReadFile(path + "/data.json")
	//var data []interface{}
	var data []map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}