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
	"os"
	"syscall"
)

var Path string
var Project string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Download all the objects in your project",
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		export(cmd, args)
	},
}

func init() {

	RootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	exportCmd.PersistentFlags().StringVarP(&Path, "path","", "./templates", "path where export the templates")
	exportCmd.PersistentFlags().StringVarP(&Project, "project", "p", "myproject", "name of the Openshift project")


	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func export(cmd *cobra.Command, args []string){

	fmt.Println("exporting ...")
	binary, lookErr := exec.LookPath("oc")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("Using the binary in " + binary)
	args1 := []string{"oc", "project", Project}
	env := os.Environ()
	execErr := syscall.Exec(binary, args1, env)
	if execErr != nil {
		panic(execErr)
	}

	args1 := []string{"oc", "export", ""}
}