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

			//Take the name of the object
			for i := range items {
				var namePod string
				nameObjectOsAux, ok :=
					items[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)
				if ok {
					namePod = nameObjectOsAux
				} else {
					namePod = typeObject + string(i)

				}
				//Create a folder for each deployment
				os.Mkdir(PathData+"/"+namePod, os.FileMode(0777))
				//fmt.Println(namePod)
				var volumeName string
				volumesAux, ok :=
					items[i].(map[string]interface{})["spec"].(map[string]interface{})["volumes"].([]interface{})
				if ok {
					for j := range volumesAux {
						volumeName = volumesAux[j].(map[string]interface{})["name"].(string)
						//fmt.Println(volumeName)
						descriptionVolume := volumesAux[j]
						fmt.Println("-------")
						//fmt.Println(descriptionVolume)
						volumesMountAuxs, ok1 := items[i].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})
						var volumesMountAux = volumesAux
						for u := range volumesMountAuxs {
							if ok1 {
								volumesMountAux := volumesMountAuxs[u].(map[string]interface{})["volumeMounts"].([]interface{})
								fmt.Println(volumesMountAux)
							}

						}
						if ok1 {
							 for k := range volumesMountAux {
								nameVolumeMount := volumesMountAux[k].(map[string]interface{})["name"].(string)
								if nameVolumeMount == volumeName {
									fmt.Println("-----")
									fmt.Println(volumeName)
									fmt.Println(descriptionVolume)
									descriptionVolumeMount := volumesMountAux[k].(map[string]interface{})
									fmt.Println(descriptionVolumeMount)
									fmt.Println(namePod)
									mountPath := volumesMountAux[k].(map[string]interface{})["mountPath"].(string)
									fmt.Println(mountPath)
								}
							}
						}

					}
				} else {
					volumeName = ""
				}

			}
		}else {
			fmt.Println("No objects for the type " + typeObject)
		}
		fmt.Println("-----------")
	}



}