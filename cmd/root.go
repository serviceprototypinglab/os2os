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
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ClusterFrom  string
	ClusterTo    string
	ProjectFrom  string
	ProjectTo    string
	PathTemplate string
	PathData     string
	UsernameFrom string
	UsernameTo   string
	PasswordFrom string
	PasswordTo   string
	cfgFile      string
)
var ObjectsOc []string

/*var ObjectsOc = []string{"service", "buildconfig", "build", "configmap", "daemonset","daemonset","deployment",
	"deploymentconfig",
	"event","endpoints","horizontalpodautoscaler","imagestream","imagestreamtag","ingress","group","job",
	"limitrange","node","namespace","pod","persistentvolume","persistentvolumeclaim","policy","project","quota",
	"resourcequota","replicaset","replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}
/* "cluster", "imagestreamimage", "petset", "componentstatus"
objectsOc := []string{"buildconfig", "build", "configmap", "daemonset","daemonset","deployment", "deploymentconfig",
	"event","endpoints","horizontalpodautoscaler","imagestream","imagestreamtag","ingress","group","job",
	"limitrange","node","namespace","pod","persistentvolume","persistentvolumeclaim","policy","project","quota",
	"resourcequota","replicaset","replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}*/
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "os2os",
	Short: "Migrate your Openshift application between different Openshift clusters",
	Long: `
os2os is a tool for help you to migrate a Openshift project between different Openshift clusters.
You can download all your templates, convert and rigth size the application to fix in the new cluster,
migrate the data, deploy your app in the new cluster and delete your project in the old cluster`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	/*
	  TODO Read config file and set the default value if exits:
	  TODO using the config file value
	  TODO in other case:
	  TODO using a hardcoder value */

	//fmt.Println(viper.AllKeys())
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.os2os.yaml)")
	RootCmd.PersistentFlags().StringVarP(&ClusterFrom, "clusterFrom", "", "https://192.168.99.100:8443", "Cluster where is the project that you want to migrate")
	RootCmd.PersistentFlags().StringVarP(&ClusterTo, "clusterTo", "", "https://127.0.0.1:8443", "Cluster where you want to migrate the project")
	RootCmd.PersistentFlags().StringVarP(&ProjectFrom, "projectFrom", "", "myproject", "name of the old Openshift project")
	RootCmd.PersistentFlags().StringVarP(&ProjectTo, "projectTo", "", "myproject", "name of the new Openshift project")
	RootCmd.PersistentFlags().StringVarP(&UsernameFrom, "usernameFrom", "", "developer", "username in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&UsernameTo, "usernameTo", "", "developer", "username in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PasswordFrom, "passwordFrom", "", "developer", "password in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&PasswordTo, "passwordTo", "", "developer", "password in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PathTemplate, "pathTemplate","", "./templates", "path where export the templates")
	RootCmd.PersistentFlags().StringVarP(&PathData, "pathData","", "./volumes", "path where export the volumes")
	defaultValue := []string{"default"}
	RootCmd.PersistentFlags().StringArrayVarP(&ObjectsOc, "objects", "o", defaultValue, "list of objects to export" )
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("clusterFrom", RootCmd.PersistentFlags().Lookup("clusterFrom"))
	viper.BindPFlag("ClusterTo", RootCmd.PersistentFlags().Lookup("ClusterTo"))
	viper.BindPFlag("ProjectFrom", RootCmd.PersistentFlags().Lookup("ProjectFrom"))
	viper.BindPFlag("ProjectTo", RootCmd.PersistentFlags().Lookup("ProjectTo"))
	viper.BindPFlag("UsernameFrom", RootCmd.PersistentFlags().Lookup("UsernameFrom"))
	viper.BindPFlag("UsernameTo", RootCmd.PersistentFlags().Lookup("UsernameTo"))
	viper.BindPFlag("PasswordFrom", RootCmd.PersistentFlags().Lookup("PasswordFrom"))
	viper.BindPFlag("PasswordTo", RootCmd.PersistentFlags().Lookup("PasswordTo"))
	viper.BindPFlag("PathTemplate", RootCmd.PersistentFlags().Lookup("PathTemplate"))
	viper.BindPFlag("PathData", RootCmd.PersistentFlags().Lookup("PathData"))
	viper.BindPFlag("ObjectsOc", RootCmd.PersistentFlags().Lookup("ObjectsOc"))

	fmt.Println(viper.AllKeys())



}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".os2os" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".os2os")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		initParametersFromConfigFile()

	} else {
		fmt.Println("Error reading config file")
	}
}

func initParametersFromConfigFile() {
	for _, keyConfig := range viper.AllKeys() {
		switch keyConfig {
		case "pathtemplate":
			PathTemplate = viper.GetString(keyConfig)
		case "pathdata":
			PathData = viper.GetString(keyConfig)
		case "clusterto":
			ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			ProjectFrom = viper.GetString(keyConfig)
		case "usernamefrom":
			UsernameFrom = viper.GetString(keyConfig)
		case "usernameto":
			UsernameTo = viper.GetString(keyConfig)
		case "passwordfrom":
			PasswordFrom = viper.GetString(keyConfig)
		case "passwordto":
			PasswordTo = viper.GetString(keyConfig)
		case "objectsoc":
			ObjectsOc = viper.GetStringSlice(keyConfig)
		}
	}
}
