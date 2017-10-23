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
)

// migrateDataCmd represents the migrateData command
var migrateDataCmd = &cobra.Command{
	Use:   "migrateData",
	Short: "migrate a persistant volume between the two OpenShift clusters",
	Long: `migrate a persistant volume between the two OpenShift clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateData called")
	},
}




func init() {
	RootCmd.AddCommand(migrateDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
