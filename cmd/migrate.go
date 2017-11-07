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
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Full migration",
	Long: `All the steps in a full migration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		migrate(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func migrate(cmd *cobra.Command, args []string){
	fmt.Println("Starting the migration")
	fmt.Println("1. Exporting the templates ...")
	export(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("2. Exporting the data ...")
	exportData(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("3. Creating the objects ...")
	up(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("4. Uploading the data ...")
	upData(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("5. Deleting the data ...")
	downData(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("6. Deleting the objects...")
	down(cmd, args)
	fmt.Println("------------*---------------")
	fmt.Println("Finished the migration")
}