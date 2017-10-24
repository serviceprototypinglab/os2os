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
	"os"
	//"os/exec"
	"encoding/json"
	//"strings"
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
	loginCluster(ClusterTo, UsernameFrom, PasswordFrom)
	os.Mkdir(PathData, os.FileMode(0777)) //All permision??
	changeProject(ProjectFrom)

	var dat map[string]interface{}
	typeObject := "pods"
	typeString := getObjects(typeObject)

	fmt.Println(typeString)

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
				os.Mkdir(PathData+"/"+podName, os.FileMode(0777))
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
										//fmt.Println("-----1------")
										//fmt.Println(volumeName)
										//fmt.Println(descriptionVolume)
										descriptionVolumeMount := volumesMountAux[k].(map[string]interface{})
										//fmt.Println(descriptionVolumeMount)
										//fmt.Println(podName)
										mountPath := volumesMountAux[k].(map[string]interface{})["mountPath"].(string)
										//fmt.Println(mountPath)
										rsName := getReplicaSet(podName)
										deploymentName := getDeployment(podName)
										//fmt.Println(rsName)
										//fmt.Println(deploymentName)
										pathVolume := PathData+"/"+podName + "/" + volumeName
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
				}
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}

//TODO
func getDeployment(pod string) string {
	return "todo"
}

//TODO
func getReplicaSet(pod string) string {
	return "todo"
}

func exportDataFromVolume(pod string, path string, mountPath string) {
	a := "oc rsync " + pod + ":" + mountPath +  " " + path
	fmt.Println(a)

}

func createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName string,
	descriptionVolume, descriptionVolumeMount map[string]interface{}){


	/*m := make(map[string]string)

	m["pathVolume"] = pathVolume
	fmt.Println(volumeName)
	f, err3 := os.Create(pathVolume + "/data.json")
	//Copy the json to a file
	type DataJson struct {
		podName string
	}



	dataJson := DataJson{podName}
	fmt.Println(dataJson)
	objectOs, err2 := json.Marshal(dataJson)
	fmt.Println(err2)
	fmt.Println(objectOs)
	objectOs, err2 := json.Marshal(m)
	fmt.Println(err2)
	if err3 != nil {
		fmt.Println("Error creating data.json")
	} else {
		f.WriteString(string(objectOs))
		//f.Write(dataJson)
		f.Sync()
		fmt.Println("Created  data.json in " + pathVolume)
	}*/

}