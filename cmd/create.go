/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var InputFile string

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new resource",
	Long:  `Allows to create a new resource of a specific type`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"create", "new"},
	ValidArgs:  []string{"tag"},
}

func init() {
	rootCmd.AddCommand(CreateCmd)
	CreateCmd.PersistentFlags().
		StringVarP(&InputFile, "file", "f", "", `file or stdin (-) as input for resource creation. Input type might be either json or yaml.`)
}

func ValidateCreateInput() {
	if InputFile != "" && InputFile != "-" {
		if _, err := os.Stat(InputFile); os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("File '%v' for input cannot be accessed due to error %v", InputFile, err))
		}
	}
}

func OpenStdinOrFile() []byte {
	if InputFile == "-" {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(fmt.Sprintf("An error occured while reading from stdin: %v", err))
		}
		return probeYamlToJsonConverion(bytes)
	}
	if InputFile != "" {
		reader, err := os.Open(InputFile)
		if err != nil {
			log.Fatal(fmt.Sprintf("could not open file '%v' due to error %v", InputFile, err))
		}
		bytes, err := io.ReadAll(reader)
		if err != nil {
			log.Fatal(fmt.Sprintf("An error occured while reading from file: %v", err))
		}
		return probeYamlToJsonConverion(bytes)
	}
	return nil
}

func probeYamlToJsonConverion(content []byte) []byte {
	convertedFromYamlToJson, err := yaml.YAMLToJSON(content)
	if err != nil {
		return content
	}
	return convertedFromYamlToJson
}
