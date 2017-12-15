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
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"encoding/json"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert your template to adapt to your new cluster",
	Long: `Convert your template to adapt to your new cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert called")
		convert()
	},
}

func init() {
	RootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func convert(){
	getAllValue()
	convert_project()
}

func convert_project() {

	if ProjectFrom == ProjectTo {
		fmt.Println("Same project")
	} else {
		PathTemplateTo := PathTemplate + "/" + ProjectTo
		os.Mkdir(PathTemplateTo, os.FileMode(0777))

		PathTemplate += "/" + ProjectFrom


		ObjectsOc = getTypeObjects(ObjectsOc)
		for _, object := range ObjectsOc {
			//fmt.Println(object)
			files, err := ioutil.ReadDir(PathTemplate + "/" + object)
			if err != nil {
				fmt.Println("not " + PathTemplate + "/" + object)
			}
			os.Mkdir(PathTemplateTo + "/" + object, os.FileMode(0777))
			for _, f := range files {

				fmt.Println(f.Name())
				//metadata.namespace = projectTo
				file, e := ioutil.ReadFile(PathTemplate + "/" + object + "/" + f.Name())
				if e != nil {
					fmt.Printf("File error: %v\n", e)
					os.Exit(1)
				}
				typeString := string(file)
				//fmt.Printf("%s\n", string(file))
				byt := []byte(string(file))
				var dat map[string]interface{}
				if err1 := json.Unmarshal(byt, &dat); err1 != nil {
					fmt.Println("Error with the objects with type " + object)
					fmt.Println("-------")
					if typeString != "" {
						fmt.Println(typeString)
					}
				} else {
					dat["metadata"].(map[string]interface{})["namespace"] = ProjectTo
					os.Mkdir(PathTemplateTo + "/" + object, os.FileMode(0777))

					f1, err3 := os.Create(PathTemplateTo + "/" + object + "/" + f.Name())
					//checkError(err)
					if err3 != nil {
						fmt.Println("Error with " + f.Name())
					} else {
						objectOs, err2 := json.Marshal(dat)
						if err2 != nil {
							fmt.Println("erro json marshal")
						} else {
							f1.WriteString(string(objectOs))
							f1.Sync()
						}
					}
				}
			}
		}
	}

}